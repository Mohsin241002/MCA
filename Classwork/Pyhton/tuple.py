
# tuple1 = (1,2,3)
# tuple2 = (4,5,6)
# tuple3 = tuple1 + tuple2
# print(tuple3)


# tuple1 = (1,2,3,4,5,6)
# slliced_tuple = tuple1[1:4]
# print(slliced_tuple)

# tuple1 = (3,4,5,6)
# print(len(tuple1))

# tuple1 = (1)
# print(type(tuple1))

#### create a tuple named fruits containing the follwing fruits apple banana cherry date

fruits= ('apple','banana','cherry', 'date')
repeated_fruits = fruits *3
print(repeated_fruits)                                 #repeat each fruit 3 times

if(fruits.count('apple')):
    print("Apple is prsent")                            #check if apple is present

for x in fruits:
    print('fruit = ',x)                                 #print each element seperate