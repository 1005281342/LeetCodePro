class Solution:
    def xorOperation(self, n: int, start: int) -> int:
        ans = 0
        for i in range(start, start + 2 * n, 2):
            ans ^= i
        return ans

