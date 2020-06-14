from typing import List


class TreeAncestor:

    def __init__(self, n: int, parent: List[int]):
        self.d = {i: parent[i] for i in range(n)}
        self.d[-1] = -1
        self.n = n

    def getKthAncestor(self, node: int, k: int) -> int:

        if (k - 1) * 2 > self.n:
            return -1

        while k > 0:
            if node < 0:
                return -1

            k -= 1
            node = self.d[node]
        return node
# Your TreeAncestor object will be instantiated and called as such:
# obj = TreeAncestor(n, parent)
# param_1 = obj.getKthAncestor(node,k)
