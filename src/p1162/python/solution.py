from typing import List
from math import inf
from collections import deque


class Solution:
    def maxDistance(self, grid: List[List[int]]) -> int:
        # 地图规模
        n, m = len(grid), len(grid[0])
        # 每个点到陆地的曼哈顿距离
        dist = [[inf for _ in range(m)] for _ in range(n)]
        # 该点是否被访问过
        visited = [[False for _ in range(m)] for _ in range(n)]
        # 队列
        dq = deque()
        # 陆地计数
        cnt = 0
        ans = 0
        tot = n * m
        for i in range(n):
            for j in range(m):
                if grid[i][j]:
                    dist[i][j] = 0
                    visited[i][j] = True
                    dq.append((i, j))
                    cnt += 1
        # 如果都是陆地或者都是海洋，则返回-1
        if cnt == tot or cnt == 0:
            return -1
        while dq:
            x, y = dq.popleft()  # 出列
            for i, j in [(x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1)]:
                # 如果坐标合法并且没被访问过
                if 0 <= i < n and 0 <= j < m and not visited[i][j]:
                    dist[i][j] = min(dist[i][j], dist[x][y] + 1)  # 取最小值
                    ans = max(ans, dist[i][j])  # 更新答案
                    visited[i][j] = True
                    dq.append((i, j))  # 入列
        return ans

# class Solution:
#     def maxDistance(self, grid: List[List[int]]) -> int:
#         le = len(grid)
#         luDi = list()
#         haiYan = list()
#         for i in range(le):
#             for j in range(le):
#                 if grid[i][j]:
#                     luDi.append((i, j))
#                 else:
#                     haiYan.append((i, j))
#         if not luDi or not haiYan:
#             return -1
#
#         ts = 0
#         ans = 205
#         for x, y in haiYan:
#             tts = 0
#             tans = 205
#             for a, b in luDi:
#                 mhd = abs(x - a) + abs(y - b)
#                 tts += mhd
#                 tans = min(mhd, tans)
#             if tts > ts:
#                 ts = tts
#                 ans = tans
#         return ans
