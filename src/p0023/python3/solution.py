from typing import List

import heapq


# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def mergeKLists(self, lists: List[ListNode]):

        rq = list()
        for node in lists:
            while node:
                heapq.heappush(rq, node.val)
                node = node.next

        dummy = ListNode(None)
        cur = dummy
        # 依次将堆中的元素取出(因为是小顶堆，所以每次出来的都是目前堆中值最小的元素），然后重新构建一个列表返回
        while rq:
            temp_node = ListNode(heapq.heappop(rq))
            cur.next = temp_node
            cur = temp_node
        return dummy.next
