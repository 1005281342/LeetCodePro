from typing import List
from collections import defaultdict


class Solution:
    def largestOverlap(self, A: List[List[int]], B: List[List[int]]) -> int:
        dd = defaultdict(int)
        for i in range(len(B)):
            for j in range(len(B[0])):
                if not B[i][j]:
                    continue
                for x in range(len(A)):
                    for y in range(len(A)):
                        if not A[x][y]:
                            continue
                        dd[(i - x, j - y)] += 1

        return max(dd.values() or [0])