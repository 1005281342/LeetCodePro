from typing import List

class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:

        ans = [0]

        for n in nums:
            ans.append(ans[-1]+n)

        return ans[1:]