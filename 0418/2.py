from typing import List

from collections import defaultdict, deque


class Solution:
    def numWays(self, n: int, relation: List[List[int]], k: int) -> int:
        start = 0
        end = n - 1
        dd = defaultdict(set)
        for a, b in relation:
            dd[a].add(b)
        mark = set()
        dq = deque()
        dq.append(((0,), 0))
        while dq:
            node, d = dq.popleft()
            if d == k and len(node) == k+1:
                if node[0] == 0 and node[-1] == n -1 :
                    mark.add(node)
                continue
            dq.extend([(tuple(list(node) + [a]), d + 1) for a in dd[node[-1]]])

        return len(mark)
