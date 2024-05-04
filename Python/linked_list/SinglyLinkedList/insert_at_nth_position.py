class Node:
  def __init__(self, data, next=None):
    self.data = data
    self.next = next

class LinkedList:
  def __init__(self, head=None):
    self.head = head

  def insertAt(self, data, position):
    if position == 1:
      node = Node(data, self.head)
      self.head = node
    else:
      i = 1
      ptr = self.head
      while ((i < position - 1) and (ptr.next is not None)):
        ptr = ptr.next
        i += 1

      node = Node(data, ptr.next)
      ptr.next = node

    self.printList()

  def printList(self):
    temp = self.head
    string = ""
    while temp is not None:
      string = string + str(temp.data) + "--->"
      temp = temp.next

    print(string)

if __name__ == "__main__":
  list = LinkedList()
  list.insertAt(12, 1)
  list.insertAt(23, 2)
  list.insertAt(15, 2)
  list.insertAt(15, 12)
  list.insertAt(11, 5)
