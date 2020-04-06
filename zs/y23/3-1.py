# from math import sqrt
#
#
# class Solution:
#     def checkOverlap(self, radius: int, x: int, y: int, x1: int, y1: int, x2: int, y2: int) -> bool:
#
#         if x1 > x2:
#             x1, x2 = x2, x1
#         if y1 > y2:
#             y1, y2 = y2, y1
#         # 1. x in (x1, x2) or y in (y1, y2)
#         # 里面 相交
#         if x1 <= x <= x2 and y1 <= y <= y2:
#             return True
#
#         # 圆 左右
#         if x2 < x or x1 > x:
#             r_min = min(abs(x - x1), abs(x - x2))
#             return radius >= r_min
#
#         # 圆 上下
#         if y2 < y or y > y1:
#             r_min = min(abs(y - y1), abs(y - y2),
#                         )
#             return radius >= r_min
#         # 不相交
#         r_min = min(
#             sqrt((x - x1) ** 2 + (y - y1) ** 2),
#             sqrt((x - x2) ** 2 + (y - y2) ** 2),
#             sqrt((x - x1) ** 2 + (y - y2) ** 2),
#             sqrt((x - x2) ** 2 + (y - y1) ** 2),
#         )
#         return radius >= r_min
#
#
# if __name__ == '__main__':
#     S = Solution()
#     ans = S.checkOverlap(radius=1, x=1, y=1, x1=1, y1=-3, x2=2, y2=-1)
#     print(ans)
