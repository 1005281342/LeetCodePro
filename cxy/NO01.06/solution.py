class Solution:
    def compressString(self, S: str) -> str:
        if len(S) <= 1:
            return S
        ans = ""
        c = S[0]
        cnt = 1
        for v in S[1:]:
            if v != c:
                ans += c + str(cnt)
                cnt = 1
                c = v
                continue
            cnt += 1
        ans += c + str(cnt)
        return ans if len(ans) < len(S) else S


if __name__ == '__main__':
    S = Solution()

    a = ["aabcccccaa",
         "abbccd",
         "aabcccccaaa"]
    for s in a:
        print(S.compressString(s))
