#include <stdio.h>
#include <stdlib.h>

typedef struct AdjListNode {
    int dest;
    struct AdjListNode* next;
} AdjListNode;

typedef struct AdjList {
    AdjListNode *head;  // Pointer to head node of list
} AdjList;

typedef struct Graph {
    int V;              // Number of vertices
    AdjList* array;     // An array of adjacency lists
} Graph;

// Function prototypes
Graph* createGraph(int V);
void addEdge(Graph* graph, int src, int dest);
void BFS(Graph* graph, int startVertex, int searchVertex);
void DFS(Graph* graph, int startVertex, int searchVertex);
AdjListNode* newAdjListNode(int dest);
int readInt(const char* prompt);

int main() {
    int V = readInt("Enter the number of vertices: ");
    Graph* graph = createGraph(V);

    int choice, src, dest;
    while (1) {
        printf("\nMenu:\n");
        printf("1. Add Edge\n");
        printf("2. BFS\n");
        printf("3. DFS\n");
        printf("4. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                src = readInt("Enter source vertex: ");
                dest = readInt("Enter destination vertex: ");
                if (src >= 0 && src < V && dest >= 0 && dest < V) {
                    addEdge(graph, src, dest);
                } else {
                    printf("Invalid vertex number. Vertex must be between 0 and %d\n", V-1);
                }
                break;
            case 2:
                src = readInt("Enter starting vertex for BFS: ");
                dest = readInt("Enter vertex to search using BFS: ");
                if (src >= 0 && src < V) {
                    BFS(graph, src, dest);
                } else {
                    printf("Invalid starting vertex.\n");
                }
                break;
            case 3:
                src = readInt("Enter starting vertex for DFS: ");
                dest = readInt("Enter vertex to search using DFS: ");
                if (src >= 0 && src < V) {
                    DFS(graph, src, dest);
                } else {
                    printf("Invalid starting vertex.\n");
                }
                break;
            case 4:
                printf("Exiting...\n");
                return 0;
            default:
                printf("Invalid choice.\n");
        }
    }

    return 0;
}

// Create a graph with V vertices
Graph* createGraph(int V) {
    Graph* graph = (Graph*) malloc(sizeof(Graph));
    graph->V = V;
    graph->array = (AdjList*) malloc(V * sizeof(AdjList));

    for (int i = 0; i < V; ++i) {
        graph->array[i].head = NULL;
    }

    return graph;
}

// Adds an edge to an undirected graph
void addEdge(Graph* graph, int src, int dest) {
    // Add an edge from src to dest. A new node is added to the adjacency
    // list of src. The node is added at the beginning
    AdjListNode* newNode = newAdjListNode(dest);
    newNode->next = graph->array[src].head;
    graph->array[src].head = newNode;

    // Since graph is undirected, add an edge from dest to src also
    newNode = newAdjListNode(src);
    newNode->next = graph->array[dest].head;
    graph->array[dest].head = newNode;
}

AdjListNode* newAdjListNode(int dest) {
    AdjListNode* newNode = (AdjListNode*) malloc(sizeof(AdjListNode));
    newNode->dest = dest;
    newNode->next = NULL;
    return newNode;
}

int readInt(const char* prompt) {
    int value;
    printf("%s", prompt);
    while (scanf("%d", &value) != 1) {
        while (getchar() != '\n');  // clear buffer
        printf("Invalid input. Please enter an integer: ");
    }
    return value;
}

void BFS(Graph* graph, int startVertex, int searchVertex) {
    int V = graph->V;
    int visited[V];
    int queue[V], front = 0, rear = -1, found = 0;

    for (int i = 0; i < V; i++) {
        visited[i] = 0;
    }

    visited[startVertex] = 1;
    queue[++rear] = startVertex;

    printf("BFS Traversal: ");
    while (front <= rear) {
        int currentVertex = queue[front++];
        printf("%d ", currentVertex);

        for (AdjListNode* crawl = graph->array[currentVertex].head; crawl != NULL; crawl = crawl->next) {
            int adjVertex = crawl->dest;
            if (!visited[adjVertex]) {
                visited[adjVertex] = 1;
                queue[++rear] = adjVertex;
                if (adjVertex == searchVertex) {
                    printf("\nVertex %d found in BFS.\n", searchVertex);
                    found = 1;
                    break;
                }
            }
        }
        if (found) break;
    }

    if (!found)
        printf("\nVertex %d not found in BFS.\n", searchVertex);
}

void DFS(Graph* graph, int startVertex, int searchVertex) {
    int V = graph->V;
    int visited[V];
    for (int i = 0; i < V; i++) {
        visited[i] = 0;
    }

    printf("DFS Traversal: ");
    DFSUtil(graph, startVertex, visited, searchVertex);
}

void DFSUtil(Graph* graph, int v, int visited[], int searchVertex) {
    visited[v] = 1;
    printf("%d ", v);

    if (v == searchVertex) {
        printf("\nVertex %d found in DFS.\n", searchVertex);
        return;
    }

    for (AdjListNode* crawl = graph->array[v].head; crawl != NULL; crawl = crawl->next) {
        if (!visited[crawl->dest]) {
            DFSUtil(graph, crawl->dest, visited, searchVertex);
            if (visited[searchVertex]) return; // If found, stop further search
        }
    }
}
