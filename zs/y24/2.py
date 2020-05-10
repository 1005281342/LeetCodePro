from bisect import bisect_left as search

fib_list = [1, 1]
for i in range(2, 50):
    fib_list.append(fib_list[-1] + fib_list[-2])


class Solution:

    def findMinFibonacciNumbers(self, k: int) -> int:
        n = k
        index_a = search(fib_list, k)
        if fib_list[index_a] == k:
            return 1
        cnt = 0
        while n > 0:
            index_a = search(fib_list, n)

            if fib_list[index_a] > n:
                index_a -= 1
                cnt += 1
                n -= fib_list[index_a]
            elif fib_list[index_a] == n:
                return cnt + 1


if __name__ == '__main__':
    S = Solution()
    print(S.findMinFibonacciNumbers(7))
    print(S.findMinFibonacciNumbers(10))
    print(S.findMinFibonacciNumbers(19))
