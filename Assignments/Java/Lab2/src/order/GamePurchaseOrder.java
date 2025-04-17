package order;

import model.*;

public class GamePurchaseOrder {
    private Game game;
    private AdditionalContent content;
    private GameAccessory accessory;

    // Default constructor for standard order
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
                           " (" + game.getBasePrice() + ")");
        game.printAddOns();
        System.out.println("Additional Content: " + content.getName() +
                           " (" + content.getEditionPrice() + ")");
        System.out.println("Gaming Accessory: " + accessory.getName() +
                           " (" + accessory.getBasePrice() + ")");
    }

    public double getTotalPrice() {
        return game.getTotalPrice() + content.getEditionPrice() + accessory.getBasePrice();
    }
}
