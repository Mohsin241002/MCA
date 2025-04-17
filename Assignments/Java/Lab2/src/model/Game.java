package model;

public class Game extends Item {
    protected String[] addOns;
    protected double[] addOnPrices;
    protected int maxAddOns;
    private int currentAddOns;

    // Constructor for a game
    public Game(String name, double price) {
        super(name, "GAME", price);
        this.maxAddOns = 3;
        this.addOns = new String[maxAddOns];
        this.addOnPrices = new double[maxAddOns];
        this.currentAddOns = 0;
    }

    // Method Overloading: Allow different ways to add an add-on
    public void addGameAddOn(String name, double price) {
        if (currentAddOns < maxAddOns) {
            addOns[currentAddOns] = name;
            addOnPrices[currentAddOns] = price;
            currentAddOns++;
        } else {
            System.out.println("Cannot add more add-ons to " + getName());
        }
    }

    // Overloaded version for add-on with just a name, default price
    public void addGameAddOn(String name) {
        addGameAddOn(name, 5.0);  // Default price for unspecified add-ons
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
