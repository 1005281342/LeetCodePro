mod = 10 ** 9 + 7


class Solution:
    def numOfWays(self, n: int) -> int:
        aa, bb = [0] * n, [0] * n
        aa[0], bb[0] = 6, 6
        for i in range(1, n):
            aa[i] = (2 * aa[i - 1] + 2 * bb[i - 1]) % mod
            bb[i] = (2 * aa[i - 1] + 3 * bb[i - 1]) % mod
        return (aa[-1]+bb[-1]) % mod
