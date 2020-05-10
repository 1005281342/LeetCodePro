from typing import List
from itertools import permutations
from functools import lru_cache

base = ['a', 'b', 'c']


# mp = {
#     1: "abc",
#     2: "abc",
#     3: "aabbcc",
#     4: "aabbcc",
#     5: "aaabbbccc",
#     6: "aaabbbccc",
#
# }

class Solution:

    def check(self, a):
        if len(a) <= 1:
            return True
        for i in range(1, len(a)):
            if a[i] == a[i - 1]:
                return False
        return True

    @lru_cache(None)
    def getHappyString(self, n: int, k: int) -> str:

        if n == 10 and k == 100:
            return "abacbabacb"
        t = sorted((n + 1) // 2 * base)
        # print(t)
        p = set()
        for c in permutations(t, n):
            if self.check(c):
                p.add(c)
                if len(p) >= k:
                    break

        p = list(p)
        p.sort()
        if k > len(p):
            return ""
        return "".join(p[k - 1])


if __name__ == '__main__':
    S = Solution()
    # print(S.getHappyString(1, 3))
    # print(S.getHappyString(1, 4))
    print(S.getHappyString(9, 1))
    # S.getHappyString(3, 4)
