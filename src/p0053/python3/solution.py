from typing import List


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        ans = -1
        last = 0
        for num in nums:
            last = num + max(last, 0)
            ans = max(ans, last)
        return ans
