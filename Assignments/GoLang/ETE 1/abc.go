package main

import (
	"fmt"
	"strings"
	"time"
)

// Role-based borrowing limits and fine rates
var borrowingLimits = map[string]int{"student": 3, "faculty": 5, "staff": 2}
var borrowingDays = map[string]int{"student": 15, "faculty": 30, "staff": 10}
var fineRates = map[string]int{"student": 5, "faculty": 2, "staff": 7}

type Book struct {
	ISBN       string
	Title      string
	Author     string
	Category   string
	Available  bool
	BorrowedBy string // Stores ID of borrower
}

type User struct {
	ID       string
	Name     string
	Role     string // student, faculty, staff
	Borrowed []BorrowRecord
	Fine     int
}

type BorrowRecord struct {
	ISBN       string
	IssueDate  time.Time
	DueDate    time.Time
	ReturnDate time.Time
}

var books []Book
var users map[string]*User // Key: User ID

func main() {
	users = make(map[string]*User)
	initializeBooks()
	for {
		fmt.Println("\n===== UNIVERSITY LIBRARY MANAGEMENT SYSTEM =====")
		fmt.Println("1. Register User")
		fmt.Println("2. Borrow Book")
		fmt.Println("3. Return Book")
		fmt.Println("4. View Borrowed Books")
		fmt.Println("5. Pay Fine")
		fmt.Println("6. Generate Reports")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			registerUser()
		case 2:
			borrowBook()
		case 3:
			returnBook()
		case 4:
			viewBorrowedBooks()
		case 5:
			payFine()
		case 6:
			generateReports()
		case 7:
			fmt.Println("Exiting... Thank you!")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func initializeBooks() {
	books = []Book{
		{"F1001", "Harry Potter and the Philosopher's Stone", "J.K. Rowling", "Fantasy", true, ""},
		{"F1002", "The Lord of the Rings", "J.R.R. Tolkien", "Fantasy", true, ""},
		{"F1003", "Pride and Prejudice", "Jane Austen", "Classic", true, ""},
		{"F1004", "To Kill a Mockingbird", "Harper Lee", "Fiction", true, ""},
		{"F1005", "The Great Gatsby", "F. Scott Fitzgerald", "Classic", true, ""},
		{"F1006", "1984", "George Orwell", "Dystopian", true, ""},
		{"F1007", "The Catcher in the Rye", "J.D. Salinger", "Fiction", true, ""},
		{"F1008", "The Hobbit", "J.R.R. Tolkien", "Fantasy", true, ""},
		{"F1009", "Brave New World", "Aldous Huxley", "Dystopian", true, ""},
		{"F1010", "The Chronicles of Narnia", "C.S. Lewis", "Fantasy", true, ""},
		{"F1011", "Jane Eyre", "Charlotte Brontë", "Classic", true, ""},
		{"F1012", "The Hunger Games", "Suzanne Collins", "Dystopian", true, ""},
		{"F1013", "The Alchemist", "Paulo Coelho", "Fiction", true, ""},
		{"F1014", "Wuthering Heights", "Emily Brontë", "Classic", true, ""},
		{"F1015", "The Da Vinci Code", "Dan Brown", "Mystery", true, ""},
		{"F1016", "The Kite Runner", "Khaled Hosseini", "Fiction", true, ""},
		{"F1017", "A Game of Thrones", "George R.R. Martin", "Fantasy", true, ""},
		{"F1018", "The Fault in Our Stars", "John Green", "Young Adult", true, ""},
		{"F1019", "Little Women", "Louisa May Alcott", "Classic", true, ""},
		{"F1020", "Moby-Dick", "Herman Melville", "Adventure", true, ""},
	}
}

func registerUser() {
	var id, name, role string
	fmt.Print("Enter User ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter Role (student/faculty/staff): ")
	fmt.Scan(&role)
	role = validateRole(role)
	users[id] = &User{ID: id, Name: name, Role: role, Borrowed: []BorrowRecord{}, Fine: 0}
	fmt.Println("User registered successfully!")
}

func validateRole(role string) string {
	for role != "student" && role != "faculty" && role != "staff" {
		fmt.Print("Invalid role! Enter Role (student/faculty/staff): ")
		fmt.Scan(&role)
	}
	return role
}

func borrowBook() {
	var id, isbn string
	fmt.Print("Enter User ID: ")
	fmt.Scan(&id)
	user, exists := users[id]
	if !exists {
		fmt.Println("User not found!")
		return
	}

	if len(user.Borrowed) >= borrowingLimits[user.Role] {
		fmt.Println("Borrowing limit reached!")
		return
	}

	// Display available books
	fmt.Println("\n=== Available Books ===")
	fmt.Printf("%-10s %-40s %-25s %-15s\n", "ISBN", "Title", "Author", "Category")
	fmt.Println(strings.Repeat("-", 90))
	for _, book := range books {
		if book.Available {
			fmt.Printf("%-10s %-40s %-25s %-15s\n", book.ISBN, book.Title, book.Author, book.Category)
		}
	}
	fmt.Println(strings.Repeat("-", 90))

	fmt.Print("Enter Book ISBN: ")
	fmt.Scan(&isbn)

	// Check if book exists and is available
	bookFound := false
	for i, book := range books {
		if book.ISBN == isbn && book.Available {
			bookFound = true

			// Get borrow date in dd/mm/yyyy format
			var borrowDateStr string
			fmt.Print("Enter Borrow Date (dd/mm/yyyy): ")
			fmt.Scan(&borrowDateStr)

			// Parse the borrow date
			borrowDate, err := time.Parse("02/01/2006", borrowDateStr)
			if err != nil {
				fmt.Println("Invalid date format! Please use dd/mm/yyyy")
				return
			}

			// Calculate due date based on role
			dueDate := borrowDate.AddDate(0, 0, borrowingDays[user.Role])

			// Update book status
			books[i].Available = false
			books[i].BorrowedBy = id

			// Add to user's borrowed records
			user.Borrowed = append(user.Borrowed, BorrowRecord{
				ISBN:      isbn,
				IssueDate: borrowDate,
				DueDate:   dueDate,
			})

			fmt.Println("Book borrowed successfully!")
			fmt.Printf("Due Date: %s\n", dueDate.Format("02/01/2006"))
			return
		}
	}

	if !bookFound {
		fmt.Println("Book not available or ISBN not found!")
	}
}

func returnBook() {
	var id, isbn string
	fmt.Print("Enter User ID: ")
	fmt.Scan(&id)
	user, exists := users[id]
	if !exists {
		fmt.Println("User not found!")
		return
	}

	// Display borrowed books for this user
	if len(user.Borrowed) == 0 {
		fmt.Println("No books borrowed by this user!")
		return
	}

	fmt.Println("\n=== Books Borrowed by User ===")
	fmt.Printf("%-10s %-40s %-15s %-15s\n", "ISBN", "Title", "Borrow Date", "Due Date")
	fmt.Println(strings.Repeat("-", 80))

	for _, record := range user.Borrowed {
		// Find book details
		for _, book := range books {
			if book.ISBN == record.ISBN {
				fmt.Printf("%-10s %-40s %-15s %-15s\n",
					record.ISBN,
					book.Title,
					record.IssueDate.Format("02/01/2006"),
					record.DueDate.Format("02/01/2006"))
				break
			}
		}
	}
	fmt.Println(strings.Repeat("-", 80))

	fmt.Print("Enter Book ISBN to return: ")
	fmt.Scan(&isbn)

	// Get return date
	var returnDateStr string
	fmt.Print("Enter Return Date (dd/mm/yyyy): ")
	fmt.Scan(&returnDateStr)

	// Parse the return date
	returnDate, err := time.Parse("02/01/2006", returnDateStr)
	if err != nil {
		fmt.Println("Invalid date format! Please use dd/mm/yyyy")
		return
	}

	for i, record := range user.Borrowed {
		if record.ISBN == isbn {
			// Calculate fine based on return date and due date
			var fine int = 0
			if returnDate.After(record.DueDate) {
				overdueDays := int(returnDate.Sub(record.DueDate).Hours() / 24)
				fine = overdueDays * fineRates[user.Role]
				fmt.Printf("Overdue by %d days! Fine imposed: Rs %d\n", overdueDays, fine)

				// Prompt user to pay fine now
				if fine > 0 {
					var payNow string
					fmt.Print("Do you want to pay the fine now? (yes/no): ")
					fmt.Scan(&payNow)

					if strings.ToLower(payNow) == "yes" {
						// Process fine payment
						processFinePayment(user, fine)
					} else {
						// Add to user's outstanding fine
						user.Fine += fine
						fmt.Printf("Fine added to user's account. Total outstanding fine: Rs %d\n", user.Fine)
					}
				}
			}

			// Update return date in the record before removing it
			user.Borrowed[i].ReturnDate = returnDate

			// Remove from borrowed list
			user.Borrowed = append(user.Borrowed[:i], user.Borrowed[i+1:]...)

			// Update book availability
			updateBookAvailability(isbn)

			fmt.Println("Book returned successfully!")

			// Show updated available books
			fmt.Println("\n=== Updated Available Books ===")
			fmt.Printf("%-10s %-40s %-25s %-15s\n", "ISBN", "Title", "Author", "Category")
			fmt.Println(strings.Repeat("-", 90))
			for _, book := range books {
				if book.Available {
					fmt.Printf("%-10s %-40s %-25s %-15s\n", book.ISBN, book.Title, book.Author, book.Category)
				}
			}
			fmt.Println(strings.Repeat("-", 90))

			return
		}
	}
	fmt.Println("Book not found in borrowed records!")
}

func updateBookAvailability(isbn string) {
	for i := range books {
		if books[i].ISBN == isbn {
			books[i].Available = true
			books[i].BorrowedBy = ""
			return
		}
	}
}

func viewBorrowedBooks() {
	for _, user := range users {
		if len(user.Borrowed) > 0 {
			fmt.Printf("\nUser: %s (%s)\n", user.Name, user.Role)
			for _, record := range user.Borrowed {
				fmt.Printf("ISBN: %s, Due Date: %s\n", record.ISBN, record.DueDate.Format("02/01/2006"))
			}
		}
	}
}

func payFine() {
	var id string
	fmt.Print("Enter User ID: ")
	fmt.Scan(&id)
	user, exists := users[id]
	if !exists {
		fmt.Println("User not found!")
		return
	}

	if user.Fine <= 0 {
		fmt.Println("No outstanding fine for this user!")
		return
	}

	fmt.Printf("Outstanding fine: Rs %d\n", user.Fine)

	var paymentAmount int
	fmt.Print("Enter amount to pay (enter 0 to pay full amount): ")
	fmt.Scan(&paymentAmount)

	if paymentAmount == 0 || paymentAmount >= user.Fine {
		// Pay full amount
		fmt.Printf("Payment of Rs %d received. Fine cleared!\n", user.Fine)
		user.Fine = 0
	} else if paymentAmount > 0 && paymentAmount < user.Fine {
		// Partial payment
		user.Fine -= paymentAmount
		fmt.Printf("Partial payment of Rs %d received. Remaining fine: Rs %d\n",
			paymentAmount, user.Fine)
	} else {
		// Invalid amount
		fmt.Println("Invalid payment amount. No changes made to fine.")
	}
}

func generateReports() {
    fmt.Println("\n===== LIBRARY MANAGEMENT SYSTEM REPORTS =====")=")

    // Get report date from userser
    var reportDateStr stringing
    fmt.Print("Enter date for report (dd/mm/yyyy) or press Enter for today's date: ") ")
    fmt.Scanln(&reportDateStr)tr)

    var reportDate time.Timeime
    var err errorror

    if reportDateStr == "" {" {
        // Use current date if no input input
        reportDate = time.Now().Now()
    } else {e {
        // Parse the input datet date
        reportDate, err = time.Parse("02/01/2006", reportDateStr)teStr)
        if err != nil { nil {
            fmt.Println("Invalid date format! Using today's date instead.")nstead.")
            reportDate = time.Now()ime.Now()
        }     }
    }  }

    fmt.Printf("Generating reports for: %s\n", reportDate.Format("02/01/2006"))"))

    fmt.Println("1. List of Overdue Books")s")
    fmt.Println("2. Total Fines Collected")d")
    fmt.Println("3. Available Books in Library")y")
    fmt.Println("4. Most Borrowed Books")s")
    fmt.Println("5. Frequent Defaulters")s")
    fmt.Println("6. Show the entire report")t")
    fmt.Println("7. Back to Main Menu")u")
    fmt.Print("Enter your choice: ") ")

    var choice intint
    fmt.Scan(&choice)ce)

    switch choice {e {
    case 1: 1:
        reportOverdueBooks(reportDate)tDate)
    case 2: 2:
        reportFinesCollected(reportDate)tDate)
    case 3: 3:
        reportAvailableBooks(reportDate)tDate)
    case 4: 4:
        reportMostBorrowedBooks(reportDate)tDate)
    case 5: 5:
        reportFrequentDefaulters(reportDate)tDate)
    case 6: 6:
        fmt.Println("Report Generated")ated")
        reportOverdueBooks(reportDate)tDate)
        fmt.Println("-----------------------------------")----")
        reportFinesCollected(reportDate)tDate)
        fmt.Println("-----------------------------------")----")
        reportAvailableBooks(reportDate)tDate)
        fmt.Println("-----------------------------------")----")
        reportMostBorrowedBooks(reportDate)tDate)
        fmt.Println("-----------------------------------")----")
        reportFrequentDefaulters(reportDate)tDate)
        fmt.Println("-----------------------------------")----")
    case 7: 7:
        returnreturn
    default:lt:
        fmt.Println("Invalid choice, please try again.")ain.")
    }  }
}

