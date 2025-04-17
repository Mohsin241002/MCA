#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

#define MAX 20

// Function declarations
void inputGraph(int graph[MAX][MAX], int *vertices);
void printGraph(int graph[MAX][MAX], int vertices);
void bfs(int graph[MAX][MAX], int vertices, int start, int searchVertex);
void dfsUtil(int graph[MAX][MAX], int vertices, int start, int searchVertex, bool visited[], bool *found);
void dfs(int graph[MAX][MAX], int vertices, int start, int searchVertex);
int validateVertex(int vertices);

int main() {
    int graph[MAX][MAX] = {0};
    int vertices = 0;
    int choice, startVertex, searchVertex;

    printf("Gaming Network Graph Traversal Program\n");

    while (1) {
        printf("\nMenu:\n");
        printf("1. Input Gaming Network\n");
        printf("2. Display Gaming Network\n");
        printf("3. BFS Search (Find Player)\n");
        printf("4. DFS Search (Find Player)\n");
        printf("5. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                inputGraph(graph, &vertices);
                break;
            case 2:
                if (vertices == 0) {
                    printf("Gaming network not initialized. Please input the network first.\n");
                } else {
                    printGraph(graph, vertices);
                }
                break;
            case 3:
                if (vertices == 0) {
                    printf("Gaming network not initialized. Please input the network first.\n");
                } else {
                    printf("Enter the starting player for BFS: ");
                    startVertex = validateVertex(vertices);
                    printf("Enter the player to search for: ");
                    searchVertex = validateVertex(vertices);
                    bfs(graph, vertices, startVertex, searchVertex);
                }
                break;
            case 4:
                if (vertices == 0) {
                    printf("Gaming network not initialized. Please input the network first.\n");
                } else {
                    printf("Enter the starting player for DFS: ");
                    startVertex = validateVertex(vertices);
                    printf("Enter the player to search for: ");
                    searchVertex = validateVertex(vertices);
                    dfs(graph, vertices, startVertex, searchVertex);
                }
                break;
            case 5:
                printf("Exiting program.\n");
                exit(0);
            default:
                printf("Invalid choice. Please try again.\n");
        }
    }
    return 0;
}

// Function to input the gaming network as an adjacency matrix
void inputGraph(int graph[MAX][MAX], int *vertices) {
    printf("Enter the number of players in the gaming network (max %d): ", MAX);
    scanf("%d", vertices);

    if (*vertices <= 0 || *vertices > MAX) {
        printf("Invalid number of players. Please try again.\n");
        *vertices = 0;
        return;
    }

    printf("Enter the adjacency matrix (connections between players):\n");
    for (int i = 0; i < *vertices; i++) {
        for (int j = 0; j < *vertices; j++) {
            scanf("%d", &graph[i][j]);
            if (graph[i][j] != 0 && graph[i][j] != 1) {
                printf("Invalid input. Connections should only contain 0 (no connection) or 1 (connection).\n");
                j--;
            }
        }
    }
    printf("Gaming network successfully inputted.\n");
}

// Function to print the adjacency matrix
void printGraph(int graph[MAX][MAX], int vertices) {
    printf("\nAdjacency Matrix (Gaming Network):\n");
    for (int i = 0; i < vertices; i++) {
        for (int j = 0; j < vertices; j++) {
            printf("%d ", graph[i][j]);
        }
        printf("\n");
    }
}

// BFS traversal and search
void bfs(int graph[MAX][MAX], int vertices, int start, int searchVertex) {
    bool visited[MAX] = {false};
    int queue[MAX], front = 0, rear = 0;
    bool found = false;

    queue[rear++] = start;
    visited[start] = true;

    printf("BFS Traversal (Player Connections): ");

    while (front < rear) {
        int current = queue[front++];
        printf("Player %d ", current);

        if (current == searchVertex) {
            found = true;
            break;
        }

        for (int i = 0; i < vertices; i++) {
            if (graph[current][i] && !visited[i]) {
                queue[rear++] = i;
                visited[i] = true;
            }
        }
    }

    printf("\n");
    if (found) {
        printf("Player %d found during BFS.\n", searchVertex);
    } else {
        printf("Player %d not found during BFS.\n", searchVertex);
    }
}

// DFS traversal and search
void dfs(int graph[MAX][MAX], int vertices, int start, int searchVertex) {
    bool visited[MAX] = {false};
    bool found = false;

    printf("DFS Traversal (Player Connections): ");
    dfsUtil(graph, vertices, start, searchVertex, visited, &found);
    printf("\n");

    if (found) {
        printf("Player %d found during DFS.\n", searchVertex);
    } else {
        printf("Player %d not found during DFS.\n", searchVertex);
    }
}

void dfsUtil(int graph[MAX][MAX], int vertices, int start, int searchVertex, bool visited[], bool *found) {
    visited[start] = true;
    printf("Player %d ", start);

    if (start == searchVertex) {
        *found = true;
        return;
    }

    for (int i = 0; i < vertices; i++) {
        if (graph[start][i] && !visited[i]) {
            dfsUtil(graph, vertices, i, searchVertex, visited, found);
            if (*found) return; // Stop further traversal if found
        }
    }
}

// Validate player input
int validateVertex(int vertices) {
    int vertex;
    while (1) {
        scanf("%d", &vertex);
        if (vertex >= 0 && vertex < vertices) {
            return vertex;
        } else {
            printf("Invalid player. Enter a value between 0 and %d: ", vertices - 1);
        }
    }
}
