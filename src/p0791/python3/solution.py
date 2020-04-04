from collections import Counter


class Solution:
    def customSortString(self, S: str, T: str) -> str:
        ans = ""
        t = ""
        for c in T:
            if c not in S:
                ans += c
            else:
                t += c
        cd = Counter(t)
        for c in S:
            ans += (cd.get(c) or 0)*c
        return ans