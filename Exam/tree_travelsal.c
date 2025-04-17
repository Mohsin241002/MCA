#include <stdio.h>
#include <stdlib.h>

// Define the structure for the binary tree node
typedef struct node {
    int data;
    struct node* left;
    struct node* right;
} Node;

// Function prototypes
Node* insertAtRoot(Node* root, int value);
Node* deleteRoot(Node* root);
void preorderTraversal(Node* root);
void inorderTraversal(Node* root);
void postorderTraversal(Node* root);
void levelOrderTraversal(Node* root);
Node* newNode(int data);
void printLevel(Node* root, int level);
int height(Node* root);

int main() {
    Node* root = NULL;
    int choice, value;

    while (1) {
        printf("\nMenu:\n");
        printf("1. Insert at Root\n");
        printf("2. Delete Root\n");
        printf("3. Preorder Traversal\n");
        printf("4. Inorder Traversal\n");
        printf("5. Postorder Traversal\n");
        printf("6. Level Order Traversal\n");
        printf("7. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                printf("Enter value to insert: ");
                scanf("%d", &value);
                root = insertAtRoot(root, value);
                break;
            case 2:
                root = deleteRoot(root);
                if (root == NULL) {
                    printf("Tree is empty now.\n");
                }
                break;
            case 3:
                preorderTraversal(root);
                break;
            case 4:
                inorderTraversal(root);
                break;
            case 5:
                postorderTraversal(root);
                break;
            case 6:
                levelOrderTraversal(root);
                break;
            case 7:
                return 0;
            default:
                printf("Invalid choice. Please try again.\n");
        }
    }
    return 0;
}

// Insert a new node at the root
Node* insertAtRoot(Node* root, int value) {
    Node* new_node = newNode(value);
    new_node->left = root;
    return new_node;
}

// Delete the root node
Node* deleteRoot(Node* root) {
    if (root == NULL) {
        printf("The tree is already empty.\n");
        return NULL;
    }
    Node* temp = root;
    root = root->left; // Simple example, generally more complex handling needed
    free(temp);
    return root;
}

// Create a new node
Node* newNode(int data) {
    Node* node = (Node*)malloc(sizeof(Node));
    node->data = data;
    node->left = NULL;
    node->right = NULL;
    return node;
}

// Preorder traversal
void preorderTraversal(Node* root) {
    if (root != NULL) {
        printf("%d ", root->data);
        preorderTraversal(root->left);
        preorderTraversal(root->right);
    }
}

// Inorder traversal
void inorderTraversal(Node* root) {
    if (root != NULL) {
        inorderTraversal(root->left);
        printf("%d ", root->data);
        inorderTraversal(root->right);
    }
}

// Postorder traversal
void postorderTraversal(Node* root) {
    if (root != NULL) {
        postorderTraversal(root->left);
        postorderTraversal(root->right);
        printf("%d ", root->data);
    }
}

// Level Order Traversal
void levelOrderTraversal(Node* root) {
    int h = height(root);
    for (int i = 1; i <= h; i++) {
        printLevel(root, i);
        printf("\n");
    }
}

void printLevel(Node* root, int level) {
    if (root == NULL) return;
    if (level == 1)
        printf("%d ", root->data);
    else if (level > 1) {
        printLevel(root->left, level-1);
        printLevel(root->right, level-1);
    }
}

int height(Node* root) {
    if (root == NULL)
        return 0;
    else {
        int lheight = height(root->left);
        int rheight = height(root->right);
        if (lheight > rheight) return(lheight + 1);
        else return(rheight + 1);
    }
}
