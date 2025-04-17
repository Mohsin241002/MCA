package main

import (
	"errors"
	"fmt"
)

// Game struct represents a game in the store.
type Game struct {
	Name  string
	Price float64
}

// Custom errors
var (
	ErrGameNotFound        = errors.New("error: game not found")
	ErrInsufficientBalance = errors.New("error: insufficient balance")
	ErrInvalidGamePrice    = errors.New("error: invalid game price")
	ErrInvalidInput        = errors.New("error: invalid input provided")
)

// Game store map
var gameStore = make(map[string]Game)

// AddGame function adds a game with validation
func AddGame(name string, price float64) error {
	if price <= 0 {
		return fmt.Errorf("%w: price must be greater than zero", ErrInvalidGamePrice)
	}
	if len(name) == 0 {
		return fmt.Errorf("%w: game name cannot be empty", ErrInvalidInput)
	}

	gameStore[name] = Game{Name: name, Price: price}
	fmt.Println("✅ Game added successfully:", name)
	return nil
}

// SearchGame function searches for a game
func SearchGame(name string) (Game, error) {
	game, exists := gameStore[name]
	if !exists {
		return Game{}, fmt.Errorf("%w: %s", ErrGameNotFound, name)
	}
	return game, nil
}

// PurchaseGame function handles game purchases with balance validation
func PurchaseGame(name string, balance float64) error {
	game, err := SearchGame(name)
	if err != nil {
		return err
	}

	if balance < game.Price {
		return fmt.Errorf("%w: need more $%.2f", ErrInsufficientBalance, game.Price-balance)
	}

	fmt.Printf("🎮 Purchase successful! You bought '%s' for $%.2f\n", game.Name, game.Price)
	return nil
}

// GetUserInput function reads user input and returns a trimmed string
func GetUserInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}

func main() {
	var balance float64 = 100.0 // Default user balance
	var choice int

	for {
		// Display menu
		fmt.Println("\n🎮 Gaming Store Menu:")
		fmt.Println("1. Add Game")
		fmt.Println("2. Search Game")
		fmt.Println("3. Purchase Game")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("❌ Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			// Adding a game
			name := GetUserInput("Enter game name: ")
			var price float64
			fmt.Print("Enter game price: ")
			_, err := fmt.Scanln(&price)

			if err != nil {
				fmt.Println("❌ Invalid price. Please enter a valid number.")
				continue
			}

			if err := AddGame(name, price); err != nil {
				fmt.Println(err)
			}

		case 2:
			// Searching for a game
			name := GetUserInput("Enter game name to search: ")
			game, err := SearchGame(name)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("✅ Game found: %s - Price: $%.2f\n", game.Name, game.Price)
			}

		case 3:
			// Purchasing a game
			name := GetUserInput("Enter game name to purchase: ")
			err := PurchaseGame(name, balance)
			if err != nil {
				fmt.Println(err)
			} else {
				balance -= gameStore[name].Price
				fmt.Printf("💰 New Balance: $%.2f\n", balance)
			}

		case 4:
			fmt.Println("👋 Exiting Gaming Store. See you soon!")
			return

		default:
			fmt.Println("❌ Invalid choice. Please enter a valid option.")
		}
	}
}
