from itertools import zip_longest


class Solution:
    def compareVersion(self, version1: str, version2: str) -> int:
        f = lambda x: map(int, x.split('.'))
        for v1, v2 in zip_longest(f(version1), f(version2), fillvalue=0):
            if v1 != v2:
                return 1 if v1 > v2 else -1
        return 0
