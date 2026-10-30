package stack

import "testing"

func TestStackPushPop(t *testing.T) {
	s := New[int]()

	s.Push(10)
	s.Push(20)

	if s.Size() != 2 {
		t.Errorf("Ожидался размер 2, получено %d", s.Size())
	}

	// Тест извлечения
	val, ok := s.Pop()
	if !ok || val != 20 {
		t.Errorf("Ожидалось 20, получено %v", val)
	}

	val, ok = s.Pop()
	if !ok || val != 10 {
		t.Errorf("Ожидалось 10, получено %v", val)
	}

	if !s.IsEmpty() {
		t.Errorf("Стек должен быть пустым")
	}
}

func TestStackPopEmpty(t *testing.T) {
	s := New[string]()
	_, ok := s.Pop()

	if ok {
		t.Errorf("Pop из пустого стека должен возвращать false")
	}
}

func BenchmarkStackPush(b *testing.B) {
	s := New[int]()

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
}
