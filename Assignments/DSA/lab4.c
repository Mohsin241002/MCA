#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Node structure for queue representing a player
struct Player {
    char name[50];
    int level;    // Player level, higher level indicates more experienced
    int priority; // For priority queue (e.g., VIP status)
    struct Player* next;
};

// Function to create a new player node
struct Player* createPlayer(const char* name, int level, int priority) {
    struct Player* newPlayer = (struct Player*)malloc(sizeof(struct Player));
    strcpy(newPlayer->name, name);
    newPlayer->level = level;
    newPlayer->priority = priority;
    newPlayer->next = NULL;
    return newPlayer;
}

// Queue structure to represent waiting line for game actions
struct Queue {
    struct Player *front, *rear;
};

// Initialize a new empty queue
struct Queue* createQueue() {
    struct Queue* q = (struct Queue*)malloc(sizeof(struct Queue));
    q->front = q->rear = NULL;
    return q;
}

// Function to enqueue a player for regular game actions
void enqueue(struct Queue* q, const char* name, int level) {
    struct Player* temp = createPlayer(name, level, 0);
    if (q->rear == NULL) {
        q->front = q->rear = temp;
        return;
    }
    q->rear->next = temp;
    q->rear = temp;
}

// Function to dequeue a player for regular game actions
void dequeue(struct Queue* q) {
    if (q->front == NULL) {
        printf("No players waiting.\n");
        return;
    }
    struct Player* temp = q->front;
    printf("Player %s with level %d is taking action.\n", temp->name, temp->level);
    q->front = q->front->next;
    if (q->front == NULL)
        q->rear = NULL;
    free(temp);
}

// Function to enqueue a player in priority queue (VIP players)
void priorityEnqueue(struct Queue* q, const char* name, int level, int priority) {
    struct Player* temp = createPlayer(name, level, priority);
    if (q->front == NULL || q->front->priority > priority) {
        temp->next = q->front;
        q->front = temp;
    } else {
        struct Player* start = q->front;
        while (start->next != NULL && start->next->priority <= priority)
            start = start->next;
        temp->next = start->next;
        start->next = temp;
    }
}

// Function to dequeue a player from priority queue
void priorityDequeue(struct Queue* q) {
    if (q->front == NULL) {
        printf("No VIP players waiting.\n");
        return;
    }
    struct Player* temp = q->front;
    printf("VIP Player %s with level %d is taking priority action.\n", temp->name, temp->level);
    q->front = q->front->next;
    free(temp);
}

// Display queue for regular and priority queues
void display(struct Queue* q) {
    struct Player* temp = q->front;
    if (temp == NULL) {
        printf("No players in queue.\n");
        return;
    }
    printf("Players in queue:\n");
    while (temp != NULL) {
        printf("Player: %s, Level: %d", temp->name, temp->level);
        if (temp->priority > 0)
            printf(" (VIP Priority: %d)", temp->priority);
        printf("\n");
        temp = temp->next;
    }
}

int main() {
    struct Queue* actionQueue = createQueue();
    struct Queue* vipQueue = createQueue();

    printf("Adding regular players to action queue:\n");
    display(actionQueue);
    enqueue(actionQueue, "PlayerA", 5);
    enqueue(actionQueue, "PlayerB", 3);
    enqueue(actionQueue, "PlayerC", 4);
    display(actionQueue);

    printf("\nDequeuing regular player:\n");
    dequeue(actionQueue);
    display(actionQueue);

    printf("\nAdding VIP players to priority queue:\n");
    priorityEnqueue(vipQueue, "VIP_Player1", 10, 1);
    priorityEnqueue(vipQueue, "VIP_Player2", 8, 2);
    priorityEnqueue(vipQueue, "VIP_Player3", 7, 1);
    display(vipQueue);

    printf("\nDequeuing VIP player with highest priority:\n");
    priorityDequeue(vipQueue);
    display(vipQueue);

    return 0;
}
