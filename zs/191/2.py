from typing import List
import heapq


class Solution:
    def getStrongest(self, arr: List[int], k: int) -> List[int]:
        arr.sort()
        mid = arr[(len(arr) - 1) >> 1]
        ans = list()
        heapq.heapify(ans)

        for num in arr:
            heapq.heappush(ans, (abs(num - mid), num))

        return [b for _, b in heapq.nlargest(k, ans)]
