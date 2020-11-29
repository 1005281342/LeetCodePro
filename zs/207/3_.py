from typing import List


class Solution:
    def maxProductPath(self, grid: List[List[int]]) -> int:
        if grid[0][0] == 0 or grid[-1][-1] == 0:
            return 0
        m = len(grid)
        n = len(grid[0])
        dp = [[0 for _ in range(n)] for _ in range(m)]
        dp[0][0] = grid[0][0]
        for i in range(1, m):
            dp[i][0] = grid[i][0] = grid[i][0] * grid[i - 1][0]
        for j in range(1, n):
            dp[0][j] = grid[0][j] = grid[0][j] * grid[0][j - 1]
        for i in range(1, m):
            for j in range(1, n):
                if grid[i][j] < 0:
                    dp[i][j] = min(grid[i][j] * grid[i - 1][j], grid[i][j] * grid[i][j - 1])
                    grid[i][j] = max(grid[i][j] * dp[i - 1][j], grid[i][j] * dp[i][j - 1])
                elif grid[i][j] > 0:
                    dp[i][j] = min(grid[i][j] * dp[i - 1][j], grid[i][j] * dp[i][j - 1])
                    grid[i][j] = max(grid[i][j] * grid[i - 1][j], grid[i][j] * grid[i][j - 1])
        if grid[-1][-1] < 0:
            return -1
        return grid[-1][-1] % (10 ** 9 + 7)
