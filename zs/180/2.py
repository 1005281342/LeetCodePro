class CustomStack:

    def __init__(self, maxSize: int):
        self.max_size = maxSize
        self.cnt = 0
        self.lst = list()

    def push(self, x: int) -> None:
        if self.cnt < self.max_size:
            self.lst.append(x)
            self.cnt += 1

    def pop(self) -> int:
        if self.cnt <= 0:
            return -1
        t = self.lst.pop()
        self.cnt -= 1
        return t

    def increment(self, k: int, val: int) -> None:
        k = min(self.cnt, k)
        for i in range(k):
            self.lst[i] += val

# Your CustomStack object will be instantiated and called as such:
# obj = CustomStack(maxSize)
# obj.push(x)
# param_2 = obj.pop()
# obj.increment(k,val)
