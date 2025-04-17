class Game {
    String title = "Untitled Game";
    String genre = "Unknown Genre";
}

class Player {
    String playerName = "Unnamed Player";
    int playerScore = 0;
}

class Developer {
    String devName = "Unknown Developer";
    String company = "Unknown Company";
}

public class GameDetails {
    public static void main(String[] args) {
        Game game1 = new Game();
        Player player1 = new Player();
        Developer dev1 = new Developer();

        game1.title = "Cyber Quest";
        game1.genre = "Action-Adventure";

        player1.playerName = "Mohsin24";
        player1.playerScore = 1200;

        dev1.devName = "Ezio";
        dev1.company = "NextGen Studios";

        System.out.println("Game Details:");
        System.out.println("Title: " + game1.title);
        System.out.println("Genre: " + game1.genre);

        System.out.println("\nPlayer Details:");
        System.out.println("Player Name: " + player1.playerName);
        System.out.println("Score: " + player1.playerScore);

        System.out.println("\nDeveloper Details:");
        System.out.println("Developer Name: " + dev1.devName);
        System.out.println("Company: " + dev1.company);
    }
}
