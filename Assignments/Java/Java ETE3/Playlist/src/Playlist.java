import java.util.LinkedList;

public class Playlist {
    public static void main(String[] args) {
        Album album = new Album("Greatest Hits", "Artist Name");

        // Adding songs to the album
        album.addSong("Song 1", 4.05);
        album.addSong("Song 2", 3.45);
        album.addSong("Song 3", 5.20);

        // Playlist
        LinkedList<Song> playList = new LinkedList<>();

        // Adding songs to the playlist by track number
        album.addToPlayList(1, playList);
        album.addToPlayList(2, playList);

        // Adding songs to the playlist by title
        album.addToPlayList("Song 3", playList);

        // Attempt to add a song not in the album
        if (!album.addToPlayList("Nonexistent Song", playList)) {
            System.out.println("Song not found in the album.");
        }

        // Display playlist
        System.out.println("Playlist:");
        for (Song song : playList) {
            System.out.println(song);
        }
    }
}
