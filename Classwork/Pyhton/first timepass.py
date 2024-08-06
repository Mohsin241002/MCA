print("Hello")
pi = 3.14
print(type(pi))                   #find type of a data type

num_1= 1234.5678
z=3+4j
num_2= 234.34                  #data type in python
print(z)
print(type(z))

# slicing in stings
## slicing is a way of creating a substring from a given string
### slicing to extract a portion of a string by specyfying a start and end indexe

string='name'           #positive index increments from left to right
print(string[-2])           #negetive indexing  : the character at the end has -1 index



fname= 'Mohsin'
lname = 'Chunawala'
fullname= fname+lname     #add strings using + operator
print(fullname)


print(len(fullname))        #prints length in nos of string

number = 123
num_str = str(number)            #convert number to string
print(num_str)

print(fullname[2:])    #string slicing


full_name = fname + ' ' + lname           #adding space in between
print(full_name)

print(fullname.lower())
print(fullname.upper())                         #methods in python
print(fullname.capitalize())
print(fullname.split())

a = 12
b = 23
print(a+b)

age = int(input("Enter ur age"))          #taking input from user
print(age)
    