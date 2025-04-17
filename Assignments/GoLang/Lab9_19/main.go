package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync" // Package sync provides synchronization primitives like WaitGroup
	"time"
)

// Game represents a video game
type Game struct {
	ID            int           `json:"id"`
	Title         string        `json:"title"`
	Genre         string        `json:"genre"`
	Publisher     string        `json:"publisher"`
	ReleaseYear   int           `json:"releaseYear"`
	Rating        float64       `json:"rating"`
	Price         float64       `json:"price"`
	IsMultiplayer bool          `json:"isMultiplayer"`
	FetchTime     time.Duration `json:"fetchTime"`              // Tracks how long it took to fetch this game
	BufferReused  bool          `json:"bufferReused,omitempty"` // Indicates if a buffer was reused for this game
}

// Result combines a game with any error that occurred
// Used for communication between goroutines via channels
type Result struct {
	Game Game
	Err  error
}

// Template data for rendering
type PageData struct {
	Games          []Game
	Timing         time.Duration // Total execution time
	Message        string        // For success/error messages
	IsJSONImported bool          // Flag to indicate if JSON is imported
	BuffersReused  int           // Count of buffers reused (for leaky buffer demo)
}

// Simulated database of games - will be replaced by JSON data when imported
var games = []Game{
	{ID: 1, Title: "Elden Ring", Genre: "Action RPG", Publisher: "FromSoftware", ReleaseYear: 2022, Rating: 9.5, Price: 59.99, IsMultiplayer: true},
	{ID: 2, Title: "God of War: Ragnarok", Genre: "Action-Adventure", Publisher: "Sony", ReleaseYear: 2022, Rating: 9.4, Price: 69.99, IsMultiplayer: false},
	{ID: 3, Title: "Cyberpunk 2077", Genre: "RPG", Publisher: "CD Projekt", ReleaseYear: 2020, Rating: 7.2, Price: 49.99, IsMultiplayer: false},
	{ID: 4, Title: "The Legend of Zelda: TOTK", Genre: "Action-Adventure", Publisher: "Nintendo", ReleaseYear: 2023, Rating: 9.8, Price: 59.99, IsMultiplayer: false},
	{ID: 5, Title: "Call of Duty: Modern Warfare III", Genre: "FPS", Publisher: "Activision", ReleaseYear: 2023, Rating: 8.2, Price: 69.99, IsMultiplayer: true},
	{ID: 6, Title: "FIFA 24", Genre: "Sports", Publisher: "EA Sports", ReleaseYear: 2023, Rating: 8.0, Price: 69.99, IsMultiplayer: true},
	{ID: 7, Title: "Starfield", Genre: "RPG", Publisher: "Bethesda", ReleaseYear: 2023, Rating: 8.7, Price: 69.99, IsMultiplayer: false},
	{ID: 8, Title: "Baldur's Gate 3", Genre: "RPG", Publisher: "Larian Studios", ReleaseYear: 2023, Rating: 9.6, Price: 59.99, IsMultiplayer: true},
}

// JSON file path for persistent storage
const jsonFilePath = "games_database.json"

// Mutex to protect games slice during concurrent operations
var gamesMutex sync.RWMutex

// BufferPool represents a pool of byte buffers
// This is the core of the leaky buffer pattern
type BufferPool struct {
	buffers chan []byte     // Channel to hold available buffers
	size    int             // Size of each buffer
	mutex   sync.Mutex      // Mutex to protect metrics
	stats   BufferPoolStats // Statistics for the buffer pool
}

// BufferPoolStats tracks statistics for the buffer pool
type BufferPoolStats struct {
	Created int // Number of buffers created
	Reused  int // Number of times buffers were reused
}

// Global buffer pool with a fixed capacity and buffer size
var bufferPool = NewBufferPool(10, 4096) // 10 buffers of 4KB each

// NewBufferPool creates a new buffer pool with the specified capacity and buffer size
func NewBufferPool(capacity, size int) *BufferPool {
	return &BufferPool{
		buffers: make(chan []byte, capacity),
		size:    size,
	}
}

// Get retrieves a buffer from the pool or creates a new one if none are available
func (p *BufferPool) Get() []byte {
	select {
	case buffer := <-p.buffers:
		// Reuse an existing buffer from the pool
		p.mutex.Lock()
		p.stats.Reused++
		p.mutex.Unlock()
		return buffer
	default:
		// No buffers available in the pool, create a new one
		p.mutex.Lock()
		p.stats.Created++
		p.mutex.Unlock()
		return make([]byte, p.size)
	}
}

