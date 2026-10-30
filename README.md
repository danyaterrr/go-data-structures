# 🧩 Go Data Structures

Коллекция фундаментальных структур данных на Go с использованием Generics.

## 📊 Статус проекта

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-blue.svg)


## 🚀 Примеры использования

### Стек
```go
s := stack.New[int]()
s.Push(10)
s.Push(20)

val, ok := s.Pop() // val = 20, ok = true

### Очередь

q := queue.New[string]()
q.Enqueue("первый")
q.Enqueue("второй")

val, ok := q.Dequeue() // val = "первый", ok = true

### Связный список

l := linkedlist.New[int]()
l.Append(10)
l.Prepend(5)
l.Delete(10)

exists := l.Contains(5) // true