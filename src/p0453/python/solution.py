from typing import List


class Solution:
    def minMoves(self, nums: List[int]) -> int:
        mi = min(nums)
        return sum(nums) - mi * len(nums)