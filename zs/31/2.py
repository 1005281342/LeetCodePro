from typing import List


class Solution:
    def numOfSubarrays(self, arr: List[int]) -> int:
        res = 0
        d = [1, 0]
        tmp = 0
        for i in range(len(arr)):
            tmp += arr[i]
            if tmp % 2:
                res += d[0]
                res = res % (10 ** 9 + 7)
                d[1] += 1
            else:
                res += d[1]
                res = res % (10 ** 9 + 7)
                d[0] += 1
        return res
