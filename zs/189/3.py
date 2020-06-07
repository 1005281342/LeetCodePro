from typing import List


class Solution:
    def peopleIndexes(self, favoriteCompanies: List[List[str]]) -> List[int]:
        index = 0
        s = set()

        ans = list()
        while index < len(favoriteCompanies):

            s.add(index)
            t = set(favoriteCompanies[index])

            flag = False

            for i in range(len(favoriteCompanies)):
                if i == index:
                    continue
                flag = True
                ii = set(favoriteCompanies[i])
                # if ii.issubset(t):
                #     s.add(i)
                #     print(index, i)
                if t.issubset(ii):
                    flag = False
                    break
            if flag:
                # print(index, s)
                # print('*' * 10)
                ans.append(index)
            index += 1
        # print(ans)
        return ans


if __name__ == '__main__':
    S = Solution()
    S.peopleIndexes(
        [["leetcode", "google", "facebook"], ["google", "microsoft"], ["google", "facebook"], ["google"], ["amazon"]])
