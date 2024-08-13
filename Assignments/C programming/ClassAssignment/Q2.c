// Write a program in C Program to find multiplication of two matrices

#include <stdio.h>

#define MAX 100

void multiplyMatrices(int mat1[MAX][MAX], int mat2[MAX][MAX], int result[MAX][MAX], int row1, int col1, int row2, int col2)
{
    for (int i = 0; i < row1; i++)
    {
        for (int j = 0; j < col2; j++)
        {
            result[i][j] = 0;
            for (int k = 0; k < col1; k++)
            {
                result[i][j] += mat1[i][k] * mat2[k][j];
            }
        }
    }
}

int main()
{
    int mat1[MAX][MAX], mat2[MAX][MAX], result[MAX][MAX];
    int row1, col1, row2, col2;

    printf("Enter rows and columns for first matrix: ");
    scanf("%d %d", &row1, &col1);

    printf("Enter elements of first matrix:\n");
    for (int i = 0; i < row1; i++)
    {
        for (int j = 0; j < col1; j++)
        {
            scanf("%d", &mat1[i][j]);
        }
    }

    printf("Enter rows and columns for second matrix: ");
    scanf("%d %d", &row2, &col2);

    if (col1 != row2)
    {
        printf("Matrix multiplication not possible.\n");
        return 1;
    }

    printf("Enter elements of second matrix:\n");
    for (int i = 0; i < row2; i++)
    {
        for (int j = 0; j < col2; j++)
        {
            scanf("%d", &mat2[i][j]);
        }
    }

    multiplyMatrices(mat1, mat2, result, row1, col1, row2, col2);

    printf("Product of the matrices:\n");
    for (int i = 0; i < row1; i++)
    {
        for (int j = 0; j < col2; j++)
        {
            printf("%d ", result[i][j]);
        }
        printf("\n");
    }

    return 0;
}
