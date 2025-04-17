package gaming.stats;

// import gaming.GameData;

public class Player implements ScoringSystem {
    private String playerName;

    public Player(String playerName) {
        this.playerName = playerName;
    }

    @Override
    public void fetchGameDetails(String gameName) {
        System.out.println(playerName + " is playing " + gameName);
    }

    @Override
    public void displayGameDetails() {
        System.out.println("Player Name: " + playerName);
    }

    @Override
    public int calculatePlayerStats(int kills, int assists, int deaths) {
        return (kills * 15) + (assists * 7) - (deaths * 2); // Player-specific formula
    }

    @Override
    public int calculatePlayerScore(int level, int achievements) {
        return (level * 50) + (achievements * 20);
    }

    @Override
    public int calculateHighScore(int[] scores) {
        int highScore = 0;
        for (int score : scores) {
            if (score > highScore) highScore = score;
        }
        return highScore;
    }
}
