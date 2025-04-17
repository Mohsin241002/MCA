# Go Banking System GUI

A graphical user interface for the banking system implemented using Fyne, a cross-platform UI toolkit for Go.

## Features

- Modern, user-friendly GUI interface
- All features of the terminal-based banking system:
  - Multiple account types (Savings and Current)
  - Balance tracking with pointers
  - Transaction history
  - Error handling
  - Multiple user accounts

## Requirements

- Go 1.16 or higher
- Fyne toolkit and its dependencies

## Installation

### 1. Install Go

If you haven't installed Go, download and install it from [golang.org](https://golang.org/dl/).

### 2. Install Fyne

```bash
go get fyne.io/fyne/v2
```

### 3. Install required dependencies

#### For Windows:

1. Install GCC
   - Download and install MinGW-w64 from [mingw-w64.org](https://www.mingw-w64.org/) or through MSYS2
   - Add MinGW bin directory to your PATH

2. Install dependencies for building GUI applications
   ```bash
   go get fyne.io/fyne/v2/cmd/fyne
   ```

#### For macOS:

```bash
brew install gcc
go get fyne.io/fyne/v2/cmd/fyne
```

#### For Linux:

```bash
sudo apt-get install gcc libgl1-mesa-dev xorg-dev
go get fyne.io/fyne/v2/cmd/fyne
```

## How to Run

1. Navigate to the project directory
2. Build and run the application:

```bash
go run banking_gui.go
```

## How to Build an Executable

To create a standalone executable:

```bash
go build -o banking_app banking_gui.go
```

For a more optimized build with Fyne packaging:

```bash
fyne package -os windows -icon icon.png
```

## Usage

1. **Launch the application**:
   - The welcome screen appears with options to create or select an account

2. **Create an account**:
   - Click "Create Account"
   - Enter account holder name
   - Select account type (Savings/Current)
   - Enter initial balance
   - Click "Submit"

3. **Select an account**:
   - Click "Select Account"
   - Choose an account from the list

4. **Perform transactions**:
   - Click "Deposit" to add funds
   - Click "Withdraw" to remove funds
   - View transaction history
   - Check current balance

5. **Navigate between screens**:
   - Use "Back" buttons to return to previous screens
   - Use "Main Menu" to return to the welcome screen

## Screenshots

(Screenshots would be included here once the application is running)

## Troubleshooting

### Common Issues:

1. **Missing dependencies**: Ensure all required packages are installed using the commands above.

2. **Compilation errors**: Make sure you have the correct version of Go and Fyne installed.

3. **Display issues**: If the application doesn't display correctly, try updating your graphics drivers.

4. **Build errors**:
   - For Windows users, ensure MinGW is properly installed and in your PATH
   - For Linux users, make sure the required development packages are installed 