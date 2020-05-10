from typing import List
from collections import defaultdict


class Solution:
    def minJump(self, jump: List[int]) -> int:
        if len(jump) == 1:
            return 0

        if jump[0] >= len(jump):
            return 1

        dd = defaultdict(set)
        mark = set()
        mark.add(0)
        i = 1
        dd[i].add(0 + jump[0])
        while dd.get(i) and i <= len(jump):
            v = dd[i]
            d = i
            for a in v:
                # mark.add(a)
                if a + jump[a] >= len(jump):
                    return d
                else:
                    if a + jump[a] in mark:
                        continue
                    dd[d + 1].add(a + jump[a])
                if a - jump[a] >= 0 and a - jump[a] not in mark:
                    dd[d + 1].add(a - jump[a])
            i += 1