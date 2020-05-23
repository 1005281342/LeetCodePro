from typing import List
from collections import defaultdict, Counter


class Solution:
    def displayTable(self, orders: List[List[str]]) -> List[List[str]]:

        cai = set()
        dd = defaultdict(list)

        for a, b, c in orders:
            dd[int(b)].append(c)
            cai.add(c)
        cai = sorted(list(cai))
        ans = list()
        ans.append(["Table"] + cai)

        for k in sorted(dd.keys()):
            v = dd[k]
            cd = Counter(v)
            t = list()
            for i in cai:
                if not cd.get(i):
                    t.append('0')
                else:
                    t.append(str(cd[i]))

            ans.append([str(k)] + t)
        return ans


if __name__ == '__main__':
    S = Solution()
    S.displayTable(
        [["James", "12", "Fried Chicken"], ["Ratesh", "12", "Fried Chicken"], ["Amadeus", "12", "Fried Chicken"],
         ["Adam", "1", "Canadian Waffles"], ["Brianna", "1", "Canadian Waffles"]])
