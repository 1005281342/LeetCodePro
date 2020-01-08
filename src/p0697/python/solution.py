from typing import List
from collections import Counter


class Solution:
    def findShortestSubArray(self, nums: List[int]) -> int:

        cd = Counter(nums)
        t = cd.most_common(len(cd))[0][1]
        lst = list()
        for k, v in cd.most_common(len(cd)):
            if v >= t:
                lst.append(k)

        ans = len(nums)
        for k in lst:
            for i in range(0, len(nums)):
                if nums[i] == k:
                    break

            for j in range(len(nums) - 1, -1, -1):
                if nums[j] == k:
                    break

            ans = min(j - i + 1, ans)

        return ans
