import functools
from typing import List


class Solution:
    def sort_rule(self, x, y):
        a, b = x + y, y + x
        if a > b:
            return 1
        elif a < b:
            return -1
        else:
            return 0

    def minNumber(self, nums: List[int]) -> str:

        return "".join(sorted([str(num) for num in nums], key=functools.cmp_to_key(self.sort_rule)))


"""
class cmpSmaller(str):
    def __lt__(self, y):
        return self + y < y + self  # 字符串拼接比较(两两比较)
    # 按由小到大来排列

class Solution:
    def minNumber(self, nums: List[int]) -> str:
        res=sorted(map(str, nums),key=cmpSmaller)
        smallest = ''.join(res)
        return smallest
"""
