package main

import "fmt"

// Node represents an element in the circular doubly linked list
type Node struct {
    data int
    next *Node
    prev *Node
}

// CircularDoublyLinkedList represents the circular doubly linked list
type CircularDoublyLinkedList struct {
    head *Node
}

// Insert a new node at the end of the list
func (list *CircularDoublyLinkedList) insertNode(data int) {
    newNode := &Node{data: data}

    if list.head == nil {
        list.head = newNode
        newNode.next = list.head
        newNode.prev = list.head
    } else {
        tail := list.head.prev
        tail.next = newNode
        newNode.prev = tail
        newNode.next = list.head
        list.head.prev = newNode
    }
}

// Print all nodes in the list
func (list *CircularDoublyLinkedList) printList() {
    if list.head == nil {
        fmt.Println("List is empty")
        return
    }

    current := list.head
    for {
        fmt.Print(current.data, " ")
        current = current.next
        if current == list.head {
            break
        }
    }
    fmt.Println()
}

func main() {
    list := CircularDoublyLinkedList{}

    list.insertNode(10)
    list.insertNode(20)
    list.insertNode(30)
    list.printList() // Output: 10 20 30
}
