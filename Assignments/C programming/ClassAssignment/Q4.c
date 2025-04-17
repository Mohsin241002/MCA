// Write a program in C program to interchange the rows in the matrix

#include <stdio.h>

#define MAX 100

void interchangeRows(int matrix[MAX][MAX], int row1, int row2, int n)
{
    for (int i = 0; i < n; i++)
    {
        int temp = matrix[row1][i];
        matrix[row1][i] = matrix[row2][i];
        matrix[row2][i] = temp;
    }
}

int main()
{
    int matrix[MAX][MAX];
    int n, row1, row2;

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

    printf("Enter the rows to interchange (0-indexed): ");
    scanf("%d %d", &row1, &row2);

    interchangeRows(matrix, row1, row2, n);

    printf("Matrix after interchanging rows:\n");
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < n; j++)
        {
            printf("%d ", matrix[i][j]);
        }
        printf("\n");
    }

    return 0;
}
