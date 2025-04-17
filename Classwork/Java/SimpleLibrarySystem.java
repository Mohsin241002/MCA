import java.util.ArrayList;
import java.util.List;

public class SimpleLibrarySystem {

    // Book class
    static class Book {
        int bookID;
        String title;
        String author;
        boolean isAvailable;

        public Book(int bookID, String title, String author) {
            this.bookID = bookID;
            this.title = title;
            this.author = author;
            this.isAvailable = true;
        }

        @Override
        public String toString() {
            return "BookID: " + bookID + ", Title: " + title + ", Author: " + author + ", Available: " + isAvailable;
        }
    }

    // Member class
    static class Member {
        int memberID;
        String name;
        List<Book> borrowedBooks;

        public Member(int memberID, String name) {
            this.memberID = memberID;
            this.name = name;
            this.borrowedBooks = new ArrayList<>();
        }

        public void borrowBook(Book book) {
            if (book.isAvailable) {
                borrowedBooks.add(book);
                book.isAvailable = false;
                System.out.println(name + " borrowed: " + book.title);
            } else {
                System.out.println(book.title + " is not available.");
            }
        }

        public void returnBook(Book book) {
            if (borrowedBooks.contains(book)) {
                borrowedBooks.remove(book);
                book.isAvailable = true;
                System.out.println(name + " returned: " + book.title);
            } else {
                System.out.println("This book is not borrowed by " + name);
            }
        }

        public void listBorrowedBooks() {
            System.out.println("Books borrowed by " + name + ":");
            if (borrowedBooks.isEmpty()) {
                System.out.println("No books borrowed.");
            } else {
                for (Book book : borrowedBooks) {
                    System.out.println(book.title);
                }
            }
        }
    }

    // Library class
    static class Library {
        List<Book> books;
        List<Member> members;

        public Library() {
            books = new ArrayList<>();
            members = new ArrayList<>();
        }

        public void addBook(int bookID, String title, String author) {
            Book book = new Book(bookID, title, author);
            books.add(book);
            System.out.println("Added book: " + title);
        }

        public void registerMember(int memberID, String name) {
            Member member = new Member(memberID, name);
            members.add(member);
            System.out.println("Registered member: " + name);
        }

        public void borrowBook(int memberID, int bookID) {
            Member member = findMemberByID(memberID);
            Book book = findBookByID(bookID);

            if (member != null && book != null) {
                member.borrowBook(book);
            } else {
                System.out.println("Invalid member ID or book ID.");
            }
        }

        public void returnBook(int memberID, int bookID) {
            Member member = findMemberByID(memberID);
            Book book = findBookByID(bookID);

            if (member != null && book != null) {
                member.returnBook(book);
            } else {
                System.out.println("Invalid member ID or book ID.");
            }
        }

        public void listAvailableBooks() {
            System.out.println("Available books:");
            for (Book book : books) {
                if (book.isAvailable) {
                    System.out.println(book);
                }
            }
        }

        public void listBorrowedBooks(int memberID) {
            Member member = findMemberByID(memberID);
            if (member != null) {
                member.listBorrowedBooks();
            } else {
                System.out.println("Invalid member ID.");
            }
        }

        private Member findMemberByID(int memberID) {
            for (Member member : members) {
                if (member.memberID == memberID) {
                    return member;
                }
            }
            return null;
        }

        private Book findBookByID(int bookID) {
            for (Book book : books) {
                if (book.bookID == bookID) {
                    return book;
                }
            }
            return null;
        }
    }

    // Main method
    public static void main(String[] args) {
        // Create a library
        Library library = new Library();

        library.addBook(1, "The Great Gatsby", "F. Scott Fitzgerald");
        library.addBook(2, "1984", "George Orwell");
        library.addBook(3, "To Kill a Mockingbird", "Harper Lee");

        library.registerMember(101, "Alice");
        library.registerMember(102, "Bob");

        library.borrowBook(101, 1); // Alice borrows "The Great Gatsby"
        library.borrowBook(102, 2); // Bob borrows "1984"

        library.listAvailableBooks();

        library.listBorrowedBooks(101);

        library.returnBook(101, 1); // Alice returns "The Great Gatsby"

        library.listAvailableBooks();

        library.listBorrowedBooks(101);
    }
}
