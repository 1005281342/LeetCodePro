from typing import List


class Solution:
    def stringMatching(self, words: List[str]) -> List[str]:

        ans = list()
        for i in range(len(words)):
            for j in range(len(words)):
                if i == j:
                    continue
                if words[i] in words[j]:
                    ans.append(words[i])
                    break
        return ans


if __name__ == '__main__':
    S = Solution()
    print(S.stringMatching(["leetcode", "et", "code"]))
    print(S.stringMatching(["leetcoder","leetcode","od","hamlet","am"]))
