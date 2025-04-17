# class person:
#     def __init__(self, parameeter1_name,parameter2_age):
#         self.name = parameeter1_name
#         self.age= parameter2_age

# p1= person("john", 36)
# p2 = person("Riya", 21)

# print(type(p1))
# print(p1.name)
# print(p1.age)


# class car:
#     def __init__(self,make,model,year):
#         self.make = make
#         self.model = model
#         self.year = year
    
#     def display_info(self):
#         print(f"Car information :\nMake : {self.make}\n Model : {self.model}\n Year: {self.year}")

# my_car = car("Toyota","corolla",2020)

# my_car.display_info()


class account:
    def __init__(self,name,balance):
        self.name = name
        self.balance = balance
    def deposit(self,value):
        self.balance = value


acc1 = account('HSBC', 1000)
acc1.deposit(1000)
print(acc1.balance)


 