class Solution:
    def buddyStrings(self, A: str, B: str) -> bool:

        if len(A) != len(B):
            return False

        if A == B and len(A) != len(set(A)):
            return True

        sa = set()
        sb = set()
        for i in range(len(A)):
            if A[i] != B[i]:
                sa.add(A[i])
                sb.add(B[i])
        return sa == sb and len(sa) == 2
