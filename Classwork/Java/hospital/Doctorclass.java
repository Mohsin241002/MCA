package hospital;

public class Doctorclass {
    public String name = "Dr. John";           // Public - accessible from anywhere
    protected String department = "Cardiology"; // Protected - accessible within the package and subclasses
    String hospitalName = "City Hospital";      // Default (package-private) - accessible only within the package
    private int doctorId = 12345;               // Private - accessible only within the class

    public Doctorclass() {
        System.out.println("Doctor constructor called.");
    }

    public void showDoctorInfo() {
        System.out.println("Doctor Name: " + name);
        System.out.println("Department: " + department);
        System.out.println("Hospital Name: " + hospitalName);
        System.out.println("Doctor ID: " + doctorId);
    }

    protected void displayDepartment() {
        System.out.println("Department: " + department);
    }
}