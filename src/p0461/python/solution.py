class Solution:
    def hammingDistance(self, x: int, y: int) -> int:
        if x > y:
            x, y = y, x

        xs = bin(x)[2:]
        ys = bin(y)[2:]

        xs = '0' * (len(ys) - len(xs)) + xs

        ans = 0
        for i in range(len(xs)):
            if xs[i] != ys[i]:
                ans += 1

        return ans
