# print("Hello")
# pi = 3.14
# print(type(pi))                   #find type of a data type

# digit="5"
# print(digit)
# print(type(digit))

# str1="Hello world"
# print(str1)
# print(type(str1))

# alp_1="abcd"
# alp_2="efgh"
# alp_3="ijkl"
# alp_4="mnop"
# print(alp_1+alp_2+alp_3+alp_4)

# eng_alp=(alp_1+alp_2+alp_3+alp_4)
# print(len(eng_alp))

# string='I returned a bag of groceries'
# print(string)
# print(string[0])
# print(string[1])

# Str='HELLO PYTHON'
# print(len(Str))
# A=Str[0]
# print(A)

# # slicing in stings
# ## slicing is a way of creating a substring from a given string
# ### slicing to extract a portion of a string by specyfying a start and end indexe
# print(Str[0:5])
# print(Str[:5])
# print(Str[:])
# print(Str[::])
# print(Str[6:])
# print(Str[8:10])
# print(Str[-1])
# print(Str[-12:7])

# fname="Kumkum"
# lname="Rathi"
# fullname=fname+lname
# print(fullname)

# fname="Kumkum "
# lname="Rathi"
# fullname=fname+lname
# print(fullname)
# full_name=fname+' '+lname
# print(full_name)

# print(fullname.lower())
# print(fullname.upper())                         #methods in python
# print(fullname.capitalize())
# print(fullname.split())

# number = 123
# num_str = str(number)            #convert number to string
# print(num_str)

# a=7
# b=3
# sum=a+b
# print(sum)

# diff=a=b
# print(diff)

# pro=a*b
# print(pro)

# ratio=a/b
# print(ratio)

# rem=a%b
# print(rem)

# quo=a//b
# print(quo)

# name=input("Enter your name: ")
# print("Your name is",name)

# age=int(input("Enter your age: "))
# print(age)
# print(type(age))

# wei=float(input("Enter your weight: "))
# print(wei)
# print(type(wei))

# a = 12
# b = 23
# print(a+b)

#list
Songs=['I should be allowed to think','Birdhouse in your soul']
Songs[0]
print(Songs[1])
print(type(Songs[1]))

#Nested lists
Hex=[[0,1,2,3,4,5],['A','B','C','D','E']]
print(Hex[0])

print('Numbers:',Hex[0])

Chara_fra_List=Hex[1][3]
print(Chara_fra_List)
print(Hex[1][0:3])

List1=['abc',1,2,3,3.14]
List2=['ABCDEFGH',8,'OCTAGON']
List_Concat=List1+List2
print(List_Concat)

List=[1,2,3]
List_Concat2=Hex[1]+List
print(List_Concat2)

Songs+=['AKA Driver']
print(Songs)

List_mul=List*3
print(List_mul)

#mutable
print(42)
print(isinstance(42,int))
print(3.14)
print(isinstance(3.14,float))

#finding memory address
print(id(42))

#list and mutability
Str='Narrow your Eyes'
print(len(Str))
print(Str[7:11])
#Str[7:11]='her '

#change value
num1=[1,2,314]
num1[2]=3
print(num1)

letters=["A","B","C","D"]
letters[2:]=["c","d"]
print(letters)

#delete letter
del letters[2]
print(letters)

#list method
#sort
num_list=[3,5,7,2,0]
num_list.sort()
print(num_list)
lst1=[1,2,3]
lst2=[4,5,6]
#append
lst1.append(7)
print(lst1)
#extend
lst1.extend(lst2)
print(lst1)
#index
lst1.index(7)
print(lst1)
#insert
lst1.insert(4,'Five')
print(lst1)
#remove
lst1.remove('Five')
print(lst1)
#pop
lst1.pop()
print(lst1)
#reverse
lst1.reverse()
print(lst1)

#tuples
months=('jan','feb','mar','apr','may','june','july','aug','sept','oct','nov','dec')
#meaningless if we change the order

# months[3]='Monday'
# print(months)

typ1=(1,)
tup2=1,2,3,4,5
print(type(tup2))