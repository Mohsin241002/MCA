# Exception handling

# types of KeyError
# name error


# handling exception
# python provides a way to handle exception so that the program can continue running even after encountering an error

# a= int(input("Enter a number"))
# b= int(input("Enter a number"))
# try:
#     result = a/b
#     print(result)
# except ZeroDivisionError:
#     print("Cant be divided")


# x=5
# y="hello"
# try:
#     z=x+y
# except TypeError:
#     print("Eroor: cant add integer and string")



# try:
#     result = 10/2
# except ZeroDivisionError:
#     print("you cant divide by zero")
# else:
#     print("the result is ", result)
# finally:
#     print("this will be executed")




# def function_specific_exception(a):
#     if a<4:
#         b=a/(a-3)
#     print("value of b= ",b)
# try:
#     function_specific_exception(3)
#     function_specific_exception(5)

# except ZeroDivisionError:
#     print("Zero division error occoured")
# except NameError:
#     print("Nameerror has occoured")




# Raising an Exception
# raising an exception allows the programmer to indicate that some thing went wrong or that an unusal condition has been encountered 

# def divide(a,b):
#     if b==0:
#         raise ZeroDivisionError("You cannot divide by zero")
#     return a/b
# try:
#     result = divide(10,0)
# except ZeroDivisionError as err:
#     print(err)


# Raising ca custom Exception
# to create a custom exception define a new call that inherits from python built in exception class or any of its subclass
# then raise the exception using the raise statement

class NegativeValueError(Exception):
    pass

def square_root(x):
    if x<0:
        raise NegativeValueError("cannot calculate the square root of negative no")
    return x**0.5

try:
    result = square_root(-4)
except NegativeValueError as e:
    print(e)

