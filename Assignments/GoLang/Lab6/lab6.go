package main

import (
	"errors"
	"fmt"
	"strings"
)

// Base Interface
type GameInterface interface {
	GetGameDetails() string
}

// Nested Interface extending GameInterface
type AdvancedGameInterface interface {
	GameInterface
	Play() string
	Stop() string
}

// Struct implementing both interfaces
type Game struct {
	Name        string
	Genre       string
	Players     int
	ReleaseYear int
}

// Constructor function with validation
func NewGame(name, genre string, players, year int) (*Game, error) {
	if name == "" || genre == "" {
		return nil, errors.New("game name and genre cannot be empty")
	}
	if players <= 0 {
		return nil, errors.New("players count must be greater than zero")
	}
	if year < 1980 || year > 2025 {
		return nil, errors.New("release year must be between 1980 and 2025")
	}

	return &Game{Name: name, Genre: genre, Players: players, ReleaseYear: year}, nil
}

// Implementing GameInterface method
func (g Game) GetGameDetails() string {
	return fmt.Sprintf("Game: %s\nGenre: %s\nPlayers: %d\nRelease Year: %d", g.Name, g.Genre, g.Players, g.ReleaseYear)
}

// Implementing AdvancedGameInterface methods
func (g Game) Play() string {
	return fmt.Sprintf("Starting game: %s! Enjoy playing.", g.Name)
}

func (g Game) Stop() string {
	return fmt.Sprintf("Stopping game: %s. Come back soon!", g.Name)
}

func main() {
	var games []AdvancedGameInterface

	for {
		var name, genre string
		var players, year int

		fmt.Println("\nEnter Game Details:")

		fmt.Print("Game Name: ")
		fmt.Scanln(&name)
		name = strings.TrimSpace(name)

		fmt.Print("Genre: ")
		fmt.Scanln(&genre)
		genre = strings.TrimSpace(genre)

		fmt.Print("Number of Players: ")
		_, err := fmt.Scanln(&players)
		if err != nil {
			fmt.Println("Error: Invalid input for players. Please enter a valid number.")
			continue
		}

		fmt.Print("Release Year: ")
		_, err = fmt.Scanln(&year)
		if err != nil {
			fmt.Println("Error: Invalid input for release year. Please enter a valid number.")
			continue
		}

		// Creating Game object with validation
		game, err := NewGame(name, genre, players, year)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Storing game in slice
		games = append(games, game)
		fmt.Println("Game added successfully!")

		// Ask if the user wants to add another game
		var choice string
		fmt.Print("\nDo you want to add another game? (yes/no): ")
		fmt.Scanln(&choice)
		choice = strings.ToLower(strings.TrimSpace(choice))

		if choice != "yes" {
			break
		}
	}

	// Check if games list is empty
	if len(games) == 0 {
		fmt.Println("No games available to play.")
		return
	}

	// Game selection menu
	for {
		fmt.Println("\nAvailable Games:")
		for i, game := range games {
			fmt.Printf("%d. %s\n", i+1, game.GetGameDetails())
		}

		fmt.Print("\nEnter the number of the game you want to play (or 0 to exit): ")
		var gameChoice int
		_, err := fmt.Scanln(&gameChoice)
		if err != nil || gameChoice < 0 || gameChoice > len(games) {
			fmt.Println("Invalid selection. Please enter a valid number.")
			continue
		}

		if gameChoice == 0 {
			fmt.Println("Exiting game selection. Have a great day!")
			break
		}

		selectedGame := games[gameChoice-1]

		// Play or Stop options
		for {
			fmt.Print("\nDo you want to 'play' or 'stop' the game? (play/stop/exit): ")
			var action string
			fmt.Scanln(&action)
			action = strings.ToLower(strings.TrimSpace(action))

			if action == "play" {
				fmt.Println(selectedGame.Play())
			} else if action == "stop" {
				fmt.Println(selectedGame.Stop())
			} else if action == "exit" {
				break
			} else {
				fmt.Println("Invalid input. Please enter 'play', 'stop', or 'exit'.")
			}
		}
	}
}
