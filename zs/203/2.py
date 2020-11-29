from typing import List


class Solution:
    def maxCoins(self, piles: List[int]) -> int:
        piles.sort(reverse=True)
        cnt = 0
        ans = 0
        index = 1
        while cnt < len(piles) // 3:
            ans += piles[index]
            index += 2
            cnt += 1
        return ans
