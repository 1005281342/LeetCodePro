from collections import Counter


class Solution:
    def judgeCircle(self, moves: str) -> bool:
        cd = Counter(moves)
        return cd.get("U") == cd.get("D") and cd.get("L") == cd.get("R")


if __name__ == '__main__':

    S = Solution()
    ans = S.judgeCircle("RLUURDDDLU")
    print(ans)