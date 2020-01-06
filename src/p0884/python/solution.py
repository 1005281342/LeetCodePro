from typing import List
from collections import Counter


class Solution:
    def uncommonFromSentences(self, A: str, B: str) -> List[str]:
        ac = Counter(A.split(" "))
        bc = Counter(B.split(" "))
        kk = set()
        a = set()
        b = set()
        for k, v in ac.items():
            if v == 1:
                a.add(k)
            else:
                kk.add(k)

        for k, v in bc.items():
            if v == 1:
                b.add(k)
            else:
                kk.add(k)

        ans = list()
        for c in a ^ b:
            if c not in kk:
                ans.append(c)
        return ans
