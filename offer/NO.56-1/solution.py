from typing import List


class Solution:
    def singleNumbers(self, nums: List[int]) -> List[int]:
        xor, ans = 0, [0, 0]
        for num in nums:
            xor ^= num
        low = xor ^ (xor - 1) & xor
        print(low)
        for num in nums:
            print(not num & low)
            ans[not num & low] ^= num
        return ans


if __name__ == '__main__':

    S = Solution()
    ans = S.singleNumbers([4,4,2,6])
    print(ans)