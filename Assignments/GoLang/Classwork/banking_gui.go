package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"banking/banking"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// GUI Application
func main() {
	// Initialize bank
	bank := banking.NewBank()

	// Create new Fyne application
	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())

	// Create main window
	w := a.NewWindow("Go Banking System")
	w.Resize(fyne.NewSize(800, 600))

	// Create the content for the initial welcome screen
	welcomeScreen := createWelcomeScreen(bank, w)

	// Set the content for the main window
	w.SetContent(welcomeScreen)

	// Show and run
	w.ShowAndRun()
}

// createWelcomeScreen creates the initial welcome screen
func createWelcomeScreen(bank *banking.Bank, w fyne.Window) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("Welcome to Go Banking System", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Create buttons for initial options
	createAccountBtn := widget.NewButton("Create Account", func() {
		w.SetContent(createAccountScreen(bank, w))
	})

	selectAccountBtn := widget.NewButton("Select Account", func() {
		w.SetContent(selectAccountScreen(bank, w))
	})

	exitBtn := widget.NewButton("Exit", func() {
		w.Close()
	})

	// Create vertical container with buttons
	content := container.NewVBox(
		title,
		widget.NewLabel(""), // spacing
		createAccountBtn,
		selectAccountBtn,
		exitBtn,
	)

	return container.NewCenter(container.NewVBox(
		layout.NewSpacer(),
		content,
		layout.NewSpacer(),
	))
}

// createAccountScreen creates the account creation screen
func createAccountScreen(bank *banking.Bank, w fyne.Window) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("Create New Account", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Input fields
	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Enter Account Holder Name")

	accountTypeSelect := widget.NewSelect([]string{"Savings", "Current"}, nil)
	accountTypeSelect.SetSelectedIndex(0) // Default to Savings

	initialBalanceEntry := widget.NewEntry()
	initialBalanceEntry.SetPlaceHolder("Enter Initial Balance")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Account Holder Name", Widget: nameEntry},
			{Text: "Account Type", Widget: accountTypeSelect},
			{Text: "Initial Balance ($)", Widget: initialBalanceEntry},
		},
		OnSubmit: func() {
			// Validate input
			if nameEntry.Text == "" {
				dialog.ShowError(errors.New("account holder name cannot be empty"), w)
				return
			}

			accountType := accountTypeSelect.Selected
			if accountType == "" {
				dialog.ShowError(errors.New("please select an account type"), w)
				return
			}

			initialBalanceStr := initialBalanceEntry.Text
			initialBalance, err := strconv.ParseFloat(initialBalanceStr, 64)
			if err != nil {
				dialog.ShowError(errors.New("invalid initial balance"), w)
				return
			}

			if initialBalance < 0 {
				dialog.ShowError(errors.New("initial balance cannot be negative"), w)
				return
			}

			// Create account based on type
			if strings.ToLower(accountType) == "savings" {
				bank.AddAccount(nameEntry.Text, banking.NewSavingsAccount(initialBalance))
				dialog.ShowInformation("Success", fmt.Sprintf("Savings account created for %s with balance $%.2f", nameEntry.Text, initialBalance), w)
			} else {
				bank.AddAccount(nameEntry.Text, banking.NewCurrentAccount(initialBalance))
				dialog.ShowInformation("Success", fmt.Sprintf("Current account created for %s with balance $%.2f", nameEntry.Text, initialBalance), w)
			}

			// Return to welcome screen
			w.SetContent(createWelcomeScreen(bank, w))
		},
	}

	backBtn := widget.NewButton("Back", func() {
		w.SetContent(createWelcomeScreen(bank, w))
	})

	content := container.NewVBox(
		title,
		widget.NewLabel(""), // spacing
		form,
		backBtn,
	)

	return container.NewCenter(container.NewVBox(
		layout.NewSpacer(),
		content,
		layout.NewSpacer(),
	))
}

// selectAccountScreen creates the account selection screen
func selectAccountScreen(bank *banking.Bank, w fyne.Window) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("Select Account", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Display accounts
	accountNames := bank.GetAllAccountNames()

	content := container.NewVBox(title, widget.NewLabel("")) // spacing

	if len(accountNames) == 0 {
		noAccountsLabel := widget.NewLabel("No accounts found. Please create an account first.")
		content.Add(noAccountsLabel)
	} else {
		accountList := widget.NewList(
			func() int {
				return len(accountNames)
			},
			func() fyne.CanvasObject {
				return container.NewHBox(widget.NewLabel(""), widget.NewLabel(""))
			},
			func(id widget.ListItemID, item fyne.CanvasObject) {
				name := accountNames[id]
				account, _ := bank.GetAccount(name)

				items := item.(*fyne.Container).Objects
				nameLabel := items[0].(*widget.Label)
				balanceLabel := items[1].(*widget.Label)

				nameLabel.SetText(fmt.Sprintf("%s (%s)", name, account.AccountType()))
				balanceLabel.SetText(fmt.Sprintf("$%.2f", account.CheckBalance()))
			},
		)

		accountList.OnSelected = func(id widget.ListItemID) {
			selectedName := accountNames[id]
			w.SetContent(accountDashboard(bank, selectedName, w))
		}

		content.Add(container.NewVScroll(accountList))
	}

	backBtn := widget.NewButton("Back", func() {
		w.SetContent(createWelcomeScreen(bank, w))
	})

	content.Add(backBtn)

	return container.NewCenter(container.NewVBox(
		layout.NewSpacer(),
		content,
		layout.NewSpacer(),
	))
}

