class Solution:
    def cuttingRope(self, n: int) -> int:
        if n <= 3:
            return n - 1

        dp = [i for i in range(0, n+1)]
        for i in range(4, n + 1):
            for j in range(1, (i >> 1) + 1):
                dp[i] = max(dp[j] * dp[i - j], dp[i])

        return dp[n]
