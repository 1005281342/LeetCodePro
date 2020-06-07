from collections import defaultdict


class Solution:
    def arrangeWords(self, text: str) -> str:
        dd = defaultdict(list)
        lst = text.lower().split(" ")
        for c in lst:
            dd[len(c)].append(c)

        ans = list()
        for k in sorted(dd.keys()):
            ans.extend(dd[k])

        if ans and ans[0]:
            ans[0] = str(ans[0]).capitalize()

        return " ".join(ans)