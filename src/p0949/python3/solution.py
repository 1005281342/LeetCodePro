from typing import List
from itertools import permutations


class Solution:
    def largestTimeFromDigits(self, A: List[int]) -> str:
        a1 = list({cc[0] * 10 + cc[1] for cc in permutations(A, 2) if cc[0] * 10 + cc[1] <= 23})
        a1.sort()
        print(a1)
        while a1:
            c = a1.pop()
            tc = str(c)
            if len(tc) == 1:
                tc = '0' + tc
            tt = tc
            t = list()
            for ca in A:
                if str(ca) in tc:
                    tc = tc.replace(str(ca), "")
                    continue
                t.append(ca)
            ans = [10 * a + b for a, b in {(t[0], t[1]), (t[1], t[0])} if 10 * a + b <= 59]
            if ans:
                at = str(max(ans))
                at = "0" + at if len(at) == 1 else at
                return tt + ":" + at
        return ""
