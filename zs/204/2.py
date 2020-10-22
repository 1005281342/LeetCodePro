from typing import List


class Solution:
    def getMaxLen(self, nums: List[int]) -> int:



        return max(self.x(nums), self.x(nums[::-1]))

    def x(self, nums):

        tns = [1] * (len(nums) + 1)
        for i, num in enumerate(nums):
            if tns[i] == 0:
                tns[i + 1] = num
            elif num == 0:
                tns[i + 1] = 0
            else:
                tns[i + 1] *= tns[i] * num
        # print(tns)
        ans = 0
        s = 0
        for i in range(1, len(tns)):
            if tns[i] == 0:
                s = i
            elif tns[i] > 0:
                ans = max(ans, i - s)

        if ans == 0:
            for num in nums:
                if num > 0:
                    return 1

        return ans

