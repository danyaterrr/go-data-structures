package main

import (
	"fmt"

	"github.com/danyaterrr/linkedlist"
	"github.com/danyaterrr/queue"
	"github.com/danyaterrr/stack"
)

func main() {
	fmt.Println("Демонстрация работы структур данных на Go")

	fmt.Println("\nStack:")
	s := stack.New[string]()
	s.Push("Первый")
	s.Push("Второй")
	s.Push("Третий")

	fmt.Println("Добавили в стек: Первый, Второй, Третий")

	for !s.IsEmpty() {
		val, _ := s.Pop()
		fmt.Println("Достали из стека:", val)
	}

	fmt.Println("\nQueue:")
	q := queue.New[int]()
	q.Enqueue(100)
	q.Enqueue(200)
	q.Enqueue(300)

	fmt.Println("Добавили в очередь: 100, 200, 300")

	for {
		val, ok := q.Dequeue()
		if !ok {
			break
		}
		fmt.Println("Извлекли из очереди:", val)
	}

	fmt.Println("\nLinked List:")
	list := linkedlist.New[string]()

	list.Append("Яблоко")
	list.Append("Банан")
	list.Prepend("Апельсин")
	list.Append("Груша")

	fmt.Println("Список после добавлений:", list.Display())

	list.Delete("Банан")
	fmt.Println("Список после удаления 'Банан':", list.Display())

	fmt.Println("\n=== Готово! Можно заливать на GitHub ===")
}
