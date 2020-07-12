from typing import List
from bisect import bisect

from itertools import permutations
from collections import defaultdict

dd = defaultdict(str)
# print(len(list(permutations(range(1, 100 + 1), 2))))
for a, b in permutations(range(1, 100 + 1), 2):
    if dd.get(a / b):
        continue
    dd[a / b] = str(str(a) + '/' + str(b))
vs = sorted(dd.keys())[:3044]


# print(vs)

class Solution:
    def simplifiedFractions(self, n: int) -> List[str]:
        if n == 1:
            return []

        # print(dd)
        index = bisect(vs, 1)
        # print(index)
        # print(dd[vs[index]])
        ans = list()
        for i in range(3044):
            if vs[i] >= n:
                continue
            v = dd[vs[i]]
            a, b = [int(x) for x in v.split('/')]
            if a < n and b <= n:
                ans.append(v)

        return ans



if __name__ == '__main__':
    S = Solution()
    S.simplifiedFractions(3)
