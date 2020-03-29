from typing import List


class Solution:
    def numTeams(self, rating: List[int]) -> int:

        s = set()

        for i in range(len(rating)-2):
            for j in range(i+1, len(rating)-1):
                for z in range(j+1, len(rating)):
                    if rating[i] < rating[j] < rating[z] or rating[i] > rating[j] > rating[z]:
                        s.add((i, j, z))
        return len(s)