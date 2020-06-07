class BrowserHistory:

    def __init__(self, homepage: str):
        self.stk = list()
        self.forw = list()
        # self.forw.append(homepage)
        self.now = homepage

    def visit(self, url: str) -> None:
        self.stk.append(self.now)
        self.now = url
        self.forw = list()

    def back(self, steps: int) -> str:
        ans = self.now
        while self.stk and steps > 0:
            steps -= 1
            ans = self.stk.pop()
            self.forw.append(ans)
        self.now = ans
        return ans

    def forward(self, steps: int) -> str:
        ans = self.now
        while self.forw and steps > 0:
            steps -= 1
            ans = self.forw.pop()
            self.stk.append(ans)
        self.now = ans
        return ans

# Your BrowserHistory object will be instantiated and called as such:
# obj = BrowserHistory(homepage)
# obj.visit(url)
# param_2 = obj.back(steps)
# param_3 = obj.forward(steps)Ò
