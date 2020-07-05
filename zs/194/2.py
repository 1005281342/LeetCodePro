from typing import List


class Solution:
    def getFolderNames(self, names: List[str]) -> List[str]:

        dd = dict()
        ans = list()
        for name in names:
            if dd.get(name) is None:
                ans.append(name)
                dd[name] = 1
            else:
                new_name = name + "(" + str(dd[name]) + ")"
                while dd.get(new_name) is not None:
                    dd[name] += 1
                    new_name = name + "(" + str(dd[name]) + ")"
                ans.append(new_name)
                dd[name] += 1
                dd[new_name] = 1
        return ans

