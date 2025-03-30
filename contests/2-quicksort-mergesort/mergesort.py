n = int(input())
lst = list(map(int, input().split()))

def merge(left, right, buffer):
        i = j = k = 0
        while i < len(left) and j < len(right):
            if left[i] < right[j]:
                buffer[k] = left[i]
                i += 1
            else:
                buffer[k] = right[j]
                j += 1
            k += 1
        while i < len(left):
            buffer[k] = left[i]
            i += 1
            k += 1
        while j < len(right):
            buffer[k] = right[j]
            j += 1
            k += 1
        return buffer[:k]
    
def merge_sort_recursive(lst, buffer):
    if len(lst) <= 1:
        return lst
    mid = len(lst) // 2
    left = merge_sort_recursive(lst[:mid], buffer[:mid])
    right = merge_sort_recursive(lst[mid:], buffer[mid:])
    return merge(left, right, buffer)

def merge_sort(lst):
    n = len(lst)
    buffer = [0 for _ in range(n)]
    return merge_sort_recursive(lst, buffer)

print(*merge_sort(lst))

# python3 run_tests.py contests/2-quicksort-mergesort/mergesort.py contests/2-quicksort-mergesort/merge_sort_tests.zip 
# Extracted tests to /tmp/tmpyjvto3t5/tests
# Test test1 PASSED
# Test test2 PASSED
# Test test3 PASSED
# Test test4 PASSED
# Test test5 PASSED
# Test test6 PASSED
# Test test7 PASSED
# Test test8 PASSED
# Test test9 PASSED
# All 9/9 tests passed!