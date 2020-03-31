from typing import List
from itertools import permutations


class Solution:
    def permutation(self, s: str) -> List[str]:
        return list({"".join(cs) for cs in permutations(s)})
