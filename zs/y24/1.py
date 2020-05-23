from typing import List


class Solution:
    def minStartValue(self, nums: List[int]) -> int:
        ss = list()
        cnt = 0
        for a in nums:
            cnt += a
            ss.append(cnt)

        s = min(ss)
        if s >= 0:
            return 1
        return 1 - s