#include <stdio.h>
#include <string.h>

// Function prototype for a function that takes a pointer as a parameter
void displayMessage(char *message);

// Function prototype for a function pointer example
int add(int a, int b);

int main() {
    // Simple pointer example
    int x = 10;
    int *ptr = &x;

    printf("Value of x: %d\n", x);
    printf("Address of x: %p\n", &x);
    printf("Value of ptr (Address of x): %p\n", (void*)ptr);
    printf("Value pointed to by ptr: %d\n\n", *ptr);

    // Pointer to pointer example
    int **pptr = &ptr;

    printf("Value of pptr (Address of ptr): %p\n", pptr);
    printf("Value pointed to by pptr (Value of ptr): %p\n", *pptr);
    printf("Value pointed to by the pointer pointed to by pptr: %d\n\n", **pptr);

    // Array pointer example
    int arr[5] = {1, 2, 3, 4, 5};
    int *arrPtr = arr; // or int *arrPtr = &arr[0];

    for (int i = 0; i < 5; i++) {
        printf("arr[%d] = %d, *(arrPtr + %d) = %d\n", i, arr[i], i, *(arrPtr + i));
    }
    printf("\n");

    // Function pointer example
    int (*funcPtr)(int, int) = add;
    int result = funcPtr(5, 3);
    printf("Result of add(5, 3) using function pointer: %d\n\n", result);

    // String concept with examples of string.h functions
    
    char str1[20] = "Hello";
    char str2[20] = "World";
    char str3[40];

    // strcpy example
    strcpy(str3, str1);
    printf("strcpy: str3 = %s\n", str3);

    // strcat example
    strcat(str3, " ");
    strcat(str3, str2);
    printf("strcat: str3 = %s\n", str3);

    // strlen example
    printf("strlen: Length of str3 = %lu\n", strlen(str3));

    // strcmp example
    int cmp = strcmp(str1, str2);
    printf("strcmp: Comparison of str1 and str2 = %d\n", cmp);

    // strchr example
    char *charPtr = strchr(str3, 'W');
    if (charPtr) {
        printf("strchr: 'W' found in str3 at position = %ld\n", charPtr - str3);
    } else {
        printf("strchr: 'W' not found in str3\n");
    }

    // strstr example
    char *substrPtr = strstr(str3, "World");
    if (substrPtr) {
        printf("strstr: 'World' found in str3 at position = %ld\n", substrPtr - str3);
    } else {
        printf("strstr: 'World' not found in str3\n");
    }

    return 0;
}

// Function definition for displayMessage
void displayMessage(char *message) {
    printf("Message: %s\n", message);
}

// Function definition for add
int add(int a, int b) {
    return a + b;
}
