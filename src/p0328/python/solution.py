# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


"""
示例 1:
输入: 1->2->3->4->5->NULL
02
输出: 1->3->5->2->4->NULL
示例 2:
输入: 2->1->3->5->6->4->7->NULL 
输出: 2->3->6->7->1->5->4->NULL
说明:
应当保持奇数节点和偶数节点的相对顺序。
链表的第一个节点视为奇数节点，第二个节点视为偶数节点，以此类推。
"""


class Solution:
    def oddEvenList(self, head: ListNode) -> ListNode:
        e, o = head, head.next
        t1, t2 = e, o
        while t2 and t2.next is not None:
            t1.next = t2.next
            t2.next = t1.next.next
            t1 = t1.next
            t2 = t2.next
        e.next = o
        return e
