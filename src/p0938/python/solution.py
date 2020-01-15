# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def __init__(self):
        self.ans = 0
        self.le = 0
        self.ri = 0

    def helper(self, r: TreeNode):
        if r:
            self.helper(r.left)
            if self.le <= r.val <= self.ri:
                self.ans += r.val
            self.helper(r.right)

    def rangeSumBST(self, root: TreeNode, L: int, R: int) -> int:
        self.le = L
        self.ri = R
        self.helper(root)
        return self.ans