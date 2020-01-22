from typing import List
from collections import defaultdict


class Solution:
    def findRepeatedDnaSequences(self, s: str) -> List[str]:

        dd = defaultdict(int)
        for i in range(len(s) - 9):  # len(s) - 10 + 1
            dd[s[i:i + 10]] += 1
        ans = list()
        for k, v in dd.items():
            if v > 1:
                ans.append(k)
        return ans
