class Solution:
    def isUgly(self, num: int) -> bool:
        if 0 >= num:
            return False
        s = True
        while s:
            s = False
            if num % 5 == 0:
                num /= 5
                s = True
            if num % 3 == 0:
                num /= 3
                s = True
            if num % 2 == 0:
                num /= 2
                s = True
        return num == 1