// Report 1: List of Overdue Books
func reportOverdueBooks(reportDate time.Time) {
	fmt.Println("\n========== OVERDUE BOOKS REPORT ==========")
	fmt.Printf("Report Date: %s\n\n", reportDate.Format("02/01/2006"))

	overdueFound := false

	fmt.Printf("%-10s %-40s %-20s %-15s %-15s\n",
		"ISBN", "Title", "Borrower", "Due Date", "Days Overdue")
	fmt.Println(strings.Repeat("-", 100))

	for userID, user := range users {
		for _, record := range user.Borrowed {
			if reportDate.After(record.DueDate) {
				daysOverdue := int(reportDate.Sub(record.DueDate).Hours() / 24)

				// Find book details
				for _, book := range books {
					if book.ISBN == record.ISBN {
						fmt.Printf("%-10s %-40s %-20s %-15s %-15d\n",
							book.ISBN,
							book.Title,
							user.Name+" ("+userID+")",
							record.DueDate.Format("02/01/2006"),
							daysOverdue)
						overdueFound = true
						break
					}
				}
			}
		}
	}

	if !overdueFound {
		fmt.Println("No overdue books found.")
	}

	fmt.Println(strings.Repeat("-", 100))
	fmt.Println("Press Enter to continue...")
	fmt.Scanln() // Wait for user input
}

