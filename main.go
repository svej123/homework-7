package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}
type LinkedList struct {
	Head *Node
}

func (list *LinkedList) Append(value int) {
	newNode := &Node{Value: value}
	if list.Head == nil {
		list.Head = newNode
		return
	}
	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}
func (list *LinkedList) Print() {
	current := list.Head
	for current != nil {
		fmt.Println("%d -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}
func (list *LinkedList) Reverse() {
	var prev *Node
	current := list.Head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	list.Head = prev
}

func main() {
	list := &LinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)
	fmt.Println("before reverse")
	list.Print()

	list.Reverse()

	fmt.Println("after reverse")
	list.Print()
}
