from typing import List


class Solution:
    def kLengthApart(self, nums: List[int], k: int) -> bool:

        if len(nums) <= 1:
            return True
        cnt = 0
        try:
            index = nums.index(1)
        except:
            return True
        for i, v in enumerate(nums[index + 1:]):
            if v == 0:
                cnt += 1
            else:
                if cnt < k:
                    return False
                cnt = 0
            # print(cnt, i, v)
        return True


if __name__ == '__main__':
    S = Solution()
    print(S.kLengthApart([1, 0, 0, 0, 1, 0, 0, 1], 2))
    print(S.kLengthApart([1, 1, 1, 0], 3))
