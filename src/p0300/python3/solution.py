from typing import List


class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        if len(nums) <= 1:
            return len(nums)

        dp = [1] * len(nums)
        max_length = 1

        for i in range(1, len(nums)):
            for j in range(0, i):
                if nums[i] > nums[j] and dp[i] < dp[j] + 1:
                    dp[i] = dp[j] + 1
            if dp[i] > max_length:
                max_length = dp[i]

        return max_length


if __name__ == '__main__':
    S = Solution()
    a = S.lengthOfLIS([10, 9, 2, 5, 3, 7, 101, 18])
    print(a)
