#include <stdio.h>
#include <stdlib.h>
#include <time.h> // For clock() function to measure efficiency

// Function Prototypes
void mergeSort(int arr[], int l, int r);
void merge(int arr[], int l, int m, int r);
void quickSort(int arr[], int low, int high);
int partition(int arr[], int low, int high);
void displayArray(int arr[], int n);
int getInt();
void copyArray(int src[], int dest[], int n);

int main() {
    int n, choice;
    printf("Enter the number of elements in the array: ");
    n = getInt();

    if (n <= 0) {
        printf("Invalid array size.\n");
        return 1;
    }

    int arr[n], temp[n];
    printf("Enter the elements of the array:\n");
    for (int i = 0; i < n; i++) {
        arr[i] = getInt();
    }

    while (1) {
        printf("\nMenu:\n");
        printf("1. Merge Sort\n");
        printf("2. Quick Sort\n");
        printf("3. Compare Sorts\n");
        printf("4. Exit\n");
        printf("Enter your choice: ");
        choice = getInt();

        switch (choice) {
            case 1:
                copyArray(arr, temp, n); // Copy original array to temp
                mergeSort(temp, 0, n - 1);
                displayArray(temp, n);
                break;
            case 2:
                copyArray(arr, temp, n); // Copy original array to temp
                quickSort(temp, 0, n - 1);
                displayArray(temp, n);
                break;
            case 3:
                compareSorts(arr, n);
                break;
            case 4:
                printf("Exiting program.\n");
                return 0;
            default:
                printf("Invalid choice. Please try again.\n");
        }
    }
}

void mergeSort(int arr[], int l, int r) {
    if (l < r) {
        int m = l + (r - l) / 2; // Same as (l+r)/2 but avoids overflow
        mergeSort(arr, l, m);
        mergeSort(arr, m + 1, r);
        merge(arr, l, m, r);
    }
}

void merge(int arr[], int l, int m, int r) {
    int i, j, k;
    int n1 = m - l + 1;
    int n2 = r - m;
    int L[n1], R[n2];

    for (i = 0; i < n1; i++)
        L[i] = arr[l + i];
    for (j = 0; j < n2; j++)
        R[j] = arr[m + 1 + j];

    i = 0; j = 0; k = l;
    while (i < n1 && j < n2) {
        if (L[i] <= R[j]) {
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

void quickSort(int arr[], int low, int high) {
    if (low < high) {
        int pi = partition(arr, low, high);
        quickSort(arr, low, pi - 1);
        quickSort(arr, pi + 1, high);
    }
}

int partition(int arr[], int low, int high) {
    int pivot = arr[high];
    int i = (low - 1);

    for (int j = low; j <= high - 1; j++) {
        if (arr[j] < pivot) {
            i++;
            int t = arr[i];
            arr[i] = arr[j];
            arr[j] = t;
        }
    }
    int t = arr[i + 1];
    arr[i + 1] = arr[high];
    arr[high] = t;
    return (i + 1);
}

void displayArray(int arr[], int n) {
    printf("Array: ");
    for (int i = 0; i < n; i++)
        printf("%d ", arr[i]);
    printf("\n");
}

int getInt() {
    int num;
    while (scanf("%d", &num) != 1) {
        printf("Invalid input. Please enter a valid integer: ");
        while (getchar() != '\n');
    }
    return num;
}

void copyArray(int src[], int dest[], int n) {
    for (int i = 0; i < n; i++)
        dest[i] = src[i];
}

void compareSorts(int arr[], int n) {
    int temp[n];
    clock_t start, end;
    double cpu_time_used;

    copyArray(arr, temp, n);
    start = clock();
    mergeSort(temp, 0, n - 1);
    end = clock();
    cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;
    printf("Merge Sort took %f seconds.\n", cpu_time_used);

    copyArray(arr, temp, n);
    start = clock();
    quickSort(temp, 0, n - 1);
    end = clock();
    cpu_time_used = ((double) (end - start)) / CLOCKS_PER_SEC;
    printf("Quick Sort took %f seconds.\n", cpu_time_used);
}
