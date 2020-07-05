class Solution:
    def kthFactor(self, n: int, k: int) -> int:

        cnt = 0
        for x in range(1, n + 1):
            if not n % x:
                cnt += 1
                if cnt == k:
                    return x

        return -1