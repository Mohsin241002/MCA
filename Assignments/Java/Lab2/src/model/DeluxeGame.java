package model;

// DeluxeGame overrides behavior of Game
public class DeluxeGame extends Game {

    public DeluxeGame() {
        super("Deluxe Game Bundle", 149.99);  // Fixed price deluxe game
        this.maxAddOns = 5;  // Deluxe allows 5 add-ons
    }

    // Polymorphism - Override method to provide specific implementation
    @Override
    public double getTotalPrice() {
        return getBasePrice();  // Fixed price regardless of add-ons
    }
}
