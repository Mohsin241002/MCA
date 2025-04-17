#include <stdio.h>

void
main (void)
{
    // Integer variables
    int numVar_a;                
    int numVar_b = 0;            
    int numVar_c, numVar_d = -10;  
    int numVar_e = 89U;          
    long int numVar_f = 99998L;  
  
    printf ("Demo - Data Types\n\n");
    printf ("Number Variable A %d\n", numVar_a);
    printf ("Address of Number Varaible A %p\n", &numVar_a);
    printf ("Number Variable B %d\n", numVar_b);
    printf ("Number Variable D %d\n", numVar_d);
  
    printf("\n");
    // Character variables
    char charVar_a = 'a';
    char charVar_b;
    printf ("Character Variable a: %c\n", charVar_a);
    charVar_a++;  // Increment character (a -> b)
    printf ("Character Variable a after increment is: %c\n", charVar_a);
    charVar_b = 99;  // Assigned ASCII value: a-->97, b-->98, c-->99
    printf ("Character Variable b: %c\n", charVar_b);
    
    printf("\n");
    // Float variables
    float floatVar_a = 3.14;
    float floatVar_b = 2.5f;  // 'f' suffix for float literal
    printf ("Float Variable a: %f\n", floatVar_a);
    printf ("Float Variable b: %f\n", floatVar_b);
    
    printf("\n");
    // Double variables (more precise than float)
    double doubleVar_a = 3.14159265359;
    double doubleVar_b = 2.71828;
    printf ("Double Variable a: %lf\n", doubleVar_a);
    printf ("Double Variable b: %lf\n", doubleVar_b);
    
    printf("\n");
    // Short variables (smaller range than int)
    short shortVar_a = 32767;    // Maximum value for short
    short shortVar_b = -32768;   // Minimum value for short
    printf ("Short Variable a: %hd\n", shortVar_a);
    printf ("Short Variable b: %hd\n", shortVar_b);
    
    printf("\n");
    // Unsigned int variables (only positive values)
    unsigned int uintVar_a = 4294967295U;  // Maximum value for unsigned int
    unsigned int uintVar_b = 42U;
    printf ("Unsigned Int Variable a: %u\n", uintVar_a);
    printf ("Unsigned Int Variable b: %u\n", uintVar_b);
    
    printf("\n");
    // Long long variables (very large range)
    long long longLongVar_a = 9223372036854775807LL;  // Maximum value
    long long longLongVar_b = -9223372036854775808LL; // Minimum value
    printf ("Long Long Variable a: %lld\n", longLongVar_a);
    printf ("Long Long Variable b: %lld\n", longLongVar_b);
}