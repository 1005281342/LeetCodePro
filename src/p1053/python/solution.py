from typing import List


class Solution:
    def prevPermOpt1(self, A: List[int]) -> List[int]:
        left = len(A)

        for i in range(len(A) - 1, 0, -1):
            if A[i] < A[i - 1]:
                left = i - 1
                break

        if left < len(A):
            right = left + 1
            for i in range(left + 2, len(A)):
                if A[left] > A[i] > A[i - 1]:
                    right = i
            A[right], A[left] = A[left], A[right]

        return A
