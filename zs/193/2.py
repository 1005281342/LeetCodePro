from typing import List

from collections import Counter
import heapq


class Solution:
    def findLeastNumOfUniqueInts(self, arr: List[int], k: int) -> int:
        cd = Counter(arr)
        lst = list()
        heapq.heapify(lst)
        for kk, v in cd.items():
            heapq.heappush(lst, (v, kk))

        while k > 0:
            v, t = heapq.heappop(lst)
            if k >= v:
                k -= v
                continue
            heapq.heappush(lst, (v-k, t))
            k -= v

        return len(lst)