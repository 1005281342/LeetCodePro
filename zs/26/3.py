from typing import List

from collections import deque


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:

    def __init__(self):
        self.cnt = 1

    def goodNodes(self, root: TreeNode) -> int:
        dq = deque()
        if root.left:
            dq.append((root.left, root.val))
        if root.right:
            dq.append((root.right, root.val))

        while dq:
            node, v = dq.popleft()
            if node.val >= v:
                self.cnt += 1
            if node.left:
                dq.append((node.left, max(v, node.val)))
            if node.right:
                dq.append((node.right, max(v, node.val)))
        return self.cnt