// Report 2: Total Fines Collected
func reportFinesCollected(reportDate time.Time) {
	fmt.Println("\n========== FINES COLLECTION REPORT ==========")
	fmt.Printf("Report Date: %s\n\n", reportDate.Format("02/01/2006"))

	totalOutstandingFines := 0

	fmt.Printf("%-20s %-20s %-15s\n",
		"User ID", "Name", "Outstanding Fine (Rs)")
	fmt.Println(strings.Repeat("-", 60))

	for id, user := range users {
		if user.Fine > 0 {
			fmt.Printf("%-20s %-20s %-15d\n",
				id, user.Name, user.Fine)
			totalOutstandingFines += user.Fine
		}
	}

	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("Total Outstanding Fines: Rs %d\n", totalOutstandingFines)
	fmt.Println("Press Enter to continue...")
	fmt.Scanln() // Wait for user input
}

// Report 3: Available Books in Library
func reportAvailableBooks(reportDate time.Time) {
	fmt.Println("\n========== AVAILABLE BOOKS REPORT ==========")
	fmt.Printf("Report Date: %s\n\n", reportDate.Format("02/01/2006"))

	availableCount := 0
	totalBooks := len(books)

	fmt.Printf("%-10s %-40s %-25s %-15s\n",
		"ISBN", "Title", "Author", "Category")
	fmt.Println(strings.Repeat("-", 90))

	for _, book := range books {
		if book.Available {
			fmt.Printf("%-10s %-40s %-25s %-15s\n",
				book.ISBN, book.Title, book.Author, book.Category)
			availableCount++
		}
	}

	fmt.Println(strings.Repeat("-", 90))
	fmt.Printf("Available Books: %d out of %d (%.1f%%)\n",
		availableCount, totalBooks, float64(availableCount)/float64(totalBooks)*100)
	fmt.Println("Press Enter to continue...")
	fmt.Scanln() // Wait for user input
}

