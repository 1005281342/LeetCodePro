from typing import List
from collections import deque


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:

    def pathSum(self, root: TreeNode, su: int) -> List[List[int]]:

        if not root:
            return list()

        ans = list()
        dq = deque()
        dq.append((root, root.val, [root.val]))
        while dq:
            node, d, path = dq.popleft()
            if not node.left and not node.right:
                if d == su:
                    ans.append(path)
            if node.left:
                k = path[:]
                k.append(node.left.val)
                dq.append((node.left, d + node.left.val, k))
            if node.right:
                k = path[:]
                k.append(node.right.val)
                dq.append((node.right, d + node.right.val, k))
        return ans
