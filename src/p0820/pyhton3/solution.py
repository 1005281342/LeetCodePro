from typing import List


class Solution:
    def minimumLengthEncoding(self, words: List[str]) -> int:
        words.sort(key=lambda i: len(i), reverse=True)

        s = ""
        for w in words:
            if w in s and w + "#" in s:
                continue
            s += w + "#"
        return len(s)
