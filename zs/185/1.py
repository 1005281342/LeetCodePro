from typing import List


class Solution:
    def reformat(self, s: str) -> str:
        n = list()
        bb = list()
        for c in s:
            if c in '0123456789':
                n.append(c)
            else:
                bb.append(c)
        if abs(len(bb) - len(n)) > 1:
            return ""

        if len(n) > len(bb):
            bb, n = n, bb

        ans = ""
        for i in range(len(n)):
            ans += bb[i]
            ans += n[i]
        if len(ans) < len(n) + len(bb):
            ans += bb[-1]
        return ans
