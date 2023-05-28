class BinarySearchTreeNode:
  def __init__(self, data = None, parent = None, left = None, right = None) -> None:
    self.data = data
    self.left = left
    self.right = right
    self.parent = parent

  def add_left_node(self, data):
    if self.left:
      left_node = self.left
      left_node.add_node(data)
      return
    node = BinarySearchTreeNode(data, self)
    self.left = node


  def add_right_node(self, data):
    if self.right:
      left_node = self.right
      left_node.add_node(data)
      return
    node = BinarySearchTreeNode(data, self)
    self.right = node


  def add_node(self, data):
    if data == self.data:
      return

    if data < self.data:
      self.add_left_node(data)
    else:
      self.add_right_node(data)

  def print_tree(self):
    padding = self.node_level() * " " * 3
    padding += ("|_ _" if self.parent else "")
    print(padding + str(self.data))
    if self.left or self.right:
      if self.left:
        self.left.print_tree()
      if self.right:
        self.right.print_tree()

  def node_level(self):
    parent = self.parent
    level = 0
    while parent:
      level += 1
      parent = parent.parent

    return level

if __name__ == "__main__":
  root = BinarySearchTreeNode(8)
  list = [3, 10, 1,6,14, 4,7,13]
  for data in list:
    root.add_node(data)

  root.print_tree()
