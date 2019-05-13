package main

import (
	"fmt"
	"queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())

	// another type
	r := queue.Queue{"a"}
	r.Push("b")
	r.Push("c")
	fmt.Println(r.Pop())
	fmt.Println(r.Pop())
	fmt.Println(r.IsEmpty())
	fmt.Println(r.Pop())
	fmt.Println(r.IsEmpty())
	fmt.Println(r.Pop())
}
