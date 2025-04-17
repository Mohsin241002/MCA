#include <stdio.h>
#include <stdlib.h>
#include <limits.h>

#define MAX 50  // Maximum number of gaming nodes (vertices)

/* ----------- Disjoint Set (Union-Find) for Kruskal ----------- */
int parent[MAX]; 
int rankArr[MAX];

/* Initialize Disjoint Sets (Union-Find) */
void makeSet(int n) {
    for(int i = 0; i < n; i++) {
        parent[i] = i;
        rankArr[i] = 0;
    }
}

/* Find with path compression */
int findSet(int v) {
    if (parent[v] != v) {
        parent[v] = findSet(parent[v]);
    }
    return parent[v];
}

/* Union by rank */
void unionSet(int root1, int root2) {
    if (rankArr[root1] < rankArr[root2]) {
        parent[root1] = root2;
    } else if (rankArr[root1] > rankArr[root2]) {
        parent[root2] = root1;
    } else {
        parent[root2] = root1;
        rankArr[root1]++;
    }
}

/* ----------- Data Structures for MST and Graph ----------- */
int graph[MAX][MAX];  // Adjacency Matrix
int numNodes;         // Number of gaming nodes
int numEdges;         // Number of edges

/* Function to display the graph as an adjacency matrix */
void displayGraph() {
    if (numNodes <= 0) {
        printf("No valid gaming nodes.\n");
        return;
    }
    printf("\nCurrent Graph (Adjacency Matrix):\n    ");
    for(int i = 0; i < numNodes; i++) {
        printf("N%d ", i);
    }
    printf("\n");
    for(int i = 0; i < numNodes; i++) {
        printf("N%d ", i);
        for(int j = 0; j < numNodes; j++) {
            if (graph[i][j] == INT_MAX) {
                printf("INF ");
            } else {
                printf("%3d ", graph[i][j]);
            }
        }
        printf("\n");
    }
}

/* ----------- Prim's Algorithm ----------- */
void primMST() {
    if (numNodes <= 0) {
        printf("Graph not initialized. Please build the graph first.\n");
        return;
    }
    
    int key[MAX];          // Used to pick minimum weight edge
    int mstSet[MAX];       // To track if node is included in MST
    int parentPrim[MAX];   // To store MST
    
    // Initialize all keys as infinite and mstSet[] as false
    for(int i = 0; i < numNodes; i++) {
        key[i] = INT_MAX;
        mstSet[i] = 0;
        parentPrim[i] = -1;
    }

    // Start with node 0 (arbitrary choice)
    key[0] = 0;
    parentPrim[0] = -1; // 0 is the root of the MST

    // The MST will have numNodes vertices
    for(int count = 0; count < numNodes - 1; count++) {
        // Find the minimum key vertex from the set of vertices not yet included in MST
        int minKey = INT_MAX;
        int u = -1;

        for(int v = 0; v < numNodes; v++) {
            if (!mstSet[v] && key[v] < minKey) {
                minKey = key[v];
                u = v;
            }
        }

        // Add the picked vertex to the MST set
        mstSet[u] = 1;

        // Update key value and parent index of the adjacent vertices of the picked vertex.
        for(int w = 0; w < numNodes; w++) {
            // Update the key only if graph[u][w] is smaller than key[w]
            if (graph[u][w] != INT_MAX && !mstSet[w] && graph[u][w] < key[w]) {
                parentPrim[w] = u;
                key[w] = graph[u][w];
            }
        }
    }

    // Print the MST
    printf("\nMST using Prim's Algorithm:\n");
    printf("Edge   Weight\n");
    int totalWeight = 0;
    for(int i = 1; i < numNodes; i++) {
        if (parentPrim[i] != -1) {
            printf("%d - %d    %d\n", parentPrim[i], i, graph[i][parentPrim[i]]);
            totalWeight += graph[i][parentPrim[i]];
        }
    }
    printf("Total Weight of MST (Prim): %d\n", totalWeight);
}

/* ----------- Kruskal's Algorithm ----------- */

// Structure to store edges
typedef struct Edge {
    int src;
    int dest;
    int weight;
} Edge;

/* Compare function for qsort (sort edges by weight) */
int compareEdges(const void *a, const void *b) {
    Edge *e1 = (Edge *)a;
    Edge *e2 = (Edge *)b;
    return e1->weight - e2->weight;
}

