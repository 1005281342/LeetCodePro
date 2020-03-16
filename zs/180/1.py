from typing import List


class Solution:
    def luckyNumbers(self, matrix: List[List[int]]) -> List[int]:

        ans = list()
        # 找到每一行最小的元素
        for vs in matrix:
            t = 1000000
            ti = -1
            for i, a in enumerate(vs):
                if a < t:
                    t = a
                    ti = i
            # print(t, ti)
            flag = True
            for k in range(len(matrix)):
                # if t == 12:
                #     print(k, matrix[k][ti], ti)
                if matrix[k][ti] > t:
                    # print(k, matrix[k][ti], ti)
                    flag = False
                    break
            # print(t, flag)
            if flag:
                # print(t)
                ans.append(t)
        # print(ans)
        return list(set(ans))


if __name__ == '__main__':
    a = [[1, 10, 4, 2], [9, 3, 8, 7], [15, 16, 17, 12]]

    S = Solution()
    kk = S.luckyNumbers(a)
    print(kk)
