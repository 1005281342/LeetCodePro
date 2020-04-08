class Solution:
    def findComplement(self, num: int) -> int:
        a = bin(num)[2:]
        b = ""
        for c in a:
            if c == "1":
                b += "0"
            else:
                b += "1"
        return int(b, 2)


if __name__ == '__main__':
    S = Solution()
    aa = S.findComplement(5)
    print(aa)
