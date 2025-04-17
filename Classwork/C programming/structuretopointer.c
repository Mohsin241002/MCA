#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_NAME_LENGTH 100

// Structure definition for a bank account
typedef struct {
    int accountNumber;
    char accountHolderName[MAX_NAME_LENGTH];
    float balance;
    char accountType[MAX_NAME_LENGTH];
} BankAccount;

// Function prototypes
void createAccount(BankAccount *account);
void displayAccount(const BankAccount *account);
void deposit(BankAccount *account, float amount);
void withdraw(BankAccount *account, float amount);

int main() {
    BankAccount account; // Create a BankAccount structure
    BankAccount *pAccount = &account; // Pointer to the BankAccount structure
    int choice;
    float amount;

    while (1) {
        printf("\nBanking System Menu:\n");
        printf("1. Create Account\n");
        printf("2. Display Account\n");
        printf("3. Deposit\n");
        printf("4. Withdraw\n");
        printf("5. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                createAccount(pAccount);
                break;
            case 2:
                displayAccount(pAccount);
                break;
            case 3:
                printf("Enter amount to deposit: ");
                scanf("%f", &amount);
                deposit(pAccount, amount);
                break;
            case 4:
                printf("Enter amount to withdraw: ");
                scanf("%f", &amount);
                withdraw(pAccount, amount);
                break;
            case 5:
                printf("Exiting...\n");
                exit(0);
            default:
                printf("Invalid choice. Please try again.\n");
        }
    }
    
    return 0;
}

void createAccount(BankAccount *account) {
    printf("Enter account number: ");
    scanf("%d", &account->accountNumber);
    printf("Enter account holder's name: ");
    scanf("%s", account->accountHolderName);
    printf("Enter balance: ");
    scanf("%f", &account->balance);
    printf("Enter account type: ");
    scanf("%s", account->accountType);
    printf("Account created successfully!\n");
}

void displayAccount(const BankAccount *account) {
    printf("\nAccount Details:\n");
    printf("Account Number: %d\n", account->accountNumber);
    printf("Account Holder's Name: %s\n", account->accountHolderName);
    printf("Balance: %.2f\n", account->balance);
    printf("Account Type: %s\n", account->accountType);
}

void deposit(BankAccount *account, float amount) {
    if (amount > 0) {
        account->balance += amount;
        printf("Deposited %.2f successfully. New balance is %.2f\n", amount, account->balance);
    } else {
        printf("Invalid deposit amount.\n");
    }
}

void withdraw(BankAccount *account, float amount) {
    if (amount > 0 && amount <= account->balance) {
        account->balance -= amount;
        printf("Withdrew %.2f successfully. New balance is %.2f\n", amount, account->balance);
    } else if (amount > account->balance) {
        printf("Insufficient balance.\n");
    } else {
        printf("Invalid withdrawal amount.\n");
    }
}
