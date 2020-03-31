from typing import List


class Solution:
    def createTargetArray(self, nums: List[int], index: List[int]) -> List[int]:
        ans = list()
        for i, v in zip(index, nums):
            ans.insert(i, v)
        print(ans)
        return ans


if __name__ == '__main__':

    S =Solution()
    S.createTargetArray(nums = [0,1,2,3,4], index = [0,1,2,2,1])