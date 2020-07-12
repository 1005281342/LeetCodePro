from typing import List


class Solution:
    def buildArray(self, target: List[int], n: int) -> List[str]:
        ans = list()
        # print(target[0])
        v = set(target)
        if not v:
            return ans

        for i in range(1, n + 1):
            if i not in v:
                ans.append("Push")
                ans.append("Pop")
            else:
                v.remove(i)
                ans.append("Push")
        return ans
