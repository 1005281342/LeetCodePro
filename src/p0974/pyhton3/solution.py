from typing import List


# class Solution:
#     def subarraysDivByK(self, A: List[int], K: int) -> int:
#
#         ms = [0] * (len(A) + 1)
#         for i, num in enumerate(A):
#             ms[i + 1] = ms[i] + num
#
#         i = 0
#         cnt = 0
#         while i < len(A):
#
#             for j in range(i + 1, len(ms)):
#                 cnt += (ms[j] - ms[i]) % K == 0
#             i += 1
#
#         return cnt
#
