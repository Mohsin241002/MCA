// Create a C program to perform the following matrix manipulation operations on two
// matrices A and B of size N x M.
// 1. Row Sum Calculation:
// Write a function to calculate and print the sum of all elements in each row of matrix A.
// 2. Column Sum Calculation:
// Write a function to calculate and print the sum of all elements in each column of matrix A.
// 3. Matrix Addition:
// Write a function to add matrices A and B. Print the resultant matrix.
// 4. Matrix Subtraction:
// Write a function to subtract matrix B from matrix A. Print the resultant matrix.
// 5. Matrix Multiplication:
// Write a function to multiply matrix A with matrix B. Validate the matrix dimensions to ensure multiplication is possible.

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

int main() {
    int A[MAX][MAX], B[MAX][MAX];
    int rowsA, colsA, rowsB, colsB;

    // Input matrix dimensions
    printf("Enter number of rows and columns for matrix A: ");
    scanf("%d %d", &rowsA, &colsA);

    printf("Enter number of rows and columns for matrix B: ");
    scanf("%d %d", &rowsB, &colsB);

    // Input matrices
    inputMatrix(A, rowsA, colsA, 'A');
    inputMatrix(B, rowsB, colsB, 'B');

    // 1. Row Sum Calculation
    printf("\nRow sums for matrix A:\n");
    rowSum(A, rowsA, colsA);

    // 2. Column Sum Calculation
    printf("\nColumn sums for matrix A:\n");
    columnSum(A, rowsA, colsA);

    // 3. Matrix Addition (possible only if A and B have same dimensions)
    if (rowsA == rowsB && colsA == colsB) {
        printf("\nMatrix Addition (A + B):\n");
        matrixAddition(A, B, rowsA, colsA);
    } else {
        printf("\nMatrix addition not possible: Dimensions do not match.\n");
    }

    // 4. Matrix Subtraction (possible only if A and B have same dimensions)
    if (rowsA == rowsB && colsA == colsB) {
        printf("\nMatrix Subtraction (A - B):\n");
        matrixSubtraction(A, B, rowsA, colsA);
    } else {
        printf("\nMatrix subtraction not possible: Dimensions do not match.\n");
    }

    // 5. Matrix Multiplication (possible only if colsA == rowsB)
    if (colsA == rowsB) {
        printf("\nMatrix Multiplication (A * B):\n");
        matrixMultiplication(A, B, rowsA, colsA, colsB);
    } else {
        printf("\nMatrix multiplication not possible: Invalid dimensions.\n");
    }

    return 0;
}
