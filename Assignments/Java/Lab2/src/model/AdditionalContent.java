package model;

public class AdditionalContent extends Item {
    private String edition;

    public AdditionalContent(String name, String edition, double basePrice) {
        super(name, "CONTENT", basePrice);
        this.edition = edition;
    }

    public double getEditionPrice() {
        return switch (edition.toUpperCase()) {
            case "PREMIUM" -> getBasePrice() * 1.5;
            case "ULTIMATE" -> getBasePrice() * 2.0;
            default -> getBasePrice();
        };
    }

    public void upgradeEdition(String newEdition) {
        this.edition = newEdition;
    }
}
