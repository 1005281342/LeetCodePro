from typing import List

from collections import defaultdict


class Solution:
    def groupThePeople(self, groupSizes: List[int]) -> List[List[int]]:
        dd = defaultdict(list)
        for i, v in enumerate(groupSizes):
            dd[v].append(i)

        ans = list()
        for d, vs in dd.items():
            while len(vs) > d:
                ans.append(vs[:d])
                vs = vs[d:]
            ans.append(vs)
        return ans