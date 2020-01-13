from typing import List


class Solution:
    def minIncrementForUnique(self, A: List[int]) -> int:

        if len(A) <= 1:
            return 0
        A.sort()
        ans = 0
        for i, v in enumerate(A[1:]):
            if v <= A[i]:
                ans += A[i] + 1 - v
                A[i + 1] = A[i] + 1
        return ans
