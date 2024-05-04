class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def deleteNode(self, node):
        if node.val == self.head.val:
          self.head = self.head.next
        else:
          itr = self.head
          while itr.next.val != node.val:
            itr = itr.next
          itr.next = itr.next.next
