public class Games {

   String title;
   String genre;
   int releaseYear;
   double rating;

   // Constructor for the Games class
   public Games(String title) {
      this.title = title;
   }

   // Assign the genre of the Games to the variable genre
   public void setGenre(String GamesGenre) {
      genre = GamesGenre;
   }

   // Assign the release year to the variable releaseYear
   public void setReleaseYear(int year) {
      releaseYear = year;
   }

   // Assign the rating to the variable rating
   public void setRating(double GamesRating) {
      rating = GamesRating;
   }

   // Print the Games details
   public void printGamesDetails() {
      System.out.println("Title: " + title);
      System.out.println("Genre: " + genre);
      System.out.println("Release Year: " + releaseYear);
      System.out.println("Rating: " + rating);
   }

   // Main method to test the class
   public static void main(String[] args) {
      // Creating a new Games object
      Games myGames = new Games("Legend of Adventure");

      // Setting attributes
      myGames.setGenre("Action-Adventure");
      myGames.setReleaseYear(2023);
      myGames.setRating(4.5);

      // Printing Games details
      myGames.printGamesDetails();
   }
}
