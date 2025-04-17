package gaming.purchase;

public interface PurchaseProcessing {
    double calculateTotalCost(double[] prices);
    boolean processPurchase(String playerName, double amount);
    void displayPurchaseDetails();
}
