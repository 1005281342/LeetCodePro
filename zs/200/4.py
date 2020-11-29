from typing import List
from collections import defaultdict
from bisect import bisect


def l_bin_search(arr, target):
    left = 0
    right = len(arr) - 1
    while left < right:
        mid = left + ((right - left) >> 1)
        if target < arr[mid]:
            right = mid
        elif target > arr[mid]:
            left = mid + 1
        else:
            return mid
    return -1


class Solution:
    def maxSum(self, nums1: List[int], nums2: List[int]) -> int:
        return max(self.h(nums1, nums2), self.h(nums2, nums1)) % (10 ** 9 + 7)

    def h(self, nums1, nums2):
        s1 = [0] * (len(nums1) + 1)
        s2 = [0] * (len(nums2) + 1)

        for i, a in enumerate(nums1):
            s1[i + 1] = s1[i] + a

        for i, a in enumerate(nums2):
            s2[i + 1] = s2[i] + a

        idx = list()
        for i, a in enumerate(nums1):
            j = l_bin_search(nums2, a)
            if j == -1:
                continue
            idx.append((i, j))

        if not idx:
            return max(s1[-1], s2[-1])

        if len(idx) == 1:
            return s1[idx[0][0]] + max(s2[-1] - s2[idx[0][1]], s1[-1] - s1[idx[0][0]])

        print(idx)
        cnt = 0
        t1, t2 = idx[cnt]
        ans = max(s1[t1], s2[t2])
        while cnt < len(idx) - 1:
            cnt += 1
            a, b = idx[cnt]
            ans += max(s1[a] - s1[t1], s2[b] - s2[t2])
            ans %= (10 ** 9 + 7)
            print(a, b, t1, t2, ans)
            t1, t2 = a, b
        ans += max(s1[-1] - s1[t1], s2[-1] - s2[t2])
        return ans
