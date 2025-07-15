package linkedList

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

func (n Node[T]) HasNext() bool {
	if n.Next == nil {
		return false
	}
	return true
}

func (n Node[T]) GetNext() *Node[T] {
	if n.HasNext() {
		return n.Next
	}
	return nil
}
