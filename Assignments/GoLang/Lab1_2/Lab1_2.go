package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to generate a math problem based on difficulty
func generateMathProblem(rng *rand.Rand, difficulty int) (string, int) {
	num1 := rng.Intn(difficulty*5) + 1
	num2 := rng.Intn(difficulty*5) + 1
	operators := []rune{'+', '-', '*'}
	op := operators[rng.Intn(len(operators))]

	var correctAnswer int
	switch op {
	case '+':
		correctAnswer = num1 + num2
	case '-':
		correctAnswer = num1 - num2
	case '*':
		correctAnswer = num1 * num2
	}

	problem := fmt.Sprintf("%d %c %d =", num1, op, num2)
	return problem, correctAnswer
}

// Function to display game rules
func showGameRules() {
	fmt.Println("\n📜 WELCOME TO THE MATH CHALLENGE GAME! 🎮")
	fmt.Println("========================================")
	fmt.Println("📝 RULES:")
	fmt.Println("1️⃣ Solve math problems to gain XP and level up.")
	fmt.Println("2️⃣ You earn **+20 XP** for each correct answer.")
	fmt.Println("3️⃣ You lose **-40 XP** for each incorrect answer.")
	fmt.Println("4️⃣ **3 correct answers** = **LEVEL UP!** 🚀")
	fmt.Println("5️⃣ If you reach **3 total incorrect answers**, it's **GAME OVER!** 💀")
	fmt.Println("6️⃣ The difficulty increases as you progress! 🔥")
	fmt.Println("========================================")
}

// Function to get valid integer input
func getValidIntegerInput() int {
	var input int
	for {
		fmt.Print("Your Answer: ")
		_, err := fmt.Scanln(&input) // Using Scanln() for safer input
		if err != nil {
			fmt.Println("⚠️ Invalid input! Please enter a valid number.")
			var clearBuffer string
			fmt.Scanln(&clearBuffer) // Clear incorrect input
		} else {
			return input
		}
	}
}

func main() {
	// Create a properly seeded random number generator
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Show game rules at the start
	showGameRules()

	// Player stats
	level := 1
	xp := 0
	correctAnswers := 0
	requiredCorrect := 3 // Number of correct answers needed to level up
	incorrectCount := 0  // Total incorrect attempts
	maxIncorrect := 3    // Maximum allowed incorrect attempts before Game Over
	correctXP := 20      // XP for correct answer
	incorrectXP := -40   // XP penalty for incorrect answer

	// Start the game
	fmt.Print("Enter your Player ID: ")
	var playerid string
	fmt.Scanln(&playerid)

	fmt.Printf("\n🎮 Welcome, %s! Let the game begin! 🚀\n", playerid)

	// Game loop (runs until game over)
	for incorrectCount < maxIncorrect {
		fmt.Printf("\n🎯 Level %d | XP: %d | Incorrect Attempts: %d/%d\n", level, xp, incorrectCount, maxIncorrect)
		problem, correctAnswer := generateMathProblem(rng, level)
		fmt.Println(problem)

		// Get user input with error handling
		userAnswer := getValidIntegerInput()

		// Check if the answer is correct
		if userAnswer == correctAnswer {
			fmt.Println("✅ Correct! +20 XP")
			xp += correctXP
			correctAnswers++

			// Level up condition
			if correctAnswers >= requiredCorrect {
				level++
				correctAnswers = 0
				fmt.Println("🎉 Congratulations! You've advanced to the next level! 🚀")
			}
		} else {
			fmt.Println("❌ Incorrect. -40 XP")
			incorrectCount++
			xp += incorrectXP // Deduct XP for incorrect answer

			// Game over if max incorrect attempts are reached
			if incorrectCount >= maxIncorrect {
				fmt.Println("\n💀 GAME OVER! You reached Level", level, "with XP:", xp)
				break
			} else {
				fmt.Printf("⚠️ You have %d attempts left before GAME OVER!\n", maxIncorrect-incorrectCount)
			}
		}
	}

	// Final message
	fmt.Println("\n🎮 Thanks for playing! See you next time! 🚀")
}
