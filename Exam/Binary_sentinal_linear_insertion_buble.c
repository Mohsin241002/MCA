#include <stdio.h>
#include <limits.h> // For INT_MAX

// Function Prototypes
void linearSearch(int arr[], int n, int key);
void sentinelSearch(int arr[], int n, int key);
void binarySearch(int arr[], int n, int key);
void bubbleSort(int arr[], int n);
void insertionSort(int arr[], int n);
void displayArray(int arr[], int n);
int getInt();

int main() {
    int choice, n, key;
    printf("Enter the number of elements in the array: ");
    n = getInt();

    // Input validation for array size
    if (n <= 0) {
        printf("Invalid array size.\n");
        return 1;
    }

    int arr[n];
    printf("Enter the elements of the array: \n");
    for (int i = 0; i < n; i++) {
        arr[i] = getInt();
    }

    while (1) {
        printf("\nMenu:\n");
        printf("1. Linear Search\n");
        printf("2. Sentinel Search\n");
        printf("3. Binary Search\n");
        printf("4. Bubble Sort\n");
        printf("5. Insertion Sort\n");
        printf("6. Exit\n");
        printf("Enter your choice: ");
        choice = getInt();

        switch (choice) {
            case 1:
                printf("Enter the key to search: ");
                key = getInt();
                linearSearch(arr, n, key);
                break;
            case 2:
                printf("Enter the key to search: ");
                key = getInt();
                sentinelSearch(arr, n, key);
                break;
            case 3:
                bubbleSort(arr, n); // Ensuring the array is sorted for binary search
                printf("Enter the key to search: ");
                key = getInt();
                binarySearch(arr, n, key);
                break;
            case 4:
                bubbleSort(arr, n);
                displayArray(arr, n);
                break;
            case 5:
                insertionSort(arr, n);
                displayArray(arr, n);
                break;
            case 6:
                printf("Exiting program.\n");
                return 0;
            default:
                printf("Invalid choice. Please try again.\n");
        }
    }
}

// Safe integer input
int getInt() {
    int num;
    while (scanf("%d", &num) != 1) {
        printf("Invalid input. Please enter a valid integer: ");
        // Clear the input buffer
        while (getchar() != '\n');
    }
    return num;
}

// Search and Sort functions as defined previously...

void linearSearch(int arr[], int n, int key) {
    for (int i = 0; i < n; i++) {
        if (arr[i] == key) {
            printf("Key %d found at index %d.\n", key, i);
            return;
        }
    }
    printf("Key %d not found.\n", key);
}

void sentinelSearch(int arr[], int n, int key) {
    int last = arr[n-1];
    arr[n-1] = key;
    int i = 0;
    while (arr[i] != key) {
        i++;
    }
    arr[n-1] = last;
    if (i < n-1 || arr[n-1] == key) {
        printf("Key %d found at index %d.\n", key, i);
    } else {
        printf("Key %d not found.\n", key);
    }
}

void binarySearch(int arr[], int n, int key) {
    int low = 0, high = n - 1, mid;
    while (low <= high) {
        mid = (low + high) / 2;
        if (arr[mid] == key) {
            printf("Key %d found at index %d.\n", key, mid);
            return;
        } else if (arr[mid] < key) {
            low = mid + 1;
        } else {
            high = mid - 1;
        }
    }
    printf("Key %d not found.\n", key);
}

// Sorting functions as previously defined...

void bubbleSort(int arr[], int n) {
    int temp;
    for (int i = 0; i < n-1; i++) {
        for (int j = 0; j < n-i-1; j++) {
            if (arr[j] > arr[j+1]) {
                temp = arr[j];
                arr[j] = arr[j+1];
                arr[j+1] = temp;
            }
        }
    }
}

void insertionSort(int arr[], int n) {
    int i, key, j;
    for (i = 1; i < n; i++) {
        key = arr[i];
        j = i - 1;
        while (j >= 0 && arr[j] > key) {
            arr[j + 1] = arr[j];
            j = j - 1;
        }
        arr[j + 1] = key;
    }
}

void displayArray(int arr[], int n) {
    printf("Sorted array: ");
    for (int i = 0; i < n; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}
