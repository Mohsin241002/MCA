/** 
1. Design and implement and Employee Performance Management System using circular singly linked list to manage employee records . each node in list will represent and employee and contain the following details:
        Name
        Preformance Score
        Employee ID
    
    Requirements:
        Read Emplyee Details and create a Circular Singly Linked List 
            Create a circular singly linked list where each node represents an employee.
            For each employee:  
                1.Add the new employee node at the end of the sxisting list.
                2.Vlaidate the EmployeeID as a positive integer
                3.Ensure the Performance Score is between 0 and 100.

        Display All Employee Details:
            Implement a function to display the details of all employees in the list.
        
        Find the Employee with the Highest Performance Score:
            Write a function to indentify and display the details of the employee with the highest performance score.

        Split Employees Based on Performance:
            Split circular singly linked list into two Seperate circular singly linked list:
                low performance list:
                    contains employees with a performance score less than 40
                High Performance list:
                    Contains employees with a performance score of 40 and above
            Display both List separately.

 */           



#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Define the Employee structure
struct Employee {
    int id;
    char name[50];
    int performanceScore;
    struct Employee* next;
};

// Function to create a new Employee node
struct Employee* createEmployee(int id, const char* name, int performanceScore) {
    struct Employee* newEmployee = (struct Employee*)malloc(sizeof(struct Employee));
    newEmployee->id = id;
    strcpy(newEmployee->name, name);
    newEmployee->performanceScore = performanceScore;
    newEmployee->next = newEmployee; // Initially points to itself
    return newEmployee;
/**
 * Creates a new Employee node with the given id, name, and performance score.
 * Allocates memory for the Employee structure and initializes its fields.
 * The 'next' pointer is initially set to point to the node itself.
 *
 * @param id The unique identifier for the employee.
 * @param name The name of the employee.
 * @param performanceScore The performance score of the employee.
 * @return A pointer to the newly created Employee node.
 */
}

// Function to add an employee at the end of the list


struct Employee* addEmployee(struct Employee* head, int id, const char* name, int performanceScore) {
    if (id <= 0) {
        printf("Invalid Employee ID. ID must be a positive integer.\n");
        return head;
    }
    if (performanceScore < 0 || performanceScore > 100) {
        printf("Invalid Performance Score. Score must be between 0 and 100.\n");
        return head;
    }

    struct Employee* newEmployee = createEmployee(id, name, performanceScore);

    if (head == NULL) {
        head = newEmployee;
    } else {
        struct Employee* temp = head;
        while (temp->next != head) {
            temp = temp->next;
        }
        temp->next = newEmployee;
        newEmployee->next = head;
    }

    printf("Employee added successfully!\n");
    return head;

/**
 * Adds a new employee to the end of a circular linked list of employees.
 *
 * This function creates a new employee node with the given ID, name, and performance score,
 * and appends it to the end of the list. It ensures that the employee ID is a positive integer
 * and the performance score is within the range of 0 to 100. If the list is empty, the new
 * employee becomes the head of the list.
 *
 * @param head Pointer to the head of the employee list.
 * @param id The unique identifier for the employee. Must be a positive integer.
 * @param name The name of the employee.
 * @param performanceScore The performance score of the employee, ranging from 0 to 100.
 * @return Pointer to the head of the updated employee list.
 */
}

// Function to display all employee details
void displayEmployees(struct Employee* head) {
    if (head == NULL) {
        printf("No employees in the list.\n");
        return;
    }

    struct Employee* temp = head;
    printf("Employee Details:\n");
    do {
        printf("ID: %d, Name: %s, Performance Score: %d\n", temp->id, temp->name, temp->performanceScore);
        temp = temp->next;
    } while (temp != head);

/**
 * Displays the details of all employees in a circular linked list.
 *
 * This function iterates through a circular linked list of Employee
 * structures, starting from the given head node, and prints the ID,
 * name, and performance score of each employee. If the list is empty,
 * it prints a message indicating that there are no employees.
 *
 * @param head Pointer to the head node of the Employee linked list.
 */
}

