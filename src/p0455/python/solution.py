from typing import List


class Solution:
    def findContentChildren(self, g: List[int], s: List[int]) -> int:
        g.sort()
        s.sort()
        idx = 0
        for i in range(len(s)):
            if s[i] < g[idx]:
                continue
            idx += 1
            if idx >= len(g):
                break
        return idx
