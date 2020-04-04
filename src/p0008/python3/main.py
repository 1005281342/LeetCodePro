class Solution:
    def myAtoi(self, a: str) -> int:
        a = a.strip()
        if len(a) <= 0:
            return 0
        if a[0] not in "+-0123456789":
            return 0
        f = True
        if a[0] == "-":
            f = False
        index = 0
        if a[0] in "+-":
            index = 1
        ans_s = ""
        for i in range(index, len(a)):
            if a[i] in "0123456789":
                ans_s += a[i]
            else:
                break
        if not ans_s:
            return 0
        if f:
            return min(int(ans_s), 2147483647)
        else:
            return max(-int(ans_s), -2147483648)
