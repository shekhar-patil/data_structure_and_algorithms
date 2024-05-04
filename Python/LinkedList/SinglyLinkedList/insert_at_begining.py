class Node:
  def __init__(self, data=None, next=None):
    self.data = data
    self.next = next

class LinkedList:
  def __init__(self, head=None):
    self.head = head

  def insert_at_begining(self, data):
    node = Node(data, self.head)
    self.head = node

  def print(self):
    if self.head is None:
      print("Linked list is empty.")
      return

    itr = self.head
    string = ""
    while itr is not None:
      string += str(itr.data) + "-->"
      itr = itr.next
    print(string)

if __name__ == '__main__':
  ll = LinkedList()
  ll.insert_at_begining(5)
  ll.insert_at_begining(15)
  ll.insert_at_begining(52)
  ll.print()
