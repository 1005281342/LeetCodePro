from typing import List
from collections import Counter


class Solution:
    def findLucky(self, arr: List[int]) -> int:
        if not arr:
            return -1
        cd = Counter(arr)
        ans = -1
        for k, v in cd.items():
            if k == v and k > ans:
                ans = k
        return ans
