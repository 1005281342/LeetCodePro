from typing import List


class Solution:
    def destCity(self, paths: List[List[str]]) -> str:
        d = dict()
        s = set()
        for a, b in paths:
            d[a] = b
            s.add(a)
            s.add(b)

        for cc in s:
            if d.get(cc) is None:
                return cc

