// Write a program in C Program to Read a Matrix and Print Diagonals

#include <stdio.h>

#define MAX 100

void printDiagonals(int matrix[MAX][MAX], int n)
{
    printf("Main Diagonal:\n");
    for (int i = 0; i < n; i++)
    {
        printf("%d ", matrix[i][i]);
    }
    printf("\n");

    printf("Anti-Diagonal:\n");
    for (int i = 0; i < n; i++)
    {
        printf("%d ", matrix[i][n - i - 1]);
    }
    printf("\n");
}

int main()
{
    int matrix[MAX][MAX];
    int n;

    printf("Enter the size of the matrix (n x n): ");
    scanf("%d", &n);

    printf("Enter the elements of the matrix:\n");
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
        {
            scanf("%d", &matrix[i][j]);
        }
    }

    printDiagonals(matrix, n);

    return 0;
}
