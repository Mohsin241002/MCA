// Write a program in C to check if a given number is even or odd using the function

#include<stdio.h>

int evenodd(int no){
    if (no%2==0)
        return 0;
    else
        return 1;

}

void main(){
    int no, ans;
    printf("Enter a no");
    scanf("%d", &no);
    ans=evenodd(no);
    if(ans == 0)
        printf("The no is Even");
    else
        printf("The no is odd");
}
