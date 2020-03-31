from collections import Counter


def helper(k, v):
    ans = [k] * v
    left = right = v >> 1
    t = k
    while left > 0:
        left -= 1
        t -= 1
        ans[left] = t
    t = k
    while right < v - 1:
        right += 1
        t += 1
        ans[right] = t
    return ans


def demo(a):
    cd = Counter(a)
    dq = cd.most_common(len(cd))
    k, v = dq[0]
    print(helper(k, v))
    for k, v in dq[1:]:
        print(k, v)


if __name__ == '__main__':
    s = [1, 2, 3, 4, 4, 5]
    demo(s)
