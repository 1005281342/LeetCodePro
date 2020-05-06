from typing import List


class Solution:
    def maxScore(self, s: str) -> int:
        ans = 1
        for i in range(1, len(s)):
            a = s[:i]
            b = s[i:]
            ans = max(a.count('0') + b.count('1'), ans)
        return ans