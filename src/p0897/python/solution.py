class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


def inorder(root):
    if not root:
        return []
    else:
        return inorder(root.left) + [root.val] + inorder(root.right)


class Solution:
    def increasingBST(self, root):
        target = inorder(root)
        n = len(target)
        root = TreeNode(target[0])
        head = root
        for i in range(1, n):
            root.right = TreeNode(target[i])
            root = root.right
        return head