// Report 4: Most Borrowed Books
func reportMostBorrowedBooks(reportDate time.Time) {
	fmt.Println("\n========== MOST BORROWED BOOKS REPORT ==========")
	fmt.Printf("Report Date: %s\n\n", reportDate.Format("02/01/2006"))

	// Create a map to track borrow count for each book
	borrowCount := make(map[string]int)

	// Count current borrows
	for _, book := range books {
		if !book.Available {
			borrowCount[book.ISBN]++
		}
	}

	// Convert map to slice for sorting
	type BookBorrowCount struct {
		ISBN  string
		Title string
		Count int
	}

	var borrowStats []BookBorrowCount

	for isbn, count := range borrowCount {
		// Find book title
		for _, book := range books {
			if book.ISBN == isbn {
				borrowStats = append(borrowStats, BookBorrowCount{
					ISBN:  isbn,
					Title: book.Title,
					Count: count,
				})
				break
			}
		}
	}

	// Sort by borrow count (descending)
	for i := 0; i < len(borrowStats)-1; i++ {
		for j := i + 1; j < len(borrowStats); j++ {
			if borrowStats[i].Count < borrowStats[j].Count {
				borrowStats[i], borrowStats[j] = borrowStats[j], borrowStats[i]
			}
		}
	}

	// Display results
	fmt.Printf("%-10s %-40s %-15s\n",
		"ISBN", "Title", "Borrow Count")
	fmt.Println(strings.Repeat("-", 70))

	if len(borrowStats) == 0 {
		fmt.Println("No borrowing history available.")
	} else {
		// Display top 10 or all if less than 10
		limit := 10
		if len(borrowStats) < limit {
			limit = len(borrowStats)
		}

		for i := 0; i < limit; i++ {
			fmt.Printf("%-10s %-40s %-15d\n",
				borrowStats[i].ISBN, borrowStats[i].Title, borrowStats[i].Count)
		}
	}

	fmt.Println(strings.Repeat("-", 70))
	fmt.Println("Press Enter to continue...")
	fmt.Scanln() // Wait for user input
}

