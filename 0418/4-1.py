from typing import List
from collections import deque


def aaa(jump):
    tt = 0
    start = 0
    while start < len(jump):
        if start + jump[start] >= len(jump):
            return tt
        tt += 1
        start += jump[start]

    return tt


class Solution:
    def minJump(self, jump: List[int]) -> int:
        if len(jump) == 1:
            return 0

        if jump[0] >= len(jump):
            return 1

        dq = deque()
        dq.append((0 + jump[0], 1))

        mark = set()
        mark.add(0)
        while dq:
            a, d = dq.popleft()
            mark.add(a)
            if a + jump[a] >= len(jump):
                return min(d, aaa(jump))
            else:
                if a + jump[a] in mark:
                    continue
                dq.append((a + jump[a], d + 1))
            if a - jump[a] >= 0 and a - jump[a] not in mark:
                dq.append((a - jump[a], d + 1))


if __name__ == '__main__':
    S = Solution()
    print(S.minJump([2, 5, 1, 1, 1, 1]))
