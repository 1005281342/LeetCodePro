from typing import List


class Solution:

    def minCount(self, coins: List[int]) -> int:

        cnt = 0
        for c in coins:
            cnt += (c+1) // 2

        return cnt



