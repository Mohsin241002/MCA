# module_SetFunction.py

def add_element(s, element):
    """Adds an element to the set."""
    s.add(element)
    return s

def remove_element(s, element):
    """Removes an element from the set."""
    s.discard(element)
    return s

def union_intersection(s1, s2):
    """Returns the union and intersection of two sets."""
    union = s1.union(s2)
    intersection = s1.intersection(s2)
    return union, intersection

def difference(s1, s2):
    """Returns the difference of two sets (s1 - s2)."""
    return s1.difference(s2)

def is_subset(s1, s2):
    """Checks if s1 is a subset of s2."""
    return s1.issubset(s2)

def set_length(s):
    """Finds the length of the set without using len() function."""
    count = 0
    for _ in s:
        count += 1
    return count

def unique_subsets(s):
    """Finds all unique subsets of the given set."""
    from itertools import chain, combinations
    s_list = list(s)
    return list(chain.from_iterable(combinations(s_list, r) for r in range(len(s_list) + 1)))
