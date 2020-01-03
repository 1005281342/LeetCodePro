from typing import List

MOVE = {
    (0, 1), (1, 0), (-1, 0), (0, -1),
    (-1, -1), (-1, 1), (1, -1), (1, 1),
}


class Solution:
    def gameOfLife(self, board: List[List[int]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        ans = [[0 for _ in board[0]] for _ in board]
        for i in range(len(board)):
            for j in range(len(board[0])):
                cnt = 0
                for x, y in MOVE:
                    x += i
                    y += j
                    if 0 <= x < len(board) and 0 <= y < len(board[0]):
                        cnt += board[x][y]
                if board[i][j] == 0:
                    if cnt == 3:
                        ans[i][j] = 1
                else:
                    if cnt == 3 or cnt == 2:
                        ans[i][j] = 1
        for i in range(len(ans)):
            for j in range(len(ans[0])):
                board[i][j] = ans[i][j]
