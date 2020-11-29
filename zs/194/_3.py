# from typing import List
# from collections import defaultdict, Counter
# from bisect import bisect_left
#
#
# class Solution:
#     def avoidFlood(self, rains: List[int]) -> List[int]:
#         dd = defaultdict(int)
#
#         lst = list()
#         for i, v in enumerate(rains):
#             if v == 0:
#                 lst.append(i)
#
#         ans = [-1] * len(rains)
#         start = 0
#         use = True
#         for x in lst:
#             for i in range(start, x):
#                 print(dd)
#                 if rains[i] == 0:
#                     use = False
#                     continue
#
#                 dd[rains[i]] += 1
#                 if dd[rains[i]] > 1:
#                     if use:
#                         return []
#                     else:
#                         use = True
#                         dd[rains[i]] -= 1
#                         ans[start] = rains[i]
#
#             start = x
#             use = False
#
#         for i in range(start, len(rains)):
#             if rains[i] == 0:
#                 use = False
#                 continue
#
#             dd[rains[i]] += 1
#             if dd[rains[i]] > 1:
#                 if use:
#                     return []
#                 else:
#                     use = True
#                     dd[rains[i]] -= 1
#                     ans[start] = rains[i]
#
#         print(dd)
#         for v in dd.values():
#             if v > 1:
#                 return []
#
#         for x in lst:
#             if ans[x] == -1:
#                 ans[x] = 1
#
#         return ans
#
#
# if __name__ == '__main__':
#     S = Solution()
#     S.avoidFlood([1, 2, 0, 0, 2, 1])
