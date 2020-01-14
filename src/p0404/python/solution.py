# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:

    def __init__(self):
        self.ans = 0

    def helper(self, r: TreeNode):
        if r:
            if r.left and r.left.left is None and r.left.right is None:
                self.ans += r.left.val
            else:
                self.helper(r.left)
            self.helper(r.right)

    def sumOfLeftLeaves(self, root: TreeNode) -> int:
        self.helper(root)
        return self.ans
