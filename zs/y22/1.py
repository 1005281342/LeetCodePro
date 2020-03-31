from typing import List


class Solution:
    def findTheDistanceValue(self, arr1: List[int], arr2: List[int], d: int) -> int:

        ans = 0
        for v in arr1:
            flag = True
            for vv in arr2:
                if abs(v-vv) <= d:
                    flag = False
                    break
            ans += flag
        return ans