// Function to find and display the employee with the highest performance score
void displayTopPerformer(struct Employee* head) {
    if (head == NULL) {
        printf("No employees in the list.\n");
        return;
    }

    struct Employee* temp = head;
    struct Employee* topPerformer = head;
    do {
        if (temp->performanceScore > topPerformer->performanceScore) {
            topPerformer = temp;
        }
        temp = temp->next;
    } while (temp != head);

    printf("Top Performer:\nID: %d, Name: %s, Performance Score: %d\n",
           topPerformer->id, topPerformer->name, topPerformer->performanceScore);
        
/**
 * Displays the employee with the highest performance score from a circular linked list.
 * 
 * This function traverses a circular linked list of Employee structures, identifies
 * the employee with the highest performance score, and prints their details. If the
 * list is empty, it outputs a message indicating that there are no employees.
 * 
 * @param head Pointer to the head of the circular linked list of employees.
 */
}

// Function to split employees into high and low performance lists
void splitEmployeeList(struct Employee* head, struct Employee** lowPerformers, struct Employee** highPerformers) {
    if (head == NULL) return;

    struct Employee* temp = head;
    do {
        if (temp->performanceScore < 40) {
            *lowPerformers = addEmployee(*lowPerformers, temp->id, temp->name, temp->performanceScore);
        } else {
            *highPerformers = addEmployee(*highPerformers, temp->id, temp->name, temp->performanceScore);
        }
        temp = temp->next;
    } while (temp != head);

/**
 * Splits a circular linked list of employees into two separate lists based on performance scores.
 *
 * This function iterates through the given circular linked list of employees and divides them
 * into two separate lists: one for low performers and another for high performers. An employee
 * is considered a low performer if their performance score is less than 40, otherwise they are
 * considered a high performer. The function updates the provided pointers to the head of the
 * low and high performer lists.
 *
 * @param head Pointer to the head of the circular linked list of employees.
 * @param lowPerformers Pointer to the head of the list for low-performing employees.
 * @param highPerformers Pointer to the head of the list for high-performing employees.
 */
}

// Menu-driven program
int main() {
    struct Employee* head = NULL;
    struct Employee* lowPerformers = NULL;
    struct Employee* highPerformers = NULL;
    int choice, id, performanceScore;
    char name[50];

    while (1) {
        printf("\nEmployee Performance Management System\n");
        printf("1. Add Employee\n");
        printf("2. Display All Employees\n");
        printf("3. Display Top Performer\n");
        printf("4. Split Employees by Performance\n");
        printf("5. Display Low Performers\n");
        printf("6. Display High Performers\n");
        printf("0. Exit\n");
        printf("Enter your choice: ");
        
        if (scanf("%d", &choice) != 1) {
            printf("Invalid input. Please enter a number.\n");
            while (getchar() != '\n'); // Clear input buffer
            continue;
        }

        switch (choice) {
            case 1:
                printf("Enter Employee ID: ");
                if (scanf("%d", &id) != 1 || id <= 0) {
                    printf("Invalid ID. Please enter a positive integer.\n");
                    while (getchar() != '\n');
                    continue;
                }

                printf("Enter Employee Name: ");
                scanf(" %[^\n]", name); // Read until newline

                printf("Enter Performance Score (0-100): ");
                if (scanf("%d", &performanceScore) != 1 || performanceScore < 0 || performanceScore > 100) {
                    printf("Invalid Performance Score. Must be between 0 and 100.\n");
                    while (getchar() != '\n');
                    continue;
                }

                head = addEmployee(head, id, name, performanceScore);
                break;

            case 2:
                displayEmployees(head);
                break;

            case 3:
                displayTopPerformer(head);
                break;

            case 4:
                lowPerformers = NULL;
                highPerformers = NULL;
                splitEmployeeList(head, &lowPerformers, &highPerformers);
                printf("Employees have been split based on performance.\n");
                break;

            case 5:
                printf("Low Performance Employees (Score < 40):\n");
                displayEmployees(lowPerformers);
                break;

            case 6:
                printf("High Performance Employees (Score >= 40):\n");
                displayEmployees(highPerformers);
                break;

            case 0:
                printf("Exiting the program.\n");
                // Free memory here
                return 0;

            default:
                printf("Invalid choice. Please select a valid option.\n");
        }
    }

    return 0;

/**
 * Main function for the Employee Performance Management System.
 * 
 * This menu-driven program allows users to manage employee data, including:
 * - Adding new employees with ID, name, and performance score.
 * - Displaying all employees.
 * - Displaying the top performer based on performance score.
 * - Splitting employees into high and low performance lists.
 * - Displaying employees in each performance category.
 * 
 * The program continuously prompts the user for actions until the exit option is selected.
 * Input validation is performed for menu choices, employee IDs, and performance scores.
 */
}
