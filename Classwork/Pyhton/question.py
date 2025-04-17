# write a program to print the sum and average of 3 numbers by entering the nos directly from the console
num1 = int(input("Enter first no"))
num2 = int(input("Enter second no"))
num3 = int(input("Enter third no"))
print(" Average = ",(num1 + num2 + num3)/3)
print(" sSum = ",(num1 + num2 + num3))


#write a short program that computes the length of the hypotenuse of a right triangle given the sides
import math
side1 = int(input("Enter first side"))
side2 = int(input("Enter second side"))
side3 = math.sqrt(pow(side1,2)+pow(side2,2))
print(side3)


# Write a programme that converts the euro into inr you can do this by finding the exchange rate of Internet and then prompting for exchange rate in your program
inr = int(input("Enter the rupees"))
euro = inr/90.34
print(euro)
