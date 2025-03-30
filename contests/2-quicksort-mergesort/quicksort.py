n = int(input())
lst = list(map(int, input().split()))

def quicksort(lst: list):
    if len(lst) <= 1:
        return lst
    
    pivot = lst[len(lst)//2]
    lt = [num for num in lst if num<pivot]
    md = [num for num in lst if num==pivot]
    rt = [num for num in lst if num>pivot]
    
    return quicksort(lt)+md+quicksort(rt)

print(*quicksort(lst))

# python3 run_tests.py contests/2-quicksort-mergesort/quicksort.py contests/2-quicksort-mergesort/quicksort_tests.zip 
# Extracted tests to /tmp/tmpbytiyc7e/tests
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