import java.util.LinkedList;
import java.util.Queue;

class Game {
    private final int gameId;
    private final String gameName;

    public Game(int gameId, String gameName) {
        this.gameId = gameId;
        this.gameName = gameName;
    }

    public int getGameId() {
        return gameId;
    }

    public String getGameName() {
        return gameName;
    }

    @Override
    public String toString() {
        return "Game{" + "ID=" + gameId + ", Name='" + gameName + '\'' + '}';
    }
}

class GameManager {
    private final Queue<Game> gameQueue = new LinkedList<>();
    private final Object lock = new Object();
    private boolean synchronizedMode;

    public GameManager(boolean synchronizedMode) {
        this.synchronizedMode = synchronizedMode;
    }

    public void addGame(Game game) {
        if (synchronizedMode) {
            synchronized (lock) {
                gameQueue.offer(game);
                System.out.println("Game added: " + game);
                lock.notifyAll(); // Notify waiting players
            }
        } else {
            gameQueue.offer(game);
            System.out.println("Game added (Unsynchronized): " + game);
        }
    }

    public Game playGame() {
        if (synchronizedMode) {
            synchronized (lock) {
                while (gameQueue.isEmpty()) {
                    try {
                        System.out.println(Thread.currentThread().getName() + " is waiting for a game...");
                        lock.wait();
                    } catch (InterruptedException e) {
                        Thread.currentThread().interrupt();
                        System.out.println("Thread interrupted.");
                    }
                }
                return gameQueue.poll();
            }
        } else {
            if (gameQueue.isEmpty()) {
                System.out.println(Thread.currentThread().getName() + " found no game (Unsynchronized).");
                return null;
            }
            return gameQueue.poll();
        }
    }
}

class Player implements Runnable {
    private final String playerName;
    private final GameManager gameManager;

    public Player(String playerName, GameManager gameManager) {
        this.playerName = playerName;
        this.gameManager = gameManager;
    }

    @Override
    public void run() {
        try {
            Thread.sleep((long) (Math.random() * 3000)); // Simulate random delay
            System.out.println(playerName + " is trying to play a game...");
            Game game = gameManager.playGame();
            if (game != null) {
                System.out.println(playerName + " started playing: " + game);
            } else {
                System.out.println(playerName + " couldn't find a game.");
            }
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            System.out.println(playerName + " was interrupted.");
        }
    }
}

public class GamingSystem {
    public static void main(String[] args) {
        boolean synchronizedMode = true; // Toggle between synchronized and non-synchronized mode
        GameManager gameManager = new GameManager(synchronizedMode);

        // Thread 1: Game Manager adding games
        Thread gameAdder = new Thread(() -> {
            for (int i = 1; i <= 5; i++) {
                try {
                    Thread.sleep(1000); // Simulate delay in adding games
                } catch (InterruptedException e) {
                    Thread.currentThread().interrupt();
                }
                gameManager.addGame(new Game(i, "Game_" + i));
            }
        }, "GameAdder");

        // Threads 2-6: Players trying to play games
        Thread player1 = new Thread(new Player("Player1", gameManager), "Player1");
        Thread player2 = new Thread(new Player("Player2", gameManager), "Player2");
        Thread player3 = new Thread(new Player("Player3", gameManager), "Player3");
        Thread player4 = new Thread(new Player("Player4", gameManager), "Player4");
        Thread player5 = new Thread(new Player("Player5", gameManager), "Player5");

        // Start all threads
        gameAdder.start();
        player1.start();
        player2.start();
        player3.start();
        player4.start();
        player5.start();

        // Wait for all threads to finish
        try {
            gameAdder.join();
            player1.join();
            player2.join();
            player3.join();
            player4.join();
            player5.join();
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }

        System.out.println("Gaming System simulation completed.");
    }
}
