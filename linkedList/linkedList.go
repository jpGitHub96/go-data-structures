package linkedList

import "fmt"

type LinkedList[T comparable] struct {
	First  *Node[T]
	Length int
}

func NewLinkedList[T comparable]() LinkedList[T] {
	return LinkedList[T]{
		First:  nil,
		Length: 0,
	}
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.Length == 0
}

func (l *LinkedList[T]) AddNode(value T) {
	if l.IsEmpty() {
		l.First = &Node[T]{Value: value, Next: nil}
		l.Length = 1
		return
	}

	current := l.First
	for current.HasNext() {
		current = current.GetNext()
	}

	newNode := &Node[T]{Value: value, Next: nil}
	current.Next = newNode
	l.Length++
}

// func add all

func (l *LinkedList[T]) GetNode(value T) *Node[T] {
	if l.Length == 0 {
		return nil
	}

	current := l.First
	for i := 0; i < l.Length; i++ {
		if current.Value == value {
			return current
		}
		current = current.GetNext()
	}

	return nil
}

func (l *LinkedList[T]) GetNodeByIndex(index int) *Node[T] {
	if index < 0 || index >= l.Length {
		return nil
	}

	current := l.First
	for i := 0; i < index; i++ {
		current = current.GetNext()
	}
	return current
}

func (l *LinkedList[T]) RemoveNode(value T) bool {
	if l.IsEmpty() {
		return false
	}

	if l.First.Value == value {
		l.First = l.First.GetNext()
		l.Length--
		return true
	}

	current := l.First
	for current.HasNext() {
		if current.GetNext().Value == value {
			current.Next = current.GetNext().GetNext()
			l.Length--
			return true
		}
		current = current.GetNext()
	}
	return false
}

func (l *LinkedList[T]) ToString() {

	current := l.First
	for current.HasNext() {
		fmt.Printf(" %v ->", current.Value)
		current = current.GetNext()
	}

	fmt.Printf(" %v -> nil\n", current.Value)
}
