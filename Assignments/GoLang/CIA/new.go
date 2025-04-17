package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter the number of transactions: ")
	fmt.Scan(&n) // Read the number of transactions from the user

	transactions := make([]float64, n) // Create a slice to store transaction amounts
	fmt.Println("Enter the transaction amounts:")
	for i := 0; i < n; i++ {
		fmt.Printf("Transaction %d: ", i+1)
		fmt.Scan(&transactions[i]) // Read each transaction amount from the user
	}

	for {
		// Display menu options
		fmt.Println("\nFinancial Data Processing Tool - Menu")
		fmt.Println("1. Calculate Total and Average Transaction")
		fmt.Println("2. Identify Highest and Lowest Transactions")
		fmt.Println("3. Categorize Transactions")
		fmt.Println("4. Search for a Specific Transaction")
		fmt.Println("5. Identify Duplicate Transactions")
		fmt.Println("6. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice) // Read the user's menu choice

		switch choice {
		case 1:
			// Calculate total and average transaction amount
			/*
				Calculate and print the total and average of a list of transaction amounts.

				Iterates over a slice of transaction amounts, summing them to compute the total.
				Then, calculates the average transaction amount and prints both the total and
				average values formatted to two decimal places.
			*/
			var total float64
			for i := 0; i < len(transactions); i++ {
				total += transactions[i] // Sum up all transactions
			}

			average := total / float64(len(transactions)) // Calculate average
			fmt.Printf("Total Transaction Amount: %.2f\n", total)
			fmt.Printf("Average Transaction Amount: %.2f\n", average)

		case 2:
			// Identify highest and lowest transaction amounts
			/*
				Iterates through a slice of transaction amounts to determine the highest and lowest values.
				Initializes the highest and lowest values with the first transaction amount.
				Updates the highest and lowest values as it iterates through the slice.
				Prints the highest and lowest transaction amounts formatted to two decimal places.
			*/
			highest := transactions[0]
			lowest := transactions[0]

			for i := 1; i < len(transactions); i++ {
				if transactions[i] > highest {
					highest = transactions[i] // Update highest transaction
				}
				if transactions[i] < lowest {
					lowest = transactions[i] // Update lowest transaction
				}
			}

			fmt.Printf("Highest Transaction Amount: %.2f\n", highest)
			fmt.Printf("Lowest Transaction Amount: %.2f\n", lowest)

		case 3:
			// Categorize transactions into lowest, medium, and highest
			/*
				Categorizes a list of transaction amounts into three categories: lowest, medium, and highest.
				Transactions less than 5000 are categorized as lowest, transactions between 5000 and 9999
				are categorized as medium, and transactions 10000 or greater are categorized as highest.
				Prints the categorized transactions.
			*/
			var lowest, medium, highest []float64
			for i := 0; i < len(transactions); i++ {
				if transactions[i] < 5000 {
					lowest = append(lowest, transactions[i]) // Add to lowest category
				} else if transactions[i] >= 5000 && transactions[i] < 10000 {
					medium = append(medium, transactions[i]) // Add to medium category
				} else {
					highest = append(highest, transactions[i]) // Add to highest category
				}
			}

			fmt.Printf("Lowest Value Transactions: %v\n", lowest)
			fmt.Printf("Medium Value Transactions: %v\n", medium)
			fmt.Printf("Highest Value Transactions: %v\n", highest)

		case 4:
			// Search for a specific transaction amount
			/*
				This code snippet prompts the user to enter a transaction amount and searches for it within a list of transactions.
				If the transaction is found, it prints the transaction amount and its position in the list.
				If the transaction is not found, it notifies the user that the transaction was not found.
			*/
			var search float64
			fmt.Print("Enter the transaction amount to search: ")
			fmt.Scan(&search)

			found := false
			for i := 0; i < len(transactions); i++ {
				if transactions[i] == search {
					fmt.Printf("Transaction %.2f found at position %d\n", search, i)
					found = true
				}
			}

			if !found {
				fmt.Println("Transaction not found.")
			}

		case 5:
			// Identify duplicate transactions
			/*
				Identifies and prints duplicate transactions from a list. Iterates through
				the list of transactions, comparing each transaction with every other
				transaction to find duplicates. If duplicates are found, they are printed
				to the console. If no duplicates are found, a message indicating this is
				printed.
			*/
			found := false
			for i := 0; i < len(transactions); i++ {
				for j := i + 1; j < len(transactions); j++ {
					if transactions[i] == transactions[j] {
						fmt.Printf("Transaction %.2f appears more than once\n", transactions[i])
						found = true
					}
				}
			}

			if !found {
				fmt.Println("No duplicate transactions found.")
			}

		case 6:
			// Exit the program
			fmt.Println("Exiting the program.")
			return

		default:
			// Handle invalid menu choice
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
