from typing import List
from bisect import insort, bisect
from collections import defaultdict


class Solution(object):
    def lengthOfLIS(self, nums: List[int]) -> int:
        if len(nums) <= 1:
            return len(nums)

        max_value = 0

        d = defaultdict(int)
        lst = list()
        for num in nums:

            index = bisect(lst, num)

            if index == 0:
                d[num] = 1

            elif index == len(lst):
                max_value += 1
                d[num] = max_value

            elif lst[index - 1] < num:
                d[num] = max(d[num], d[lst[index - 1]] + 1)
                max_value = max(max_value, d[num])
            else:
                d[num] = max(max_value, d[num])
            insort(lst, num)
        print(d)
        return max(d.values())


# class Solution:
#     def lengthOfLIS(self, nums: List[int]) -> int:
#         if len(nums) <= 1:
#             return len(nums)
#
#         dp = [1] * len(nums)
#         max_length = 1
#
#         for i in range(1, len(nums)):
#             for j in range(0, i):
#                 if nums[i] > nums[j] and dp[i] < dp[j] + 1:
#                     dp[i] = dp[j] + 1
#             if dp[i] > max_length:
#                 max_length = dp[i]
#
#         return max_length

from collections import Counter


def a(string):
    cd = Counter(string)

    if cd["1"] == 0:
        if cd["0"] > 2:
            return "Impossible"
        else:
            return string

    ans = ""
    while cd["0"] > 1 and cd["1"] > 0:
        ans = "00" + ans
        cd["0"] -= 2
        ans = "1" + ans
        cd["1"] -= 1
    if cd["0"] == 1:
        if cd["1"] == 0:
            return "0" + ans
        else:
            return "Impossible"
    return cd["1"] * "1" + ans


if __name__ == '__main__':
    print(a("01110"))
    print(a("000"))
    # S = Solution()
    # a = S.lengthOfLIS([10, 9, 2, 5, 3, 7, 101, 18])
    # print(a)

    # lst = list()
    # print(bisect(lst, 1000))
    #
    # insort(lst, 2)
    # insort(lst, 10)
    # insort(lst, 1)
    # print(lst)
    #
    # print(bisect(lst, 1000))
    # print(bisect(lst, 1))
    # print(bisect(lst, 1.5))
    # print(bisect(lst, 0))
