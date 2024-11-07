class Game {
    String title;
    String genre;
    int players;
    double rating;
    int hoursPlayed;

    public Game() {
        this.title = "Unknown Game";
        this.genre = "Unknown Genre";
        this.players = 0;
        this.rating = 0.0;
        this.hoursPlayed = 0;
    }

    public Game(String title, String genre, int players) {
        this.title = title;
        this.genre = genre;
        this.players = players;
        this.rating = 0.0;  // Default rating
        this.hoursPlayed = 0;  // Default hours played
    }

    public Game(String title, String genre, int players, double rating, int hoursPlayed) {
        this.title = title;
        this.genre = genre;
        this.players = players;
        this.rating = rating;
        this.hoursPlayed = hoursPlayed;
    }

    public int calculateTotalPlayTime() {
        return players * hoursPlayed;
    }

    public int calculateTotalPlayTime(int specificPlayers) {
        return specificPlayers * hoursPlayed;
    }

    public int calculateTotalPlayTime(int specificPlayers, double boostFactor) {
        return (int) (specificPlayers * hoursPlayed * boostFactor);
    }

    // 7. Method to display Game Information
    public void displayGameInfo() {
        System.out.println("Game Title: " + title);
        System.out.println("Genre: " + genre);
        System.out.println("Players: " + players);
        System.out.println("Rating: " + rating);
        System.out.println("Hours Played: " + hoursPlayed);
    }

    // 8. Main method to execute and display results
    public static void main(String[] args) {

        Game defaultGame = new Game();
        System.out.println("Default Game Info:");
        defaultGame.displayGameInfo();
		
        System.out.println("Total Play Time: " + defaultGame.calculateTotalPlayTime() + " hours\n");

        Game gameWithSomeAttributes = new Game("Fantasy World", "RPG", 200);
        System.out.println("Parameterized Game Info (3 attributes):");
        gameWithSomeAttributes.displayGameInfo();
        System.out.println("Total Play Time: " + gameWithSomeAttributes.calculateTotalPlayTime() + " hours\n");

        Game fullyDefinedGame = new Game("Warriors Quest", "Action", 300, 4.8, 50);
        System.out.println("Fully Defined Game Info (All attributes):");
        fullyDefinedGame.displayGameInfo();
        System.out.println("Total Play Time (All Players): " + fullyDefinedGame.calculateTotalPlayTime() + " hours");
        System.out.println("Total Play Time (200 specific players): " + fullyDefinedGame.calculateTotalPlayTime(200) + " hours");
        System.out.println("Total Play Time (200 specific players with 1.5x boost): " + fullyDefinedGame.calculateTotalPlayTime(200, 1.5) + " hours\n");
    }
}
