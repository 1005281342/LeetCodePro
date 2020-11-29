from typing import List


class Solution:
    def reorderSpaces(self, text: str) -> str:
        cnt = text.count(" ")
        t = [x for x in text.split(" ") if x]
        if len(t) == 1:
            return t[0] + (cnt * " ")

        d = cnt // (len(t) - 1)
        m = cnt % (len(t) - 1)
        ans = (d * " ").join(t) + (m * " ")
        return ans
