from typing import List

ans = list()
for i in range(1, int(1e4)):
    si = str(i)
    if "0" in si:
        continue
    flag = True
    for c in si:
        if i % int(c):
            flag = False
            break
    if flag:
        ans.append(i)


class Solution:
    def selfDividingNumbers(self, left: int, right: int) -> List[int]:
        res = list()
        for a in ans:
            if a > right:
                break
            if left <= a:
                res.append(a)
        return res
