from typing import List


class Solution:
    def minimalSteps(self, maze: List[str]) -> int:
        m, n = len(maze), len(maze[0])

        maze = [[line[i] for i in range(n)] for line in maze]

        dummpyrec = (0, 0)

        def move(i, j, a=None, b=None):
            nonlocal dummpyrec
            if not (0 <= i < m and 0 <= j < n) or maze[i][j] == " " or maze[i][j] == "#":
                return float("inf")
            if a is not None and b is not None and a == i and b == j:
                return 0
            if maze[i][j] == "O":
                dummpyrec = (i, j)
                return 0
            maze[i][j], restore = ' ', maze[i][j]
            res = 1 + min(move(i + 1, j, a, b), move(i - 1, j, a, b), move(i, j + 1, a, b), move(i, j - 1, a, b))
            maze[i][j] = restore
            return res

        pair = {}
        m_dist = []
        target = [0, 0]

        for i in range(m):
            for j in range(n):
                if maze[i][j] == "T":
                    target = (i, j)
                if maze[i][j] == "M":
                    dummpyrec = (0, 0)
                    minpath = move(i, j)
                    pair[(i, j)] = dummpyrec
                    m_dist.append([dummpyrec, minpath])

        paths = [x[1] for x in sorted(m_dist, key=lambda x: x[0][0] + x[0][1])]

        order = sorted([dot[1] for dot in pair.items()], key=lambda x: x[0] + x[1])

        for k, v in pair.items():
            if v == order[-1]:
                break

        if any([x == float("inf") for x in paths]):
            return -1

        return move(0, 0, order[0][0], order[0][1]) + 2 * sum(paths[:-1]) + paths[-1] + abs(
            target[0] - k[0]) + abs(target[1] - k[1])
