#include <stdio.h>
#include <stdlib.h>

void createFile() {
    FILE *fp;
    fp = fopen("example.txt", "w"); 
    if (fp == NULL) {
        printf("Could not open file\n");
        return;
    }
    fprintf(fp, "This is an example file.\n");
    fclose(fp);
    printf("File created successfully\n");
}

void readFile() {
    FILE *fp;
    fp = fopen("example.txt", "r"); 
    if (fp == NULL) {
        printf("Could not open file\n");
        return;
    }
    char buffer[256];
    while (fgets(buffer, 256, fp) != NULL) {
        printf("%s", buffer);
    }
    fclose(fp);
}

void appendToFile() {
    FILE *fp;
    fp = fopen("example.txt", "a"); 
    if (fp == NULL) {
        printf("Could not open file\n");
        return;
    }
    fprintf(fp, "This is an appended line.\n");
    fclose(fp);
    printf("File appended successfully\n");
}

void updateFile() {
    FILE *fp;
    fp = fopen("example.txt", "r+"); 
    if (fp == NULL) {
        printf("Could not open file\n");
        return;
    }
    fseek(fp, 0, SEEK_SET); // move to the beginning of the file
    fprintf(fp, "This is the updated first line.\n");
    fclose(fp);
    printf("File updated successfully\n");
}

void deleteFile() {
    if (remove("example.txt") == 0) {
        printf("File deleted successfully\n");
    } else {
        printf("Error deleting file\n");
    }
}

int main() {
    int choice;
    while (1) {
        printf("File Handling Operations\n");
        printf("1. Create a new file\n");
        printf("2. Read from a file\n");
        printf("3. Append to a file\n");
        printf("4. Update a file\n");
        printf("5. Delete a file\n");
        printf("6. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                createFile();
                break;
            case 2:
                readFile();
                break;
            case 3:
                appendToFile();
                break;
            case 4:
                updateFile();
                break;
            case 5:
                deleteFile();
                break;
            case 6:
                printf("Exiting...\n");
                return 0;
            default:
                printf("Invalid choice. Please try again.\n");
        }
    }

    return 0;
}