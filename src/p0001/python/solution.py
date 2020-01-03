from typing import List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:

        d = dict()
        for i, v in enumerate(nums):
            d[v] = i
  
        for i, v in enumerate(nums):
            it = d.get(target - v)
            if it is not None and it != i:
                return [i, it] if i < it else [it, i]

        return list()
