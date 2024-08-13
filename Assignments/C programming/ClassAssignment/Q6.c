// Write a program in C to print individual characters of a string in reverse order.

#include <stdio.h>

int stringLength(char str[])
{
    int length = 0;
    while (str[length] != '\0')
    {
        length++;
    }
    return length;
}

int main()
{
    char str[100];
    int length;

    printf("Enter a string: ");
    fgets(str, sizeof(str), stdin);

    length = stringLength(str);

    printf("String in reverse order:\n");
    for (int i = length - 1; i >= 0; i--)
    {
        printf("%c", str[i]);
    }
    printf("\n");

    return 0;
}
