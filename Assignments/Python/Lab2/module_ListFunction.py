# module_ListFunction.py

def max_value(lst):
    """Returns the maximum value in the list."""
    if not lst:
        return None
    return max(lst)

def min_value(lst):
    """Returns the minimum value in the list."""
    if not lst:
        return None
    return min(lst)

def sum_values(lst):
    """Returns the sum of all elements in the list."""
    return sum(lst)

def average_value(lst):
    """Returns the average of the list."""
    if not lst:
        return None
    return sum(lst) / len(lst)

def median_value(lst):
    """Returns the median of the list."""
    if not lst:
        return None
    sorted_lst = sorted(lst)
    n = len(sorted_lst)
    middle = n // 2
    if n % 2 == 0:
        return (sorted_lst[middle - 1] + sorted_lst[middle]) / 2
    else:
        return sorted_lst[middle]
