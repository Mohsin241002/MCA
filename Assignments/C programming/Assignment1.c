#include <stdio.h>

int main() {
  
  float mark;
  char grade;

  
  printf("Enter the mark obtained by the student in computer science: ");
  scanf("%f", &mark);

  
  if (mark > 90) {
    grade = 'A';
  } else if (mark >= 70) {
    grade = 'B';
  } else if (mark >= 50) {
    grade = 'C';
  } else {
    grade = 'D';
  }

  
  printf("The grade is: %c\n", grade);

  return 0;
}


/////////////////////////////////////////////////////////////
//Q2

#include <stdio.h>

int main() {
  
  float total_cost, discount, amount_to_pay;

  
  printf("Enter the total cost of the items purchased: ");
  scanf("%f", &total_cost);

  
  if (total_cost < 2000) {
    discount = 0.05;
  } else if (total_cost >= 2001 && total_cost <= 5000) {
    discount = 0.25;
  } else if (total_cost >= 5001 && total_cost <= 10000) {
    discount = 0.35;
  } else {
    discount = 0.50;
  }

  
  amount_to_pay = total_cost - (total_cost * discount);

  
  printf("Amount to pay after discount: Rs. %.2f\n", amount_to_pay);

  return 0;
}

///////////////////////////////////////////////////////////////////
//Q3

#include <stdio.h>

int main() {
    int sum = 0;
    int count = 0;
    int number = 1;

    while (count < 10) {
        if (number % 2 != 0) {
            sum += number;
            count++;
        }
        number++;
    }

    printf("The sum of the first 10 odd natural numbers is: %d\n", sum);

    return 0;
}


////////////////////////////////////////////////////////////////////////////
//Q4

#include <stdio.h>

int main() {
    int sum = 0;
    int i;

    for (i = 100; i <= 998; i += 2) {
        sum += i;
    }

    printf("The sum of all 3-digit even natural numbers is: %d\n", sum);

    return 0;
}








//////////////////////////////////////////////////////////////////////////////////
//Q5

#include <stdio.h>

int main() {
    int sum = 0;
    int i;

    for (i = 105; i <= 995; i += 10) {
        if (i % 2 != 0) {
            sum += i;
        }
    }

    printf("The sum of all 3-digit odd natural numbers which are multiples of 5 is: %d\n", sum);

    return 0;
}