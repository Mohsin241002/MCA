// Operators in c
#include<stdio.h>

void main() {
    // Arithmetic Operators
    int num1 = 54, num2 = 2;
    printf("\nNum1 = 54, Num2 = 2");
    printf("\n\nArithmetic Operators");
    printf("\nAddition = %d", num1 + num2);
    printf("\nSubtraction = %d", num1 - num2);
    printf("\nMultiplication = %d", num1 * num2);
    printf("\nDivision = %d", num1 / num2);
    printf("\nModulus = %d", num1 % num2);
    printf("\nUnary plus = %d", +num1);
    printf("\nUnary Minus = %d", -num1);
    printf("\nIncrement (postfix) = %d", num1++);
    printf("\nDecrement (postfix) = %d", num1--);

    // Relational Operators
    printf("\n\n\n\n\nRelational Operators (True = 1, False = 0)");
    printf("\nIs num1 less than num2 = %d", num1 < num2);
    printf("\nIs num1 greater than num2 = %d", num1 > num2);
    printf("\nIs num1 equal to num2 = %d", num1 == num2);
    printf("\nIs num1 greater than or equal to num2 = %d", num1 >= num2);
    printf("\nIs num1 less than or equal to num2 = %d", num1 <= num2);
    printf("\nIs num1 not equal to num2 = %d", num1 != num2);

    // Logical Operators
    printf("\n\n\n\n\nLogical Operators (True = 1, False = 0)");
    printf("\nLogical AND (num1 && num2) = %d", num1 && num2);
    printf("\nLogical OR (num1 || num2) = %d", num1 || num2);
    printf("\nLogical NOT (!num1) = %d", !num1);

    // Bitwise Operators
    printf("\n\n\n\n\nBitwise Operators");
    printf("\nBitwise AND (num1 & num2) = %d", num1 & num2);
    printf("\nBitwise OR (num1 | num2) = %d", num1 | num2);
    printf("\nBitwise XOR (num1 ^ num2) = %d", num1 ^ num2);
    printf("\nBitwise NOT (~num1) = %d", ~num1);
    printf("\nLeft Shift (num1 << num2) = %d", num1 << num2);
    printf("\nRight Shift (num1 >> num2) = %d", num1 >> num2);

    // Assignment Operators
    int result = 2;
    printf("\n\n\n\n\nAssignment Operators(Result = 2)");
    result += num1;
    printf("\nAddition Assignment (result += num1), result = %d", result);
    result = 2;
    result -= num1;
    printf("\nSubtraction Assignment (result -= num1), result = %d", result);
    result = 2;
    result *= num1;
    printf("\nMultiplication Assignment (result *= num1), result = %d", result);
    result = 2;
    result /= num1;
    printf("\nDivision Assignment (result /= num1), result = %d", result);
    result = 2;
    result %= num1;
    printf("\nModulus Assignment (result = num1), result = %d", result);
    result = 2;
    result &= num1;
    printf("\nBitwise AND Assignment (result &= num1), result = %d", result);
    result = 2;
    result |= num1;
    printf("\nBitwise OR Assignment (result |= num1), result = %d", result);
    result = 2;
    result ^= num1;
    printf("\nBitwise XOR Assignment (result ^= num1), result = %d", result);
    result = 2;
    result <<= num1;
    printf("\nLeft Shift Assignment (result <<= num1), result = %d", result);
    result = 2;
    result >>= num1;
    printf("\nRight Shift Assignment (result >>= num1), result = %d", result);

}
