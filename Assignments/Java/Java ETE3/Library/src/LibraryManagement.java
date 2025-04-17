import java.sql.*;
import java.util.Scanner;

public class LibraryManagement {

    private static final String URL = "jdbc:mysql://localhost:3306/library";
    private static final String USER = "root";
    private static final String PASSWORD = "admin";

    public static void main(String[] args) {
        try (Connection conn = DriverManager.getConnection(URL, USER, PASSWORD);
             Scanner scanner = new Scanner(System.in)) {

            System.out.println("Library Management System");
            while (true) {
                System.out.println("\nMenu:");
                System.out.println("1. Add Book");
                System.out.println("2. Update Book");
                System.out.println("3. Delete Book");
                System.out.println("4. View Books (Submenu)");
                System.out.println("5. Exit");
                System.out.print("Enter your choice: ");
                int choice = scanner.nextInt();
                scanner.nextLine(); // Consume newline

                switch (choice) {
                    case 1:
                        addBook(conn, scanner);
                        break;
                    case 2:
                        updateBook(conn, scanner);
                        break;
                    case 3:
                        deleteBook(conn, scanner);
                        break;
                    case 4:
                        viewBooksSubMenu(conn, scanner);
                        break;
                    case 5:
                        System.out.println("Exiting...");
                        return;
                    default:
                        System.out.println("Invalid choice. Try again.");
                }
            }
        } catch (SQLException e) {
            e.printStackTrace();
        }
    }

    private static void addBook(Connection conn, Scanner scanner) throws SQLException {
        System.out.print("Enter book ID: ");
        int id = scanner.nextInt();
        scanner.nextLine(); // Consume newline

        System.out.print("Enter title: ");
        String title = scanner.nextLine();
        System.out.print("Enter author: ");
        String author = scanner.nextLine();
        System.out.print("Enter genre: ");
        String genre = scanner.nextLine();
        System.out.print("Enter price: ");
        double price = scanner.nextDouble();
        scanner.nextLine(); // Consume newline

        String sql = "INSERT INTO books (id, title, author, genre, price) VALUES (?, ?, ?, ?, ?)";
        try (PreparedStatement pstmt = conn.prepareStatement(sql)) {
            pstmt.setInt(1, id);
            pstmt.setString(2, title);
            pstmt.setString(3, author);
            pstmt.setString(4, genre);
            pstmt.setDouble(5, price);
            int rowsInserted = pstmt.executeUpdate();
            System.out.println(rowsInserted + " book(s) added.");
        }
    }

    private static void updateBook(Connection conn, Scanner scanner) throws SQLException {
        System.out.print("Enter book ID to update: ");
        int id = scanner.nextInt();
        scanner.nextLine(); // Consume newline

        System.out.print("Enter new title: ");
        String title = scanner.nextLine();
        System.out.print("Enter new author: ");
        String author = scanner.nextLine();
        System.out.print("Enter new genre: ");
        String genre = scanner.nextLine();
        System.out.print("Enter new price: ");
        double price = scanner.nextDouble();
        scanner.nextLine(); // Consume newline

        String sql = "UPDATE books SET title = ?, author = ?, genre = ?, price = ? WHERE id = ?";
        try (PreparedStatement pstmt = conn.prepareStatement(sql)) {
            pstmt.setString(1, title);
            pstmt.setString(2, author);
            pstmt.setString(3, genre);
            pstmt.setDouble(4, price);
            pstmt.setInt(5, id);
            int rowsUpdated = pstmt.executeUpdate();
            System.out.println(rowsUpdated + " book(s) updated.");
        }
    }

    private static void deleteBook(Connection conn, Scanner scanner) throws SQLException {
        System.out.print("Enter book ID to delete: ");
        int id = scanner.nextInt();
        scanner.nextLine(); // Consume newline

        String sql = "DELETE FROM books WHERE id = ?";
        try (PreparedStatement pstmt = conn.prepareStatement(sql)) {
            pstmt.setInt(1, id);
            int rowsDeleted = pstmt.executeUpdate();
            System.out.println(rowsDeleted + " book(s) deleted.");
        }
    }

    private static void viewBooksSubMenu(Connection conn, Scanner scanner) throws SQLException {
        System.out.println("\nView Books Submenu:");
        System.out.println("1. View All Books");
        System.out.println("2. Search Book by Title");
        System.out.print("Enter your choice: ");
        int choice = scanner.nextInt();
        scanner.nextLine(); // Consume newline

        switch (choice) {
            case 1:
                viewAllBooks(conn);
                break;
            case 2:
                searchBookByTitle(conn, scanner);
                break;
            default:
                System.out.println("Invalid choice. Returning to main menu.");
        }
    }

    private static void viewAllBooks(Connection conn) throws SQLException {
        String sql = "SELECT * FROM books";
        try (Statement stmt = conn.createStatement();
             ResultSet rs = stmt.executeQuery(sql)) {

            System.out.println("\nBooks in Library:");
            System.out.println("ID | Title       | Author     | Genre       | Price");
            while (rs.next()) {
                System.out.printf("%d | %s | %s | %s | %.2f%n",
                        rs.getInt("id"),
                        rs.getString("title"),
                        rs.getString("author"),
                        rs.getString("genre"),
                        rs.getDouble("price"));
            }
        }
    }

    private static void searchBookByTitle(Connection conn, Scanner scanner) throws SQLException {
        System.out.print("Enter the title of the book to search: ");
        String title = scanner.nextLine();

        String sql = "SELECT * FROM books WHERE title LIKE ?";
        try (PreparedStatement pstmt = conn.prepareStatement(sql)) {
            pstmt.setString(1, "%" + title + "%");
            ResultSet rs = pstmt.executeQuery();

            System.out.println("\nSearch Results:");
            System.out.println("ID | Title       | Author     | Genre       | Price");
            boolean found = false;
            while (rs.next()) {
                found = true;
                System.out.printf("%d | %s | %s | %s | %.2f%n",
                        rs.getInt("id"),
                        rs.getString("title"),
                        rs.getString("author"),
                        rs.getString("genre"),
                        rs.getDouble("price"));
            }
            if (!found) {
                System.out.println("No books found with the given title.");
            }
        }
    }
}
