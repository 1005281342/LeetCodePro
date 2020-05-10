from typing import List
from math import gcd
from collections import deque


class Solution:
    def splitArray(self, nums: List[int]) -> int:
        dq = deque()
        dq.append([(0, len(nums) - 1), ])
        while dq:
            x = dq.popleft()
            flag = True
            aa, bb = x[0]
            for a, b in x:
                if gcd(nums[a], nums[b]) <= 1:
                    flag = False
                    break
            if flag:
                return min(len(x), len(nums))
            dq.append([(aa, aa + 1), ()])
