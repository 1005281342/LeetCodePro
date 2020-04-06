from collections import defaultdict


class Solution:

    def my_sum(self, i):
        ans = 0
        while i >= 10:
            ans += i % 10
            i //= 10
        ans += i
        return ans

    def countLargestGroup(self, n: int) -> int:
        dd = defaultdict(list)
        for i in range(1, 1 + n):
        # dd[i % 9].append(i)
            dd[self.my_sum(i)].append(i)
        # print(dd)
        max_length = 0
        for v in dd.values():
            if len(v) > max_length:
                max_length = len(v)
        ans = 0
        for v in dd.values():
            if len(v) == max_length:
                ans += 1
        return ans


if __name__ == '__main__':
    S = Solution()
    S.countLargestGroup(24)