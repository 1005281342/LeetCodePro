from typing import List


class Solution:
    def busyStudent(self, startTime: List[int], endTime: List[int], queryTime: int) -> int:
        cnt = 0
        for a, b in zip(startTime, endTime):
            if a<=queryTime<=b:
                cnt += 1
        return cnt
