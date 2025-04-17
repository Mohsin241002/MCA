// Write a program in C to count the total number of alphabets, digits and special characters in
// a string.

#include <stdio.h>
#include <ctype.h>

int main()
{
    char str[100];
    int alphabets = 0, digits = 0, specials = 0;

    printf("Enter a string: ");

    scanf("%99[^\n]", str);

    for (int i = 0; str[i] != '\0'; i++)
    {
        if (isalpha(str[i]))
        {
            alphabets++;
        }
        else if (isdigit(str[i]))
        {
            digits++;
        }
        else if (!isspace(str[i]))
        {
            specials++;
        }
    }

    // Print the counts
    printf("Alphabets: %d\n", alphabets);
    printf("Digits: %d\n", digits);
    printf("Special Characters: %d\n", specials);

    return 0;
}
