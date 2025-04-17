#include <stdio.h>
#include <stdlib.h>
#include <string.h>


/**
 * This program manages a collection of games, allowing the user to input game details,
 * view the list of games, and sort them by score using either Merge Sort or Quick Sort.
 * The user can choose from a menu to perform these actions repeatedly until they decide to exit.
 * 
 * The program defines a `Game` structure to store game information, including name, score,
 * and release year. It provides functions for sorting the games using Merge Sort and Quick Sort,
 * printing the list of games, and inputting game data from the user.
 */
// Structure to store gaming information

typedef struct {
    char name[50];
    int score;
    int releaseYear;
} Game;

// Function prototypes
void mergeSort(Game arr[], int l, int r);
void merge(Game arr[], int l, int m, int r);
void quickSort(Game arr[], int low, int high);
int partition(Game arr[], int low, int high);
void printGames(Game arr[], int size);
void inputGames(Game arr[], int *size);

int main() {
    Game games[100];
    int size = 0;
    int choice;

    printf("Enter the number of games: ");
    scanf("%d", &size);
    inputGames(games, &size);

    while (1) {
        printf("\nMenu:\n");
        printf("1. View Games\n");
        printf("2. Apply Merge Sort by Score\n");
        printf("3. Apply Quick Sort by Score\n");
        printf("4. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                printGames(games, size);
                break;
            case 2: {
                Game tempGames[100];
                memcpy(tempGames, games, sizeof(Game) * size);
                mergeSort(tempGames, 0, size - 1);
                printf("\nGames sorted by Merge Sort:\n");
                printGames(tempGames, size);
                break;
            }
            case 3: {
                Game tempGames[100];
                memcpy(tempGames, games, sizeof(Game) * size);
                quickSort(tempGames, 0, size - 1);
                printf("\nGames sorted by Quick Sort:\n");
                printGames(tempGames, size);
                break;
            }
            case 4:
                exit(0);
            default:
                printf("Invalid choice! Please try again.\n");
        }
    }

    return 0;
}

// Merge Sort functions
// Time Complexity: O(n log n) in all cases (best, average, worst)


/**
 * Sorts an array of Game structures by score using the Merge Sort algorithm.
 *
 * This function recursively divides the array into two halves, sorts each half,
 * and then merges the sorted halves. The merge operation ensures that the
 * resulting array is sorted by the score of each Game.
 *
 * @param arr The array of Game structures to be sorted.
 * @param l The starting index of the array segment to be sorted.
 * @param r The ending index of the array segment to be sorted.
 */
void mergeSort(Game arr[], int l, int r) {
    if (l < r) {
        int m = l + (r - l) / 2;
        mergeSort(arr, l, m);
        mergeSort(arr, m + 1, r);
        merge(arr, l, m, r);
    }
}

void merge(Game arr[], int l, int m, int r) {
    int n1 = m - l + 1;
    int n2 = r - m;
    Game L[n1], R[n2];

    for (int i = 0; i < n1; i++)
        L[i] = arr[l + i];
    for (int j = 0; j < n2; j++)
        R[j] = arr[m + 1 + j];

    int i = 0, j = 0, k = l;
    while (i < n1 && j < n2) {
        if (L[i].score <= R[j].score) {
            arr[k] = L[i];
            i++;
        } else {
            arr[k] = R[j];
            j++;
        }
        k++;
    }

    while (i < n1) {
        arr[k] = L[i];
        i++;
        k++;
    }

    while (j < n2) {
        arr[k] = R[j];
        j++;
        k++;
    }
}

// Quick Sort functions
// Time Complexity: Best/Average Case: O(n log n), Worst Case: O(n^2)


/**
 * Sorts an array of Game structures by score using the Quick Sort algorithm.
 *
 * This function implements the Quick Sort algorithm to sort an array of Game
 * structures in ascending order based on their score. It recursively partitions
 * the array around a pivot element and sorts the partitions.
 *
 * @param arr The array of Game structures to be sorted.
 * @param low The starting index of the array segment to be sorted.
 * @param high The ending index of the array segment to be sorted.
 */
void quickSort(Game arr[], int low, int high) {
    if (low < high) {
        int pi = partition(arr, low, high);
        quickSort(arr, low, pi - 1);
        quickSort(arr, pi + 1, high);
    }
}

int partition(Game arr[], int low, int high) {
    int pivot = arr[high].score;
    int i = (low - 1);

    for (int j = low; j <= high - 1; j++) {
        if (arr[j].score <= pivot) {
            i++;
            Game temp = arr[i];
            arr[i] = arr[j];
            arr[j] = temp;
        }
    }
    Game temp = arr[i + 1];
    arr[i + 1] = arr[high];
    arr[high] = temp;
    return (i + 1);
}

// Utility function to print game data
void printGames(Game arr[], int size) {
    printf("\n%-20s %-10s %-10s\n", "Name", "Score", "Year");
    printf("----------------------------------------\n");
    for (int i = 0; i < size; i++) {
        printf("%-20s %-10d %-10d\n", arr[i].name, arr[i].score, arr[i].releaseYear);
    }
}

// Function to input game data
void inputGames(Game arr[], int *size) {
    for (int i = 0; i < *size; i++) {
        printf("\nEnter details for game %d:\n", i + 1);
        printf("Name: ");
        scanf(" %[^\n]%*c", arr[i].name); // Reads string with spaces
        printf("Score: ");
        scanf("%d", &arr[i].score);
        printf("Release Year: ");
        scanf("%d", &arr[i].releaseYear);
    }
}
