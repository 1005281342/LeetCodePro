from typing import List

import heapq


class Solution:

    def __init__(self):
        self.lst = list()
        heapq.heapify(self.lst)

    def count(self, n):
        ans = 0
        while n > 1:
            if n & 1:
                n = 3 * n + 1
            else:
                n = n >> 1
            ans += 1
        return ans

    def getKth(self, lo: int, hi: int, k: int) -> int:

        for v in range(lo, hi+1):
            c = self.count(v)
            heapq.heappush(self.lst, (c, v))

        ans = heapq.nsmallest(k, self.lst)
        return ans[-1][-1]

if __name__ == '__main__':
    S = Solution()
    S.getKth(12, 15, 2)
