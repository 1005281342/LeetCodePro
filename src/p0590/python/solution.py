"""
# Definition for a Node.
class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children
"""
from typing import List


class Node:
    def __init__(self, val=None, children=None):
        self.val = val
        self.children = children


class Solution:

    def __init__(self):
        self.ans = list()

    def post_order(self, r: Node):
        if r:
            for node in r.children:
                self.post_order(node)
            self.ans.append(r.val)

    def postorder(self, root: 'Node') -> List[int]:
        self.post_order(root)
        return self.ans
