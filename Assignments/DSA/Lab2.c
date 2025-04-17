#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

// Problem Domain: Gaming
// Scenario: We are developing a system to manage a collection of video games using linked lists.
// The user can choose between different types of linked lists (singly, doubly, circular) to manage games.
// Each game has a name, release date, and price. Prices are generated randomly in some cases.
// The system should allow adding, deleting, displaying games, and saving/loading data to/from files.

// Structure for Singly Linked List Node (Game)
typedef struct GameNode {
    char game_name[50];
    char release_date[12];
    double price;
    struct GameNode* next;
} GameNode;

// Structure for Doubly Linked List Node (Game)
typedef struct DoublyGameNode {
    char game_name[50];
    char release_date[12];
    double price;
    struct DoublyGameNode* next;
    struct DoublyGameNode* prev;
} DoublyGameNode;

// Structure for Circular Linked List Node (Game)
typedef struct CircularGameNode {
    char game_name[50];
    char release_date[12];
    double price;
    struct CircularGameNode* next;
} CircularGameNode;

// Function prototypes
void singly_linked_list_menu();
void doubly_linked_list_menu();
void circular_linked_list_menu();
void save_games_to_file(GameNode* head, const char* filename);
void load_games_from_file(GameNode** head, const char* filename);
void save_doubly_games_to_file(DoublyGameNode* head, const char* filename);
void load_doubly_games_from_file(DoublyGameNode** head, const char* filename);
void save_circular_games_to_file(CircularGameNode* head, const char* filename);
void load_circular_games_from_file(CircularGameNode** head, const char* filename);
double generate_random_price();  // Function to generate random game prices

// Create a new game node for the singly linked list
GameNode* create_game(char* name, char* date, double price) {
    GameNode* new_node = (GameNode*)malloc(sizeof(GameNode));
    strcpy(new_node->game_name, name);
    strcpy(new_node->release_date, date);
    new_node->price = price;
    new_node->next = NULL;
    return new_node;
}

// Generate random prices between $10 and $100
double generate_random_price() {
    return (rand() % 91) + 10;
}

// Insert at the beginning of the singly linked list
void insert_at_beginning(GameNode** head, char* name, char* date, double price) {
    GameNode* new_node = create_game(name, date, price);
    new_node->next = *head;
    *head = new_node;
}

// Insert at the end of the singly linked list
void insert_at_end(GameNode** head, char* name, char* date, double price) {
    GameNode* new_node = create_game(name, date, price);
    if (*head == NULL) {
        *head = new_node;
        return;
    }
    GameNode* last = *head;
    while (last->next != NULL)
        last = last->next;
    last->next = new_node;
}

// Delete a game by name from the singly linked list
void delete_game(GameNode** head, char* name) {
    GameNode* temp = *head;
    GameNode* prev = NULL;

    if (temp != NULL && strcmp(temp->game_name, name) == 0) {
        *head = temp->next;
        free(temp);
        return;
    }

    while (temp != NULL && strcmp(temp->game_name, name) != 0) {
        prev = temp;
        temp = temp->next;
    }

    if (temp == NULL) {
        printf("Game not found.\n");
        return;
    }

    prev->next = temp->next;
    free(temp);
}

// Display the list of games in the singly linked list
void display_games(GameNode* head) {
    if (head == NULL) {
        printf("No games available.\n");
        return;
    }
    GameNode* current = head;
    while (current != NULL) {
        printf("Game: %s, Release Date: %s, Price: $%.2f\n", current->game_name, current->release_date, current->price);
        current = current->next;
    }
}

// Save games to file
void save_games_to_file(GameNode* head, const char* filename) {
    FILE* file = fopen(filename, "w");
    if (file == NULL) {
        printf("Error opening file for writing.\n");
        return;
    }
    GameNode* current = head;
    while (current != NULL) {
        fprintf(file, "%s %s %.2f\n", current->game_name, current->release_date, current->price);
        current = current->next;
    }
    fclose(file);
    printf("Games saved to file.\n");
}

