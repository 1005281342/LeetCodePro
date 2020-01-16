class Solution:
    def reverseOnlyLetters(self, S: str) -> str:

        si = set()
        slst = list(S)
        a = set(range(65, 91)) | set(range(91, 123))
        for i, v in enumerate(S):
            if ord(v) in a:
                slst.append(v)
            else:
                si.add(i)

        ans = ""
        for i in range(len(S)):
            if i not in si:
                ans += slst.pop()
            else:
                ans += S[i]
        return ans

