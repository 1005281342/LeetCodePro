from typing import List
from collections import defaultdict, deque


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def levelOrder(self, root: TreeNode) -> List[List[int]]:

        if root is None:
            return list()

        dd = defaultdict(list)

        dq = deque()
        dq.append((root, 0))
        while dq:
            node, d = dq.popleft()
            dd[d].append(node.val)

            if node.left is not None:
                dq.append((node.left, d + 1))
            if node.right is not None:
                dq.append((node.right, d + 1))
        return list(dd.values())