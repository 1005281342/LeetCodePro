from typing import List


class Solution:
    def maxUniqueSplit(self, s: str) -> int:
        ss = set()
        i = 0
        d = 1
        while i < len(s):
            if s[i] not in ss:
                ss.add(s[i])
                i += 1
                d = 1
                continue

            while s[i:d + i] in ss:
                d += 1

            ss.add(s[i:d + i])
            i = d + i
            d = 1

        return len(s)
