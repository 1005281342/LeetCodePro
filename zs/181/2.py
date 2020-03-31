from typing import List
import math


class Solution:

    def __init__(self):
        self.ans = 0

    def count(self, a):
        s = set()
        for i in range(1, int(math.sqrt(a)) + 1):
            if a % i == 0:
                s.add(i)
                s.add(a // i)
        if len(s) == 4:
            # print(self.ans)
            self.ans += sum(s)

    def sumFourDivisors(self, nums: List[int]) -> int:

        for num in nums:
            self.count(num)
        # print(self.ans)
        return self.ans


if __name__ == '__main__':
    S = Solution()
    # S.sumFourDivisors([1, 2, 3, 4, 5])
    S.sumFourDivisors([21, 21])
