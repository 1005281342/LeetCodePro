from typing import List


class Solution:
    def diStringMatch(self, S: str) -> List[int]:
        left = 0
        right = len(S)
        ans = list()
        for c in S:
            if c == "I":
                ans.append(left)
                left += 1
            elif c == "D":
                ans.append(right)
                right -= 1
        ans.append(right)

        return ans
