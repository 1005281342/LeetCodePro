from collections import Counter


class Solution:
    def canConstruct(self, s: str, k: int) -> bool:

        if len(s) < k:
            return False
        if len(s) == k:
            return True

        cd = Counter(s)
        if len(cd) <= k:
            return True

        cnt = 0
        for x in cd.values():
            if x & 1:
                cnt += 1
        return cnt <= k