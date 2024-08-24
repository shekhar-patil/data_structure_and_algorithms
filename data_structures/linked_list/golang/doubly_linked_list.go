package main

import (
  "fmt"
  "strconv"
)

// LinkedList represents a doubly linked list with a head pointer.
type LinkedList struct {
  head *Node
}

// Node represents a node in the doubly linked list with pointers to the next and previous nodes, and an integer data field.
type Node struct {
  next *Node
  prev *Node
  data int
}

// insertNode adds a new node with the given data at the end of the linked list.
func (l *LinkedList) insertNode(data int) {
  // Create a new node with the provided data
  node := &Node{data: data}

  // If the list is empty, set the new node as the head
  if l.head == nil {
    l.head = node
    return
  }

  // Traverse the list to find the last node
  ptr := l.head
  for ptr.next != nil {
    ptr = ptr.next
  }

  // Set the new node as the next node of the last node and update the previous pointer of the new node
  node.prev = ptr
  ptr.next = node
}

// deleteNode removes the node with the given data from the linked list, if it exists.
func (l *LinkedList) deleteNode(data int) {
  // Check if the list is empty
  if l.head == nil {
    fmt.Println("List is empty")
    return
  }

  // Pointer to traverse the list
  ptr := l.head

  // Check if the node to delete is the head node
  if ptr.data == data {
    l.head = ptr.next
    if l.head != nil {
      l.head.prev = nil // Update the previous pointer of the new head
    }
    return
  }

  // Traverse the list to find the node with the matching data
  for ptr != nil && ptr.data != data {
    ptr = ptr.next
  }

  // If the node was not found
  if ptr == nil {
    fmt.Println("Number not found")
    return
  }

  // If the node to delete is the last node
  if ptr.next == nil {
    ptr.prev.next = nil
  } else {
    // If the node to delete is in the middle, update the pointers of the neighboring nodes
    ptr.next.prev = ptr.prev
    ptr.prev.next = ptr.next
  }
}

// printNodes prints all the nodes in the linked list in a formatted manner.
func (l *LinkedList) printNodes() {
  // Pointer to traverse the list
  ptr := l.head
  var list string

  // Traverse and concatenate the data of each node to the list string
  for ptr != nil {
    list += strconv.Itoa(ptr.data) // Convert integer data to string and add to the list

    // If not the last node, add an arrow to the string
    if ptr.next != nil {
      list += "--->"
    }
    ptr = ptr.next
  }

  // Print the final concatenated string representing the linked list
  fmt.Println(list)
}

// main function to demonstrate the linked list operations
func main() {
  // Create an empty linked list
  l := LinkedList{}

  // Insert nodes into the linked list
  l.insertNode(1)
  l.insertNode(2)
  l.insertNode(3)
  l.insertNode(4)

  // Print the linked list
  l.printNodes() // Output: 1--->2--->3--->4

  // Delete a node from the linked list
  l.deleteNode(1)
  l.printNodes() // Output: 2--->3--->4

  // Try to delete a non-existent node
  l.deleteNode(10) // Output: Number not found
}