// Put returns a buffer to the pool if there's room, otherwise it's discarded
// and will be garbage collected (this is the "leaky" part)
func (p *BufferPool) Put(buffer []byte) {
	select {
	case p.buffers <- buffer:
		// Buffer successfully returned to the pool
	default:
		// Pool is full, buffer is "leaked" (will be garbage collected)
	}
}

// GetStats returns the current statistics for the buffer pool
func (p *BufferPool) GetStats() BufferPoolStats {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.stats
}

// Initialize games from JSON file if it exists
func initGamesFromJSON() {
	if _, err := os.Stat(jsonFilePath); err == nil {
		// File exists, load games from JSON
		loadGamesFromJSON()
	} else {
		// File doesn't exist, create it with default games
		saveGamesToJSON()
	}
}

// Load games from JSON file
func loadGamesFromJSON() error {
	file, err := os.Open(jsonFilePath)
	if err != nil {
		return fmt.Errorf("error opening JSON file: %v", err)
	}
	defer file.Close()

	var loadedGames []Game
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&loadedGames); err != nil {
		return fmt.Errorf("error decoding JSON: %v", err)
	}

	// Set the IDs based on position
	for i := range loadedGames {
		loadedGames[i].ID = i + 1
	}

	// Lock mutex before updating games
	gamesMutex.Lock()
	games = loadedGames
	gamesMutex.Unlock()

	return nil
}

