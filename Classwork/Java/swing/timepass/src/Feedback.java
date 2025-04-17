public class Feedback {
    private String userName;
    private String movieName;
    private int rating;
    private String comments;

    public Feedback(String userName, String movieName, int rating, String comments) {
        this.userName = userName;
        this.movieName = movieName;
        this.rating = rating;
        this.comments = comments;
    }

    // Getters and toString method
    public String getUserName() {
        return userName;
    }

    public String getMovieName() {
        return movieName;
    }

    public int getRating() {
        return rating;
    }

    public String getComments() {
        return comments;
    }

    @Override
    public String toString() {
        return "User: " + userName + ", Movie: " + movieName +
                ", Rating: " + rating + ", Comments: " + comments;
    }
}
