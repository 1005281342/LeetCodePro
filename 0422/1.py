def count_primes(n: int) -> int:
    if n < 3:
        return 0
    results = [1] * n
    results[0], results[1] = 0, 1
    for i in range(2, int(n ** 0.5) + 1):
        if results[i] == 1:
            results[i * 2:n:i] = [0] * len(results[i * 2:n:i])
    return sum(results) - 1


def func1():
    print(x / 2)


if __name__ == '__main__':
    n = int(input())
    while n:
        x = int(input())
        n -= 1
        if x <= 2:
            print(1)
            continue
        func1()
