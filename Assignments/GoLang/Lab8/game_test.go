package main

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestGameJSONMarshal(t *testing.T) {
	// Create a test game
	releaseDate := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	testGame := Game{
		ID:            1,
		Title:         "Test Game",
		Genre:         "Action",
		ReleaseDate:   releaseDate,
		Rating:        8.5,
		Platforms:     []string{"PC", "PlayStation", "Xbox"},
		Price:         59.99,
		IsMultiplayer: true,
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(testGame)
	if err != nil {
		t.Fatalf("Failed to marshal game to JSON: %v", err)
	}

	// Unmarshal back to verify
	var unmarshaledGame Game
	err = json.Unmarshal(jsonData, &unmarshaledGame)
	if err != nil {
		t.Fatalf("Failed to unmarshal game from JSON: %v", err)
	}

	// Verify the fields match
	if testGame.ID != unmarshaledGame.ID {
		t.Errorf("ID mismatch: expected %d, got %d", testGame.ID, unmarshaledGame.ID)
	}
	if testGame.Title != unmarshaledGame.Title {
		t.Errorf("Title mismatch: expected %s, got %s", testGame.Title, unmarshaledGame.Title)
	}
	if testGame.Genre != unmarshaledGame.Genre {
		t.Errorf("Genre mismatch: expected %s, got %s", testGame.Genre, unmarshaledGame.Genre)
	}
	if !testGame.ReleaseDate.Equal(unmarshaledGame.ReleaseDate) {
		t.Errorf("ReleaseDate mismatch: expected %v, got %v", testGame.ReleaseDate, unmarshaledGame.ReleaseDate)
	}
	if testGame.Rating != unmarshaledGame.Rating {
		t.Errorf("Rating mismatch: expected %.1f, got %.1f", testGame.Rating, unmarshaledGame.Rating)
	}
	if testGame.Price != unmarshaledGame.Price {
		t.Errorf("Price mismatch: expected %.2f, got %.2f", testGame.Price, unmarshaledGame.Price)
	}
	if testGame.IsMultiplayer != unmarshaledGame.IsMultiplayer {
		t.Errorf("IsMultiplayer mismatch: expected %v, got %v", testGame.IsMultiplayer, unmarshaledGame.IsMultiplayer)
	}
	if !reflect.DeepEqual(testGame.Platforms, unmarshaledGame.Platforms) {
		t.Errorf("Platforms mismatch: expected %v, got %v", testGame.Platforms, unmarshaledGame.Platforms)
	}
}

func TestGameLibraryMarshalUnmarshal(t *testing.T) {
	// Create a test library
	releaseDate := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	testLibrary := GameLibrary{
		Games: []Game{
			{
				ID:            1,
				Title:         "Game One",
				Genre:         "Adventure",
				ReleaseDate:   releaseDate,
				Rating:        9.0,
				Platforms:     []string{"PC", "Xbox"},
				Price:         49.99,
				IsMultiplayer: false,
			},
			{
				ID:            2,
				Title:         "Game Two",
				Genre:         "RPG",
				ReleaseDate:   releaseDate,
				Rating:        8.5,
				Platforms:     []string{"PlayStation", "Switch"},
				Price:         59.99,
				IsMultiplayer: true,
			},
		},
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(testLibrary)
	if err != nil {
		t.Fatalf("Failed to marshal game library to JSON: %v", err)
	}

	// Unmarshal back to verify
	var unmarshaledLibrary GameLibrary
	err = json.Unmarshal(jsonData, &unmarshaledLibrary)
	if err != nil {
		t.Fatalf("Failed to unmarshal game library from JSON: %v", err)
	}

	// Verify the library length
	if len(testLibrary.Games) != len(unmarshaledLibrary.Games) {
		t.Errorf("Library size mismatch: expected %d, got %d",
			len(testLibrary.Games), len(unmarshaledLibrary.Games))
		return
	}

	// Verify each game
	for i, originalGame := range testLibrary.Games {
		unmarshaledGame := unmarshaledLibrary.Games[i]
		if originalGame.ID != unmarshaledGame.ID {
			t.Errorf("Game %d: ID mismatch: expected %d, got %d",
				i, originalGame.ID, unmarshaledGame.ID)
		}
		if originalGame.Title != unmarshaledGame.Title {
			t.Errorf("Game %d: Title mismatch: expected %s, got %s",
				i, originalGame.Title, unmarshaledGame.Title)
		}
		// Additional field checks could be added here
	}
}

func TestSaveLoadFile(t *testing.T) {
	// Create a temporary file for testing
	testFile := "test_games.json"
	defer os.Remove(testFile) // Clean up after test

	// Create a library
	library := NewGameLibrary()
	releaseDate := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	game := Game{
		ID:            1,
		Title:         "Test Game",
		Genre:         "Action",
		ReleaseDate:   releaseDate,
		Rating:        8.5,
		Platforms:     []string{"PC", "PlayStation", "Xbox"},
		Price:         59.99,
		IsMultiplayer: true,
	}
	library.AddGame(game)

	// Save to file
	err := library.SaveToFile(testFile)
	if err != nil {
		t.Fatalf("Failed to save library to file: %v", err)
	}

	// Load from file
	loadedLibrary := NewGameLibrary()
	err = loadedLibrary.LoadFromFile(testFile)
	if err != nil {
		t.Fatalf("Failed to load library from file: %v", err)
	}

	// Verify library content
	if len(loadedLibrary.Games) != 1 {
		t.Errorf("Expected 1 game in loaded library, got %d", len(loadedLibrary.Games))
		return
	}

	loadedGame := loadedLibrary.Games[0]
	if loadedGame.ID != game.ID {
		t.Errorf("ID mismatch: expected %d, got %d", game.ID, loadedGame.ID)
	}
	if loadedGame.Title != game.Title {
		t.Errorf("Title mismatch: expected %s, got %s", game.Title, loadedGame.Title)
	}
	if loadedGame.Genre != game.Genre {
		t.Errorf("Genre mismatch: expected %s, got %s", game.Genre, loadedGame.Genre)
	}
	if !loadedGame.ReleaseDate.Equal(game.ReleaseDate) {
		t.Errorf("ReleaseDate mismatch: expected %v, got %v", game.ReleaseDate, loadedGame.ReleaseDate)
	}
	if loadedGame.Rating != game.Rating {
		t.Errorf("Rating mismatch: expected %.1f, got %.1f", game.Rating, loadedGame.Rating)
	}
	if loadedGame.Price != game.Price {
		t.Errorf("Price mismatch: expected %.2f, got %.2f", game.Price, loadedGame.Price)
	}
	if loadedGame.IsMultiplayer != game.IsMultiplayer {
		t.Errorf("IsMultiplayer mismatch: expected %v, got %v", game.IsMultiplayer, loadedGame.IsMultiplayer)
	}
	if !reflect.DeepEqual(loadedGame.Platforms, game.Platforms) {
		t.Errorf("Platforms mismatch: expected %v, got %v", game.Platforms, loadedGame.Platforms)
	}
}

func TestJSONIndentation(t *testing.T) {
	// Test that JSON is properly indented when saved to file
	testFile := "test_indentation.json"
	defer os.Remove(testFile)

	// Create a simple library
	library := NewGameLibrary()
	library.AddGame(Game{
		ID:    1,
		Title: "Test Game",
	})

	// Save to file
	err := library.SaveToFile(testFile)
	if err != nil {
		t.Fatalf("Failed to save library to file: %v", err)
	}

	// Read the file contents
	fileContent, err := os.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Failed to read saved file: %v", err)
	}

	// Check that the file contains indented JSON
	// The indented JSON should have newlines and spaces
	if len(fileContent) == 0 {
		t.Errorf("File is empty")
	}

	// Count the number of newlines
	newlineCount := 0
	for _, b := range fileContent {
		if b == '\n' {
			newlineCount++
		}
	}

	// An indented JSON with one game should have multiple newlines
	if newlineCount < 5 {
		t.Errorf("JSON appears not to be properly indented, found only %d newlines", newlineCount)
	}
}
