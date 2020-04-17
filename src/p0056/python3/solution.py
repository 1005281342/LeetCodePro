from typing import List


class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort(reverse=True)
        if len(intervals) == 1:
            return intervals

        ans = list()
        x, y = intervals.pop()
        while intervals:
            a, b = intervals.pop()
            if a <= x:
                x = a
                y = max(b, y)
            elif a <= y:
                y = max(b, y)
            else:
                ans.append([x, y])
                x, y = a, b
        ans.append([x, y])
        return ans