// Load games from file
void load_games_from_file(GameNode** head, const char* filename) {
    FILE* file = fopen(filename, "r");
    if (file == NULL) {
        printf("Error opening file for reading.\n");
        return;
    }
    char name[50], date[12];
    double price;
    while (fscanf(file, "%s %s %lf", name, date, &price) != EOF) {
        insert_at_end(head, name, date, price);
    }
    fclose(file);
    printf("Games loaded from file.\n");
}

// Menu for singly linked list operations
void singly_linked_list_menu() {
    GameNode* singly_head = NULL;
    int choice;
    char name[50], date[12];
    double price;

    do {
        printf("\nSingly Linked List Menu:\n");
        printf("1. Insert Game at Beginning\n");
        printf("2. Insert Game at End\n");
        printf("3. Delete Game\n");
        printf("4. Display Games\n");
        printf("5. Save Games to File\n");
        printf("6. Load Games from File\n");
        printf("7. Exit Singly Linked List Menu\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                printf("Enter Game Name: ");
                scanf("%s", name);
                printf("Enter Release Date (YYYY-MM-DD): ");
                scanf("%s", date);
                price = generate_random_price();
                insert_at_beginning(&singly_head, name, date, price);
                break;
            case 2:
                printf("Enter Game Name: ");
                scanf("%s", name);
                printf("Enter Release Date (YYYY-MM-DD): ");
                scanf("%s", date);
                price = generate_random_price();
                insert_at_end(&singly_head, name, date, price);
                break;
            case 3:
                printf("Enter Game Name to Delete: ");
                scanf("%s", name);
                delete_game(&singly_head, name);
                break;
            case 4:
                display_games(singly_head);
                break;
            case 5:
                save_games_to_file(singly_head, "games.txt");
                break;
            case 6:
                load_games_from_file(&singly_head, "games.txt");
                break;
            case 7:
                return;
            default:
                printf("Invalid choice! Please try again.\n");
        }
    } while (choice != 7);
}

// Create a new game node for the doubly linked list
DoublyGameNode* create_doubly_game(char* name, char* date, double price) {
    DoublyGameNode* new_node = (DoublyGameNode*)malloc(sizeof(DoublyGameNode));
    strcpy(new_node->game_name, name);
    strcpy(new_node->release_date, date);
    new_node->price = price;
    new_node->next = NULL;
    new_node->prev = NULL;
    return new_node;
}

// Insert at the beginning of the doubly linked list
void insert_at_beginning_dll(DoublyGameNode** head, char* name, char* date, double price) {
    DoublyGameNode* new_node = create_doubly_game(name, date, price);
    if (*head == NULL) {
        *head = new_node;
        return;
    }
    new_node->next = *head;
    (*head)->prev = new_node;
    *head = new_node;
}

// Insert at the end of the doubly linked list
void insert_at_end_dll(DoublyGameNode** head, char* name, char* date, double price) {
    DoublyGameNode* new_node = create_doubly_game(name, date, price);
    if (*head == NULL) {
        *head = new_node;
        return;
    }
    DoublyGameNode* last = *head;
    while (last->next != NULL)
        last = last->next;
    last->next = new_node;
    new_node->prev = last;
}

// Delete a game by name from the doubly linked list
void delete_game_dll(DoublyGameNode** head, char* name) {
    DoublyGameNode* temp = *head;

    while (temp != NULL && strcmp(temp->game_name, name) != 0) {
        temp = temp->next;
    }

    if (temp == NULL) {
        printf("Game not found.\n");
        return;
    }

    if (temp->prev != NULL) {
        temp->prev->next = temp->next;
    } else {
        *head = temp->next;  // Change head if the node to be deleted is the first node
    }

    if (temp->next != NULL) {
        temp->next->prev = temp->prev;
    }

    free(temp);
}

