#include<stdio.h>

float InitalCharge = 2.00;
float AdditionalCharge = 0.50;
float MaxCharge= 10.00;

float calculateCharge(float hrs){
    float totalcharge = (InitalCharge + (hrs*AdditionalCharge));
    if(hrs<3){
        return InitalCharge;
    }
    else if(hrs>3 && hrs<=24){
        if(totalcharge > MaxCharge){
            return MaxCharge;
        }
        else
            return totalcharge;
    }

}

void main(){
    float car1,car2,car3,charge1,charge2,charge3;

    printf("Enter the No of hrs u need to park ur car1 :");
    scanf("%f",&car1);
    printf("Enter the No of hrs u need to park ur car1 :");
    scanf("%f",&car2);
    printf("Enter the No of hrs u need to park ur car1 :");
    scanf("%f",&car3);
    charge1 = calculateCharge(car1);
    charge2 = calculateCharge(car2);
    charge3 = calculateCharge(car3);


    printf("Car\tHours\tCharge\n");
    printf("1\t %f\t %f\n",car1,charge1);
    printf("1\t %f\t %f\n",car2,charge2);
    printf("1\t %f\t %f\n",car3,charge3);
}