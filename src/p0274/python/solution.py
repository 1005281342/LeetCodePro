from typing import List


class Solution:
    def hIndex(self, citations: List[int]) -> int:
        citations.sort()  # 排序
        for i in range(0, len(citations)):
            if len(citations) - i <= citations[i]:
                return len(citations) - i
