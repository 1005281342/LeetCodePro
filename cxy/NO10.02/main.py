from typing import List
from collections import defaultdict


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        dd = defaultdict(list)
        for s in strs:
            k = "".join(sorted(s))
            dd[k].append(s)

        return [*dd.values()]


if __name__ == '__main__':
    S = Solution()
    ans = S.groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"])
    print(ans)
