from random import choice


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:

    def __init__(self, head: ListNode):
        """
        @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node.
        """
        self.head = head
        self.ans = list()

    def helper(self, node: ListNode):
        while node:
            self.ans.append(node.val)
            node = node.next

    def getRandom(self) -> int:
        """
        Returns a random node's value.
        """
        self.helper(self.head)
        return choice(self.ans)
