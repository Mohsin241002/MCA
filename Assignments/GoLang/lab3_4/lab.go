package main

import (
	"fmt"
	"strings"
)

// 🎯 Define a Struct (Game)
type Game struct {
	name     string
	platform string
	rating   float32
}

// 🎯 Function outside main() - Processes game details
func printGameDetails(g Game, action func(string)) {
	fmt.Printf("\n🎮 Game Name: %s\n🕹️ Platform: %s\n⭐ Rating: %.1f\n", g.name, g.platform, g.rating)
	action(g.name) // Call function passed as a parameter
}

// 🎯 Function to get valid string input (non-empty)
func getValidStringInput(prompt string) string {
	var input string
	for {
		fmt.Print(prompt)
		fmt.Scanln(&input)
		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("⚠️ Invalid input! Please enter a valid value.")
		} else {
			return input
		}
	}
}

// 🎯 Function to get valid integer input (must be a positive number)
func getValidIntInput(prompt string) int {
	var input int
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanln(&input)
		if err != nil || input <= 0 {
			fmt.Println("⚠️ Invalid input! Please enter a valid positive number.")
			var clearBuffer string
			fmt.Scanln(&clearBuffer) // Clear incorrect input
		} else {
			return input
		}
	}
}

// 🎯 Function to get valid float input (rating)
func getValidFloatInput(prompt string) float32 {
	var rating float32
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanln(&rating)
		if err != nil || rating < 0 || rating > 10 {
			fmt.Println("⚠️ Invalid input! Please enter a rating between 0 and 10.")
			var clearBuffer string
			fmt.Scanln(&clearBuffer) // Clear incorrect input
		} else {
			return rating
		}
	}
}

func main() {
	// 📌 Ask user for number of games (with validation)
	numGames := getValidIntInput("Enter the number of games you want to add: ")

	// 📌 Array Example: Fixed-size collection of games
	gamesArray := make([]string, numGames)

	// 📌 Slice Example: Dynamic collection of platforms
	platformsSlice := []string{}

	// 📌 Map Example: Game ratings
	gameRatings := make(map[string]float32)

	// 📌 Struct Example: Storing game objects
	gameList := []Game{}

	// 📌 User Input for Games
	for i := 0; i < numGames; i++ {
		fmt.Printf("\n🎮 Enter details for Game %d:\n", i+1)
		gameName := getValidStringInput("Game Name: ")
		gamePlatform := getValidStringInput("Platform: ")
		gameRating := getValidFloatInput("Rating (0-10): ")

		// 📌 Store in array, slice, map, and struct
		gamesArray[i] = gameName
		platformsSlice = append(platformsSlice, gamePlatform)
		gameRatings[gameName] = gameRating
		gameList = append(gameList, Game{name: gameName, platform: gamePlatform, rating: gameRating})
	}

	// 📌 Function Inside main()
	sayHello := func(playerName string) {
		fmt.Println("👋 Hello, " + playerName + "! Welcome to the Gaming System.")
	}

	// 📌 Variable as Function Name
	display := sayHello // Assign function to a variable
	playerName := getValidStringInput("\nEnter your name: ")
	display(playerName) // Call the function using variable

	// 📌 Loop through Array
	fmt.Println("\n🎮 Available Games:")
	for _, game := range gamesArray {
		fmt.Println("-", game)
	}

	// 📌 Loop through Slice
	fmt.Println("\n🕹️ Supported Platforms:")
	for _, platform := range platformsSlice {
		fmt.Println("-", platform)
	}

	// 📌 Loop through Map
	fmt.Println("\n⭐ Game Ratings:")
	for game, rating := range gameRatings {
		fmt.Printf("%s: %.1f\n", game, rating)
	}

	// 📌 Using Struct and Function with Function Parameter
	fmt.Println("\n🎲 Game Details:")
	for _, game := range gameList {
		printGameDetails(game, func(gameName string) {
			fmt.Println("📢 " + gameName + " is an amazing game!")
		})
	}
}
