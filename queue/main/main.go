package main

import (
	"fmt"
	"golang_learn/queue"
)

func main() {
	var q queue.Queue
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	q.Push("hello world")
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	q.Push("hello world")
}
