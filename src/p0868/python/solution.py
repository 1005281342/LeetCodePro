class Solution:
    def binaryGap(self, N: int) -> int:
        t = [i for i, v in enumerate(bin(N)) if v == '1']
        ans = 0
        for i in range(len(t) - 1):
            ans = max(t[i + 1] - t[i], ans)
        return ans
