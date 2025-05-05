from typing import List
def increasing_array(arr: List[int]) -> int:
    # write your code here ^_^
    count = 0
    for i in range(len(arr)):
        if i+1 == len(arr):
            return count
        if arr[i] > arr[i+1]:
            while arr[i] > arr[i+1]:
                arr[i+1] +=1
                count +=1

print(increasing_array([3, 2, 5, 1, 7]))