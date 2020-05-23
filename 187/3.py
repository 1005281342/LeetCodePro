from typing import List

from collections import deque


class Solution:

    def longestSubarray(self, nums: List[int], limit: int) -> int:
        min_q, max_q = deque(), deque()
        left = 0
        ans = 0

        for i, num in enumerate(nums):
            while min_q and min_q[-1] > num:
                min_q.pop()
            min_q.append(num)

            while max_q and max_q[-1] < num:
                max_q.pop()
            max_q.append(num)

            print(min_q)
            print(max_q)

            while left < len(nums) and (
                    (min_q and abs(num - min_q[0]) > limit) or (max_q and abs(num - max_q[0]) > limit)):
                if min_q and nums[left] == min_q[0]:
                    min_q.popleft()
                if max_q and nums[left] == max_q[0]:
                    max_q.popleft()
                left += 1
            ans = max(ans, i + 1 - left)
        return ans


# import bisect
#
#
# class Solution:
#     def longestSubarray(self, nums: List[int], limit: int) -> int:
#         l, r, A, n, ret = 0, 0, [], len(nums), 0
#         while r < n:
#             bisect.insort(A, nums[r])
#             print(A)
#             while A[-1] - A[0] > limit:
#                 print(len(A), bisect.bisect(A, nums[l]))
#                 del A[bisect.bisect(A, nums[l]) - 1]
#                 l += 1
#             ret = max(ret, r - l + 1)
#             r += 1
#         return ret


if __name__ == '__main__':
    S = Solution()
    S.longestSubarray([10, 1, 2, 4, 7, 2], 5)
