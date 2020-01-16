class Solution:
    def subtractProductAndSum(self, n: int) -> int:
        a = 1
        b = 0
        while n > 0:
            n, v = n // 10, n % 10
            a *= v
            b += v
        return a - b
