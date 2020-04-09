# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def rob(self, root: TreeNode) -> int:
        return max(self.helper(root))

    def helper(self, node):
        res = [0, 0]
        if not node:
            return res
        left = self.helper(node.left)
        right = self.helper(node.right)
        res[1] += node.val + left[0] + right[0]
        res[0] += max(left[0], left[1]) + max(right[0], right[1])

        return res
