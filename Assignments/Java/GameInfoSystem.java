import java.util.*;

// Define a generic interface for rankable entities
interface Rankable {
    String getName(); // Player name
}

// Player class implementing Rankable
class Player implements Rankable {
    private final String name;
    private final int score;
    private final int level;
    private final double playtime;

    public Player(String name, int score, int level, double playtime) {
        this.name = name;
        this.score = score;
        this.level = level;
        this.playtime = playtime;
    }

    public String getName() {
        return name;
    }

    public int getScore() {
        return score;
    }

    public int getLevel() {
        return level;
    }

    public double getPlaytime() {
        return playtime;
    }

    @Override
    public String toString() {
        return String.format("Player{name='%s', score=%d, level=%d, playtime=%.2f}", 
                             name, score, level, playtime);
    }
}

// Generic Leaderboard class
class Leaderboard<T extends Rankable> {
    private final List<T> players;

    public Leaderboard() {
        players = new ArrayList<>();
    }

    public void addPlayer(T player) {
        players.add(player);
    }

    public void rankPlayers(Comparator<T> comparator) {
        players.sort(comparator);
    }

    public void display() {
        System.out.println("=== Leaderboard ===");
        for (T player : players) {
            System.out.println(player);
        }
    }
}

// Main class
public class GameInfoSystem {
    public static void main(String[] args) {
        // Create a leaderboard for players
        Leaderboard<Player> leaderboard = new Leaderboard<>();

        // Add players
        leaderboard.addPlayer(new Player("Rohan", 1500, 20, 50.5));
        leaderboard.addPlayer(new Player("Anirban", 2000, 18, 75.0));
        leaderboard.addPlayer(new Player("Mohsin", 1800, 25, 40.0));

        // Rank by score
        System.out.println("Ranking by Score:");
        leaderboard.rankPlayers(Comparator.comparingInt(Player::getScore).reversed());
        leaderboard.display();

        // Rank by level
        System.out.println("\nRanking by Level:");
        leaderboard.rankPlayers(Comparator.comparingInt(Player::getLevel).reversed());
        leaderboard.display();

        // Rank by playtime
        System.out.println("\nRanking by Playtime:");
        leaderboard.rankPlayers(Comparator.comparingDouble(Player::getPlaytime).reversed());
        leaderboard.display();
    }
}
