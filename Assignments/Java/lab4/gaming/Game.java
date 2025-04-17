package gaming;

public class Game implements GameData {
    private String gameName;
    private String genre;
    private int totalPlayers;

    @Override
    public void fetchGameDetails(String gameName) {
        this.gameName = gameName;
        this.genre = "Action";
        this.totalPlayers = 1000; // Example data
    }

    @Override
    public void displayGameDetails() {
        System.out.println("Game Name: " + gameName);
        System.out.println("Genre: " + genre);
        System.out.println("Total Players: " + totalPlayers);
    }

    @Override
    public int calculatePlayerStats(int kills, int assists, int deaths) {
        return (kills * 10) + (assists * 5) - (deaths * 3); // Example formula
    }
}
