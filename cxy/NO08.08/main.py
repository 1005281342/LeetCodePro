from typing import List
from itertools import permutations


class Solution:
    def permutation(self, S: str) -> List[str]:

        # return list({"".join(cs) for cs in permutations(S, len(S))})

        ans = []
        S = sorted(S)

        def backtrack(r, s):
            if not len(s):
                ans.append(r)
            else:
                pre = ''
                for i in range(len(s)):
                    if s[i] != pre:
                        backtrack(r + s[i], s[:i] + s[i + 1:])
                    pre = s[i]

        backtrack('', S)
        return ans
