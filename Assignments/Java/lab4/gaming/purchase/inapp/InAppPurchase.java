package gaming.purchase.inapp;

import gaming.purchase.PurchaseProcessing;

public class InAppPurchase implements PurchaseProcessing {
    @Override
    public double calculateTotalCost(double[] prices) {
        double total = 0;
        for (double price : prices) {
            total += price * 1.05; // 5% tax for in-app purchases
        }
        return total;
    }

    @Override
    public boolean processPurchase(String playerName, double amount) {
        if (amount > 0) {
            System.out.println(playerName + " made an in-app purchase worth $" + amount);
            return true;
        }
        return false;
    }

    @Override
    public void displayPurchaseDetails() {
        System.out.println("In-app purchase details are displayed here.");
    }
}
