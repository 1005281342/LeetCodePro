from typing import List
from collections import defaultdict

left = {2, 3, 4, 5}
mid = {4, 5, 6, 7}
right = {6, 7, 8, 9}


class Solution:
    def maxNumberOfFamilies(self, n: int, reservedSeats: List[List[int]]) -> int:
        ans = 0
        di = defaultdict(set)
        for i, j in reservedSeats:
            di[i].add(j)
        print(di)

        x = len(di)
        for v in di.values():
            if not (v & (left | right)):
                ans += 2
            elif (not (v & mid)) or (not (v & left)) or (not (v & right)):
                ans += 1
            print(v, ans)

        print(ans + (n - x) * 2)
        return ans + (n - x) * 2


if __name__ == '__main__':
    S = Solution()
    # S.maxNumberOfFamilies(3, [[1, 2], [1, 3], [1, 8], [2, 6], [3, 1], [3, 10]])
    S.maxNumberOfFamilies(2,
                          [[1, 6], [1, 8], [1, 3], [2, 3], [1, 10], [1, 2], [1, 5], [2, 2], [2, 4], [2, 10], [1, 7],
                           [2, 5]])
