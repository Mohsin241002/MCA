#include <stdio.h>
#include <string.h>

#define MAX 100

// Define a structure for game information
typedef struct {
    int gameID;
    char gameName[50];
} Game;

// Function to display game data
void displayGames(Game games[], int n) {
    printf("Game List:\n");
    for (int i = 0; i < n; i++) {
        printf("Game ID: %d, Game Name: %s\n", games[i].gameID, games[i].gameName);
    }
}

// Linear Search
int linearSearch(Game games[], int n, int targetID) {
    for (int i = 0; i < n; i++) {
        if (games[i].gameID == targetID) {
            return i; // Found
        }
    }
    return -1; // Not found
}

// Sentinel Search
int sentinelSearch(Game games[], int n, int targetID) {
    int last = games[n - 1].gameID;
    games[n - 1].gameID = targetID; // Set sentinel

    int i = 0;
    while (games[i].gameID != targetID) {
        i++;
    }

    games[n - 1].gameID = last; // Restore the last element

    if (i < n - 1 || games[n - 1].gameID == targetID) {
        return i; // Found
    }
    return -1; // Not found
}

// Binary Search (requires sorted array)
int binarySearch(Game games[], int n, int targetID) {
    int left = 0, right = n - 1;
    while (left <= right) {
        int mid = (left + right) / 2;
        if (games[mid].gameID == targetID) {
            return mid;
        } else if (games[mid].gameID < targetID) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }
    return -1; // Not found
}

// Bubble Sort
void bubbleSort(Game games[], int n, int *comparisons, int *transfers) {
    for (int i = 0; i < n - 1; i++) {
        for (int j = 0; j < n - i - 1; j++) {
            (*comparisons)++;
            if (games[j].gameID > games[j + 1].gameID) {
                Game temp = games[j];
                games[j] = games[j + 1];
                games[j + 1] = temp;
                (*transfers)++;
            }
        }
    }
}

// Insertion Sort
void insertionSort(Game games[], int n, int *comparisons, int *transfers) {
    for (int i = 1; i < n; i++) {
        Game key = games[i];
        int j = i - 1;

        while (j >= 0 && games[j].gameID > key.gameID) {
            (*comparisons)++;
            games[j + 1] = games[j];
            (*transfers)++;
            j--;
        }
        games[j + 1] = key;
    }
}

// Main function
int main() {
    Game games[MAX];
    int n, targetID, choice, result;
    int comparisons = 0, transfers = 0;

    // Sample data
    printf("Enter number of games: ");
    scanf("%d", &n);

    for (int i = 0; i < n; i++) {
        printf("Enter Game ID and Game Name for game %d: ", i + 1);
        scanf("%d %[^\n]s", &games[i].gameID, games[i].gameName);
    }

    int exitFlag = 0; // To control the loop

    while (!exitFlag) {
        printf("\nChoose an option:\n");
        printf("1. Linear Search\n");
        printf("2. Sentinel Search\n");
        printf("3. Binary Search\n");
        printf("4. Display Games\n");
        printf("5. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        if (choice == 3) {
            // Sort data before binary search
            printf("Choose sorting method:\n1. Bubble Sort\n2. Insertion Sort\n");
            int sortChoice;
            scanf("%d", &sortChoice);

            comparisons = 0;
            transfers = 0;

            if (sortChoice == 1) {
                bubbleSort(games, n, &comparisons, &transfers);
            } else {
                insertionSort(games, n, &comparisons, &transfers);
            }
            printf("Sorted games for Binary Search:\n");
            displayGames(games, n);
        }

        if (choice >= 1 && choice <= 3) {
            printf("Enter Game ID to search: ");
            scanf("%d", &targetID);

            switch (choice) {
                case 1:
                    result = linearSearch(games, n, targetID);
                    break;
                case 2:
                    result = sentinelSearch(games, n, targetID);
                    break;
                case 3:
                    result = binarySearch(games, n, targetID);
                    break;
            }

            if (result != -1) {
                printf("Game found at index %d: Game ID: %d, Game Name: %s\n", result, games[result].gameID, games[result].gameName);
            } else {
                printf("Game not found.\n");
            }

            printf("Comparisons: %d, Data Transfers: %d\n", comparisons, transfers);
        } else if (choice == 4) {
            displayGames(games, n);
        } else if (choice == 5) {
            exitFlag = 1; // Set flag to exit the loop
        } else {
            printf("Invalid choice. Please try again.\n");
        }
    }

    printf("Exiting program.\n");
    return 0;
}
