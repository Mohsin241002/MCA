package banking

import (
	"errors"
	"fmt"
	"time"
)

// Account interface defines the required methods for bank accounts
type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	CheckBalance() float64
	AccountType() string
}

// SavingsAccount represents a savings account
type SavingsAccount struct {
	balance       *float64
	withdrawLimit float64
}

// CurrentAccount represents a current/checking account
type CurrentAccount struct {
	balance        *float64
	overdraftLimit float64
}

// Transaction represents a single transaction
type Transaction struct {
	TransactionType string
	Amount          float64
	Timestamp       string
}

// Bank represents a collection of accounts with transaction history
type Bank struct {
	accounts     map[string]Account
	transactions map[string][]Transaction
}

// NewSavingsAccount creates a new savings account with an initial balance
func NewSavingsAccount(initialBalance float64) *SavingsAccount {
	balance := initialBalance
	return &SavingsAccount{
		balance:       &balance,
		withdrawLimit: 500.0, // $500 withdrawal limit per transaction
	}
}

// NewCurrentAccount creates a new current account with an initial balance
func NewCurrentAccount(initialBalance float64) *CurrentAccount {
	balance := initialBalance
	return &CurrentAccount{
		balance:        &balance,
		overdraftLimit: 200.0, // $200 overdraft limit
	}
}

// Deposit adds money to the SavingsAccount
func (sa *SavingsAccount) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("deposit amount cannot be negative")
	}
	*sa.balance += amount
	return nil
}

// Withdraw takes money from the SavingsAccount
func (sa *SavingsAccount) Withdraw(amount float64) error {
	if amount < 0 {
		return errors.New("withdrawal amount cannot be negative")
	}
	if amount > sa.withdrawLimit {
		return fmt.Errorf("withdrawal amount exceeds limit of $%.2f", sa.withdrawLimit)
	}
	if amount > *sa.balance {
		return errors.New("insufficient funds")
	}
	*sa.balance -= amount
	return nil
}

// CheckBalance returns the current balance of the SavingsAccount
func (sa *SavingsAccount) CheckBalance() float64 {
	return *sa.balance
}

// AccountType returns the type of account
func (sa *SavingsAccount) AccountType() string {
	return "Savings"
}

// Deposit adds money to the CurrentAccount
func (ca *CurrentAccount) Deposit(amount float64) error {
	if amount < 0 {
		return errors.New("deposit amount cannot be negative")
	}
	*ca.balance += amount
	return nil
}

// Withdraw takes money from the CurrentAccount
func (ca *CurrentAccount) Withdraw(amount float64) error {
	if amount < 0 {
		return errors.New("withdrawal amount cannot be negative")
	}
	// Allow withdrawals up to balance + overdraft limit
	if amount > *ca.balance+ca.overdraftLimit {
		return errors.New("amount exceeds balance and overdraft limit")
	}
	*ca.balance -= amount
	return nil
}

// CheckBalance returns the current balance of the CurrentAccount
func (ca *CurrentAccount) CheckBalance() float64 {
	return *ca.balance
}

// AccountType returns the type of account
func (ca *CurrentAccount) AccountType() string {
	return "Current"
}

// NewBank creates a new bank with empty accounts
func NewBank() *Bank {
	return &Bank{
		accounts:     make(map[string]Account),
		transactions: make(map[string][]Transaction),
	}
}

// AddAccount adds a new account to the bank
func (b *Bank) AddAccount(name string, account Account) {
	b.accounts[name] = account
	b.transactions[name] = []Transaction{}
}

// GetAccount retrieves an account from the bank
func (b *Bank) GetAccount(name string) (Account, bool) {
	account, exists := b.accounts[name]
	return account, exists
}

// RecordTransaction adds a transaction to the history
func (b *Bank) RecordTransaction(accountName, transactionType string, amount float64) {
	transaction := Transaction{
		TransactionType: transactionType,
		Amount:          amount,
		Timestamp:       time.Now().Format(time.RFC3339),
	}
	b.transactions[accountName] = append(b.transactions[accountName], transaction)
}

// GetTransactionHistory retrieves the transaction history for an account
func (b *Bank) GetTransactionHistory(accountName string) []Transaction {
	return b.transactions[accountName]
}

// GetAllAccountNames returns a list of all account names
func (b *Bank) GetAllAccountNames() []string {
	var names []string
	for name := range b.accounts {
		names = append(names, name)
	}
	return names
}
