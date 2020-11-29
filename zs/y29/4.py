from typing import List
from collections import defaultdict, deque
from bisect import insort


class Solution:
    def minNumberOfSemesters(self, n: int, dependencies: List[List[int]], k: int) -> int:
        dd = defaultdict(set)
        dc = defaultdict(int)
        dk = defaultdict(int)
        for a, b in dependencies:
            dd[a].add(b)
            dc[b] += 1
            dk[a] += 1

        dq = deque()
        dq.extend([(i, 0) for i in range(1, n + 1) if dc[i] == 0])

        df = defaultdict(deque)
        mark = set()
        while dq:
            node, cas = dq.popleft()
            if node in mark:
                continue
            insort(df[cas], (dk[node], node))
            mark.add(node)
            for t in dd[node]:
                dc[t] -= 1
                if dc[t] == 0:
                    dq.append((t, cas + 1))
        # TODO