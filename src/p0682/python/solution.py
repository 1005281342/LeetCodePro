from typing import List


class Solution:
    def calPoints(self, ops: List[str]) -> int:

        lst = list()
        for c in ops:
            if c == "C" and lst:
                lst.pop()
            elif c == "D" and lst:
                lst.append(lst[-1] * 2)
            elif c == "+" and len(lst) >= 2:
                lst.append(lst[-2] + lst[-1])
            elif c not in {"C", "D", "+"}:
                lst.append(int(c))

        return sum(lst)
