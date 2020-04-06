class Solution:
    def numSteps(self, s: str) -> int:
        num = int(s, 2)
        cnt = 0
        while num > 1:
            if num & 1:
                num += 1
            else:
                num >>= 1
            cnt += 1
        return cnt


if __name__ == '__main__':
    S = Solution()
    S.numSteps("")