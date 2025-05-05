from typing import List
def repetitions(s: str) -> int:
    # write your code here ^_^
    total = [1]
    count = 0

    for i in range(1, len(s)):
        if s[i] == s[i - 1]: 
            total[count] += 1
        else:
            total.append(1)
            count += 1

    return max(total)

print(repetitions(s = 'AMMMDDCCCMD'))