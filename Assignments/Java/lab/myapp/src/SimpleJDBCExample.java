import java.sql.*;

public class SimpleJDBCExample {
    public static void main(String[] args) {
        // Database URL, username, and password
        String jdbcURL = "jdbc:mysql://localhost:3306/testdb";
        String username = "root";
        String password = "admin";

        // Establish a connection
        try (Connection connection = DriverManager.getConnection(jdbcURL, username, password)) {
            System.out.println("Connected to the database!");

            // Create a new record
            String insertSQL = "INSERT INTO users (name, email) VALUES (?, ?)";
            try (PreparedStatement insertStmt = connection.prepareStatement(insertSQL)) {
                insertStmt.setString(1, "John Doe");
                insertStmt.setString(2, "johndoe@example.com");
                insertStmt.executeUpdate();
                System.out.println("Record inserted successfully!");
            }

            // Read records
            String selectSQL = "SELECT * FROM users";
            try (Statement selectStmt = connection.createStatement();
                 ResultSet resultSet = selectStmt.executeQuery(selectSQL)) {
                System.out.println("User data:");
                while (resultSet.next()) {
                    System.out.println(resultSet.getInt("id") + ", " +
                                       resultSet.getString("name") + ", " +
                                       resultSet.getString("email"));
                }
            }

            // Update a record
            String updateSQL = "UPDATE users SET email = ? WHERE name = ?";
            try (PreparedStatement updateStmt = connection.prepareStatement(updateSQL)) {
                updateStmt.setString(1, "newemail@example.com");
                updateStmt.setString(2, "John Doe");
                int rowsUpdated = updateStmt.executeUpdate();
                System.out.println(rowsUpdated + " record(s) updated.");
            }

            // Delete a record
            String deleteSQL = "DELETE FROM users WHERE name = ?";
            try (PreparedStatement deleteStmt = connection.prepareStatement(deleteSQL)) {
                deleteStmt.setString(1, "John Doe");
                int rowsDeleted = deleteStmt.executeUpdate();
                System.out.println(rowsDeleted + " record(s) deleted.");
            }
        } catch (SQLException e) {
            System.out.println("Error connecting to the database:");
            e.printStackTrace();
        }
    }
}
