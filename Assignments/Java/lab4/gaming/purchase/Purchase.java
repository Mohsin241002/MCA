package gaming.purchase;

public class Purchase implements PurchaseProcessing {
    @Override
    public double calculateTotalCost(double[] prices) {
        double total = 0;
        for (double price : prices) {
            total += price;
        }
        return total;
    }

    @Override
    public boolean processPurchase(String playerName, double amount) {
        if (amount > 0) {
            System.out.println(playerName + " purchased items worth $" + amount);
            return true;
        }
        return false;
    }

    @Override
    public void displayPurchaseDetails() {
        System.out.println("Purchase details are displayed here.");
    }
}
