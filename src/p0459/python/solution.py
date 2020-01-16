class Solution:
    def repeatedSubstringPattern(self, s: str) -> bool:
        return s in (s + s)[1: len(s) * 2 - 1]
