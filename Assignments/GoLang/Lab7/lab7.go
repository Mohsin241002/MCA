package main

import (
	"fmt"
	"strings"
)

// Game represents a game with a name and score
type Game struct {
	Name  string
	Score int
}

// updateScoreByValue updates the score of the game by value (does not affect the original game)
func updateScoreByValue(game Game, newScore int) {
	game.Score = newScore
	fmt.Println("Score updated by value (original game not affected):", game)
}

// updateScoreByReference updates the score of the game by reference (affects the original game)
func updateScoreByReference(game *Game, newScore int) {
	game.Score = newScore
}

// printGameDetails prints the details of multiple games using a variadic function
func printGameDetails(games ...Game) {
	for _, game := range games {
		fmt.Printf("Game: %s, Score: %d\n", game.Name, game.Score)
	}
}

// getUserInput gets input from the user and validates it
func getUserInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return strings.TrimSpace(input)
}

// getValidatedInt gets a validated integer input from the user
func getValidatedInt(prompt string) int {
	for {
		input := getUserInput(prompt)
		var value int
		_, err := fmt.Sscanf(input, "%d", &value)
		if err == nil {
			return value
		}
		fmt.Println("Invalid input. Please enter a valid integer.")
	}
}

// demonstratePointerManipulation demonstrates basic pointer manipulation
func demonstratePointerManipulation() {
	var x int = 10
	var p *int = &x

	fmt.Println("Initial value of x:", x)
	fmt.Println("Pointer p points to value:", *p)

	*p = 20
	fmt.Println("Value of x after updating through pointer p:", x)
}

func main() {
	var games []Game

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add a new game")
		fmt.Println("2. Update game score by value")
		fmt.Println("3. Update game score by reference")
		fmt.Println("4. Print game details")
		fmt.Println("5. Demonstrate pointer manipulation")
		fmt.Println("6. Exit")

		choice := getValidatedInt("Enter your choice: ")

		switch choice {
		case 1:
			name := getUserInput("Enter game name: ")
			score := getValidatedInt("Enter game score: ")
			games = append(games, Game{Name: name, Score: score})
		case 2:
			if len(games) == 0 {
				fmt.Println("No games available to update.")
				break
			}
			index := getValidatedInt("Enter game index to update by value: ")
			if index < 0 || index >= len(games) {
				fmt.Println("Invalid index.")
				break
			}
			newScore := getValidatedInt("Enter new score: ")
			updateScoreByValue(games[index], newScore)
		case 3:
			if len(games) == 0 {
				fmt.Println("No games available to update.")
				break
			}
			index := getValidatedInt("Enter game index to update by reference: ")
			if index < 0 || index >= len(games) {
				fmt.Println("Invalid index.")
				break
			}
			newScore := getValidatedInt("Enter new score: ")
			updateScoreByReference(&games[index], newScore)
		case 4:
			if len(games) == 0 {
				fmt.Println("No games available.")
			} else {
				printGameDetails(games...)
			}
		case 5:
			demonstratePointerManipulation()
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
