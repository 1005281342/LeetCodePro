from collections import deque

base = ['a', 'b', 'c']
mp = {
    'a': ['b', 'c'],
    'b': ['a', 'c'],
    'c': ['a', 'b']
}


class Solution:

    def getHappyString(self, n: int, k: int) -> str:

        # if n < 8:
        #     return self.helper(n, k)

        dq = deque()
        dq.extend(base)
        while dq and k > 0:
            a = dq.popleft()
            if len(a) == n:
                # print(a)
                k -= 1
                continue
            dq.append(a + mp[a[-1]][0])
            dq.append(a + mp[a[-1]][1])

        if k > 0:
            return ""

        return a


if __name__ == '__main__':
    S = Solution()
    print(S.getHappyString(2, 3))
    print(S.getHappyString(1, 4))
    print(S.getHappyString(10, 100))
    print(S.getHappyString(7, 65))
    # S.getHappyString(3, 4)
