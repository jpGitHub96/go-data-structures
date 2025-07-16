package linkedList

import "fmt"

type LinkedList[T comparable] struct {
	First  *node[T]
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
		l.First = &node[T]{Value: value, Next: nil}
		l.Length = 1
		return
	}

	current := l.First
	for current.hasNext() {
		current = current.getNext()
	}

	newNode := &node[T]{Value: value, Next: nil}
	current.Next = newNode
	l.Length++
}

func (l *LinkedList[T]) GetNode(value T) *node[T] {
	if l.Length == 0 {
		return nil
	}

	current := l.First
	for i := 0; i < l.Length; i++ {
		if current.Value == value {
			return current
		}
		current = current.getNext()
	}

	return nil
}

func (l *LinkedList[T]) GetNodeByIndex(index int) *node[T] {
	if index < 0 || index >= l.Length {
		return nil
	}

	current := l.First
	for i := 0; i < index; i++ {
		current = current.getNext()
	}
	return current
}

func (l *LinkedList[T]) RemoveNode(value T) bool {
	if l.IsEmpty() {
		return false
	}

	if l.First.Value == value {
		l.First = l.First.getNext()
		l.Length--
		return true
	}

	current := l.First
	for current.hasNext() {
		if current.getNext().Value == value {
			current.Next = current.getNext().getNext()
			l.Length--
			return true
		}
		current = current.getNext()
	}
	return false
}

func (l *LinkedList[T]) ToString() {

	current := l.First
	for current.hasNext() {
		fmt.Printf("%v ->", current.Value)
		current = current.getNext()
	}
	fmt.Println(" nil")
}
