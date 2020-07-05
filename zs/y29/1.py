from typing import List


class Solution:
    def average(self, salary: List[int]) -> float:
        salary.sort()
        t = salary[1:len(salary)-1]
        return sum(t) / len(t)
