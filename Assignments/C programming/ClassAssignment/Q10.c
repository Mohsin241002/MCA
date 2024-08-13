#include <stdio.h>
#include <string.h>

void extractSubstring(char str[], int start, int length) {
    char substring[100];
    strncpy(substring, str + start, length);
    substring[length] = '\0';
    printf("Substring: %s\n", substring);
}

int main() {
    char str[100];
    int start, length;

    printf("Enter a string: ");
    fgets(str, sizeof(str), stdin);

    printf("Enter the start index and length of substring: ");
    scanf("%d %d", &start, &length);

    extractSubstring(str, start, length);

    return 0;
}
