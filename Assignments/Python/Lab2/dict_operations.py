# dict_operations.py

def merging_Dict(*args):
    """Merges multiple dictionaries."""
    merged_dict = {}
    for dictionary in args:
        merged_dict.update(dictionary)
    return merged_dict

def common_keys(*args):
    """Finds common keys in multiple dictionaries."""
    if not args:
        return set()
    common = set(args[0].keys())
    for dictionary in args[1:]:
        common &= set(dictionary.keys())
    return common

def invert_dict(d):
    """Inverts a dictionary by swapping its keys and values.
       If multiple keys have the same value, they are grouped in a list.
    """
    inverted_dict = {}
    for key, value in d.items():
        if value in inverted_dict:
            inverted_dict[value].append(key)
        else:
            inverted_dict[value] = [key]
    return inverted_dict

def main():
    # Sample dictionaries
    dict1 = {'a': 1, 'b': 2, 'c': 3}
    dict2 = {'b': 2, 'c': 4, 'd': 5}
    dict3 = {'c': 3, 'd': 5, 'e': 6}
    
    # Demonstrate merging_Dict
    print("Merged Dictionary:", merging_Dict(dict1, dict2, dict3))
    
    # Demonstrate common_keys
    print("Common Keys:", common_keys(dict1, dict2, dict3))
    
    # Demonstrate invert_dict
    sample_dict = {'a': 1, 'b': 2, 'c': 1, 'd': 3}
    print("Original Dictionary:", sample_dict)
    print("Inverted Dictionary:", invert_dict(sample_dict))

if __name__ == "__main__":
    main()
