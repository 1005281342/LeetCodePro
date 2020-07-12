from typing import List
from collections import defaultdict, deque
import heapq


class Solution:
    def maxProbability(self, n: int, edges: List[List[int]], succProb: List[float], start: int, end: int) -> float:
        dd = defaultdict(set)
        for i, x in enumerate(edges):
            a, b = x
            dd[a].add((b, succProb[i]))
            dd[b].add((a, succProb[i]))

        dist = [0] * (n + 1)
        st = [False] * (n + 1)

        dist[start] = 1
        dq = deque()
        dq.append(start)
        st[start] = True

        while dq:
            node = dq.popleft()
            st[node] = False

            # 处理每条边
            for new_node, wight in dd[node]:
                if dist[new_node] < dist[node] * wight:
                    dist[new_node] = dist[node] * wight
                    if not st[new_node]:
                        dq.append(new_node)
                        st[new_node] = True
        return dist[end]
