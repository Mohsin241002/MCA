
public class Game {
    private String name;
    private String genre;
    private int maxPlayers;
    private double price;

    public Game(String name, String genre, int maxPlayers, double price) {
        this.name = name;
        this.genre = genre;
        this.maxPlayers = maxPlayers;
        this.price = price;
    }

    public void displayDetails() {
        System.out.println("Game Name: " + name);
        System.out.println("Genre: " + genre);
        System.out.println("Max Players: " + maxPlayers);
        System.out.println("Price: Rs" + price);
    }

    public static void main(String[] args) {
        Game myGame = new Game("Battle Arena", "Action", 4, 2999.99);
        myGame.displayDetails();
    }
}
