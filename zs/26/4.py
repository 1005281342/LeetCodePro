from typing import List

from collections import defaultdict

# class Solution:
#     def largestNumber(self, cost: List[int], target: int) -> str:
#         dd = defaultdict(str)
#         for i in range(9):
#             dd[cost[i]] = str(i + 1)
#         # print(dd)
#         tt = dict()
#         for k, v in dd.items():
#             tt[v] = k
#
#         ans = "0"
#         ks = sorted(dd.keys())
#         for index in range(9):
#             stk = list()
#             while target and index < len(ks):
#                 if target == ks[index]:
#                     stk.append(dd[ks[index]])
#                     return "".join(stk)
#                 elif target > ks[index]:
#                     stk.append(dd[ks[index]])
#                     target -= ks[index]
#                 elif stk:
#                     t = stk.pop()
#                     target += tt[t]
#                     index += 1
#                 else:
#                     # print("***")
#                     return "0"
#         return ans
#
#
# if __name__ == '__main__':
#     S = Solution()
#     print(S.largestNumber(cost=[4, 3, 2, 5, 6, 7, 2, 5, 5], target=9))
#     print(S.largestNumber(cost=[7, 6, 5, 5, 5, 6, 8, 7, 8], target=12))
#     print(S.largestNumber([2, 4, 6, 2, 4, 6, 4, 4, 4], target=5))
#     print(S.largestNumber([6, 10, 15, 40, 40, 40, 40, 40, 40], target=47)) # "32211"
