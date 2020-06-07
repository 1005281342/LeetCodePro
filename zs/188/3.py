from typing import List
from collections import defaultdict, deque


class Solution:

    def minTime(self, n: int, edges: List[List[int]], hasApple: List[bool]) -> int:

        dd = defaultdict(set)
        has = set()
        for i, v in enumerate(hasApple):
            if v:
                has.add(i)

        for v in edges:
            dd[v[0]].add(v[1])

        def downHasApple(a) -> bool:
            ddq = deque()
            ddq.append(a)
            while ddq:
                node = ddq.popleft()
                vs = dd[node]
                if set(vs) & has:
                    return True
                ddq.extend(vs)
            return False

        dq = deque()
        dq.append((0, 0))
        cnt = 0

        anst = set()
        while dq:
            node, d = dq.popleft()
            if not downHasApple(node):
                cnt += 2 * d
                anst.add(node)
                continue

            dq.extend([(x, d + 1) for x in dd[node]])

        mark = dict()
        dq = deque()
        mark[0] = "0"

        for c in dd[0]:
            v = str(c) + "0"
            dq.append(v)
            mark[c] = v

        while len(mark) < n and dq:
            x = dq.popleft()
            for c in dd[int(x[0])]:
                v = str(c) + x
                dq.append(v)
                mark[c] = v
        print(mark)

        ans = set()
        for node in anst:
            for xx in mark[node]:
                ans.add(xx)
        ans.remove('0')
        return min(2 * len(ans), cnt)
