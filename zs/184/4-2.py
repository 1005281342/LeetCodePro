import sys


class Solution:
    def numOfWays(self, n: int) -> int:
        sys.setrecursionlimit(1000000)
        MOD = int(1e9 + 7)

        from functools import lru_cache

        def check(pre, x):
            for i in range(2):
                if x[i] == x[i + 1]: return False
            for i in range(3):
                if x[i] == pre[i]: return False
            return True

        @lru_cache(None)
        def helper(pre, i):

            ans = []
            for x in itertools.product([0, 1, 2], [0, 1, 2], [0, 1, 2]):
                if check(pre, x):
                    ans.append(x)

            return ans

        @lru_cache(None)
        def dfs(idx, pre):
            if idx == n: return 1
            if idx == 0:
                res = 0
                for cur in [(0, 1, 0), (0, 1, 2), (0, 2, 0), (0, 2, 1),
                            (1, 0, 1), (1, 0, 2), (1, 2, 0), (1, 2, 1),
                            (2, 0, 1), (2, 0, 2), (2, 1, 0), (2, 1, 2)]:
                    res += dfs(idx + 1, cur)
                return res % MOD

            res = 0
            for nxt in helper(pre, 0):
                res += dfs(idx + 1, nxt)
            return res % MOD

        return dfs(0, (0, 0, 0)) % MOD
