from typing import List


class Solution:
    def getLastMoment(self, n: int, left: List[int], right: List[int]) -> int:
        ans = 0

        if not left and not right:
            return 0
        aa = bb = 0
        if left:
            # 最右
            a = sorted(left, reverse=True)[0]
            aa = max(aa, a)
            # return aa

        if right:
            # 最左
            b = sorted(right)[0]
            bb = max(n - b, bb)
            # return bb

        return max(aa, bb)
