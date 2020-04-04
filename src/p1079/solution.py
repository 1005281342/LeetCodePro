from itertools import permutations


class Solution:
    def numTilePossibilities(self, tiles: str) -> int:
        cnt = 0
        for i in range(1, len(tiles) + 1):
            cnt += len({"".join(vs) for vs in permutations(tiles, i)})
        return cnt



