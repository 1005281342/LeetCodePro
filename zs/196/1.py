from typing import List

class Solution:
    def canMakeArithmeticProgression(self, arr: List[int]) -> bool:


        arr.sort()

        d = arr[1] - arr[0]
        pre = arr[1]
        for num in arr[2:]:
            if num - pre != d:
                return False
            pre = num
        return True
