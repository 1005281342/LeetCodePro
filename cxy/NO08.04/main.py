from typing import List
from itertools import combinations


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        ans = [[], nums]
        for i in range(1, len(nums)):
            ans.extend([list(a) for a in combinations(nums, i)])
        return ans


if __name__ == '__main__':
    S = Solution()
    ans = S.subsets([1, 2, 3])
    print(ans)