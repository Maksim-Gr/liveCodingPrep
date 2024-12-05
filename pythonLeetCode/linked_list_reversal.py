class ListNode:
    def __init__(self, val):
        self.val = val
        self.next = None
        

class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        if head is None:
            return head

        prev, curr, next = None, head, None
        while curr is not None: 
            # store the current node
            next = curr.next
            # reverse the current node
            curr.next = prev
            #point prev to the current 
            prev = curr
            # move 
            curr = next
        return prev