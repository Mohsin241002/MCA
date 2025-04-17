# Encapsulation
class Student:
    name= 'John'
    __roll = 36
    branch = 'CSE'
    def read(self):
        print('able to read')
        print(self.__roll)
    def write(self):
        print('able to read')

s1=Student()
# print(s1.roll)
s1.read()


# polymorphism
# it can be achievend in two ways overloading and overiding

# overloading refers to ability to define multiple methods dunctions or operators with th same name but with different parameters

str1 = 'hello'
str2 = 'world'
print(str1+str2)

# Method overloading
# method overloading defining many methods that have the same name but with different arguments

# u r developing simple vehicle with the following: 
#   an init method that initailizes make and model
#   a method display details that prints the vehicle make and model

# define a subclass car that inherits from vehicle with the following
#     an init method that initialises make and model and number of doors
#     overide the display_details to print the number of doors

# create an instance of vehicle and an instance of car then call their display details methods