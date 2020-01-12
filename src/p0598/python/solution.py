from typing import List


class Solution:
    def maxCount(self, m: int, n: int, ops: List[List[int]]) -> int:

        x = m
        y = n
        for a, b in ops:
            if a < x:
                x = a
            if b < y:
                y = a

        return x * y

