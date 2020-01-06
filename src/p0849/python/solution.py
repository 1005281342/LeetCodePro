from typing import List


class Solution:

    def helper(self, a):
        start = 0
        t = 1
        for i, v in enumerate(a):
            if v == 1:
                t = max((i - start + 1) // 2, t)
                start = i + 1

        t = max((i - start + 1), t)

        return t

    def maxDistToClosest(self, seats: List[int]) -> int:

        if sum(seats) == 1:
            x = seats.index(1)
            return max(x - 0, len(seats) - 1 - x)

        return max(self.helper(seats), self.helper(seats[::-1]))
