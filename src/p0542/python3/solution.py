from typing import List
from collections import deque

MOVE = {
    (0, 1),
    (0, -1),
    (1, 0),
    (-1, 0),
}


class Solution:
    def updateMatrix(self, matrix: List[List[int]]) -> List[List[int]]:

        if not matrix or not matrix[0]:
            return matrix

        dq = deque()
        for i in range(len(matrix)):
            for j in range(len(matrix[0])):
                if matrix[i][j] == 0:
                    dq.append((i, j))
                else:
                    matrix[i][j] = len(matrix) + len(matrix[0])
        while dq:
            x0, y0 = dq.popleft()
            for x, y in MOVE:
                x += x0
                y += y0
                if 0 <= x < len(matrix) and 0 <= y < len(matrix[0]) and matrix[x][y] > matrix[x0][y0] + 1:
                    matrix[x][y] = matrix[x0][y0] + 1
                    dq.append((x, y))

        return matrix
