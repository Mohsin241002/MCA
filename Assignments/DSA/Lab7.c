#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Define a Node structure for the Gaming Information System
typedef struct Node {
    char data[50]; // For storing level/achievement names
    struct Node *left, *right;
} Node;

// Queue structure for level-order traversal
typedef struct Queue {
    Node *node;
    struct Queue *next;
} Queue;

// Function Prototypes
Node* createNode(char *data);
Node* insertRoot(Node *root, char *data);
Node* deleteRoot(Node *root);
void preorderTraversal(Node *root);
void inorderTraversal(Node *root);
void postorderTraversal(Node *root);
void levelOrderTraversal(Node *root);
void enqueue(Queue **front, Queue **rear, Node *node);
Node* dequeue(Queue **front, Queue **rear);
int isEmpty(Queue *front);
void clearBuffer();

// Main function
int main() {
    Node *root = NULL;
    int choice;
    char data[50];
    
    do {
        printf("\n=== Gaming Information System Binary Tree ===\n");
        printf("1. Insert Level (Root Insertion)\n");
        printf("2. Delete Root\n");
        printf("3. Preorder Traversal\n");
        printf("4. Inorder Traversal\n");
        printf("5. Postorder Traversal\n");
        printf("6. Level-order Traversal\n");
        printf("7. Exit\n");
        printf("Enter your choice (1-7): ");
        
        // Validate menu input
        if (scanf("%d", &choice) != 1) {
            printf("Invalid input! Please enter a number between 1 and 7.\n");
            clearBuffer(); // Clear invalid input
            continue;
        }
        
        getchar(); // Consume newline character
        
        switch (choice) {
            case 1:
                printf("Enter the gaming level/achievement name (max 50 characters): ");
                fgets(data, sizeof(data), stdin);
                data[strcspn(data, "\n")] = 0; // Remove newline character
                
                // Validate non-empty input
                if (strlen(data) == 0) {
                    printf("Invalid input! Level/achievement name cannot be empty.\n");
                } else {
                    root = insertRoot(root, data);
                    printf("Added level: %s\n", data);
                }
                break;
            case 2:
                if (!root) {
                    printf("The tree is already empty. Nothing to delete.\n");
                } else {
                    root = deleteRoot(root);
                }
                break;
            case 3:
                if (!root) {
                    printf("The tree is empty! No levels to traverse.\n");
                } else {
                    printf("Preorder Traversal: ");
                    preorderTraversal(root);
                    printf("\n");
                }
                break;
            case 4:
                if (!root) {
                    printf("The tree is empty! No levels to traverse.\n");
                } else {
                    printf("Inorder Traversal: ");
                    inorderTraversal(root);
                    printf("\n");
                }
                break;
            case 5:
                if (!root) {
                    printf("The tree is empty! No levels to traverse.\n");
                } else {
                    printf("Postorder Traversal: ");
                    postorderTraversal(root);
                    printf("\n");
                }
                break;
            case 6:
                if (!root) {
                    printf("The tree is empty! No levels to traverse.\n");
                } else {
                    printf("Level-order Traversal: ");
                    levelOrderTraversal(root);
                    printf("\n");
                }
                break;
            case 7:
                printf("Exiting program. Goodbye!\n");
                break;
            default:
                printf("Invalid choice! Please enter a number between 1 and 7.\n");
        }
    } while (choice != 7);

    return 0;
}

// Function Definitions

// Clear buffer to handle invalid input
void clearBuffer() {
    int c;
    while ((c = getchar()) != '\n' && c != EOF) {}
}

// Create a new node
Node* createNode(char *data) {
    Node *newNode = (Node *)malloc(sizeof(Node));
    strcpy(newNode->data, data);
    newNode->left = newNode->right = NULL;
    return newNode;
}

// Insert node at the root
Node* insertRoot(Node *root, char *data) {
    Node *newNode = createNode(data);
    if (!root) {
        return newNode;
    }
    newNode->left = root; // Move the current root to the left
    return newNode;
}

// Delete the root node
Node* deleteRoot(Node *root) {
    if (!root) {
        printf("Tree is empty!\n");
        return NULL;
    }
    Node *newRoot = root->left ? root->left : root->right;
    printf("Deleted root: %s\n", root->data);
    free(root);
    return newRoot;
}

// Preorder traversal
void preorderTraversal(Node *root) {
    if (root) {
        printf("%s ", root->data);
        preorderTraversal(root->left);
        preorderTraversal(root->right);
    }
}

// Inorder traversal
void inorderTraversal(Node *root) {
    if (root) {
        inorderTraversal(root->left);
        printf("%s ", root->data);
        inorderTraversal(root->right);
    }
}

// Postorder traversal
void postorderTraversal(Node *root) {
    if (root) {
        postorderTraversal(root->left);
        postorderTraversal(root->right);
        printf("%s ", root->data);
    }
}

// Level-order traversal
void levelOrderTraversal(Node *root) {
    if (!root) {
        printf("Tree is empty!\n");
        return;
    }
    Queue *front = NULL, *rear = NULL;
    enqueue(&front, &rear, root);
    while (!isEmpty(front)) {
        Node *current = dequeue(&front, &rear);
        printf("%s ", current->data);
        if (current->left) enqueue(&front, &rear, current->left);
        if (current->right) enqueue(&front, &rear, current->right);
    }
}

// Helper functions for queue
void enqueue(Queue **front, Queue **rear, Node *node) {
    Queue *newNode = (Queue *)malloc(sizeof(Queue));
    newNode->node = node;
    newNode->next = NULL;
    if (*rear) (*rear)->next = newNode;
    *rear = newNode;
    if (!*front) *front = newNode;
}

Node* dequeue(Queue **front, Queue **rear) {
    if (!*front) return NULL;
    Queue *temp = *front;
    Node *node = temp->node;
    *front = (*front)->next;
    if (!*front) *rear = NULL;
    free(temp);
    return node;
}

int isEmpty(Queue *front) {
    return front == NULL;
}
