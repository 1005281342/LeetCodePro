from typing import List
from functools import lru_cache


class Solution:
    def stoneGameV(self, s: List[int]) -> int:
        n = len(s)
        pre = [0]
        for i in range(n):
            pre.append(pre[-1] + s[i])

        # print(pre)
        @lru_cache(None)
        def dp(l, r):
            if l == r:
                return 0
            res = 0
            for i in range(l, r):
                lc = pre[i + 1] - pre[l]
                rc = pre[r + 1] - pre[i + 1]
                if lc > rc:
                    res = max(res, rc + dp(i + 1, r))
                elif lc < rc:
                    res = max(lc + dp(l, i), res)
                else:
                    res = max(res, lc + max(dp(l, i), dp(i + 1, r)))
            return res

        return dp(0, n - 1)

# class Solution:
#     def stoneGameV(self, stoneValue: List[int]) -> int:
#         return self.h(stoneValue, sum(stoneValue), 0)
#
#     #
#     def h(self, stoneValue: List[int], cnt: int, a: int):
#         if len(stoneValue) <= 1:
#             return a
#
#         index = 0
#         tt = 0
#         while index < len(stoneValue):
#             kt = stoneValue[index] + tt
#             if tt > min(cnt - kt, kt):
#                 stoneValue = stoneValue[:index + 1]
#                 break
#             if cnt - kt < kt:
#                 tt = cnt - kt
#                 stoneValue = stoneValue[index + 1:]
#                 break
#             if cnt - kt == kt:
#
#                 return max(self.h(stoneValue[:index+1], kt, a+kt), self.h(stoneValue[index+1:], kt, a+kt))
#                 # if index + 1 > len(stoneValue) - index - 1:
#                 #     stoneValue = stoneValue[:index + 1]
#                 # elif index + 1 == len(stoneValue) - index - 1:
#                 #     aa = stoneValue[:index + 1]
#                 #     bb = stoneValue[index + 1:]
#                 #     if aa[-1] - aa[0] > bb[-1] - bb[0]:
#                 #         stoneValue = bb
#                 #     else:
#                 #         stoneValue = aa
#                 # else:
#                 #     stoneValue = stoneValue[index + 1:]
#             tt = kt
#             index += 1
#
#         return self.h(stoneValue, tt, a + tt)
