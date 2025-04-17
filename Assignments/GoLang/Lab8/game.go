package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Game represents a video game with its details
type Game struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Genre         string    `json:"genre"`
	ReleaseDate   time.Time `json:"release_date"`
	Rating        float64   `json:"rating"`
	Platforms     []string  `json:"platforms"`
	Price         float64   `json:"price"`
	IsMultiplayer bool      `json:"is_multiplayer"`
}

// GameLibrary manages a collection of games
type GameLibrary struct {
	Games []Game `json:"games"`
}

// NewGameLibrary creates a new empty game library
func NewGameLibrary() *GameLibrary {
	return &GameLibrary{
		Games: make([]Game, 0),
	}
}

// AddGame adds a new game to the library
func (gl *GameLibrary) AddGame(game Game) {
	gl.Games = append(gl.Games, game)
}

// RemoveGame removes a game by ID
func (gl *GameLibrary) RemoveGame(id int) bool {
	for i, game := range gl.Games {
		if game.ID == id {
			gl.Games = append(gl.Games[:i], gl.Games[i+1:]...)
			return true
		}
	}
	return false
}

// GetGameByID returns a game by its ID
func (gl *GameLibrary) GetGameByID(id int) *Game {
	for _, game := range gl.Games {
		if game.ID == id {
			return &game
		}
	}
	return nil
}

// SaveToFile saves the game library to a JSON file
func (gl *GameLibrary) SaveToFile(filename string) error {
	log.Printf("Saving library with %d games to file: %s", len(gl.Games), filename)

	// Create a proper structure that matches what we're unmarshaling
	libraryData := map[string][]Game{
		"games": gl.Games,
	}

	data, err := json.MarshalIndent(libraryData, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshaling data: %w", err)
	}

	// Ensure directory exists
	dir := filepath.Dir(filename)
	if dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %w", dir, err)
		}
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	log.Printf("Successfully saved %d bytes to %s", len(data), filename)
	return nil
}

// LoadFromFile loads the game library from a JSON file
func (gl *GameLibrary) LoadFromFile(filename string) error {
	log.Printf("Loading game library from file: %s", filename)

	// Check if file exists
	fileInfo, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file does not exist: %s", filename)
		}
		return fmt.Errorf("error checking file: %w", err)
	}

	log.Printf("File size: %d bytes, last modified: %s", fileInfo.Size(), fileInfo.ModTime())

	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	if len(data) == 0 {
		return fmt.Errorf("file is empty: %s", filename)
	}

	log.Printf("Read %d bytes from file", len(data))

	// Try to unmarshal as a library first
	var libraryData map[string][]Game
	err = json.Unmarshal(data, &libraryData)
	if err == nil && libraryData["games"] != nil {
		gl.Games = libraryData["games"]
		log.Printf("Successfully loaded %d games from library format", len(gl.Games))
		return nil
	}

	// If that fails, try to unmarshal as direct array
	var games []Game
	err = json.Unmarshal(data, &games)
	if err == nil {
		gl.Games = games
		log.Printf("Successfully loaded %d games from array format", len(gl.Games))
		return nil
	}

	// If both approaches fail, try to unmarshal as our struct
	err = json.Unmarshal(data, gl)
	if err != nil {
		// Log a sample of the data for debugging
		sampleSize := 100
		if len(data) < sampleSize {
			sampleSize = len(data)
		}
		log.Printf("JSON Unmarshal error on data (first %d bytes): %s", sampleSize, data[:sampleSize])
		return fmt.Errorf("error unmarshaling data: %w", err)
	}

	log.Printf("Successfully loaded %d games from GameLibrary format", len(gl.Games))
	return nil
}
