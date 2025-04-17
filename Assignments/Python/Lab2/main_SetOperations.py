# main_SetOperations.py

import module_SetFunction as sf

def main():
    # Sample sets
    set1 = {1, 2, 3, 4}
    set2 = {3, 4, 5, 6}

    # Demonstrating the execution of each function
    print("Original Set1:", set1)
    print("Original Set2:", set2)

    # Add element
    print("After adding element 5 to Set1:", sf.add_element(set1, 5))

    # Remove element
    print("After removing element 2 from Set1:", sf.remove_element(set1, 2))

    # Union and Intersection
    union, intersection = sf.union_intersection(set1, set2)
    print("Union of Set1 and Set2:", union)
    print("Intersection of Set1 and Set2:", intersection)

    # Difference
    print("Difference of Set1 and Set2 (Set1 - Set2):", sf.difference(set1, set2))

    # Subset check
    print("Is Set1 a subset of Set2?", sf.is_subset(set1, set2))

    # Length of set without len()
    print("Length of Set1:", sf.set_length(set1))

    # Unique subsets
    print("Unique subsets of Set1:", sf.unique_subsets(set1))

if __name__ == "__main__":
    main()
