from typing import List


class Solution:
    def findLengthOfLCIS(self, nums: List[int]) -> int:

        if not nums:
            return 0

        ans = 1
        t = 1
        for i in range(len(nums) - 1):
            if nums[i] >= nums[i+1]:
                ans = max(t, ans)
                t = 1
            else:
                t += 1
        ans = max(t, ans)
        return ans