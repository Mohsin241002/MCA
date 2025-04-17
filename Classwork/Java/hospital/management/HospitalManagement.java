package hospital.management;

import hospital.Doctorclass;

public class HospitalManagement extends Doctorclass {

    public void accessDoctorDetails() {
        // Public member - accessible from any package
        System.out.println("Accessing Doctor's Name (public): " + name);

        // Protected member - accessible in subclass or package
        System.out.println("Accessing Department (protected): " + department);

        // Default (package-private) member - not accessible outside package
        // System.out.println("Accessing Hospital Name (default): " + hospitalName); // This line would cause an error

        // Private member - not accessible outside the class
        // System.out.println("Accessing Doctor ID (private): " + doctorId); // This line would cause an error

        // Calling a protected method from superclass
        displayDepartment(); // This is allowed as `displayDepartment()` is protected
    }

    public static void main(String[] args) {
        HospitalManagement management = new HospitalManagement();
        management.accessDoctorDetails();
    }
}