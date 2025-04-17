package gaming;

import gaming.stats.Player;
import gaming.purchase.Purchase;
import gaming.purchase.inapp.InAppPurchase;

public class Main {
    public static void main(String[] args) {
        // Create a Game object
        Game game = new Game();
        game.fetchGameDetails("Battle Royale");
        game.displayGameDetails();

        // Create a Player object
        Player player = new Player("Rohan");
        player.fetchGameDetails("Battle Royale");
        player.displayGameDetails();
        System.out.println("Player Stats: " + player.calculatePlayerStats(10, 5, 3));
        System.out.println("Player Score: " + player.calculatePlayerScore(5, 20));

        // Calculate purchases
        Purchase purchase = new Purchase();
        double[] prices = {9.99, 4.99, 14.99};
        System.out.println("Total Cost: Rs" + purchase.calculateTotalCost(prices));
        purchase.processPurchase("Rohan", purchase.calculateTotalCost(prices));

        // In-app purchase
        InAppPurchase inAppPurchase = new InAppPurchase();
        System.out.println("In-App Total Cost: Rs" + inAppPurchase.calculateTotalCost(prices));
        inAppPurchase.processPurchase("Rohan", inAppPurchase.calculateTotalCost(prices));
    }
}

