package queue

import "testing"

func TestQueueEnqueueDequeue(t *testing.T) {
	q := New[int]()
	q.Enqueue(100)
	q.Enqueue(200)
	q.Enqueue(300)

	if q.Len() != 3 {
		t.Errorf("Ожидалась длина 3, получено %d", q.Len())
	}

	// Проверяем принцип FIFO (первым вошел - первым вышел)
	val, ok := q.Dequeue()
	if !ok || val != 100 {
		t.Errorf("Ожидалось 100, получено %v", val)
	}

	val, ok = q.Dequeue()
	if !ok || val != 200 {
		t.Errorf("Ожидалось 200, получено %v", val)
	}

	if q.Len() != 1 {
		t.Errorf("Ожидалась длина 1, получено %d", q.Len())
	}
}

func TestQueueDequeueEmpty(t *testing.T) {
	q := New[string]()
	_, ok := q.Dequeue()
	if ok {
		t.Error("Dequeue из пустой очереди должен возвращать false")
	}
}

func TestQueuePeek(t *testing.T) {
	q := New[int]()
	q.Enqueue(50)

	val, ok := q.Peek()
	if !ok || val != 50 {
		t.Errorf("Peek должен возвращать 50, получено %v", val)
	}

	if q.Len() != 1 {
		t.Error("Peek не должен изменять размер очереди")
	}
}

func TestQueueIsEmpty(t *testing.T) {
	q := New[int]()
	if !q.IsEmpty() {
		t.Error("Новая очередь должна быть пустой")
	}

	q.Enqueue(1)
	if q.IsEmpty() {
		t.Error("Очередь с элементом не должна быть пустой")
	}
}
