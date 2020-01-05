class Solution:
    def countSegments(self, s: str) -> int:
        return len([a for a in s.split(" ") if a != ""])