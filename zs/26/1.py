from typing import List

from collections import defaultdict


class Solution:
    def maxPower(self, s: str) -> int:
        dd = defaultdict(int)
        s += "9"
        t = 1
        for i in range(1, len(s)):
            if s[i] != s[i - 1]:
                dd[s[i]] = max(dd[s[i]], t)
                t = 1
            else:
                t += 1

        return max(dd.values())
