import math

# Define functions for each task

def sum_list(A):
    total_sum = sum(A)
    print("The sum is:", total_sum)

def multiply_list(A):
    product = 1
    for num in A:
        product *= num
    print("The product is:", product)

def largest_number_in_list(A):
    largest_number = max(A)
    print("The largest number is:", largest_number)

def smallest_number_in_list(A):
    smallest_number = min(A)
    print("The smallest number is:", smallest_number)

def find_special_strings(A):
    for index, string in enumerate(A):
        if string[0] == string[-1]:
            print(f"String: '{string}' at Index: {index}")

def print_alphabet_pattern(n):
    for i in range(n):
        for j in range(n - i - 1):
            print(" ", end="")
        for k in range(i + 1):
            print(chr(65 + k), end=" ")
        print()

def print_star_pattern(n):
    for i in range(1, n + 1):
        print('*' * i)

def convert_to_dict(color_names, color_codes):
    color_data = [{"colorName": name, "colorCode": code} for name, code in zip(color_names, color_codes)]
    print(color_data)

def print_even_numbers_and_squares(start, end):
    for num in range(start, end):
        if num % 2 == 0:
            print(f"Number: {num}, Square: {num ** 2}")

def find_sum_and_reverse():
    while True:
        num = input("Enter a four-digit number: ")
        if num.isdigit() and 1000 <= int(num) <= 9999:
            sum_digits = sum(int(digit) for digit in num)
            reverse_num = num[::-1]
            print(f"Sum of digits: {sum_digits}")
            print(f"Reverse: {reverse_num}")
            break
        else:
            print("Invalid input. Please enter a four-digit number.")

def calculate_triangle_area(a, b, c):
    s = (a + b + c) / 2
    area = math.sqrt(s * (s - a) * (s - b) * (s - c))
    return area

def triangle_area():
    a1 = float(input("Enter side 1 of triangle 1: "))
    b1 = float(input("Enter side 2 of triangle 1: "))
    c1 = float(input("Enter side 3 of triangle 1: "))
    a2 = float(input("Enter side 1 of triangle 2: "))
    b2 = float(input("Enter side 2 of triangle 2: "))
    c2 = float(input("Enter side 3 of triangle 2: "))
    area1 = calculate_triangle_area(a1, b1, c1)
    area2 = calculate_triangle_area(a2, b2, c2)
    total_area = area1 + area2
    contribution1 = (area1 / total_area) * 100
    contribution2 = (area2 / total_area) * 100
    print("Triangle 1 area: {:.2f}".format(area1))
    print("Triangle 2 area: {:.2f}".format(area2))
    print("Total area: {:.2f}".format(total_area))
    print("Triangle 1 contribution: {:.2f}%".format(contribution1))
    print("Triangle 2 contribution: {:.2f}%".format(contribution2))

def print_person_details(people):
    for person in people:
        print(f"Name: {person['name']}")
        print(f"Age: {person['age']}")
        print(f"Blood Group: {person['blood_group']}")
        print("-" * 30)

def extract_rear_letter(tuple_string):
    tuple_string = tuple_string.strip("()")
    elements = tuple_string.split(", ")
    rear_letters = [element[-1] for element in elements]
    print(rear_letters)

def days_in_month(month, year):
    if month == "February":
        if year % 4 == 0 and (year % 100 != 0 or year % 400 == 0):
            return 29
        else:
            return 28
    elif month in ("April", "June", "September", "November"):
        return 30
    else:
        return 31

def month_days():
    months = ("January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December")
    year = int(input("Enter a year: "))
    month = input("Enter a month: ")
    if month in months:
        print("The month of {} in the year {} has {} days.".format(month, year, days_in_month(month, year)))
    else:
        print("Invalid month. Please enter a valid month name.")

def main_menu():
    while True:
        print("\nMain Menu:")
        print("1. Sum and Product of List Items")
        print("2. Largest and Smallest Number in List")
        print("3. Find Special Strings in List")
        print("4. Print Alphabet Pattern")
        print("5. Print Star Pattern")
        print("6. Convert List to List of Dictionaries")
        print("7. Print Even Numbers and Their Squares")
        print("8. Sum and Reverse of Four-Digit Number")
        print("9. Area of Two Triangles")
        print("10. Print Person Details")
        print("11. Extract Rear Letters from Tuple String")
        print("12. Days in a Month")
        print("13. Exit")
        
        choice = input("Enter your choice: ")
        
        if choice == '1':
            A = [11, 22, 33, 44, 55]
            sum_list(A)
            multiply_list(A)
        elif choice == '2':
            A = [11, 22, 33, 44, 55]
            largest_number_in_list(A)
            smallest_number_in_list(A)
        elif choice == '3':
            A = ['abc', 'xyz', 'aba', '1221']
            find_special_strings(A)
        elif choice == '4':
            n = 5
            print_alphabet_pattern(n)
        elif choice == '5':
            n = 5
            print_star_pattern(n)
        elif choice == '6':
            color_names = ["Black", "Red", "Maroon", "Yellow"]
            color_codes = ["000000", "FF0000", "800000", "FFFF00"]
            convert_to_dict(color_names, color_codes)
        elif choice == '7':
            print("Range (1, 50):")
            print_even_numbers_and_squares(1, 50)
            print("Range (1, 100):")
            print_even_numbers_and_squares(1, 100)
        elif choice == '8':
            find_sum_and_reverse()
        elif choice == '9':
            triangle_area()
        elif choice == '10':
            people = [
                {"name": "John Doe", "age": 30, "blood_group": "A+"},
                {"name": "Jane Smith", "age": 25, "blood_group": "B-"},
                {"name": "Emily Davis", "age": 40, "blood_group": "O+"},
                {"name": "Michael Brown", "age": 35, "blood_group": "AB-"},
                {"name": "William Johnson", "age": 28, "blood_group": "A-"},
                {"name": "Emma Wilson", "age": 22, "blood_group": "B+"},
                {"name": "Oliver Martinez", "age": 33, "blood_group": "O-"},
                {"name": "Sophia Anderson", "age": 27, "blood_group": "AB+"},
                {"name": "James Thomas", "age": 45, "blood_group": "A+"},
                {"name": "Isabella Lee", "age": 38, "blood_group": "B-"}
            ]
            print_person_details(people)
        elif choice == '11':
            tuple_string = "(apple, banana, orange, grapes, mango)"
            extract_rear_letter(tuple_string)
        elif choice == '12':
            month_days()
        elif choice == '13':
            print("Exiting the program.")
            break
        else:
            print("Invalid choice. Please try again.")

if __name__ == "__main__":
    main_menu()
