package linkedlist

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type LinkedList[T comparable] struct {
	Head *Node[T]
	Size int
}

func New[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value}
	newNode.Next = l.Head
	l.Head = newNode
	l.Size++
}

func (l *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	if l.Head == nil {
		l.Head = newNode
		l.Size++
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	l.Size++
}

func (l *LinkedList[T]) Delete(value T) bool {
	if l.Head == nil {
		return false
	}
	if l.Head.Value == value {
		l.Head = l.Head.Next
		l.Size--
		return true
	}
	current := l.Head
	for current.Next != nil {
		if current.Next.Value == value {
			current.Next = current.Next.Next
			l.Size--
			return true
		}
		current = current.Next
	}
	return false
}

func (l *LinkedList[T]) Contains(value T) bool {
	current := l.Head
	for current != nil {
		if current.Value == value {
			return true
		}
		current = current.Next
	}
	return false
}

func (l *LinkedList[T]) Display() []T {
	var result []T
	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}
