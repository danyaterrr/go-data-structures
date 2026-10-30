## Go Data Structures

Collections of fundamental data structures on Go with using Generics.

## Project status

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-blue.svg)

## Examples for use

### Stack
```go
s := stack.New[int]()
s.Push(10)
s.Push(20)

val, ok := s.Pop() // val = 20, ok = true
```
### Queue
```go
q := queue.New[string]()
q.Enqueue("first")
q.Enqueue("second")

val, ok := q.Dequeue() // val = "первый", ok = true
```
### Linked list
```go
l := linkedlist.New[int]()
l.Append(10)
l.Prepend(5)
l.Delete(10)

exists := l.Contains(5) // true
```