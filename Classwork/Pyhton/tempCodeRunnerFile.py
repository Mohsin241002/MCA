months = ("January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December")

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

year = int(input("Enter a year: "))
month = input("Enter a month: ")

if month in months:
    print("The month of {} in the year {} has {} days.".format(month, year, days_in_month(month, year)))
else:
    print("Invalid month. Please enter a valid month name.")
