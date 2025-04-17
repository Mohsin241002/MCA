#include <stdio.h>
#include <stdlib.h>
#include <string.h>


struct Node {
    int gameID;
    struct Node *left, *right;
};

// Structure to store game metadata
struct Game {
    int gameID;
    char name[50];
    char genre[30];
    float rating;
} gameData[100];

int gameCount = 0; // Counter for the game metadata table

// Function prototypes
struct Node* insertNode(struct Node* root, int gameID);
struct Node* searchNode(struct Node* root, int gameID);
struct Node* deleteNode(struct Node* root, int gameID);
struct Node* findMin(struct Node* root);
int calculateHeight(struct Node* root);
void inorderTraversal(struct Node* root);
void preorderTraversal(struct Node* root);
void postorderTraversal(struct Node* root);
void addGameMetadata(int gameID, const char* name, const char* genre, float rating);
void displayGameMetadata(int gameID);
int getMenuChoice();

// Main function
int main() {
    struct Node* root = NULL;
    int choice;

    do {
        printf("\nGaming Information System\n");
        printf("1. Insert Game\n");
        printf("2. Search Game\n");
        printf("3. Delete Game\n");
        printf("4. Calculate Height\n");
        printf("5. Tree Traversal\n");
        printf("6. Exit\n");

        choice = getMenuChoice();

        switch (choice) {
            case 1: {
                int gameID;
                char name[50], genre[30];
                float rating;

                printf("Enter Game ID: ");
                scanf("%d", &gameID);
                printf("Enter Game Name: ");
                scanf(" %s", name);
                printf("Enter Game Genre: ");
                scanf(" %s", genre);
                printf("Enter Game Rating (0.0 to 10.0): ");
                scanf("%f", &rating);

                root = insertNode(root, gameID);
                addGameMetadata(gameID, name, genre, rating);
                printf("Game with ID %d inserted successfully.\n", gameID);
                break;
            }
            case 2: {
                int gameID;
                printf("Enter Game ID to search: ");
                scanf("%d", &gameID);

                struct Node* found = searchNode(root, gameID);
                if (found) {
                    printf("Game with ID %d found.\n", gameID);
                    displayGameMetadata(gameID);
                } else {
                    printf("Game with ID %d not found.\n", gameID);
                }
                break;
            }
            case 3: {
                int gameID;
                printf("Enter Game ID to delete: ");
                scanf("%d", &gameID);

                root = deleteNode(root, gameID);
                printf("Game with ID %d deleted if it existed.\n", gameID);
                break;
            }
            case 4: {
                int height = calculateHeight(root);
                printf("Height of the BST: %d\n", height);
                break;
            }
            case 5: {
                printf("Inorder Traversal:\n");
                inorderTraversal(root);
                printf("\nPreorder Traversal:\n");
                preorderTraversal(root);
                printf("\nPostorder Traversal:\n");
                postorderTraversal(root);
                printf("\n");
                break;
            }
            case 6:
                printf("Exiting program.\n");
                break;
            default:
                printf("Invalid choice. Try again.\n");
        }
    } while (choice != 6);

    return 0;
}

// Function to create a new node
struct Node* createNode(int gameID) {
    struct Node* newNode = (struct Node*)malloc(sizeof(struct Node));
    newNode->gameID = gameID;
    newNode->left = newNode->right = NULL;
    return newNode;
}

// Insert a new node into the BST
struct Node* insertNode(struct Node* root, int gameID) {
    if (root == NULL) {
        return createNode(gameID);
    }

    if (gameID < root->gameID) {
        root->left = insertNode(root->left, gameID);
    } else if (gameID > root->gameID) {
        root->right = insertNode(root->right, gameID);
    } else {
        printf("Game ID %d already exists in the tree.\n", gameID);
    }

    return root;
}

// Search for a node by Game ID
struct Node* searchNode(struct Node* root, int gameID) {
    if (root == NULL || root->gameID == gameID) {
        return root;
    }

    if (gameID < root->gameID) {
        return searchNode(root->left, gameID);
    } else {
        return searchNode(root->right, gameID);
    }
}

// Find the minimum node in the BST
struct Node* findMin(struct Node* root) {
    while (root && root->left != NULL) {
        root = root->left;
    }
    return root;
}

// Delete a node from the BST
struct Node* deleteNode(struct Node* root, int gameID) {
    if (root == NULL) {
        return root;
    }

    if (gameID < root->gameID) {
        root->left = deleteNode(root->left, gameID);
    } else if (gameID > root->gameID) {
        root->right = deleteNode(root->right, gameID);
    } else {
        if (root->left == NULL) {
            struct Node* temp = root->right;
            free(root);
            return temp;
        } else if (root->right == NULL) {
            struct Node* temp = root->left;
            free(root);
            return temp;
        }

        struct Node* temp = findMin(root->right);
        root->gameID = temp->gameID;
        root->right = deleteNode(root->right, temp->gameID);
    }

    return root;
}

// Calculate the height of the BST
int calculateHeight(struct Node* root) {
    if (root == NULL) {
        return -1;
    }

    int leftHeight = calculateHeight(root->left);
    int rightHeight = calculateHeight(root->right);

    return (leftHeight > rightHeight ? leftHeight : rightHeight) + 1;
}

// Inorder traversal
void inorderTraversal(struct Node* root) {
    if (root) {
        inorderTraversal(root->left);
        printf("%d ", root->gameID);
        inorderTraversal(root->right);
    }
}

// Preorder traversal
void preorderTraversal(struct Node* root) {
    if (root) {
        printf("%d ", root->gameID);
        preorderTraversal(root->left);
        preorderTraversal(root->right);
    }
}

// Postorder traversal
void postorderTraversal(struct Node* root) {
    if (root) {
        postorderTraversal(root->left);
        postorderTraversal(root->right);
        printf("%d ", root->gameID);
    }
}

// Add game metadata to the external table
void addGameMetadata(int gameID, const char* name, const char* genre, float rating) {
    gameData[gameCount].gameID = gameID;
    strcpy(gameData[gameCount].name, name);
    strcpy(gameData[gameCount].genre, genre);
    gameData[gameCount].rating = rating;
    gameCount++;
}

// Display game metadata
void displayGameMetadata(int gameID) {
    for (int i = 0; i < gameCount; i++) {
        if (gameData[i].gameID == gameID) {
            printf("Game ID: %d\n", gameID);
            printf("Name: %s\n", gameData[i].name);
            printf("Genre: %s\n", gameData[i].genre);
            printf("Rating: %.1f\n", gameData[i].rating);
            return;
        }
    }
    printf("No metadata found for Game ID %d.\n", gameID);
}

// Get menu choice with validation
int getMenuChoice() {
    int choice;
    printf("Enter your choice: ");
    scanf("%d", &choice);
    return choice;
}
