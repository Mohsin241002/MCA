class Player {
    String name;
    int level;
    int health;

    // Default Constructor
    public Player() {
        this.name = "Unknown";
        this.level = 1;
        this.health = 100;
    }

    // Parameterized Constructor with one parameter
    public Player(String name) {
        this.name = name;
        this.level = 1;
        this.health = 100;
    }

    // Parameterized Constructor with two parameters
    public Player(String name, int level) {
        this.name = name;
        this.level = level;
        this.health = 100;
    }

    // Parameterized Constructor with three parameters
    public Player(String name, int level, int health) {
        this.name = name;
        this.level = level;
        this.health = health;
    }

    // Copy Constructor
    public Player(Player otherPlayer) {
        this.name = otherPlayer.name;
        this.level = otherPlayer.level;
        this.health = otherPlayer.health;
    }

    // Display method to show basic player details
    public void displayBasicInfo() {
        System.out.println("Player Name: " + this.name);
    }

    // Display method to show full player details
    public void displayFullInfo() {
        System.out.println("Player Name: " + this.name + ", Level: " + this.level + ", Health: " + this.health);
    }

    public static void main(String[] args) {
        // Using Default Constructor
        Player player1 = new Player();
        player1.displayBasicInfo();
        player1.displayFullInfo();

        // Using Parameterized Constructor with one parameter
        Player player2 = new Player("Knight");
        player2.displayBasicInfo();
        player2.displayFullInfo();

        // Using Parameterized Constructor with two parameters
        Player player3 = new Player("Archer", 3);
        player3.displayBasicInfo();
        player3.displayFullInfo();

        // Using Parameterized Constructor with three parameters
        Player player4 = new Player("Mage", 5, 120);
        player4.displayBasicInfo();
        player4.displayFullInfo();

        // Using Copy Constructor
        Player player5 = new Player(player4);
        player5.displayBasicInfo();
        player5.displayFullInfo();
    }
}
