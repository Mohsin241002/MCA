import javax.swing.*;
import java.awt.*;
import java.util.ArrayList;
import java.util.List;

public class FeedbackApp {
    private List<Feedback> feedbackList;

    public FeedbackApp() {
        feedbackList = new ArrayList<>();
    }

    public void start() {
        JFrame frame = new JFrame("Feedback App");
        frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        frame.setSize(400, 400);

        // UI Components
        JLabel userLabel = new JLabel("User Name:");
        JTextField userField = new JTextField(15);
        JLabel movieLabel = new JLabel("Movie Name:");
        JTextField movieField = new JTextField(15);
        JLabel ratingLabel = new JLabel("Rating (1-5):");
        JTextField ratingField = new JTextField(5);
        JLabel commentLabel = new JLabel("Comments:");
        JTextField commentField = new JTextField(15);

        JButton submitButton = new JButton("Submit Feedback");
        JButton viewButton = new JButton("View Feedback");
        JButton avgButton = new JButton("Show Average Rating");

        JTextArea displayArea = new JTextArea(10, 30);
        displayArea.setEditable(false);

        JPanel panel = new JPanel();
        panel.setLayout(new GridLayout(8, 1));
        panel.add(userLabel);
        panel.add(userField);
        panel.add(movieLabel);
        panel.add(movieField);
        panel.add(ratingLabel);
        panel.add(ratingField);
        panel.add(commentLabel);
        panel.add(commentField);
        panel.add(submitButton);
        panel.add(viewButton);
        panel.add(avgButton);

        frame.add(panel, BorderLayout.NORTH);
        frame.add(new JScrollPane(displayArea), BorderLayout.CENTER);

        // Button Listeners
        submitButton.addActionListener(e -> {
            try {
                String userName = userField.getText();
                String movieName = movieField.getText();
                int rating = Integer.parseInt(ratingField.getText());
                String comments = commentField.getText();

                if (rating < 1 || rating > 5) {
                    JOptionPane.showMessageDialog(frame, "Rating must be between 1 and 5.", "Error", JOptionPane.ERROR_MESSAGE);
                    return;
                }

                Feedback feedback = new Feedback(userName, movieName, rating, comments);
                feedbackList.add(feedback);

                userField.setText("");
                movieField.setText("");
                ratingField.setText("");
                commentField.setText("");

                JOptionPane.showMessageDialog(frame, "Feedback submitted successfully!");
            } catch (NumberFormatException ex) {
                JOptionPane.showMessageDialog(frame, "Please enter a valid number for rating.", "Error", JOptionPane.ERROR_MESSAGE);
            }
        });

        viewButton.addActionListener(e -> {
            displayArea.setText("Feedback List:\n");
            for (Feedback fb : feedbackList) {
                displayArea.append(fb.toString() + "\n");
            }
        });

        avgButton.addActionListener(e -> {
            if (feedbackList.isEmpty()) {
                JOptionPane.showMessageDialog(frame, "No feedback available to calculate average rating.", "Info", JOptionPane.INFORMATION_MESSAGE);
                return;
            }
            double average = feedbackList.stream().mapToInt(Feedback::getRating).average().orElse(0);
            JOptionPane.showMessageDialog(frame, "Average Rating: " + average);
        });

        frame.setVisible(true);
    }

    public static void main(String[] args) {
        SwingUtilities.invokeLater(() -> new FeedbackApp().start());
    }
}
