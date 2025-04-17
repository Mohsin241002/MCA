import java.util.*;

// Class to represent a Game
class Game {
    private int gameId;
    private String name;
    private String genre;
    private double rating;

    public Game(int gameId, String name, String genre, double rating) {
        this.gameId = gameId;
        this.name = name;
        this.genre = genre;
        this.rating = rating;
    }

    public int getGameId() {
        return gameId;
    }

    public String getName() {
        return name;
    }

    public String getGenre() {
        return genre;
    }

    public double getRating() {
        return rating;
    }

    @Override
    public String toString() {
        return "Game{" +
                "gameId=" + gameId +
                ", name='" + name + '\'' +
                ", genre='" + genre + '\'' +
                ", rating=" + rating +
                '}';
    }
}

// Class to represent a Player
class Player {
    private int playerId;
    private String name;
    private List<Game> gamesPlayed;

    public Player(int playerId, String name) {
        this.playerId = playerId;
        this.name = name;
        this.gamesPlayed = new ArrayList<>();
    }

    public void addGame(Game game) {
        gamesPlayed.add(game);
    }

    public List<Game> getGamesPlayed() {
        return gamesPlayed;
    }

    @Override
    public String toString() {
        return "Player{" +
                "playerId=" + playerId +
                ", name='" + name + '\'' +
                ", gamesPlayed=" + gamesPlayed +
                '}';
    }
}

public class GamingInformationSystem {

    public static void main(String[] args) {
        // List to manage Players
        List<Player> players = new ArrayList<>();

        // Set to manage unique Genres
        Set<String> genres = new HashSet<>();

        // Map to manage Games with quick retrieval by ID
        Map<Integer, Game> games = new HashMap<>();

        // Queue to manage matchmaking
        Queue<Player> matchmakingQueue = new LinkedList<>();

        // Deque to manage leaderboard
        Deque<Player> leaderboard = new ArrayDeque<>();

        // Adding some Games
        Game game1 = new Game(1, "Halo", "FPS", 4.5);
        Game game2 = new Game(2, "Minecraft", "Sandbox", 4.8);
        Game game3 = new Game(3, "Fortnite", "Battle Royale", 4.3);

        games.put(game1.getGameId(), game1);
        games.put(game2.getGameId(), game2);
        games.put(game3.getGameId(), game3);

        // Add unique genres
        genres.add(game1.getGenre());
        genres.add(game2.getGenre());
        genres.add(game3.getGenre());

        // Adding Players
        Player player1 = new Player(1, "Alice");
        Player player2 = new Player(2, "Bob");
        Player player3 = new Player(3, "Charlie");

        players.add(player1);
        players.add(player2);
        players.add(player3);

        // Players playing games
        player1.addGame(game1);
        player1.addGame(game2);

        player2.addGame(game2);

        player3.addGame(game3);

        // Matchmaking
        matchmakingQueue.add(player1);
        matchmakingQueue.add(player2);
        matchmakingQueue.add(player3);

        // Adding to leaderboard
        leaderboard.addFirst(player1); // Top player
        leaderboard.addLast(player3); // Lower-ranked player

        // Displaying data
        System.out.println("All Games:");
        for (Game game : games.values()) {
            System.out.println(game);
        }

        System.out.println("\nAll Players:");
        for (Player player : players) {
            System.out.println(player);
        }

        System.out.println("\nUnique Genres:");
        for (String genre : genres) {
            System.out.println(genre);
        }

        System.out.println("\nMatchmaking Queue:");
        while (!matchmakingQueue.isEmpty()) {
            System.out.println(matchmakingQueue.poll());
        }

        System.out.println("\nLeaderboard:");
        for (Player player : leaderboard) {
            System.out.println(player);
        }
    }
}
