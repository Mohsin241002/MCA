// Base class for all purchasable items
public class Item {
    private String name;
    private String type;
    private double basePrice;
    
    public Item(String name, String type, double basePrice) {
        this.name = name;
        this.type = type;
        this.basePrice = basePrice;
    }
    
    public double getBasePrice() {
        return basePrice;
    }
    
    public String getName() {
        return name;
    }
}

// Class for gaming accessories (like Side items in original)
public class GameAccessory extends Item {
    public GameAccessory(String name, double price) {
        super(name, "ACCESSORY", price);
    }
}

// Class for additional content/services (like Drink in original)
public class AdditionalContent extends Item {
    private String edition; // STANDARD, PREMIUM, ULTIMATE (like drink sizes)
    
    public AdditionalContent(String name, String edition, double basePrice) {
        super(name, "CONTENT", basePrice);
        this.edition = edition;
    }
    
    public double getEditionPrice() {
        return switch(edition.toUpperCase()) {
            case "PREMIUM" -> getBasePrice() * 1.5;
            case "ULTIMATE" -> getBasePrice() * 2.0;
            default -> getBasePrice();
        };
    }
    
    public void upgradeEdition(String newEdition) {
        this.edition = newEdition;
    }
}

// Base class for games (like Burger in original)
public class Game extends Item {
    private String[] addOns;
    private double[] addOnPrices;
    private int maxAddOns;
    private int currentAddOns;
    
    public Game(String name, double price) {
        super(name, "GAME", price);
        this.maxAddOns = 3;
        this.addOns = new String[maxAddOns];
        this.addOnPrices = new double[maxAddOns];
        this.currentAddOns = 0;
    }
    
    public void addGameAddOn(String name, double price) {
        if (currentAddOns < maxAddOns) {
            addOns[currentAddOns] = name;
            addOnPrices[currentAddOns] = price;
            currentAddOns++;
        } else {
            System.out.println("Cannot add more add-ons to " + getName());
        }
    }
    
    public double getTotalPrice() {
        double total = getBasePrice();
        for (int i = 0; i < currentAddOns; i++) {
            total += addOnPrices[i];
        }
        return total;
    }
    
    public void printAddOns() {
        for (int i = 0; i < currentAddOns; i++) {
            System.out.println("Added " + addOns[i] + " for an extra " + addOnPrices[i]);
        }
    }
}

// Deluxe edition game with more add-on slots and fixed price
public class DeluxeGame extends Game {
    public DeluxeGame() {
        super("Deluxe Game Bundle", 99.99);
        this.maxAddOns = 5;
    }
    
    @Override
    public double getTotalPrice() {
        return getBasePrice(); // Price is fixed regardless of add-ons
    }
}

// Main order management class
public class GamePurchaseOrder {
    private Game game;
    private AdditionalContent content;
    private GameAccessory accessory;
    
    // Default constructor
    public GamePurchaseOrder() {
        this.game = new Game("Standard Game", 59.99);
        this.content = new AdditionalContent("Season Pass", "STANDARD", 29.99);
        this.accessory = new GameAccessory("Standard Controller", 49.99);
    }
    
    // Custom order constructor
    public GamePurchaseOrder(Game game, AdditionalContent content, GameAccessory accessory) {
        this.game = game;
        this.content = content;
        this.accessory = accessory;
    }
    
    public void addGameAddOn(String name, double price) {
        game.addGameAddOn(name, price);
    }
    
    public void upgradeContent(String edition) {
        content.upgradeEdition(edition);
    }
    
    public void printItemizedList() {
        System.out.println("\nORDER DETAILS");
        System.out.println("Base Game: " + game.getName() + 
                         " (€" + game.getBasePrice() + ")");
        game.printAddOns();
        System.out.println("Additional Content: " + content.getName() + 
                         " (€" + content.getEditionPrice() + ")");
        System.out.println("Gaming Accessory: " + accessory.getName() + 
                         " (€" + accessory.getBasePrice() + ")");
    }
    
    public double getTotalPrice() {
        return game.getTotalPrice() + content.getEditionPrice() + accessory.getBasePrice();
    }
}

// Example usage
public class Main {
    public static void main(String[] args) {
        // Default order
        GamePurchaseOrder defaultOrder = new GamePurchaseOrder();
        defaultOrder.addGameAddOn("Character Skin", 4.99);
        defaultOrder.addGameAddOn("Weapon Pack", 9.99);
        defaultOrder.upgradeContent("PREMIUM");
        defaultOrder.printItemizedList();
        System.out.println("Total Price: €" + defaultOrder.getTotalPrice());
        
        // Custom order with Deluxe Game
        DeluxeGame deluxeGame = new DeluxeGame();
        AdditionalContent premiumContent = new AdditionalContent("Premium Pass", "ULTIMATE", 39.99);
        GameAccessory proController = new GameAccessory("Pro Controller", 69.99);
        
        GamePurchaseOrder deluxeOrder = new GamePurchaseOrder(deluxeGame, premiumContent, proController);
        deluxeOrder.addGameAddOn("Extra Maps", 14.99);
        deluxeOrder.addGameAddOn("Character Pack", 19.99);
        deluxeOrder.addGameAddOn("Soundtrack", 9.99);
        deluxeOrder.addGameAddOn("Art Book", 24.99);
        deluxeOrder.printItemizedList();
        System.out.println("Total Price: €" + deluxeOrder.getTotalPrice());
    }
}