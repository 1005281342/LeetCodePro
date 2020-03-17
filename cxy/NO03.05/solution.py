import heapq


class SortedStack:

    def __init__(self):
        self.lst = list()
        self.cnt = 0
        heapq.heapify(self.lst)

    def push(self, val: int) -> None:
        heapq.heappush(self.lst, val)
        self.cnt += 1

    def pop(self) -> None:
        if self.cnt > 0:
            heapq.heappop(self.lst)
            self.cnt -= 1

    def peek(self) -> int:
        if self.cnt <= 0:
            return -1
        return heapq.nsmallest(1, self.lst)[0]

    def isEmpty(self) -> bool:
        return self.cnt <= 0

# Your SortedStack object will be instantiated and called as such:
# obj = SortedStack()
# obj.push(val)
# obj.pop()
# param_3 = obj.peek()
# param_4 = obj.isEmpty()
