from typing import List


class Solution:
    def pivotIndex(self, nums: List[int]) -> int:

        nums = [0] + nums + [0]
        sum_num = sum(nums)
        for i, v in enumerate(nums[1:len(nums) - 1]):
            if (sum_num - v) % 2:
                continue
            if sum(nums[1:i + 1]) == ((sum_num - v) // 2):
                return i
        return -1
