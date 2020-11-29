from collections import Counter, defaultdict


class Solution:
    def numSplits(self, s: str) -> int:
        if len(s) <= 1:
            return 0

        length = len(set(s))
        if length == 1:
            return len(s) - 1

        cd = Counter(s)
        cnt = 0
        dd = defaultdict(int)
        for i in range(0, len(s)):
            dd[s[i]] += 1
            cd[s[i]] -= 1
            if cd[s[i]] == 0:
                del cd[s[i]]
            cnt += len(dd) == len(cd)

        return cnt
