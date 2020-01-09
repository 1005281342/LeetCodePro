from typing import List

tmp = [[0 for _ in range(3)] for _ in range(3)]


class Solution:

    def is_magic(self):

        s = set()
        for a in range(3):
            for b in range(3):
                if not (0 < tmp[a][b] <= 9):
                    return False
                s.add(tmp[a][b])
        if len(s) != 9:
            return False

        ans = True

        for c in [[tmp[0][0], tmp[1][1], tmp[2][2]],
                  [tmp[2][0], tmp[1][1], tmp[0][2]],
                  tmp[0], tmp[1], tmp[2],
                  [tmp[0][0], tmp[1][0], tmp[2][0]],
                  [tmp[0][1], tmp[1][1], tmp[2][1]],
                  [tmp[0][2], tmp[1][2], tmp[2][2]],
                  ]:
            if sum(c) != 15:
                ans = False
                break
        return ans

    def numMagicSquaresInside(self, grid: List[List[int]]) -> int:

        move = {(1, 0), (0, 1), (-1, 0), (0, -1),
                (1, 1), (1, -1), (-1, 1), (-1, -1)
                }
        cnt = 0
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] != 5:
                    continue
                flag = True
                for a, b in move:
                    x = a + i
                    y = b + j
                    if not (0 <= x < len(grid) and 0 <= y < len(grid[0])):
                        flag = False
                        break
                    tmp[1 + a][1 + b] = grid[x][y]
                tmp[1][1] = 5
                if flag and self.is_magic():
                    cnt += 1
        return cnt