// Display the list of games in the doubly linked list
void display_games_dll(DoublyGameNode* head) {
    if (head == NULL) {
        printf("No games available.\n");
        return;
    }
    DoublyGameNode* current = head;
    while (current != NULL) {
        printf("Game: %s, Release Date: %s, Price: $%.2f\n", current->game_name, current->release_date, current->price);
        current = current->next;
    }
}

// Save games to file (Doubly Linked List)
void save_doubly_games_to_file(DoublyGameNode* head, const char* filename) {
    FILE* file = fopen(filename, "w");
    if (file == NULL) {
        printf("Error opening file for writing.\n");
        return;
    }
    DoublyGameNode* current = head;
    while (current != NULL) {
        fprintf(file, "%s %s %.2f\n", current->game_name, current->release_date, current->price);
        current = current->next;
    }
    fclose(file);
    printf("Doubly linked list games saved to file.\n");
}

// Load games from file (Doubly Linked List)
void load_doubly_games_from_file(DoublyGameNode** head, const char* filename) {
    FILE* file = fopen(filename, "r");
    if (file == NULL) {
        printf("Error opening file for reading.\n");
        return;
    }
    char name[50], date[12];
    double price;
    while (fscanf(file, "%s %s %lf", name, date, &price) != EOF) {
        insert_at_end_dll(head, name, date, price);
    }
    fclose(file);
    printf("Doubly linked list games loaded from file.\n");
}

// Menu for doubly linked list operations
void doubly_linked_list_menu() {
    DoublyGameNode* doubly_head = NULL;
    int choice;
    char name[50], date[12];
    double price;

    do {
        printf("\nDoubly Linked List Menu:\n");
        printf("1. Insert Game at Beginning\n");
        printf("2. Insert Game at End\n");
        printf("3. Delete Game\n");
        printf("4. Display Games\n");
        printf("5. Save Games to File\n");
        printf("6. Load Games from File\n");
        printf("7. Exit Doubly Linked List Menu\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                printf("Enter Game Name: ");
                scanf("%s", name);
                printf("Enter Release Date (YYYY-MM-DD): ");
                scanf("%s", date);
                price = generate_random_price();
                insert_at_beginning_dll(&doubly_head, name, date, price);
                break;
            case 2:
                printf("Enter Game Name: ");
                scanf("%s", name);
                printf("Enter Release Date (YYYY-MM-DD): ");
                scanf("%s", date);
                price = generate_random_price();
                insert_at_end_dll(&doubly_head, name, date, price);
                break;
            case 3:
                printf("Enter Game Name to Delete: ");
                scanf("%s", name);
                delete_game_dll(&doubly_head, name);
                break;
            case 4:
                display_games_dll(doubly_head);
                break;
            case 5:
                save_doubly_games_to_file(doubly_head, "doubly_games.txt");
                break;
            case 6:
                load_doubly_games_from_file(&doubly_head, "doubly_games.txt");
                break;
            case 7:
                return;
            default:
                printf("Invalid choice! Please try again.\n");
        }
    } while (choice != 7);
}

// Create a new game node for the circular linked list
CircularGameNode* create_circular_game(char* name, char* date, double price) {
    CircularGameNode* new_node = (CircularGameNode*)malloc(sizeof(CircularGameNode));
    strcpy(new_node->game_name, name);
    strcpy(new_node->release_date, date);
    new_node->price = price;
    new_node->next = NULL;
    return new_node;
}

// Insert at the end of the circular linked list
void insert_at_end_cll(CircularGameNode** head, char* name, char* date, double price) {
    CircularGameNode* new_node = create_circular_game(name, date, price);
    if (*head == NULL) {
        *head = new_node;
        new_node->next = *head; // Point to itself
        return;
    }
    CircularGameNode* last = *head;
    while (last->next != *head)
        last = last->next;
    last->next = new_node;
    new_node->next = *head; // Complete the circular link
}

