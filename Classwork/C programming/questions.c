// write a program in c to swap two numbers using a function

#include<stdio.h>
void swap(int num1,int num2){
    int temp = num1;
    num1 = num2;
    num2 = temp;

    printf("Fisrt no is %d second no is %d", num1, num2);
}
void main(){
    int num1, num2;
    scanf("%d %d", &num1, &num2);
    swap(num1,num2);
}
