package stack

type Stack[T any] struct {
	items []T
}

// конструктор стека
func New[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// добавляет элемент на вершину стека
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// удаляет и возвращает элемент с вершины стека
// Возвращает элемент и флаг успешности (false, если стек пуст)
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}

	item := s.items[len(s.items)-1]

	s.items = s.items[:len(s.items)-1]

	return item, true
}

// показывает элемент на вершине без удаления
func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	return s.items[len(s.items)-1], true
}

// проверка стека на пустоту
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// возвращает количество элементов
func (s *Stack[T]) Size() int {
	return len(s.items)
}
