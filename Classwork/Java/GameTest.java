public class GameTest {

   public static void main(String[] args) {
      Games gameOne = new Games("Mystic Quest");

      gameOne.setGenre("RPG");
      gameOne.setReleaseYear(2020);
      gameOne.setRating(4.7);
      gameOne.printGameDetails();
      
      System.out.println();

      Games gameTwo = new Games("Space Adventures");

      gameTwo.setGenre("Action-Adventure");
      gameTwo.setReleaseYear(2022);
      gameTwo.setRating(4.9);
      gameTwo.printGameDetails();
   }
}
