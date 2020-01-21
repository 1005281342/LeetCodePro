ht = [0]
for i in range(1, 11):
    ht.append(ht[-1] + i * 9 * (10 ** (i - 1)))


class Solution:

    def findNthDigit(self, n: int) -> int:

        if n <= 9:
            return n

        for ind, c in enumerate(ht):
            if n < c:
                ii = n - ht[ind - 1] - 1  # 在ind区间的第ii个序列
                k, v = 10 ** (ind - 1) + (ii // ind), (ii % ind)
                return int(str(k)[v])


if __name__ == '__main__':
    S = Solution()
    a = S.findNthDigit(10)
    print(a)
    a = S.findNthDigit(11)
    print(a)
    a = S.findNthDigit(12)
    print(a)
    a = S.findNthDigit(13)
    print(a)
    a = S.findNthDigit(14)
    print(a)