// accountDashboard creates the dashboard for a selected account
func accountDashboard(bank *banking.Bank, accountName string, w fyne.Window) fyne.CanvasObject {
	account, exists := bank.GetAccount(accountName)
	if !exists {
		dialog.ShowError(errors.New("account not found"), w)
		return createWelcomeScreen(bank, w)
	}

	// Display account information
	title := widget.NewLabelWithStyle(fmt.Sprintf("Account: %s", accountName), fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	accountTypeLabel := widget.NewLabel(fmt.Sprintf("Account Type: %s", account.AccountType()))
	balanceLabel := widget.NewLabel(fmt.Sprintf("Current Balance: $%.2f", account.CheckBalance()))

	// Create buttons for actions
	depositBtn := widget.NewButton("Deposit", func() {
		showDepositDialog(bank, accountName, balanceLabel, w)
	})

	withdrawBtn := widget.NewButton("Withdraw", func() {
		showWithdrawDialog(bank, accountName, balanceLabel, w)
	})

	transactionHistoryBtn := widget.NewButton("Transaction History", func() {
		w.SetContent(transactionHistoryScreen(bank, accountName, w))
	})

	backBtn := widget.NewButton("Back to Account Selection", func() {
		w.SetContent(selectAccountScreen(bank, w))
	})

	mainMenuBtn := widget.NewButton("Main Menu", func() {
		w.SetContent(createWelcomeScreen(bank, w))
	})

	content := container.NewVBox(
		title,
		widget.NewLabel(""), // spacing
		accountTypeLabel,
		balanceLabel,
		widget.NewLabel(""), // spacing
		depositBtn,
		withdrawBtn,
		transactionHistoryBtn,
		widget.NewLabel(""), // spacing
		backBtn,
		mainMenuBtn,
	)

	return container.NewCenter(container.NewVBox(
		layout.NewSpacer(),
		content,
		layout.NewSpacer(),
	))
}

// showDepositDialog shows the deposit dialog
func showDepositDialog(bank *banking.Bank, accountName string, balanceLabel *widget.Label, w fyne.Window) {
	account, _ := bank.GetAccount(accountName)

	amountEntry := widget.NewEntry()
	amountEntry.SetPlaceHolder("Enter deposit amount")

	dialogContent := container.NewVBox(
		widget.NewLabel("Enter deposit amount:"),
		amountEntry,
	)

	depositBtn := widget.NewButton("Deposit", func() {
		amountStr := amountEntry.Text
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			dialog.ShowError(errors.New("invalid amount"), w)
			return
		}

		err = account.Deposit(amount)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		bank.RecordTransaction(accountName, "Deposit", amount)
		balanceLabel.SetText(fmt.Sprintf("Current Balance: $%.2f", account.CheckBalance()))
		dialog.ShowInformation("Success", fmt.Sprintf("Successfully deposited $%.2f", amount), w)
	})

	dialogContent.Add(depositBtn)

	depositDialog := dialog.NewCustom("Deposit", "Cancel", dialogContent, w)
	depositDialog.Show()
}

// showWithdrawDialog shows the withdraw dialog
func showWithdrawDialog(bank *banking.Bank, accountName string, balanceLabel *widget.Label, w fyne.Window) {
	account, _ := bank.GetAccount(accountName)

	amountEntry := widget.NewEntry()
	amountEntry.SetPlaceHolder("Enter withdrawal amount")

	dialogContent := container.NewVBox(
		widget.NewLabel("Enter withdrawal amount:"),
		amountEntry,
	)

	withdrawBtn := widget.NewButton("Withdraw", func() {
		amountStr := amountEntry.Text
		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			dialog.ShowError(errors.New("invalid amount"), w)
			return
		}

		err = account.Withdraw(amount)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		bank.RecordTransaction(accountName, "Withdrawal", amount)
		balanceLabel.SetText(fmt.Sprintf("Current Balance: $%.2f", account.CheckBalance()))
		dialog.ShowInformation("Success", fmt.Sprintf("Successfully withdrew $%.2f", amount), w)
	})

	dialogContent.Add(withdrawBtn)

	withdrawDialog := dialog.NewCustom("Withdraw", "Cancel", dialogContent, w)
	withdrawDialog.Show()
}

// transactionHistoryScreen creates the transaction history screen
func transactionHistoryScreen(bank *banking.Bank, accountName string, w fyne.Window) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("Transaction History", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	accountLabel := widget.NewLabel(fmt.Sprintf("Account: %s", accountName))

	transactions := bank.GetTransactionHistory(accountName)

	content := container.NewVBox(
		title,
		accountLabel,
		widget.NewLabel(""), // spacing
	)

	if len(transactions) == 0 {
		content.Add(widget.NewLabel("No transactions found."))
	} else {
		transactionsList := widget.NewList(
			func() int {
				return len(transactions)
			},
			func() fyne.CanvasObject {
				return container.NewHBox(
					widget.NewLabel(""), // Type
					widget.NewLabel(""), // Amount
					widget.NewLabel(""), // Timestamp
				)
			},
			func(id widget.ListItemID, item fyne.CanvasObject) {
				transaction := transactions[id]

				items := item.(*fyne.Container).Objects
				typeLabel := items[0].(*widget.Label)
				amountLabel := items[1].(*widget.Label)
				timestampLabel := items[2].(*widget.Label)

				typeLabel.SetText(transaction.TransactionType)
				amountLabel.SetText(fmt.Sprintf("$%.2f", transaction.Amount))
				timestampLabel.SetText(transaction.Timestamp)
			},
		)

		content.Add(container.NewVScroll(transactionsList))
	}

	backBtn := widget.NewButton("Back to Account", func() {
		w.SetContent(accountDashboard(bank, accountName, w))
	})

	content.Add(backBtn)

	return container.NewCenter(container.NewVBox(
		layout.NewSpacer(),
		content,
		layout.NewSpacer(),
	))
}
