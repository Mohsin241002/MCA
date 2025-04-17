import java.util.function.*;

class Game {
    private String title;
    private String genre;
    private int playtime; // in hours

    public Game(String title, String genre, int playtime) {
        this.title = title;
        this.genre = genre;
        this.playtime = playtime;
    }

    public String getTitle() {
        return title;
    }

    public String getGenre() {
        return genre;
    }

    public int getPlaytime() {
        return playtime;
    }

    @Override
    public String toString() {
        return "Game{" +
                "title='" + title + '\'' +
                ", genre='" + genre + '\'' +
                ", playtime=" + playtime +
                '}';
    }
}

public class lab7 {
    public static void main(String[] args) {
        // Sample Game Objects
        Game game1 = new Game("Elder Quest", "RPG", 120);
        Game game2 = new Game("Speed Racer", "Racing", 15);

        // Consumer: Print game details
        Consumer<Game> printGameDetails = game -> System.out.println(game);
        System.out.println("--- Game Details ---");
        printGameDetails.accept(game1);
        printGameDetails.accept(game2);

        // Function: Map playtime to difficulty level
        Function<Integer, String> playtimeToDifficulty = playtime -> {
            if (playtime > 100) return "Hard";
            else if (playtime > 50) return "Medium";
            else return "Easy";
        };
        System.out.println("\n--- Difficulty Level ---");
        System.out.println(game1.getTitle() + " is " + playtimeToDifficulty.apply(game1.getPlaytime()));
        System.out.println(game2.getTitle() + " is " + playtimeToDifficulty.apply(game2.getPlaytime()));

        // Predicate: Check if the game belongs to a specific genre
        Predicate<Game> isRacingGame = game -> game.getGenre().equalsIgnoreCase("Racing");
        System.out.println("\n--- Genre Check ---");
        System.out.println(game1.getTitle() + " is a Racing Game? " + isRacingGame.test(game1));
        System.out.println(game2.getTitle() + " is a Racing Game? " + isRacingGame.test(game2));

        // Supplier: Provide default game data
        Supplier<Game> defaultGameSupplier = () -> new Game("Default Game", "Adventure", 50);
        System.out.println("\n--- Default Game ---");
        Game defaultGame = defaultGameSupplier.get();
        System.out.println(defaultGame);
    }
}
