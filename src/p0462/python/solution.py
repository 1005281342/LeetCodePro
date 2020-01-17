from typing import List


class Solution:
    def minMoves2(self, nums: List[int]) -> int:
        nums.sort()
        mid = len(nums) >> 1
        return sum([abs(nums[mid] - i) for i in nums])
