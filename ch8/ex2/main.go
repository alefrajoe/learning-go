package main

import (
	"fmt"
)

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type List[T comparable] struct {
	Node *Node[T]
}

func (l *List[T]) String() string {
	// Create a temporary pointer to traverse the list
	current := l.Node
	// Build a string representation
	result := "List: ["

	for current != nil {
		result += fmt.Sprintf("%v", current.Value)
		if current.Next != nil {
			result += ", "
		}
		current = current.Next
	}
	result += "]"
	return result
}

func (l *List[T]) Add(value T) {
	newNode := &Node[T]{Value: value, Next: nil}

	if l.Node == nil {
		l.Node = newNode
		return
	}

	current := l.Node
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

func (l *List[T]) Remove() {
	if l.Node == nil {
		return
	}

	l.Node = l.Node.Next
}

func (l *List[T]) Index(value T) int {
	current := l.Node
	index := 0
	for current != nil {
		if current.Value == value {
			return index
		}
		index++
		current = current.Next
	}
	return -1
}

func (l *List[T]) Contains(value T) bool {
	return l.Index(value) != -1
}

func main() {
	fmt.Println("Hello, World!")

	list := List[int]{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Remove()
	fmt.Println(list.String())
	fmt.Println(list.Contains(3))
}
