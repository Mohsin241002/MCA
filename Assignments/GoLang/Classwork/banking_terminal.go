package main

import (
	"fmt"
	"strings"

	"banking/banking"
)

// Terminal-based banking application
func main() {
	bank := banking.NewBank()
	var currentUser string

	fmt.Println("Welcome to Go Banking System (Terminal Version)")

	for {
		if currentUser == "" {
			fmt.Println("\n1. Create new account")
			fmt.Println("2. Select existing account")
			fmt.Println("3. Exit")
			fmt.Print("Enter choice: ")

			var choice int
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				createAccount(bank)
			case 2:
				currentUser = selectAccount(bank)
			case 3:
				fmt.Println("Thank you for using Go Banking System!")
				return
			default:
				fmt.Println("Invalid choice. Please try again.")
			}
		} else {
			account, _ := bank.GetAccount(currentUser)
			fmt.Printf("\nCurrent User: %s, Account Type: %s, Balance: $%.2f\n",
				currentUser, account.AccountType(), account.CheckBalance())

			fmt.Println("\n1. Deposit")
			fmt.Println("2. Withdraw")
			fmt.Println("3. Check Balance")
			fmt.Println("4. View Transaction History")
			fmt.Println("5. Switch User")
			fmt.Println("6. Exit")
			fmt.Print("Enter choice: ")

			var choice int
			fmt.Scanln(&choice)

			switch choice {
			case 1:
				deposit(bank, currentUser)
			case 2:
				withdraw(bank, currentUser)
			case 3:
				checkBalance(bank, currentUser)
			case 4:
				viewTransactions(bank, currentUser)
			case 5:
				currentUser = ""
			case 6:
				fmt.Println("Thank you for using Go Banking System!")
				return
			default:
				fmt.Println("Invalid choice. Please try again.")
			}
		}
	}
}

func createAccount(bank *banking.Bank) {
	var name string
	var accountType string
	var initialBalance float64

	fmt.Print("Enter account holder name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter account type (Savings/Current): ")
	fmt.Scanln(&accountType)

	fmt.Print("Enter initial balance: $")
	fmt.Scanln(&initialBalance)

	if initialBalance < 0 {
		fmt.Println("Error: Initial balance cannot be negative.")
		return
	}

	accountType = strings.ToLower(accountType)

	if accountType == "savings" {
		bank.AddAccount(name, banking.NewSavingsAccount(initialBalance))
		fmt.Printf("Savings account created for %s with balance $%.2f\n", name, initialBalance)
	} else if accountType == "current" {
		bank.AddAccount(name, banking.NewCurrentAccount(initialBalance))
		fmt.Printf("Current account created for %s with balance $%.2f\n", name, initialBalance)
	} else {
		fmt.Println("Error: Invalid account type. Please choose 'Savings' or 'Current'.")
	}
}

func selectAccount(bank *banking.Bank) string {
	accountNames := bank.GetAllAccountNames()

	if len(accountNames) == 0 {
		fmt.Println("No accounts found. Please create an account first.")
		return ""
	}

	fmt.Println("\nAvailable accounts:")
	for _, name := range accountNames {
		account, _ := bank.GetAccount(name)
		fmt.Printf("%s (%s): $%.2f\n", name, account.AccountType(), account.CheckBalance())
	}

	fmt.Print("Enter account holder name: ")
	var name string
	fmt.Scanln(&name)

	_, exists := bank.GetAccount(name)
	if !exists {
		fmt.Println("Account not found.")
		return ""
	}

	return name
}

func deposit(bank *banking.Bank, name string) {
	account, _ := bank.GetAccount(name)

	fmt.Print("Enter deposit amount: $")
	var amount float64
	fmt.Scanln(&amount)

	err := account.Deposit(amount)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Successfully deposited $%.2f. New balance: $%.2f\n", amount, account.CheckBalance())
		bank.RecordTransaction(name, "Deposit", amount)
	}
}

func withdraw(bank *banking.Bank, name string) {
	account, _ := bank.GetAccount(name)

	fmt.Print("Enter withdrawal amount: $")
	var amount float64
	fmt.Scanln(&amount)

	err := account.Withdraw(amount)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("Successfully withdrew $%.2f. New balance: $%.2f\n", amount, account.CheckBalance())
		bank.RecordTransaction(name, "Withdrawal", amount)
	}
}

func checkBalance(bank *banking.Bank, name string) {
	account, _ := bank.GetAccount(name)
	fmt.Printf("Current balance: $%.2f\n", account.CheckBalance())
}

func viewTransactions(bank *banking.Bank, name string) {
	transactions := bank.GetTransactionHistory(name)

	if len(transactions) == 0 {
		fmt.Println("No transactions found.")
		return
	}

	fmt.Println("\nTransaction History:")
	for i, transaction := range transactions {
		fmt.Printf("%d. %s: $%.2f (%s)\n",
			i+1, transaction.TransactionType, transaction.Amount, transaction.Timestamp)
	}
}
