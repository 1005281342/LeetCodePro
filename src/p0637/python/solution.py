from typing import List
from collections import defaultdict


# Definition for a binary tree node.
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:

    def __init__(self):
        self.dd = defaultdict(list)

    def helper(self, r: TreeNode, i=0):
        if r:
            self.helper(r.left, i + 1)
            self.dd[i].append(r.val)
            self.helper(r.right, i + 1)

    def averageOfLevels(self, root: TreeNode) -> List[float]:

        self.helper(root)
        ans = [0] * len(self.dd)
        for i, vs in self.dd.items():
            ans[i] = sum(vs)/len(vs)

        return ans

