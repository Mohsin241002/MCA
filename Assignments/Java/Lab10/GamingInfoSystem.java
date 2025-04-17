import javax.swing.*;
import java.awt.*;
import java.awt.event.*;

public class GamingInfoSystem extends JFrame {
    private JTabbedPane tabbedPane;
    private JComboBox<String> genreComboBox;
    private JTextArea infoTextArea;
    private JTextField searchField, gameNameField, gameGenreField;
    private JButton addButton, searchButton;

    public GamingInfoSystem() {
        // Frame initialization
        setTitle("Gaming Information System");
        setSize(600, 500);
        setLocationRelativeTo(null);
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        setLayout(new BorderLayout());

        // Tabbed Pane for different functionalities
        tabbedPane = new JTabbedPane();

        // Games by Genre Tab
        JPanel genrePanel = new JPanel(new BorderLayout());
        JPanel comboPanel = new JPanel();
        JLabel comboLabel = new JLabel("Select Genre:");
        genreComboBox = new JComboBox<>(new String[]{"Action", "RPG", "Strategy", "Sports"});
        genreComboBox.addActionListener(new ActionListener() {
            public void actionPerformed(ActionEvent e) {
                updateGameInfo(genreComboBox.getSelectedItem().toString());
            }
        });
        comboPanel.add(comboLabel);
        comboPanel.add(genreComboBox);
        infoTextArea = new JTextArea(10, 40);
        infoTextArea.setEditable(false);
        JScrollPane scrollPane = new JScrollPane(infoTextArea);
        scrollPane.setVerticalScrollBarPolicy(JScrollPane.VERTICAL_SCROLLBAR_ALWAYS);
        genrePanel.add(comboPanel, BorderLayout.NORTH);
        genrePanel.add(scrollPane, BorderLayout.CENTER);

        // Add New Game Tab
        JPanel addGamePanel = new JPanel(new GridLayout(3, 2));
        addGamePanel.add(new JLabel("Game Name:"));
        gameNameField = new JTextField();
        addGamePanel.add(gameNameField);
        addGamePanel.add(new JLabel("Genre:"));
        gameGenreField = new JTextField();
        addGamePanel.add(gameGenreField);
        addButton = new JButton("Add Game");
        addButton.addActionListener(new ActionListener() {
            public void actionPerformed(ActionEvent e) {
                addNewGame(gameNameField.getText(), gameGenreField.getText());
            }
        });
        addGamePanel.add(addButton);

        // Search Game Tab
        JPanel searchPanel = new JPanel(new BorderLayout());
        searchField = new JTextField();
        searchButton = new JButton("Search");
        searchButton.addActionListener(new ActionListener() {
            public void actionPerformed(ActionEvent e) {
                searchGame(searchField.getText());
            }
        });
        searchPanel.add(searchField, BorderLayout.CENTER);
        searchPanel.add(searchButton, BorderLayout.EAST);

        // Adding tabs to TabbedPane
        tabbedPane.addTab("Games by Genre", genrePanel);
        tabbedPane.addTab("Add New Game", addGamePanel);
        tabbedPane.addTab("Search Game", searchPanel);

        // Adding TabbedPane to the Frame
        add(tabbedPane, BorderLayout.CENTER);
    }

    private void updateGameInfo(String genre) {
        // Dummy data update method
        infoTextArea.setText("Displaying games for genre: " + genre);
    }

    private void addNewGame(String name, String genre) {
        // Dummy method to simulate adding a game
        JOptionPane.showMessageDialog(this, "Added new game: " + name + " under genre " + genre);
    }

    private void searchGame(String query) {
        // Dummy search method
        JOptionPane.showMessageDialog(this, "Searching for: " + query);
    }

    public static void main(String[] args) {
        SwingUtilities.invokeLater(new Runnable() {
            public void run() {
                new GamingInfoSystem().setVisible(true);
            }
        });
    }
}
