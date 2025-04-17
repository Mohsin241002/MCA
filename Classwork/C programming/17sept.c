#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Structure to store student details
struct Student {
    char name[100];
    int roll_no;
    float marks[5]; // Marks of 5 subjects
};

// Function prototypes
void createFile();
void writeStudentDetails();
void readStudentDetails();
void appendStudentDetails();
void renameFile();
void deleteFile();

int main() {
    int choice;
    
    while (1) {
        // Display menu options
        printf("\nFile Management System Menu:\n");
        printf("1. Create a file\n");
        printf("2. Write student details to a file\n");
        printf("3. Read student details from a file\n");
        printf("4. Append student details to a file\n");
        printf("5. Rename a file\n");
        printf("6. Delete a file\n");
        printf("7. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);
        
        // Handle user's choice
        switch (choice) {
            case 1:
                createFile();
                break;
            case 2:
                writeStudentDetails();
                break;
            case 3:
                readStudentDetails();
                break;
            case 4:
                appendStudentDetails();
                break;
            case 5:
                renameFile();
                break;
            case 6:
                deleteFile();
                break;
            case 7:
                printf("Exiting...\n");
                exit(0);
            default:
                printf("Invalid choice. Please try again.\n");
        }
    }

    return 0;
}

// Function to create a file
void createFile() {
    FILE *fp;
    char filename[100];
    
    printf("Enter the file name to create: ");
    scanf("%s", filename);
    
    fp = fopen(filename, "w");
    if (fp == NULL) {
        printf("Error creating file!\n");
        return;
    }
    
    printf("File created successfully: %s\n", filename);
    fclose(fp);
}

// Function to write student details to a file
void writeStudentDetails() {
    FILE *fp;
    char filename[100];
    struct Student s;
    
    printf("Enter the file name to write to: ");
    scanf("%s", filename);
    
    fp = fopen(filename, "w");
    if (fp == NULL) {
        printf("Error opening file!\n");
        return;
    }
    
    // Input student details
    printf("Enter student name: ");
    getchar(); // To consume the newline character
    fgets(s.name, sizeof(s.name), stdin);
    s.name[strcspn(s.name, "\n")] = 0; // Remove trailing newline character

    printf("Enter roll number: ");
    scanf("%d", &s.roll_no);

    for (int i = 0; i < 5; i++) {
        printf("Enter marks for subject %d: ", i + 1);
        scanf("%f", &s.marks[i]);
    }

    // Write details to the file
    fprintf(fp, "Name: %s\n", s.name);
    fprintf(fp, "Roll No: %d\n", s.roll_no);
    for (int i = 0; i < 5; i++) {
        fprintf(fp, "Marks for subject %d: %.2f\n", i + 1, s.marks[i]);
    }
    
    printf("Student details written to file successfully.\n");
    
    fclose(fp);
}

// Function to read student details from a file
void readStudentDetails() {
    FILE *fp;
    char filename[100], c;
    
    printf("Enter the file name to read from: ");
    scanf("%s", filename);
    
    fp = fopen(filename, "r");
    if (fp == NULL) {
        printf("Error opening file!\n");
        return;
    }
    
    printf("\nFile contents:\n");
    while ((c = fgetc(fp)) != EOF) {
        putchar(c);
    }
    
    fclose(fp);
}

// Function to append student details to a file
void appendStudentDetails() {
    FILE *fp;
    char filename[100];
    struct Student s;
    
    printf("Enter the file name to append to: ");
    scanf("%s", filename);
    
    fp = fopen(filename, "a");
    if (fp == NULL) {
        printf("Error opening file!\n");
        return;
    }
    
    // Input student details
    printf("Enter student name: ");
    getchar(); // To consume the newline character
    fgets(s.name, sizeof(s.name), stdin);
    s.name[strcspn(s.name, "\n")] = 0; // Remove trailing newline character

    printf("Enter roll number: ");
    scanf("%d", &s.roll_no);

    for (int i = 0; i < 5; i++) {
        printf("Enter marks for subject %d: ", i + 1);
        scanf("%f", &s.marks[i]);
    }

    // Append details to the file
    fprintf(fp, "Name: %s\n", s.name);
    fprintf(fp, "Roll No: %d\n", s.roll_no);
    for (int i = 0; i < 5; i++) {
        fprintf(fp, "Marks for subject %d: %.2f\n", i + 1, s.marks[i]);
    }
    
    printf("Student details appended to file successfully.\n");
    
    fclose(fp);
}

// Function to rename a file
void renameFile() {
    char oldname[100], newname[100];
    
    printf("Enter the current file name: ");
    scanf("%s", oldname);
    printf("Enter the new file name: ");
    scanf("%s", newname);
    
    if (rename(oldname, newname) == 0) {
        printf("File renamed successfully.\n");
    } else {
        printf("Error renaming file.\n");
    }
}

// Function to delete a file
void deleteFile() {
    char filename[100];
    
    printf("Enter the file name to delete: ");
    scanf("%s", filename);
    
    if (remove(filename) == 0) {
        printf("File deleted successfully.\n");
    } else {
        printf("Error deleting file.\n");
    }
}
