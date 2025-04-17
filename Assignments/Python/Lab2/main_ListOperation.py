# main_ListOperations.py

import module_ListFunction as lf

def main():
    # Sample list
    sample_list = [5, 3, 8, 6, 2, 7, 4, 1]
    
    # Demonstrating the execution of each function
    print("Sample List:", sample_list)
    print("Maximum Value:", lf.max_value(sample_list))
    print("Minimum Value:", lf.min_value(sample_list))
    print("Sum of Values:", lf.sum_values(sample_list))
    print("Average Value:", lf.average_value(sample_list))
    print("Median Value:", lf.median_value(sample_list))

if __name__ == "__main__":
    main()
