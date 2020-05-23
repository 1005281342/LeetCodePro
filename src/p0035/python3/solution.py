from bisect import bisect

class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:

        index = bisect(nums, target)
        if target == nums[index-1]:
            return index - 1
        return index