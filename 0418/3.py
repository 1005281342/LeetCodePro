from typing import List

from bisect import bisect_left as bisect


class Solution:
    def getTriggerTime(self, increase: List[List[int]], requirements: List[List[int]]) -> List[int]:
        a, b, c = 0, 0, 0
        ka = [0]
        kb = [0]
        kc = [0]
        for i, j, k in increase:
            a += i
            b += j
            c += k
            ka.append(a)
            kb.append(b)
            kc.append(c)

        t = list()
        for i, v in enumerate(requirements):
            c1, r, h = v
            if a >= c1 and b >= r and c >= h:
                t.append(i)
        ans = [-1] * len(requirements)
        for i in t:
            aa, bb, cc = requirements[i]
            index = max([bisect(ka, aa), bisect(kb, bb), bisect(kc, cc)])
            ans[i] = index

        for i, v in enumerate(requirements):
            aa, bb, cc = v
            if aa == 0 == bb == cc:
                ans[i] = 0
        return ans


if __name__ == '__main__':
    S = Solution()
    S.getTriggerTime([[2, 8, 4], [2, 5, 0], [10, 9, 8]]
                     , [[2, 11, 3], [15, 10, 7], [9, 17, 12], [8, 1, 14]])
