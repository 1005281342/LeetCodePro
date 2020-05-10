from typing import List
import heapq


class Solution:
    def minJump(self, jump: List[int]) -> int:
        if len(jump) == 1:
            return 0

        if jump[0] >= len(jump):
            return 1

        # tt = 0
        # start = 0
        # while start < len(jump):
        #     if start + jump[start] >= len(jump):
        #         return tt
        #     tt += 1
        #     start += jump[start]
        #
        # print(tt)

        dq = list()
        heapq.heapify(dq)
        dq.append((0 + jump[0], 1))

        mark = set()
        mark.add(0)
        while dq:
            a, d = heapq.heappop(dq)
            if a + jump[a] >= len(jump):
                return d
            else:
                if a + jump[a] in mark:
                    continue
                heapq.heappush(dq, (a + jump[a], d + 1))
            if a - jump[a] >= 0 and a - jump[a] not in mark:
                heapq.heappush(dq, (a - jump[a], d + 1))
            mark.add(a)


if __name__ == '__main__':
    S = Solution()
    print(S.minJump([2, 5, 1, 1, 1, 1]))