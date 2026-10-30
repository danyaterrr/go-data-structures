package linkedlist

import "testing"

func TestLinkedListAppendPrepend(t *testing.T) {
	list := New[int]()

	list.Append(10)
	list.Append(20)
	list.Prepend(5)

	if list.Size != 3 {
		t.Errorf("Ожидался размер 3, получено %d", list.Size)
	}

	expected := []int{5, 10, 20}
	result := list.Display()

	if len(result) != len(expected) {
		t.Fatalf("Длина среза не совпадает: ожидалось %d, получено %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Индекс %d: ожидалось %d, получено %d", i, expected[i], v)
		}
	}
}

func TestLinkedListDelete(t *testing.T) {
	list := New[string]()
	list.Append("A")
	list.Append("B")
	list.Append("C")

	ok := list.Delete("B")
	if !ok {
		t.Error("Ожидалось, что элемент 'B' будет удален")
	}

	if list.Size != 2 {
		t.Errorf("Ожидался размер 2, получено %d", list.Size)
	}

	if list.Contains("B") {
		t.Error("Элемент 'B' не должен существовать в списке")
	}

	ok = list.Delete("A")
	if !ok {
		t.Error("Ожидалось удаление головы 'A'")
	}

	if list.Head == nil || list.Head.Value != "C" {
		t.Error("Головой должен стать элемент 'C'")
	}
}

func TestLinkedListDeleteEmpty(t *testing.T) {
	list := New[int]()
	ok := list.Delete(10)
	if ok {
		t.Error("Удаление из пустого списка должно возвращать false")
	}
}

func TestLinkedListContains(t *testing.T) {
	list := New[string]()
	list.Append("Go")
	list.Append("Java")

	if !list.Contains("Go") {
		t.Error("Список должен содержать 'Go'")
	}

	if list.Contains("Python") {
		t.Error("Список не должен содержать 'Python'")
	}
}
