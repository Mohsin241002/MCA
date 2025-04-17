#include <stdio.h>
#include <ctype.h>
#include <string.h>
#include <stdlib.h>

#define MAX 100 // Maximum size for stack

// Structure to hold variables and their values
struct Variable {
    char name;
    int value;
};

// Function to find precedence of operators
int precedence(char ch) {
    switch(ch) {
        case '+':
        case '-':
            return 1;
        case '*':
        case '/':
            return 2;
        case '^':
            return 3;
    }
    return -1;
}

// Function to check if a character is an operator
int isOperator(char ch) {
    return (ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '^');
}

// Function to convert infix expression to postfix
void infixToPostfix(char* infix, char* postfix) {
    char stack[MAX];
    int top = -1;
    int k = 0;
    
    for (int i = 0; infix[i]; i++) {
        if (isalnum(infix[i])) {  // If character is alphanumeric (variable or number)
            postfix[k++] = infix[i];
        }
        else if (infix[i] == '(') {  // Left parenthesis
            stack[++top] = infix[i];
        }
        else if (infix[i] == ')') {  // Right parenthesis
            while (top != -1 && stack[top] != '(') {
                postfix[k++] = stack[top--];
            }
            top--;  // Remove left parenthesis '(' from the stack
        }
        else {  // Operator
            while (top != -1 && precedence(infix[i]) <= precedence(stack[top])) {
                postfix[k++] = stack[top--];
            }
            stack[++top] = infix[i];
        }
    }
    
    // Pop all remaining operators from the stack
    while (top != -1) {
        postfix[k++] = stack[top--];
    }
    
    postfix[k] = '\0';  // Null terminate the postfix expression
}

// Function to find the value of a variable
int findVariableValue(struct Variable vars[], int varCount, char var) {
    for (int i = 0; i < varCount; i++) {
        if (vars[i].name == var) {
            return vars[i].value;
        }
    }
    return 0;  // Default to 0 if variable not found
}

// Function to evaluate a postfix expression
int evaluatePostfix(char* postfix, struct Variable vars[], int varCount) {
    int stack[MAX];
    int top = -1;
    
    for (int i = 0; postfix[i]; i++) {
        if (isalnum(postfix[i])) {  // If operand (variable)
            stack[++top] = findVariableValue(vars, varCount, postfix[i]);
        }
        else {  // Operator
            int val1 = stack[top--];
            int val2 = stack[top--];
            
            switch (postfix[i]) {
                case '+': stack[++top] = val2 + val1; break;
                case '-': stack[++top] = val2 - val1; break;
                case '*': stack[++top] = val2 * val1; break;
                case '/': stack[++top] = val2 / val1; break;
            }
        }
    }
    
    return stack[top];  // Final result
}

int main() {
    char infix[MAX], postfix[MAX];
    struct Variable vars[MAX];
    int varCount = 0;

    // Take input of infix expression
    printf("Enter your Game code (e.g., gameno+userid*(playercode*tapeid)): ");
    scanf("%s", infix);

    // Convert infix to postfix
    infixToPostfix(infix, postfix);
    printf("Postfix expression: %s\n", postfix);
    
    // Find unique variables in the expression and ask for their values
    for (int i = 0; infix[i]; i++) {
        if (isalpha(infix[i])) {  // Check if it's a variable (A-Z, a-z)
            int exists = 0;
            for (int j = 0; j < varCount; j++) {
                if (vars[j].name == infix[i]) {
                    exists = 1;
                    break;
                }
            }
            if (!exists) {  // If variable is unique, ask for its value
                vars[varCount].name = infix[i];
                printf("Enter the value for %c: ", infix[i]);
                scanf("%d", &vars[varCount].value);
                varCount++;
            }
        }
    }

    // Evaluate the postfix expression
    int result = evaluatePostfix(postfix, vars, varCount);
    printf("Result of postfix evaluation: %d\n", result);

    return 0;
}
