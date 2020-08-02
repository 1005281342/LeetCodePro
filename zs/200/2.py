from typing import List


class Solution:
    def getWinner(self, arr: List[int], k: int) -> int:
        print(len(arr), k)
        if k >= len(arr):
            return max(arr)

        cnt = 0
        i, j = 0, 1
        while True:
            a, b = arr[i], arr[j]
            if cnt == k:
                return a
            if a > b:
                arr.append(b)
                j += 1
                cnt += 1
            else:
                cnt = 1
                arr.append(a)
                i = j
                j = j + 1