void kruskalMST() {
    if (numNodes <= 0) {
        printf("Graph not initialized. Please build the graph first.\n");
        return;
    }

    // Build a list of all edges
    Edge edges[MAX * MAX];
    int k = 0; // index for edges
    for(int i = 0; i < numNodes; i++) {
        for(int j = i + 1; j < numNodes; j++) {
            if (graph[i][j] != INT_MAX) {
                edges[k].src = i;
                edges[k].dest = j;
                edges[k].weight = graph[i][j];
                k++;
            }
        }
    }

    // Sort edges in non-decreasing order of weight
    qsort(edges, k, sizeof(Edge), compareEdges);

    // Initialize disjoint sets
    makeSet(numNodes);

    printf("\nMST using Kruskal's Algorithm:\n");
    printf("Edge   Weight\n");

    int eCount = 0;      // Count how many edges included in MST
    int totalWeight = 0; // MST total weight
    for(int i = 0; i < k && eCount < numNodes - 1; i++) {
        int root1 = findSet(edges[i].src);
        int root2 = findSet(edges[i].dest);

        // If including this edge doesn't cause a cycle, include it
        if (root1 != root2) {
            printf("%d - %d    %d\n", edges[i].src, edges[i].dest, edges[i].weight);
            unionSet(root1, root2);
            totalWeight += edges[i].weight;
            eCount++;
        }
    }
    printf("Total Weight of MST (Kruskal): %d\n", totalWeight);
}

/* ----------- Helper to Build/Update the Graph ----------- */
void buildGraph() {
    int i, j, w;
    
    // Input number of nodes
    int tempNodes;
    do {
        printf("Enter number of gaming nodes (1 - %d): ", MAX);
        scanf("%d", &tempNodes);
        if(tempNodes <= 0 || tempNodes > MAX) {
            printf("Invalid number of nodes. Please re-enter.\n");
        }
    } while(tempNodes <= 0 || tempNodes > MAX);

    numNodes = tempNodes;

    // Initialize all edges to INT_MAX (represents no direct edge)
    for(i = 0; i < numNodes; i++) {
        for(j = 0; j < numNodes; j++) {
            if(i == j) {
                graph[i][j] = 0;  // distance to itself is 0
            } else {
                graph[i][j] = INT_MAX;
            }
        }
    }

    // Input number of edges
    // Max possible edges in an undirected graph is n*(n-1)/2
    int maxEdges = numNodes*(numNodes-1)/2;
    do {
        printf("Enter number of edges (0 - %d): ", maxEdges);
        scanf("%d", &numEdges);
        if(numEdges < 0 || numEdges > maxEdges) {
            printf("Invalid number of edges. Please re-enter.\n");
        }
    } while(numEdges < 0 || numEdges > maxEdges);

    // Input edge information
    for(i = 0; i < numEdges; i++) {
        int src, dest;
        do {
            printf("\nEdge %d:\n", i+1);
            printf("Enter source node (0 to %d): ", numNodes - 1);
            scanf("%d", &src);
            printf("Enter destination node (0 to %d): ", numNodes - 1);
            scanf("%d", &dest);
            if(src < 0 || src >= numNodes || dest < 0 || dest >= numNodes) {
                printf("Invalid node indices. Please re-enter.\n");
            }
        } while(src < 0 || src >= numNodes || dest < 0 || dest >= numNodes);

        // Weight
        do {
            printf("Enter weight of edge (non-negative integer): ");
            scanf("%d", &w);
            if(w < 0) {
                printf("Invalid weight. Please re-enter.\n");
            }
        } while(w < 0);

        // Since this might be an undirected graph in MST context, set both sides
        graph[src][dest] = w;
        graph[dest][src] = w;
    }

    printf("\nGraph updated!\n");
}

/* ----------- Main (Menu-Driven) ----------- */
int main() {
    int choice;
    
    // Initialize graph to "empty"
    numNodes = 0;
    numEdges = 0;
    for(int i = 0; i < MAX; i++){
        for(int j = 0; j < MAX; j++){
            graph[i][j] = INT_MAX;
        }
    }

    do {
        printf("\n===== GAMING INFO SYSTEM (MST) =====\n");
        printf("1. Build/Update Gaming Graph\n");
        printf("2. Display Graph\n");
        printf("3. Find MST using Prim's Algorithm\n");
        printf("4. Find MST using Kruskal's Algorithm\n");
        printf("5. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch(choice) {
            case 1:
                buildGraph();
                break;
            case 2:
                displayGraph();
                break;
            case 3:
                primMST();
                break;
            case 4:
                kruskalMST();
                break;
            case 5:
                printf("Exiting the program...\n");
                break;
            default:
                printf("Invalid choice. Please try again.\n");
        }
    } while(choice != 5);

    return 0;
}
