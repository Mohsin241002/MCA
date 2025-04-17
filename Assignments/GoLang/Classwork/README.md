# Go Banking System

A banking system implemented in Go that uses pointers, interfaces, and functions while incorporating multiple account types, error handling, and user interaction. The system is available in both terminal and GUI versions.

## Features

- Account interface with Deposit, Withdraw, CheckBalance, and AccountType methods
- Two account types: Savings Account and Current Account
- Balance stored as a pointer to a float64
- Savings accounts have a withdrawal limit of $500 per transaction
- Current accounts have an overdraft limit of $200
- Multiple user accounts with a transaction history feature
- Available in both terminal and GUI versions

## Project Structure

```
banking/
├── banking/           # Core banking logic package
│   └── banking.go     # Banking interfaces and implementations
├── banking_terminal.go # Terminal-based user interface
├── banking_gui.go     # GUI-based user interface (requires Fyne)
├── go.mod             # Go module file
├── go.sum             # Go module checksums
├── README.md          # This file
└── README_GUI.md      # Detailed GUI setup instructions
```

## Requirements

- Go 1.16 or higher
- For GUI version: Fyne toolkit and its dependencies (see README_GUI.md for details)

## How to Run

### Terminal Version

```bash
go run banking_terminal.go
```

### GUI Version

First, make sure you have Fyne and its dependencies installed (see README_GUI.md).

```bash
go run banking_gui.go
```

## Usage

### Terminal Version

The terminal interface allows you to:
1. Create an account (Savings/Current)
2. Select an existing account
3. Deposit money
4. Withdraw money (with appropriate limits)
5. Check balance
6. View transaction history
7. Switch between users

### GUI Version

The GUI version provides a more user-friendly interface with:
- Welcome screen
- Account creation form
- Account selection list
- Account dashboard showing details and transactions
- Deposit and withdrawal dialogs
- Transaction history display

See README_GUI.md for detailed instructions on setting up and using the GUI version.

## Account Types

### Savings Account
- Withdrawal limit: $500 per transaction
- No overdraft facility

### Current Account
- No withdrawal limit
- Overdraft facility up to $200

## Error Handling

The system includes comprehensive error handling for:
- Negative deposit/withdrawal amounts
- Insufficient funds
- Exceeding withdrawal limits
- Invalid account types

## Architecture

The system uses:
- **Interfaces**: To define account behavior
- **Pointers**: For storing and manipulating account balances
- **Structs**: To implement different account types
- **Maps**: To store multiple user accounts
- **Slices**: To maintain transaction history
- **Packages**: To separate core banking logic from user interfaces 