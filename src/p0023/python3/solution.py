from typing import List

import heapq


# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        from heapq import heappush, heappop
        node_pools = []
        lookup = collections.defaultdict(list)
        for head in lists:
            if head:
                heappush(node_pools, head.val)
                lookup[head.val].append(head)
        dummy = cur = ListNode(None)
        while node_pools:
            smallest_val = heappop(node_pools)
            smallest_node = lookup[smallest_val].pop(0)
            cur.next = smallest_node
            cur = cur.next
            if smallest_node.next:
                heappush(node_pools, smallest_node.next.val)
                lookup[smallest_node.next.val].append(smallest_node.next)
        return dummy.next