// Delete a game by name from the circular linked list
void delete_game_cll(CircularGameNode** head, char* name) {
    if (*head == NULL) {
        printf("No games available.\n");
        return;
    }

    CircularGameNode* current = *head;
    CircularGameNode* prev = NULL;

    // Find the node to be deleted
    do {
        if (strcmp(current->game_name, name) == 0) {
            if (prev == NULL) { // Node to be deleted is the head
                CircularGameNode* last = *head;
                while (last->next != *head) {
                    last = last->next;
                }
                if (current->next == *head) {
                    free(current);
                    *head = NULL; // List becomes empty
                } else {
                    last->next = current->next;
                    *head = current->next; // Move head to the next node
                    free(current);
                }
            } else {
                prev->next = current->next;
                free(current);
            }
            return;
        }
        prev = current;
        current = current->next;
    } while (current != *head);

    printf("Game not found.\n");
}

// Display the list of games in the circular linked list
void display_games_cll(CircularGameNode* head) {
    if (head == NULL) {
        printf("No games available.\n");
        return;
    }
    CircularGameNode* current = head;
    do {
        printf("Game: %s, Release Date: %s, Price: $%.2f\n", current->game_name, current->release_date, current->price);
        current = current->next;
    } while (current != head);
}

// Save games to file (Circular Linked List)
void save_circular_games_to_file(CircularGameNode* head, const char* filename) {
    FILE* file = fopen(filename, "w");
    if (file == NULL) {
        printf("Error opening file for writing.\n");
        return;
    }
    CircularGameNode* current = head;
    if (current != NULL) {
        do {
            fprintf(file, "%s %s %.2f\n", current->game_name, current->release_date, current->price);
            current = current->next;
        } while (current != head);
    }
    fclose(file);
    printf("Circular linked list games saved to file.\n");
}

// Load games from file (Circular Linked List)
void load_circular_games_from_file(CircularGameNode** head, const char* filename) {
    FILE* file = fopen(filename, "r");
    if (file == NULL) {
        printf("Error opening file for reading.\n");
        return;
    }
    char name[50], date[12];
    double price;
    while (fscanf(file, "%s %s %lf", name, date, &price) != EOF) {
        insert_at_end_cll(head, name, date, price);
    }
    fclose(file);
    printf("Circular linked list games loaded from file.\n");
}

// Menu for circular linked list operations
void circular_linked_list_menu() {
    CircularGameNode* circular_head = NULL;
    int choice;
    char name[50], date[12];
    double price;

    do {
        printf("\nCircular Linked List Menu:\n");
        printf("1. Insert Game at End\n");
        printf("2. Delete Game\n");
        printf("3. Display Games\n");
        printf("4. Save Games to File\n");
        printf("5. Load Games from File\n");
        printf("6. Exit Circular Linked List Menu\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                printf("Enter Game Name: ");
                scanf("%s", name);
                printf("Enter Release Date (YYYY-MM-DD): ");
                scanf("%s", date);
                price = generate_random_price();
                insert_at_end_cll(&circular_head, name, date, price);
                break;
            case 2:
                printf("Enter Game Name to Delete: ");
                scanf("%s", name);
                delete_game_cll(&circular_head, name);
                break;
            case 3:
                display_games_cll(circular_head);
                break;
            case 4:
                save_circular_games_to_file(circular_head, "circular_games.txt");
                break;
            case 5:
                load_circular_games_from_file(&circular_head, "circular_games.txt");
                break;
            case 6:
                return;
            default:
                printf("Invalid choice! Please try again.\n");
        }
    } while (choice != 6);
}

// Main menu
int main() {
    srand(time(0)); // Seed for random number generation
    int choice;

    do {
        printf("\nMain Menu:\n");
        printf("1. Singly Linked List\n");
        printf("2. Doubly Linked List\n");
        printf("3. Circular Linked List\n");
        printf("4. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                singly_linked_list_menu();
                break;
            case 2:
                doubly_linked_list_menu();
                break;
            case 3:
                circular_linked_list_menu();
                break;
            case 4:
                printf("Exiting the program. Goodbye!\n");
                break;
            default:
                printf("Invalid choice! Please try again.\n");
        }
    } while (choice != 4);

    return 0;
}
