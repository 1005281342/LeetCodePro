"""
题目类似于找最大环的长度
因为没有重复元素，所以每个元素都属于一个环，每个环只计算一次，每个走过的元素置为-1（也可以置为负数）。从开头遍历数组，如果该元素为-1则跳过
时间复杂度O（n），空间复杂度O（1）
"""

from typing import List


class Solution:
    def arrayNesting(self, nums: List[int]) -> int:
        res = 0
        for i in range(len(nums)):
            if nums[i] == -1:
                continue
            temp = 1
            path_index = i
            while nums[path_index] != i:
                nums[path_index], path_index = -1, nums[path_index]
                temp += 1
            nums[path_index] = -1
            res = max(temp, res)
        return res
