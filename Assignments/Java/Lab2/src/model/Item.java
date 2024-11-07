package model;

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

    public String getType() {
        return type;
    }
}
