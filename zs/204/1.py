from typing import List


class Solution:
    def containsPattern(self, arr: List[int], m: int, k: int) -> bool:
        if len(arr) < k or len(arr) < m:
            return False

        x = "".join([str(c) for c in arr])

        for i in range(0, len(x) - m):
            if x[i: i+m] * k in x:
                return True
        return False
