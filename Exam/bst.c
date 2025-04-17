#include <stdio.h>
#include <stdlib.h>

typedef struct node {
    int data;
    struct node* left;
    struct node* right;
} Node;

Node* insert(Node* node, int value);
Node* deleteNode(Node* root, int key);
Node* search(Node* root, int key);
Node* minNode(Node* node);
void preorderTraversal(Node* root);
void inorderTraversal(Node* root);
void postorderTraversal(Node* root);
void levelOrderTraversal(Node* root);
void printLevel(Node* root, int level);
int height(Node* root);
Node* newNode(int value);

int main() {
    Node* root = NULL;
    int choice, value;

    while (1) {
        printf("\nMenu:\n");
        printf("1. Insert\n");
        printf("2. Delete\n");
        printf("3. Search\n");
        printf("4. Preorder Traversal\n");
        printf("5. Inorder Traversal\n");
        printf("6. Postorder Traversal\n");
        printf("7. Level Order Traversal\n");
        printf("8. Calculate Height\n");
        printf("9. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                printf("Enter value to insert: ");
                scanf("%d", &value);
                root = insert(root, value);
                break;
            case 2:
                printf("Enter value to delete: ");
                scanf("%d", &value);
                root = deleteNode(root, value);
                break;
            case 3:
                printf("Enter value to search: ");
                scanf("%d", &value);
                Node* found = search(root, value);
                if (found == NULL) {
                    printf("Value not found.\n");
                } else {
                    printf("Value %d found.\n", found->data);
                }
                break;
            case 4:
                preorderTraversal(root);
                printf("\n");
                break;
            case 5:
                inorderTraversal(root);
                printf("\n");
                break;
            case 6:
                postorderTraversal(root);
                printf("\n");
                break;
            case 7:
                levelOrderTraversal(root);
                printf("\n");
                break;
            case 8:
                printf("Height of the tree is %d.\n", height(root));
                break;
            case 9:
                return 0;
            default:
                printf("Invalid choice.\n");
        }
    }
    return 0;
}

// Function definitions

Node* newNode(int value) {
    Node* temp = (Node*) malloc(sizeof(Node));
    temp->data = value;
    temp->left = temp->right = NULL;
    return temp;
}

Node* insert(Node* node, int value) {
    if (node == NULL) return newNode(value);

    if (value < node->data)
        node->left = insert(node->left, value);
    else if (value > node->data)
        node->right = insert(node->right, value);

    return node;
}

Node* deleteNode(Node* root, int key) {
    if (root == NULL) return root;

    if (key < root->data)
        root->left = deleteNode(root->left, key);
    else if (key > root->data)
        root->right = deleteNode(root->right, key);
    else {
        if (root->left == NULL) {
            Node* temp = root->right;
            free(root);
            return temp;
        } else if (root->right == NULL) {
            Node* temp = root->left;
            free(root);
            return temp;
        }

        Node* temp = minNode(root->right);
        root->data = temp->data;
        root->right = deleteNode(root->right, temp->data);
    }
    return root;
}

Node* minNode(Node* node) {
    Node* current = node;
    while (current && current->left != NULL)
        current = current->left;
    return current;
}

Node* search(Node* root, int key) {
    if (root == NULL || root->data == key) return root;

    if (root->data > key)
        return search(root->left, key);
    return search(root->right, key);
}

void preorderTraversal(Node* root) {
    if (root != NULL) {
        printf("%d ", root->data);
        preorderTraversal(root->left);
        preorderTraversal(root->right);
    }
}

void inorderTraversal(Node* root) {
    if (root != NULL) {
        inorderTraversal(root->left);
        printf("%d ", root->data);
        inorderTraversal(root->right);
    }
}

void postorderTraversal(Node* root) {
    if (root != NULL) {
        postorderTraversal(root->left);
        postorderTraversal(root->right);
        printf("%d ", root->data);
    }
}

void levelOrderTraversal(Node* root) {
    int h = height(root);
    for (int i = 1; i <= h; i++) {
        printLevel(root, i);
    }
}

void printLevel(Node* root, int level) {
    if (root == NULL) return;
    if (level == 1) {
        printf("%d ", root->data);
    } else if (level > 1) {
        printLevel(root->left, level-1);
        printLevel(root->right, level-1);
    }
}

int height(Node* root) {
    if (root == NULL) return 0;
    int leftHeight = height(root->left);
    int rightHeight = height(root->right);
    return (leftHeight > rightHeight ? leftHeight : rightHeight) + 1;
}
