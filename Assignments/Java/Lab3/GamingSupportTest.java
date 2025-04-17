// Abstract class representing a person in gaming support (staff)
abstract class GamingSupportPerson {
    int id;
    String name;
    int experience; // Experience in years
    
    // Constructor
    GamingSupportPerson(int id, String name, int experience) {
        this.id = id;
        this.name = name;
        this.experience = experience;
    }
    
    // Abstract method for displaying details
    public abstract GamingSupportPerson displayDetails();
}

// Subclass representing a gaming support staff member
class GamingSupportStaff extends GamingSupportPerson {
    
    // Constructor with parameters
    public GamingSupportStaff(int id, String name, int experience) {
        super(id, name, experience);
    }
    
    // Default constructor with chaining
    public GamingSupportStaff() {
        this(111, "anjaney", 2); // Default staff member with 2 years of experience
    }
    
    // Overriding displayDetails method
    @Override
    public GamingSupportStaff displayDetails() {
        System.out.println("Support Staff ID: " + id);
        System.out.println("Support Staff Name: " + name);
        System.out.println("Experience: " + experience + " years");
        return this; // For method chaining
    }
    
    // Method to check and display role eligibility
    public GamingSupportStaff showRoleEligibility() {
        if (canHandleAdvancedSupport()) {
            System.out.println(name + " is eligible for an advanced support role.");
        } else {
            System.out.println(name + " is not eligible for an advanced support role.");
        }
        return this; // For method chaining
    }

    // Private helper method to determine eligibility for an advanced support role
    private boolean canHandleAdvancedSupport() {
        return experience >= 3; // Eligible if experience is 3 or more years
    }
}

// Main class for testing
public class GamingSupportTest {
    public static void main(String[] args) {
        
        System.out.println("Default Support Staff Test Case:");
        new GamingSupportStaff().displayDetails().showRoleEligibility();
        
        System.out.println("\nExperienced Support Staff Test Case:");
        new GamingSupportStaff(101, "Rohan", 5).displayDetails().showRoleEligibility();
        
        System.out.println("\nInexperienced Support Staff Test Case:");
        new GamingSupportStaff(102, "Mohsin", 1).displayDetails().showRoleEligibility();
    }
}
