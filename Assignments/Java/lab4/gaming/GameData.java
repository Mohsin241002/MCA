package gaming;

public interface GameData {
    void fetchGameDetails(String gameName);
    void displayGameDetails();
    int calculatePlayerStats(int kills, int assists, int deaths); // Calculation method
}
