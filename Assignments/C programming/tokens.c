
#include <stdio.h>
#include <string.h>

#define MAX_ACCOUNTS 100
#define MAX_NAME_LENGTH 50

// Structure to store account details
struct Account {
    int accountNumber;
    char name[MAX_NAME_LENGTH];
    float balance;
};

// Global array to store multiple accounts
struct Account accounts[MAX_ACCOUNTS];
int accountCount = 0;

// Function to create a new account
void createAccount(int accountNumber, char name[], float initialDeposit) {
    if (accountCount < MAX_ACCOUNTS) {
        accounts[accountCount].accountNumber = accountNumber;
        strcpy(accounts[accountCount].name, name);
        accounts[accountCount].balance = initialDeposit;
        accountCount++;
        printf("Account created successfully!\n");
    } else {
        printf("Maximum account limit reached!\n");
    }
}

// Function to deposit money into an account
void deposit(int accountNumber, float amount) {
    for (int i = 0; i < accountCount; i++) {
        if (accounts[i].accountNumber == accountNumber) {
            accounts[i].balance += amount;
            printf("Amount deposited successfully!\n");
            return;
        }
    }
    printf("Account not found!\n");
}

// Function to withdraw money from an account
void withdraw(int accountNumber, float amount) {
    for (int i = 0; i < accountCount; i++) {
        if (accounts[i].accountNumber == accountNumber) {
            if (accounts[i].balance >= amount) {
                accounts[i].balance -= amount;
                printf("Amount withdrawn successfully!\n");
            } else {
                printf("Insufficient balance!\n");
            }
            return;
        }
    }
    printf("Account not found!\n");
}

// Function to check account balance
void checkBalance(int accountNumber) {
    for (int i = 0; i < accountCount; i++) {
        if (accounts[i].accountNumber == accountNumber) {
            printf("Account Balance: $%.2f\n", accounts[i].balance);
            return;
        }
    }
    printf("Account not found!\n");
}

// Main function to demonstrate the banking system
int main() {
    createAccount(1, "John Doe", 1000.0);
    createAccount(2, "Jane Smith", 500.0);
    
    deposit(1, 200.0);
    withdraw(2, 100.0);
    checkBalance(1);
    checkBalance(2);
    
    return 0;
}
