from typing import List


# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def nextLargerNodes(self, head: ListNode) -> List[int]:
        vals, stack = [], [0]
        while head:
            vals.append(head.val)
            head = head.next
        for i in range(len(vals) - 1, -1, -1):
            while stack[-1] and vals[i] >= stack[-1]:
                stack.pop()
            stack.append(vals[i])
            vals[i] = stack[-2]
        return vals