// Save games to JSON file
func saveGamesToJSON() error {
	file, err := os.Create(jsonFilePath)
	if err != nil {
		return fmt.Errorf("error creating JSON file: %v", err)
	}
	defer file.Close()

	// Lock for reading games
	gamesMutex.RLock()
	defer gamesMutex.RUnlock()

	// Create a copy of games with IDs set correctly
	gamesToSave := make([]Game, len(games))
	for i, game := range games {
		game.ID = i + 1
		gamesToSave[i] = game
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print
	if err := encoder.Encode(gamesToSave); err != nil {
		return fmt.Errorf("error encoding JSON: %v", err)
	}

	return nil
}

// Handle JSON file upload
func handleJSONUpload(w http.ResponseWriter, r *http.Request) {
	// Check if request is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form
	if err := r.ParseMultipartForm(10 << 20); err != nil { // 10MB limit
		http.Error(w, "Unable to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Get file from form
	file, handler, err := r.FormFile("jsonFile")
	if err != nil {
		http.Error(w, "Error retrieving file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Check file extension
	if handler.Header.Get("Content-Type") != "application/json" && !strings.HasSuffix(handler.Filename, ".json") {
		http.Error(w, "Invalid file type. Please upload a JSON file.", http.StatusBadRequest)
		return
	}

	// Create temporary file to save uploaded JSON
	tempFile, err := os.CreateTemp("", "games-*.json")
	if err != nil {
		http.Error(w, "Error creating temporary file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()
	defer os.Remove(tempFile.Name())

	// Copy uploaded file to temp file
	if _, err := io.Copy(tempFile, file); err != nil {
		http.Error(w, "Error copying file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Rewind temp file for reading
	if _, err := tempFile.Seek(0, 0); err != nil {
		http.Error(w, "Error rewinding file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse JSON
	var importedGames []Game
	decoder := json.NewDecoder(tempFile)
	if err := decoder.Decode(&importedGames); err != nil {
		http.Error(w, "Error parsing JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate games have required fields
	for i, game := range importedGames {
		if game.Title == "" {
			http.Error(w, fmt.Sprintf("Game at index %d is missing a title", i), http.StatusBadRequest)
			return
		}
		if game.Genre == "" {
			http.Error(w, fmt.Sprintf("Game at index %d is missing a genre", i), http.StatusBadRequest)
			return
		}
		if game.Publisher == "" {
			http.Error(w, fmt.Sprintf("Game at index %d is missing a publisher", i), http.StatusBadRequest)
			return
		}
		if game.ReleaseYear < 1950 || game.ReleaseYear > time.Now().Year()+1 {
			http.Error(w, fmt.Sprintf("Game at index %d has an invalid release year", i), http.StatusBadRequest)
			return
		}
		if game.Rating < 0 || game.Rating > 10 {
			http.Error(w, fmt.Sprintf("Game at index %d has an invalid rating (should be 0-10)", i), http.StatusBadRequest)
			return
		}
		if game.Price < 0 {
			http.Error(w, fmt.Sprintf("Game at index %d has an invalid price", i), http.StatusBadRequest)
			return
		}
	}

	// Set IDs based on position in the array
	for i := range importedGames {
		importedGames[i].ID = i + 1
	}

	// Update games list with imported games
	gamesMutex.Lock()
	games = importedGames
	gamesMutex.Unlock()

	// Save to JSON file
	if err := saveGamesToJSON(); err != nil {
		http.Error(w, "Error saving games to JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	response := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Successfully imported %d games from JSON", len(importedGames)),
		"count":   len(importedGames),
		"games":   importedGames,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Export games to JSON (download)
func handleJSONExport(w http.ResponseWriter, r *http.Request) {
	// Set headers for file download
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Disposition", "attachment; filename=games_export.json")

	// Lock for reading games
	gamesMutex.RLock()
	defer gamesMutex.RUnlock()

	// Create a copy of games with IDs set correctly
	gamesToExport := make([]Game, len(games))
	for i, game := range games {
		game.ID = i + 1
		gamesToExport[i] = game
	}

	// Write JSON to response
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ") // Pretty print
	if err := encoder.Encode(gamesToExport); err != nil {
		http.Error(w, "Error encoding JSON: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

// fetchGame simulates fetching game data with a random delay to show concurrency
// In a real app, this would make API or database calls
func fetchGame(id int, delay bool) (Game, error) {
	// Lock for reading games
	gamesMutex.RLock()
	defer gamesMutex.RUnlock()

	// Validate ID
	if id < 1 || id > len(games) {
		return Game{}, fmt.Errorf("game with ID %d not found", id)
	}

	game := games[id-1]

	// Simulate network delay if requested
	// Each game has a different delay to demonstrate the benefits of concurrency
	if delay {
		sleepTime := time.Duration(100+id*50) * time.Millisecond
		time.Sleep(sleepTime) // Simulate network or processing delay
		game.FetchTime = sleepTime
	}

	return game, nil
}

// fetchGamesSequential fetches games one by one
// This is the traditional approach without concurrency
// Each game fetch happens one after another (sequentially)
func fetchGamesSequential(ids []int) ([]Game, time.Duration) {
	start := time.Now() // Start timing
	var results []Game

	// Process each game ID one at a time
	for _, id := range ids {
		game, err := fetchGame(id, true)
		if err == nil {
			results = append(results, game)
		}
	}

	// Return results and total execution time
	return results, time.Since(start)
}

// fetchGamesConcurrent fetches games concurrently using goroutines and channels
// This demonstrates Go's concurrency model with:
// 1. Goroutines - lightweight threads managed by Go runtime
// 2. Channels - communication pipes between goroutines
// 3. WaitGroups - synchronization primitives
func fetchGamesConcurrent(ids []int) ([]Game, time.Duration) {
	start := time.Now() // Start timing

	// Create a buffered channel to receive results from goroutines
	// The buffer size matches the number of goroutines we'll create
	resultsChan := make(chan Result, len(ids))

	// WaitGroup helps us wait for all goroutines to finish
	var wg sync.WaitGroup

	// Launch a goroutine for each game ID
	// Each goroutine will fetch a game concurrently
	for _, id := range ids {
		wg.Add(1) // Increment the WaitGroup counter

		// The go keyword creates a new goroutine (lightweight thread)
		// Each goroutine executes this anonymous function independently
		go func(gameID int) {
			defer wg.Done() // Decrement the counter when goroutine completes

			// Fetch the game in this separate goroutine
			game, err := fetchGame(gameID, true)

			// Send the result through the channel
			resultsChan <- Result{Game: game, Err: err}
		}(id) // Pass id as parameter to avoid closure issues
	}

	// Launch another goroutine to close the channel when all fetches complete
	// This ensures the range loop below will terminate
	go func() {
		wg.Wait()          // Wait for all game-fetching goroutines to finish
		close(resultsChan) // Close the channel to signal no more data is coming
	}()

	// Collect results from the channel
	// This will receive all results as they become available
	var results []Game
	for result := range resultsChan {
		if result.Err == nil {
			results = append(results, result.Game)
		}
	}

	// Return results and total execution time
	return results, time.Since(start)
}

// fetchGamesWithLeakyBuffers demonstrates the leaky buffer pattern
// This uses a buffer pool to reuse memory when processing game data
func fetchGamesWithLeakyBuffers(ids []int) ([]Game, time.Duration, int) {
	start := time.Now() // Start timing

	// Reset buffer pool statistics for this demonstration
	bufferPool.mutex.Lock()
	bufferPool.stats = BufferPoolStats{}
	bufferPool.mutex.Unlock()

	// Create a buffered channel to receive results from goroutines
	resultsChan := make(chan Result, len(ids))

	// WaitGroup to track goroutines
	var wg sync.WaitGroup

	// Launch a goroutine for each game ID
	for _, id := range ids {
		wg.Add(1)

		go func(gameID int) {
			defer wg.Done()

			// Get a buffer from the pool
			buffer := bufferPool.Get()

			// Return the buffer to the pool when done
			defer bufferPool.Put(buffer)

			// Simulate processing with the buffer
			// In a real app, this would do meaningful work with the buffer
			// (e.g., make a network request and write response to buffer,
			// or process large files without allocating new memory each time)
			time.Sleep(time.Duration(50+gameID*20) * time.Millisecond)

			// Fetch the game
			game, err := fetchGame(gameID, true)

			// If successful, mark this game as having used a reused buffer
			// (for demonstration purposes only)
			if err == nil {
				// Check if this buffer was reused or newly created
				// by comparing reused count with created count
				stats := bufferPool.GetStats()
				game.BufferReused = stats.Reused > 0
			}

			// Send result through channel
			resultsChan <- Result{Game: game, Err: err}
		}(id)
	}

	// Close channel when all goroutines are done
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Collect results
	var results []Game
	for result := range resultsChan {
		if result.Err == nil {
			results = append(results, result.Game)
		}
	}

	// Get final buffer pool statistics
	stats := bufferPool.GetStats()

	// Return results, execution time, and number of buffers reused
	return results, time.Since(start), stats.Reused
}

func main() {
	// Initialize games from JSON file if it exists
	initGamesFromJSON()

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve sample JSON file
	http.HandleFunc("/sample_games.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "sample_games.json")
	})

	// Home page handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))

		// Check if games are loaded from JSON
		data := PageData{
			IsJSONImported: true, // Always true after initialization
		}

		tmpl.Execute(w, data)
	})

	// API endpoint for sequential fetching
	// Demonstrates the traditional approach without concurrency
	http.HandleFunc("/api/games/sequential", func(w http.ResponseWriter, r *http.Request) {
		ids, err := parseGameIDs(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Fetch games sequentially (one after another)
		results, duration := fetchGamesSequential(ids)

		data := PageData{
			Games:  results,
			Timing: duration,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	// API endpoint for concurrent fetching
	// Demonstrates Go's concurrency model with goroutines and channels
	http.HandleFunc("/api/games/concurrent", func(w http.ResponseWriter, r *http.Request) {
		ids, err := parseGameIDs(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Fetch games concurrently (all at the same time)
		// This leverages Go's goroutines and channels
		results, duration := fetchGamesConcurrent(ids)

		data := PageData{
			Games:  results,
			Timing: duration,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	// API endpoint for fetching with leaky buffers
	// Demonstrates the leaky buffer pattern for memory reuse
	http.HandleFunc("/api/games/leaky-buffers", func(w http.ResponseWriter, r *http.Request) {
		ids, err := parseGameIDs(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Fetch games using the leaky buffer pattern
		results, duration, buffersReused := fetchGamesWithLeakyBuffers(ids)

		data := PageData{
			Games:         results,
			Timing:        duration,
			BuffersReused: buffersReused,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	// JSON import handler
	http.HandleFunc("/api/games/import", handleJSONUpload)

	// JSON export handler
	http.HandleFunc("/api/games/export", handleJSONExport)

	// API endpoint to get all games
	http.HandleFunc("/api/games", func(w http.ResponseWriter, r *http.Request) {
		gamesMutex.RLock()
		defer gamesMutex.RUnlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(games)
	})

	// Use a different port to avoid conflicts
	port := 8081
	fmt.Printf("Server running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

// parseGameIDs extracts and validates game IDs from request
// Demonstrates input validation
func parseGameIDs(r *http.Request) ([]int, error) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		return nil, fmt.Errorf("error parsing form data: %v", err)
	}

	// Get game IDs from form
	idStrings := r.Form["ids"]
	if len(idStrings) == 0 {
		return nil, fmt.Errorf("no game IDs provided")
	}

	// Lock for reading number of games
	gamesMutex.RLock()
	gamesCount := len(games)
	gamesMutex.RUnlock()

	var ids []int
	for _, idStr := range idStrings {
		// Convert string to integer
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, fmt.Errorf("invalid game ID: %s", idStr)
		}

		// Validate ID is within range
		if id < 1 || id > gamesCount {
			return nil, fmt.Errorf("game ID %d is out of range (1-%d)", id, gamesCount)
		}

		ids = append(ids, id)
	}

	return ids, nil
}
