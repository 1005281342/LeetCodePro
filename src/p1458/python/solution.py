from typing import List


class Solution:
    def maxDotProduct(self, nums1: List[int], nums2: List[int]) -> int:
        dp = [[-1005 for _ in range(len(nums2) + 1)] for _ in range(len(nums1) + 1)]

        for i in range(1, len(nums1) + 1):
            for j in range(1, len(nums2) + 1):
                dp[i][j] = nums1[i] * nums2[j]
                dp[i][j] = max(dp[i][j], dp[i][j - 1], dp[i - 1][j], dp[i - 1][j - 1] + dp[i][j])

        return dp[len(nums1)][len(nums2)]
