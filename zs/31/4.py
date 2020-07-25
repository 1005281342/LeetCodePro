from typing import List


class Solution:
    def minNumberOperations(self, target: List[int]) -> int:

        ans = 0
        xx = 0

        for t in target:
            if t > xx:
                ans += (t - xx)
            xx = t

        return ans