package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var library = NewGameLibrary()

func main() {
	// Create a default games.json file if it doesn't exist
	ensureDefaultFile()

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/api/games", handleGames)
	http.HandleFunc("/api/games/add", handleAddGame)
	http.HandleFunc("/api/games/delete", handleDeleteGame)
	http.HandleFunc("/api/games/save", handleSaveGames)
	http.HandleFunc("/api/games/load", handleLoadGames)

	// Print the current working directory for debugging
	dir, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting current directory: %v", err)
	} else {
		log.Printf("Current working directory: %s", dir)
	}

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ensureDefaultFile creates a default games.json file if it doesn't exist
func ensureDefaultFile() {
	defaultFile := "games.json"
	if _, err := os.Stat(defaultFile); os.IsNotExist(err) {
		// Create a sample game library
		sampleLibrary := NewGameLibrary()
		sampleLibrary.AddGame(Game{
			ID:            1,
			Title:         "Sample Game",
			Genre:         "Adventure",
			ReleaseDate:   time.Now(),
			Rating:        8.5,
			Platforms:     []string{"PC", "PlayStation", "Xbox"},
			Price:         49.99,
			IsMultiplayer: true,
		})
		// Save to file
		sampleLibrary.SaveToFile(defaultFile)

		absPath, err := filepath.Abs(defaultFile)
		if err != nil {
			log.Printf("Error getting absolute path: %v", err)
			log.Println("Created default games.json file")
		} else {
			log.Println("Created default games.json file at:", absPath)
		}
	} else {
		absPath, err := filepath.Abs(defaultFile)
		if err != nil {
			log.Printf("Error getting absolute path: %v", err)
		} else {
			log.Println("Using existing games.json file at:", absPath)
		}
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func handleGames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(library.Games)
}

func handleAddGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var game Game
	if err := json.NewDecoder(r.Body).Decode(&game); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Only set release date if not provided by client
	if game.ReleaseDate.IsZero() {
		game.ReleaseDate = time.Now()
		log.Println("No release date provided, using current time")
	} else {
		log.Printf("Using client-provided release date: %v", game.ReleaseDate)
	}

	// Check if game with this ID already exists
	for _, existingGame := range library.Games {
		if existingGame.ID == game.ID {
			http.Error(w, "Game with this ID already exists", http.StatusBadRequest)
			return
		}
	}

	library.AddGame(game)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func handleDeleteGame(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	success := library.RemoveGame(id)
	if !success {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

func handleSaveGames(w http.ResponseWriter, r *http.Request) {
	// Support both GET and POST methods for saving
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	filename := r.URL.Query().Get("filename")
	if filename == "" {
		filename = "games.json"
	}

	// Ensure the path exists
	dir := filepath.Dir(filename)
	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			http.Error(w, "Failed to create directory: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Attempt to save the file
	if err := library.SaveToFile(filename); err != nil {
		errMsg := fmt.Sprintf("Failed to save file: %v", err)
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	absPath, err := filepath.Abs(filename)
	if err != nil {
		log.Printf("Error getting absolute path: %v", err)
	} else {
		log.Printf("Saved games to file: %s", absPath)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Library saved successfully"})
}

func handleLoadGames(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for easier development
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")

	// Handle OPTIONS request for CORS preflight
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Log the request for debugging
	log.Printf("Load games request received: %s %s", r.Method, r.URL.String())

	// Explicitly allow GET method
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		errMsg := fmt.Sprintf("Method %s not allowed, only GET and POST are supported", r.Method)
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusMethodNotAllowed)
		return
	}

	filename := r.URL.Query().Get("filename")
	if filename == "" {
		filename = "games.json"
	}

	// Log the path for debugging
	absPath, err := filepath.Abs(filename)
	if err != nil {
		log.Printf("Error getting absolute path: %v", err)
	} else {
		log.Printf("Attempting to load games from: %s", absPath)
	}

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Printf("File not found: %s", filename)
		// Return empty array instead of error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode([]Game{})
		return
	}

	newLibrary := NewGameLibrary()
	err = newLibrary.LoadFromFile(filename)
	if err != nil {
		errMsg := fmt.Sprintf("Error loading file: %v", err)
		log.Println(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	// Update the global library
	library = newLibrary
	log.Printf("Successfully loaded %d games from %s", len(library.Games), filename)

	// Return just the games array, not wrapped in a library object
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(library.Games)
}
