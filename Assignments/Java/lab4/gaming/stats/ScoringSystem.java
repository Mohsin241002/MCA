package gaming.stats;

import gaming.GameData;

public interface ScoringSystem extends GameData {
    int calculatePlayerScore(int level, int achievements); // Calculation method
    int calculateHighScore(int[] scores);
}
