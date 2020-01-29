from collections import deque


class RecentCounter:

    def __init__(self):
        self.dq = deque()

    def ping(self, t: int) -> int:
        while self.dq:
            if self.dq[0] < t - 3000:
                self.dq.popleft()
            else:
                break
        self.dq.append(t)
        return len(self.dq)

# Your RecentCounter object will be instantiated and called as such:
# obj = RecentCounter()
# param_1 = obj.ping(t)