// Write a program in C to find the length of a string without using library functions.

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

    printf("Enter a string: ");
    scanf("%99[^\n]", str);

    printf("Length of the string: %d\n", stringLength(str));

    return 0;
}
