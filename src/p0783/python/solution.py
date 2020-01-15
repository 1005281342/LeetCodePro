# Definition for a binary tree node.
from math import inf


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:

    def __init__(self):
        self.ans = inf
        self.ts = [-inf]

    def in_order(self, r: TreeNode):
        if r:
            self.in_order(r.left)
            self.ts.append(r.val)
            if self.ts[-1] - self.ts[-2] < self.ans:
                self.ans = self.ts[-1] - self.ts[-2]
                if self.ans <= 1:
                    return
            self.in_order(r.right)

    def minDiffInBST(self, root: TreeNode) -> int:

        self.in_order(root)
        return self.ans
