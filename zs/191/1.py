from typing import List


class Solution:
    def shuffle(self, nums: List[int], n: int) -> List[int]:
        ans = list()
        for i in range(n, 2 * n):
            ans.append(nums[i - n])
            ans.append(nums[i])
        return ans
