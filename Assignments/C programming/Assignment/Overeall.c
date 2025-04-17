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


// convert decimal to binary

#include <stdio.h>
#include <math.h>

// function prototype
long long convert(int);

int main() {
    
  int n;
  long long bin;
  
  printf("Enter a decimal number: ");
  scanf("%d", &n);
  
  // convert to binary using the convert() function
  bin = convert(n);
  
  printf("%d in decimal =  %lld in binary", n, bin);

  return 0;
}

// function to convert decimal to binary
long long convert(int n) {

  // variable to store the result
  long long bin = 0;

  int rem, i = 1;

  // loop to convert to binary
  while (n != 0) {
    
    // get remainder of n divided by 2
    rem = n % 2;
    
    // divide n by 2
    n /= 2;
    
    // multiply remainder by i
    // add the product to bin
    bin += rem * i;
    
    // multiply i by 10
    i *= 10;
  }

  return bin;
}



// C program to demonstrate whether 
// a number is prime or not using function
#include <stdio.h> 

// Defining the function 
int primenumber(int number) 
{ 
	int i; 
	
	// Condition for checking the 
	// given number is prime or 
	// not 
	for (i = 2; i <= number / 2; i++) 
	{ 
		if (number % i != 0) 
			continue; 
		else
			return 1; 
	} 
	return 0; 
} 

// Driver code 
int main() 
{ 
	int num, res = 0; 
	printf("Enter a number");
    scanf("%d", &num);
	// Calling the function 
	res = primenumber(num); 
	if (res == 0) 
		printf("%d is a prime number", num); 
	else
		printf("%d is not a prime number", num); 
}
