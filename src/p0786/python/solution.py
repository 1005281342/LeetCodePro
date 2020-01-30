from typing import List
from itertools import permutations
import heapq


# 超时
class Solution:
    def kthSmallestPrimeFraction(self, A: List[int], K: int) -> List[int]:
        hq = list()
        heapq.heapify(hq)

        for a, b in permutations(A, 2):
            heapq.heappush(hq, (-a / b, [a, b]))

        return heapq.nlargest(K, hq)[-1][-1]
