from typing import List
from collections import Counter


class Solution:
    def mostCommonWord(self, paragraph: str, banned: List[str]) -> str:
        paragraph = paragraph.lower()
        for c in "!?',;.":
            paragraph = paragraph.replace(c, '')
        lst = [w for w in paragraph.split(" ") if w not in banned]
        cd = Counter(lst)
        return cd.most_common(1)[0][0]
