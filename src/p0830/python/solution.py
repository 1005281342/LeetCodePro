from typing import List


class Solution:
    def largeGroupPositions(self, S: str) -> List[List[int]]:
        t = S[0]
        start = 0
        end = 0
        ans = list()
        for i, v in enumerate(S[1:]):
            i += 1
            if v == t:
                end = i
            else:
                if end - start >= 2:
                    ans.append([start, end])
                start = i
                t = v

        if end - start >= 2:
            ans.append([start, end])

        return ans
