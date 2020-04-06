# import math
#
#
# class Solution:
#
#     # def get_k(self, a1, b1, a2, b2):
#     #     return (b2 - b1) / (a2 - a1)
#     #
#     # def get_b(self, a1, b1, k):
#     #     return b1 - a1 * k
#     #
#     # def get_length(self, x, y, a1, b1, a2, b2):
#     #     if a1 == a2:
#     #         c = -a1
#     #         return abs(x + c)
#     #     k = self.get_k(a1, b1, a2, b2)
#     #     b = self.get_b(a1, b1, k)
#     #     return abs((k * x - y + b) / math.sqrt(k * k + 1))
#     def is_in(self, x: int, y: int, x1: int, y1: int, x2: int, y2: int):
#         return (x1 <= x <= x2 or x2 <= x <= x1) and (y1 <= y <= y2 or y2 < y <= y1)
#
#     def checkOverlap(self, radius: int, x: int, y: int, x1: int, y1: int, x2: int, y2: int) -> bool:
#         # r_min = min(self.get_length(x_center, y_center, x1, y1, x2, y2),
#         #             self.get_length(x_center, y_center, x1, y2, x2, y1),
#         #             self.get_length(y_center, x_center, y1, x1, y2, x2),
#         #             self.get_length(y_center, x_center, y1, x2, y2, x1),
#         #             )
#         r_min = min(abs(x - x1), abs(x - x2), abs(y - y1), abs(y - y2),
#                     # abs(x-y1), abs(x-y2), abs(y-x1), abs(y-x2)
#                     )
#         return radius >= r_min or self.is_in(x, y, x1, y1, x2, y2)
#
#
# if __name__ == '__main__':
#     S = Solution()
#     ans = S.checkOverlap(1, x_center=1, y_center=1, x1=1, y1=-3, x2=2, y2=-1)
#     print(ans)
