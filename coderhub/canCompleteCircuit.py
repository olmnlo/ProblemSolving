from typing import List
def canCompleteCircuit(gas: List[int], cost: List[int]) -> int:
    if sum(gas) < sum(cost):
        return -1
    
    total_tank = 0
    current_tank = 0
    start_index = 0
    for i in range(len(gas)):
        total_tank += gas[i] - cost[i]
        current_tank += gas[i] - cost[i]
        
        if current_tank < 0:
            start_index = i + 1
            current_tank = 0
    return start_index if total_tank >= 0 else -1

gas = [2, 3, 4]
cost = [3, 4, 3]

print(canCompleteCircuit(gas=gas , cost=cost))