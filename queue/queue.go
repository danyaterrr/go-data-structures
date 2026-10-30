package queue

type Queue[T any] struct {
	items []T
}

// конструктор очереди
func New[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

// добавить элемент в конец очереди
func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

// удалить элемент из начала очереди
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item, true
}

// показать элемент в начале очереди без удаления
func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	return q.items[0], true
}

// проверка очереди на пустоту
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// возвращает количество элементов
func (q *Queue[T]) Len() int {
	return len(q.items)
}
