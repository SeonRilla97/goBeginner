package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func (q *Stack) Push(v interface{}) {
	q.v.PushBack(v)
}

func (q *Stack) Pop() interface{} {
	back := q.v.Back()
	if back != nil {
		return q.v.Remove(back)
	}
	return nil
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func main() {
	q := NewStack()
	for i := 0; i < 10; i++ {
		q.Push(i)
	}
	v := q.Pop()
	for v != nil {
		fmt.Printf("%v -> ", v)
		v = q.Pop()
	}
}
