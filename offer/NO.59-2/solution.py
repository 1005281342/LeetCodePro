import heapq


class MaxQueue:

    def __init__(self):
        self.lst = list()
        heapq.heapify(self.lst)

    def max_value(self) -> int:
        if not len(self.lst):
            return -1
        t = heapq.nlargest(1, self.lst)
        return t[0]

    def push_back(self, value: int) -> None:
        self.lst.append(value)

    def pop_front(self) -> int:
        if not len(self.lst):
            return -1
        x = self.lst[0]
        self.lst = self.lst[1:]
        return x

# Your MaxQueue object will be instantiated and called as such:
# obj = MaxQueue()
# param_1 = obj.max_value()
# obj.push_back(value)
# param_3 = obj.pop_front()
