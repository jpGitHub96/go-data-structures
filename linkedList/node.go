package linkedList

type node[T comparable] struct {
	Value T
	Next  *node[T]
}

func (n node[T]) hasNext() bool {
	if n.Next == nil {
		return false
	}
	return true
}

func (n node[T]) getNext() *node[T] {
	if n.hasNext() {
		return n.Next
	}
	return nil
}
