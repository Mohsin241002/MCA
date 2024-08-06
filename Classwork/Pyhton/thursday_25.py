# write a python Program that defines us to convert an integer from one to 50 to its corresponding to the following 
# 1.dictionary that maps integers from one to 50 to the correct corresponding Roman numeral String 
# 2.method to convert a given integer to its Roman numeral 
# 3.method should check if the given  integer is within the valid range one to fifty and return an appropriate message if it is not implemented class and demonstrate its usage


class RomanNumeralConverter:
    def __init__(self):
        self.roman_dict = self.create_roman_dict()
        
    def create_roman_dict(self):
        roman_dict = {
            1: "I", 2: "II", 3: "III", 4: "IV", 5: "V",
            6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
            11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV",
            16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX",
            21: "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV",
            26: "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX", 30: "XXX",
            31: "XXXI", 32: "XXXII", 33: "XXXIII", 34: "XXXIV", 35: "XXXV",
            36: "XXXVI", 37: "XXXVII", 38: "XXXVIII", 39: "XXXIX", 40: "XL",
            41: "XLI", 42: "XLII", 43: "XLIII", 44: "XLIV", 45: "XLV",
            46: "XLVI", 47: "XLVII", 48: "XLVIII", 49: "XLIX", 50: "L"
        }
        return roman_dict
    
    def int_to_roman(self, number):
        if 1 <= number <= 50:
            return self.roman_dict[number]
        else:
            return "Number out of range. Please enter a number between 1 and 50."
        
# Demonstrate the usage
converter = RomanNumeralConverter()

# Test cases
test_numbers = [3, 12, 25, 40, 50, 0, 51]
for number in test_numbers:
    result = converter.int_to_roman(number)
    print(f"The Roman numeral for {number} is: {result}")
