// MedicalRecord.java
   class MedicalRecord {
    // Public fields
    public String patientName;
    public int patientAge;
    public String diagnosis;
  
   
        
   
    // Public method
    public void displayMedicalRecord() {
        System.out.println("Patient Name: " + patientName);
        System.out.println("Patient Age: " + patientAge);
        System.out.println("Diagnosis: " + diagnosis);
    }
    public void setValues(String pname, int page, String pdiagnosis)
    {
     patientName = pname;
     patientAge = page;
     diagnosis = pdiagnosis;
    }
}



// Main.java
public class AccessModifier2 {
    public static void main(String[] args) {
        // Create a MedicalRecord object
        MedicalRecord patientRecord = new MedicalRecord();
        patientRecord.setValues("Aryan", 12, "Completed");

       patientRecord.displayMedicalRecord();
      //System.out.println("Patient Name: " + patientName);
    }
}