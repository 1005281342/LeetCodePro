from typing import List


class Solution:
    def minSubsequence(self, nums: List[int]) -> List[int]:
        nums.sort(reverse=True)
        cnt = 0
        s = sum(nums)
        for i in range(len(nums)):
            cnt += nums[i]
            if cnt > s - cnt:
                return nums[:i+1]
        return nums
