from typing import List
from collections import defaultdict


class Solution:
    def mostVisited(self, n: int, rounds: List[int]) -> List[int]:
        dd = defaultdict(int)
        t = [x for x in range(1, n + 1)]
        xt = list()
        for i in range(1, len(rounds)):
            xx = rounds[i]
            if xx < rounds[i - 1]:
                xx += n
            for x in range(rounds[i - 1], xx):
                dd[x % n] += 1
                xt.append(x % n)

        xt.append(rounds[-1] % n)
        dd[rounds[-1] % n] += 1
        print(xt)

        mx = max(dd.values())
        ans = list()
        for k, v in dd.items():
            if v == mx:
                if k == 0:
                    k += n
                ans.append(k)
        return sorted(ans)
