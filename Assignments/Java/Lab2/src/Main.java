import model.*;
import order.GamePurchaseOrder;

import java.util.Scanner;

public class Main {
    private static Scanner scanner = new Scanner(System.in);
    private static GamePurchaseOrder currentOrder = null;

    public static void main(String[] args) {
        boolean exit = false;
        while (!exit) {
            displayMainMenu();
            int choice = getIntInput("Enter your choice: ");

            switch (choice) {
                case 1 -> createNewOrder();
                case 2 -> createDeluxeOrder();
                case 3 -> addGameAddOn();
                case 4 -> upgradeContent();
                case 5 -> viewOrder();
                case 6 -> exit = true;
                default -> System.out.println("\nInvalid choice! Please try again.");
            }
        }
        scanner.close();
    }

    private static void displayMainMenu() {
        System.out.println("\n=== Gaming Store Menu ===");
        System.out.println("1. Create New Standard Order");
        System.out.println("2. Create New Deluxe Order");
        System.out.println("3. Add Game Add-on to Current Order");
        System.out.println("4. Upgrade Content Edition");
        System.out.println("5. View Current Order");
        System.out.println("6. Exit");
    }

    private static void createNewOrder() {
        System.out.println("\n=== Creating New Standard Order ===");
        currentOrder = new GamePurchaseOrder();
        viewOrder();
    }

    private static void createDeluxeOrder() {
        System.out.println("\n=== Creating New Deluxe Order ===");
        DeluxeGame deluxeGame = new DeluxeGame();

        System.out.println("\nSelect Additional Content Edition:");
        System.out.println("1. Standard (29.99)");
        System.out.println("2. Premium (44.99)");
        System.out.println("3. Ultimate (59.99)");
        int editionChoice = getIntInput("Enter your choice: ");
        String edition = switch (editionChoice) {
            case 2 -> "PREMIUM";
            case 3 -> "ULTIMATE";
            default -> "STANDARD";
        };

        AdditionalContent content = new AdditionalContent("Premium Pass", edition, 29.99);

        System.out.println("\nSelect Gaming Accessory:");
        System.out.println("1. Standard Controller (49.99)");
        System.out.println("2. Pro Controller (69.99)");
        System.out.println("3. Elite Controller (89.99)");
        int accessoryChoice = getIntInput("Enter your choice: ");
        String accessoryName = switch (accessoryChoice) {
            case 2 -> "Pro Controller";
            case 3 -> "Elite Controller";
            default -> "Standard Controller";
        };
        double accessoryPrice = switch (accessoryChoice) {
            case 2 -> 69.99;
            case 3 -> 89.99;
            default -> 49.99;
        };

        GameAccessory accessory = new GameAccessory(accessoryName, accessoryPrice);
        currentOrder = new GamePurchaseOrder(deluxeGame, content, accessory);
        viewOrder();
    }

    private static void addGameAddOn() {
        System.out.println("\n=== Add Game Add-on ===");
        System.out.println("Available Add-ons:");
        System.out.println("1. Character Skin (4.99)");
        System.out.println("2. Weapon Pack (9.99)");
        System.out.println("3. Extra Maps (14.99)");
        System.out.println("4. Custom Add-on");
        int choice = getIntInput("Enter your choice: ");

        String name;
        double price;
        switch (choice) {
            case 1 -> {
                name = "Character Skin";
                price = 4.99;
            }
            case 2 -> {
                name = "Weapon Pack";
                price = 9.99;
            }
            case 3 -> {
                name = "Extra Maps";
                price = 14.99;
            }
            case 4 -> {
                System.out.print("Enter add-on name: ");
                name = scanner.nextLine();
                price = getDoubleInput("Enter add-on price: ");
            }
            default -> {
                System.out.println("Invalid choice!");
                return;
            }
        }

        currentOrder.addGameAddOn(name, price);
        System.out.println("Add-on added successfully!");
    }

    private static void upgradeContent() {
        System.out.println("\n=== Upgrade Content Edition ===");
        System.out.println("1. Standard Edition");
        System.out.println("2. Premium Edition");
        System.out.println("3. Ultimate Edition");
        int choice = getIntInput("Enter your choice: ");

        String edition = switch (choice) {
            case 2 -> "PREMIUM";
            case 3 -> "ULTIMATE";
            default -> "STANDARD";
        };

        currentOrder.upgradeContent(edition);
        System.out.println("Content edition upgraded successfully!");
    }

    private static void viewOrder() {
        if (currentOrder != null) {
            currentOrder.printItemizedList();
            System.out.printf("\nTotal Price: %.2f%n", currentOrder.getTotalPrice());
        } else {
            System.out.println("\nNo active order to display!");
        }
    }

    // Utility methods for input
    private static int getIntInput(String prompt) {
        while (true) {
            try {
                System.out.print(prompt);
                return Integer.parseInt(scanner.nextLine());
            } catch (NumberFormatException e) {
                System.out.println("Please enter a valid number!");
            }
        }
    }

    private static double getDoubleInput(String prompt) {
        while (true) {
            try {
                System.out.print(prompt);
                return Double.parseDouble(scanner.nextLine());
            } catch (NumberFormatException e) {
                System.out.println("Please enter a valid number!");
            }
        }
    }
}
