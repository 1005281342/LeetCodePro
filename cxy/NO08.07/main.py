from typing import List
from itertools import permutations


class Solution:
    def permutation(self, S: str) -> List[str]:
        return ["".join(x) for x in (permutations(S))]


if __name__ == '__main__':
    S = Solution()
    print(S.permutation("qwe"))
