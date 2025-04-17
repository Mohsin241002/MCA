// Write a program in C Program to print lower diagonal of a matrix

#include <stdio.h>

#define MAX 100

void printLowerDiagonal(int matrix[MAX][MAX], int n)
{
    printf("Lower Diagonal:\n");
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j <= i; j++)
        {
            printf("%d ", matrix[i][j]);
        }
        printf("\n");
    }
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

    printLowerDiagonal(matrix, n);

    return 0;
}
