from typing import List


class Solution:
    def findPoisonedDuration(self, timeSeries: List[int], duration: int) -> int:
        start = 0
        cnt = 0
        for a in timeSeries:
            if a >= start:
                cnt += duration
            else:
                cnt += (a + duration) - start
            start = a + duration

        return cnt
