#include <stdio.h>

#define MAX 100

void inputMatrix(int matrix[MAX][MAX], int rows, int cols, char name) {
    printf("Enter elements of matrix %c:\n", name);
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            printf("Element [%d][%d]: ", i + 1, j + 1);
            scanf("%d", &matrix[i][j]);
        }
    }
}

void printMatrix(int matrix[MAX][MAX], int rows, int cols) {
    printf("Matrix:\n");
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            printf("%d ", matrix[i][j]);
        }
        printf("\n");
    }
}

void rowSum(int matrix[MAX][MAX], int rows, int cols) {
    for (int i = 0; i < rows; i++) {
        int sum = 0;
        for (int j = 0; j < cols; j++) {
            sum += matrix[i][j];
        }
        printf("Sum of row %d: %d\n", i + 1, sum);
    }
}

void columnSum(int matrix[MAX][MAX], int rows, int cols) {
    for (int j = 0; j < cols; j++) {
        int sum = 0;
        for (int i = 0; i < rows; i++) {
            sum += matrix[i][j];
        }
        printf("Sum of column %d: %d\n", j + 1, sum);
    }
}

void matrixAddition(int A[MAX][MAX], int B[MAX][MAX], int rows, int cols) {
    int result[MAX][MAX];
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            result[i][j] = A[i][j] + B[i][j];
        }
    }
    printf("Matrix A + Matrix B:\n");
    printMatrix(result, rows, cols);
}

void matrixSubtraction(int A[MAX][MAX], int B[MAX][MAX], int rows, int cols) {
    int result[MAX][MAX];
    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            result[i][j] = A[i][j] - B[i][j];
        }
    }
    printf("Matrix A - Matrix B:\n");
    printMatrix(result, rows, cols);
}

void matrixMultiplication(int A[MAX][MAX], int B[MAX][MAX], int rowsA, int colsA, int colsB) {
    int result[MAX][MAX] = {0};
    for (int i = 0; i < rowsA; i++) {
        for (int j = 0; j < colsB; j++) {
            for (int k = 0; k < colsA; k++) {
                result[i][j] += A[i][k] * B[k][j];
            }
        }
    }
    printf("Matrix A * Matrix B:\n");
    printMatrix(result, rowsA, colsB);
}

void findMinMax(int matrix[MAX][MAX], int rows, int cols, int* min, int* max) {
    // Initialize min and max with the first element of the matrix
    *min = *max = matrix[0][0];

    for (int i = 0; i < rows; i++) {
        for (int j = 0; j < cols; j++) {
            if (matrix[i][j] < *min) {
                *min = matrix[i][j];
            }
            if (matrix[i][j] > *max) {
                *max = matrix[i][j];
            }
        }
    }
}

int main() {
    int A[MAX][MAX], B[MAX][MAX];
    int rowsA, colsA, rowsB, colsB;
    int choice, min, max;

    // Input matrix dimensions
    printf("Enter number of rows and columns for matrix A: ");
    scanf("%d %d", &rowsA, &colsA);

    printf("Enter number of rows and columns for matrix B: ");
    scanf("%d %d", &rowsB, &colsB);

    // Input matrices
    inputMatrix(A, rowsA, colsA, 'A');
    inputMatrix(B, rowsB, colsB, 'B');

    do {
        printf("\nMenu:\n");
        printf("1. Row Sum of Matrix A\n");
        printf("2. Column Sum of Matrix A\n");
        printf("3. Matrix Addition (A + B)\n");
        printf("4. Matrix Subtraction (A - B)\n");
        printf("5. Matrix Multiplication (A * B)\n");
        printf("6. Find Minimum and Maximum in Matrix A\n");
        printf("7. Exit\n");
        printf("Enter your choice: ");
        scanf("%d", &choice);

        switch (choice) {
            case 1:
                printf("\nRow sums for matrix A:\n");
                rowSum(A, rowsA, colsA);
                break;

            case 2:
                printf("\nColumn sums for matrix A:\n");
                columnSum(A, rowsA, colsA);
                break;

            case 3:
                if (rowsA == rowsB && colsA == colsB) {
                    printf("\nMatrix Addition (A + B):\n");
                    matrixAddition(A, B, rowsA, colsA);
                } else {
                    printf("\nMatrix addition not possible: Dimensions do not match.\n");
                }
                break;

            case 4:
                if (rowsA == rowsB && colsA == colsB) {
                    printf("\nMatrix Subtraction (A - B):\n");
                    matrixSubtraction(A, B, rowsA, colsA);
                } else {
                    printf("\nMatrix subtraction not possible: Dimensions do not match.\n");
                }
                break;

            case 5:
                if (colsA == rowsB) {
                    printf("\nMatrix Multiplication (A * B):\n");
                    matrixMultiplication(A, B, rowsA, colsA, colsB);
                } else {
                    printf("\nMatrix multiplication not possible: Invalid dimensions.\n");
                }
                break;

            case 6:
                findMinMax(A, rowsA, colsA, &min, &max);
                printf("Minimum element in Matrix A: %d\n", min);
                printf("Maximum element in Matrix A: %d\n", max);
                break;

            case 7:
                printf("Exiting...\n");
                break;

            default:
                printf("Invalid choice! Please try again.\n");
        }
    } while (choice != 7);

    return 0;
}