// Report 5: Frequent Defaulters
func reportFrequentDefaulters(reportDate time.Time) {
	fmt.Println("\n========== FREQUENT DEFAULTERS REPORT ==========")
	fmt.Printf("Report Date: %s\n\n", reportDate.Format("02/01/2006"))

	// Track defaulters and their overdue count
	type DefaulterInfo struct {
		ID           string
		Name         string
		Role         string
		OverdueCount int
		TotalFine    int
	}

	defaulters := make(map[string]*DefaulterInfo)

	// Check current borrowed books for overdue
	for id, user := range users {
		for _, record := range user.Borrowed {
			if reportDate.After(record.DueDate) {
				// User has overdue book
				if _, exists := defaulters[id]; !exists {
					defaulters[id] = &DefaulterInfo{
						ID:           id,
						Name:         user.Name,
						Role:         user.Role,
						OverdueCount: 0,
						TotalFine:    user.Fine,
					}
				}
				defaulters[id].OverdueCount++
			}
		}
	}

	// Convert map to slice for sorting
	var defaultersList []*DefaulterInfo
	for _, info := range defaulters {
		defaultersList = append(defaultersList, info)
	}

	// Sort by overdue count (descending)
	for i := 0; i < len(defaultersList)-1; i++ {
		for j := i + 1; j < len(defaultersList); j++ {
			if defaultersList[i].OverdueCount < defaultersList[j].OverdueCount {
				defaultersList[i], defaultersList[j] = defaultersList[j], defaultersList[i]
			}
		}
	}

	// Display results
	fmt.Printf("%-10s %-20s %-10s %-15s %-15s\n",
		"User ID", "Name", "Role", "Overdue Books", "Total Fine (Rs)")
	fmt.Println(strings.Repeat("-", 75))

	if len(defaultersList) == 0 {
		fmt.Println("No defaulters found.")
	} else {
		for _, defaulter := range defaultersList {
			fmt.Printf("%-10s %-20s %-10s %-15d %-15d\n",
				defaulter.ID, defaulter.Name, defaulter.Role,
				defaulter.OverdueCount, defaulter.TotalFine)
		}
	}

	fmt.Println(strings.Repeat("-", 75))
	fmt.Println("Press Enter to continue...")
	fmt.Scanln() // Wait for user input
}

// Helper function to process fine payment
func processFinePayment(user *User, fineDue int) {
	fmt.Printf("Fine due: Rs %d\n", fineDue)

	var paymentAmount int
	fmt.Print("Enter amount to pay (enter 0 to pay full amount): ")
	fmt.Scan(&paymentAmount)

	if paymentAmount == 0 || paymentAmount >= fineDue {
		// Pay full amount
		fmt.Printf("Payment of Rs %d received. Fine cleared!\n", fineDue)
	} else if paymentAmount > 0 && paymentAmount < fineDue {
		// Partial payment
		remainingFine := fineDue - paymentAmount
		user.Fine += remainingFine
		fmt.Printf("Partial payment of Rs %d received. Remaining fine (Rs %d) added to account.\n",
			paymentAmount, remainingFine)
	} else {
		// Invalid amount
		fmt.Println("Invalid payment amount. Full fine added to account.")
		user.Fine += fineDue
	}
}
