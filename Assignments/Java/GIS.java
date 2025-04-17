public import java.sql.*;
import java.util.Scanner;

public class GIS {

    // Database connection details
    private static final String DB_URL = "jdbc:mysql://localhost:3306/GamingDB";
    private static final String DB_USER = "root"; 
    private static final String DB_PASSWORD = "admin"; 

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        try (Connection connection = DriverManager.getConnection(DB_URL, DB_USER, DB_PASSWORD)) {
            System.out.println("Connected to the database successfully!");

            while (true) {
                System.out.println("\n--- Gaming Information System ---");
                System.out.println("1. Add a Game");
                System.out.println("2. View All Games");
                System.out.println("3. Update Game Details");
                System.out.println("4. Delete a Game");
                System.out.println("5. Exit");
                System.out.print("Choose an option: ");
                int choice = scanner.nextInt();

                switch (choice) {
                    case 1:
                        addGame(connection, scanner);
                        break;
                    case 2:
                        viewGames(connection);
                        break;
                    case 3:
                        updateGame(connection, scanner);
                        break;
                    case 4:
                        deleteGame(connection, scanner);
                        break;
                    case 5:
                        System.out.println("Exiting... Goodbye!");
                        return;
                    default:
                        System.out.println("Invalid option. Please try again.");
                }
            }
        } catch (SQLException e) {
            System.out.println("Database connection failed: " + e.getMessage());
        }
    }

    // Add a new game
    private static void addGame(Connection connection, Scanner scanner) throws SQLException {
        System.out.print("Enter game name: ");
        String name = scanner.next();
        System.out.print("Enter game genre: ");
        String genre = scanner.next();
        System.out.print("Enter game price: ");
        double price = scanner.nextDouble();

        String sql = "INSERT INTO Games (name, genre, price) VALUES (?, ?, ?)";
        try (PreparedStatement statement = connection.prepareStatement(sql)) {
            statement.setString(1, name);
            statement.setString(2, genre);
            statement.setDouble(3, price);

            int rowsInserted = statement.executeUpdate();
            if (rowsInserted > 0) {
                System.out.println("Game added successfully!");
            }
        }
    }

    // View all games
    private static void viewGames(Connection connection) throws SQLException {
        String sql = "SELECT * FROM Games";
        try (Statement statement = connection.createStatement();
             ResultSet resultSet = statement.executeQuery(sql)) {

            System.out.println("\n--- Games List ---");
            while (resultSet.next()) {
                int id = resultSet.getInt("id");
                String name = resultSet.getString("name");
                String genre = resultSet.getString("genre");
                double price = resultSet.getDouble("price");

                System.out.printf("ID: %d | Name: %s | Genre: %s | Price: %.2f%n", id, name, genre, price);
            }
        }
    }

    // Update game details
    private static void updateGame(Connection connection, Scanner scanner) throws SQLException {
        System.out.print("Enter game ID to update: ");
        int id = scanner.nextInt();

        System.out.print("Enter new game name: ");
        String name = scanner.next();
        System.out.print("Enter new game genre: ");
        String genre = scanner.next();
        System.out.print("Enter new game price: ");
        double price = scanner.nextDouble();

        String sql = "UPDATE Games SET name = ?, genre = ?, price = ? WHERE id = ?";
        try (PreparedStatement statement = connection.prepareStatement(sql)) {
            statement.setString(1, name);
            statement.setString(2, genre);
            statement.setDouble(3, price);
            statement.setInt(4, id);

            int rowsUpdated = statement.executeUpdate();
            if (rowsUpdated > 0) {
                System.out.println("Game updated successfully!");
            } else {
                System.out.println("Game ID not found.");
            }
        }
    }

    // Delete a game
    private static void deleteGame(Connection connection, Scanner scanner) throws SQLException {
        System.out.print("Enter game ID to delete: ");
        int id = scanner.nextInt();

        String sql = "DELETE FROM Games WHERE id = ?";
        try (PreparedStatement statement = connection.prepareStatement(sql)) {
            statement.setInt(1, id);

            int rowsDeleted = statement.executeUpdate();
            if (rowsDeleted > 0) {
                System.out.println("Game deleted successfully!");
            } else {
                System.out.println("Game ID not found.");
            }
        }
    }
}
 {
    
}
