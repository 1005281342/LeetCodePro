from typing import List


class Solution:

    def helper(self, t):

        if len(t) < 2:
            return False

        for c in t[1:]:
            if (c > 0) != (t[0] > 0):
                return False
        return True

    def circularArrayLoop(self, nums: List[int]) -> bool:

        for i, v in enumerate(nums):
            tmp = [v, ]
            s = {i, }
            m = i + v
            m = m % len(nums)
            while m not in s:
                tmp.append(nums[m])
                s.add(m)
                m += nums[m]
                m = m % len(nums)
            print(tmp, self.helper(tmp))
            if m == i and self.helper(tmp):
                return True

        return False


if __name__ == '__main__':
    S = Solution()
    S.circularArrayLoop([2, -1, 1, 2, 2])
