from typing import List


class Solution:
    def isRectangleOverlap(self, rec1: List[int], rec2: List[int]) -> bool:
        return not (rec1[2] <= rec2[0] or  # left
                    rec1[3] <= rec2[1] or  # bottom
                    rec1[0] >= rec2[2] or  # right
                    rec1[1] >= rec2[3])  # top


if __name__ == '__main__':
    S = Solution()
    ans = S.isRectangleOverlap(rec1=[0, 0, 2, 2], rec2=[1, 1, 3, 3])
    print(ans)

    ans = S.isRectangleOverlap([5, 15, 8, 18], [0, 3, 7, 9])
    print(ans)
