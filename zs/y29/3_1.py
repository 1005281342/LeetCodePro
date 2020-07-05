from typing import List
from collections import deque


class Solution:
    def __init__(self):
        self.ans = 0

    def longestSubarray(self, nums: List[int]) -> int:
        if sum(nums) == len(nums):
            return max(0, len(nums) - 1)

        if len(nums) == 1 or sum(nums) == 0:
            return 0

        dq = deque()
        dq.append(0)
        while dq:
            left = dq.popleft()
            # print(left)
            if left >= len(nums):
                continue

            while left < len(nums):
                if not nums[left]:
                    left += 1
                else:
                    break

            flag = False

            right = left + 1
            if right >= len(nums):
                continue

            tmp_r = right

            cnt = 0

            while right < len(nums):
                if nums[right] == 1:
                    right += 1
                    continue

                if not flag:
                    tmp_r = right
                    cnt += 1
                    right += 1
                    flag = True
                    continue

                self.ans = max(right - left - cnt, self.ans)
                # print(right, left, cnt, self.ans)

                if right < len(nums):
                    dq.append(tmp_r)
                break
            self.ans = max(right - left - cnt, self.ans)
            # print(right, left, cnt, self.ans)

        # def helper(left):
        #     if left >= len(nums):
        #         return
        #
        #     while left < len(nums):
        #         if not nums[left]:
        #             left += 1
        #         else:
        #             break
        #
        #     flag = False
        #
        #     right = left + 1
        #     if right >= len(nums):
        #         return
        #
        #     tmp_r = right
        #
        #     cnt = 0
        #
        #     while right < len(nums):
        #         if nums[right] == 1:
        #             right += 1
        #             continue
        #
        #         if not flag:
        #             tmp_r = right
        #             cnt += 1
        #             right += 1
        #             flag = True
        #             continue
        #
        #         self.ans = max(right - left - cnt, self.ans)
        #         print(right, left, cnt, self.ans)
        #
        #         if right < len(nums):
        #             helper(tmp_r)
        #
        #     self.ans = max(right - left - cnt, self.ans)
        #     print(right, left, cnt, self.ans)
        #     # helper(tmp_r)
        #
        # helper(left=0)
        return self.ans


if __name__ == '__main__':
    S = Solution()
    res = S.longestSubarray([0, 1, 1, 1, 0, 1, 1, 0, 1])
    print(res)
