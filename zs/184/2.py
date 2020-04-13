from typing import List
from collections import deque


class Solution:
    def processQueries(self, queries: List[int], m: int) -> List[int]:
        t = list(range(1, m+1))
        dq = deque(queries)
        ans = list()
        while dq:
            v = dq.popleft()
            for i, vv in enumerate(t):
                if vv == v:
                    t = [t[i]] + t[:i] + t[i + 1:]
                    # print(t, "\n")
                    ans.append(i)
                    break
        return ans

if __name__ == '__main__':
    S = Solution()
    print(S.processQueries(queries = [3,1,2,1], m = 5))
    print(S.processQueries([4,1,2,2], m = 4))
    print(S.processQueries(queries = [7,5,5,8,3], m = 8))
