from typing import List
from collections import defaultdict
from random import choices


class Solution:

    def __init__(self, nums: List[int]):
        self.dd = defaultdict(list)
        for i, v in enumerate(nums):
            self.dd[v].append(i)

    def pick(self, target: int) -> int:
        return choices(self.dd[target])

# Your Solution object will be instantiated and called as such:
# obj = Solution(nums)
# param_1 = obj.pick(target)
