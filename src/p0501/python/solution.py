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
        self.dd = defaultdict(int)

    def bl(self, r: TreeNode):
        if r:
            self.dd[r.val] += 1
            self.bl(r.left)
            self.bl(r.right)

    def findMode(self, root: TreeNode) -> List[int]:

        self.bl(root)
        td = defaultdict(list)
        for k, v in self.dd.items():
            td[v].append(k)
        if td:
            return td[max(td.keys())]
        else: return list()
