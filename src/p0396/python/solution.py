from typing import List
# from math import inf


class Solution:
    def maxRotateFunction(self, A: List[int]) -> int:
        n = len(A)
        cur = sum(A[i] * i for i in range(n))
        _sum = sum(A)
        res = cur
        for i in range(1, n):
            cur = _sum - n * A[-i] + cur
            res = max(res, cur)
        return res


# 超时
# class Solution:
#
#     def __init__(self):
#         self.ans = -inf
#
#     def helper(self, A):
#         ans = 0
#         for i, v in enumerate(A):
#             ans += i * v
#         self.ans = max(ans, self.ans)
#
#     def maxRotateFunction(self, A: List[int]) -> int:
#         if not A:
#             return 0
#         for i in range(len(A)):
#             self.helper(A[i:] + A[:i])
#         return self.ans
