package main

import (
	"errors"
	"fmt"
)

type queue struct {
	data interface{}
	next *queue
	rear *queue
}
type queuer interface {
	Init()
	Clear()
	IsEmpty() bool
	GetHead() (interface{}, error)
	EnQueue(interface{})
	DeQueue() (interface{}, error)
	ElemNum() int
}

func (q *queue) Init() {
	q.next = &queue{}
	q.rear = q.next
	q.data = 0
}
func (q *queue) Clear() {
	q.Init()
}
func (q *queue) IsEmpty() bool {
	return q.next == q.rear
}
func (q *queue) GetHead() (interface{}, error) {
	if q.next == q.rear {
		return nil, errors.New("empty queue")
	}
	return q.next.data, nil
}

func (q *queue) EnQueue(e interface{}) {
	var temp = &queue{
		data: e,
		next: nil,
		rear: q.rear,
	}
	q.rear.next = temp
	q.rear = temp
	q.data = q.data.(int) + 1
}
func (q *queue) DeQueue() (interface{}, error) {
	if q.next == q.rear {
		return nil, errors.New("dequeue an empty queue")
	}
	var e = q.next.data
	q.next = q.next.next
	q.data = q.data.(int) - 1
	return e, nil
}
func (q *queue) ElemNum() int {
	return q.data.(int)
}

func (q queue) String() string {
	var s string
	s = fmt.Sprintf(`
		queue status
-------------------------------
Queue	|	%v
EleNum	|	%d
Length	|	%d
`, q.data, q.ElemNum(), q.data.(int))
	return s
}

func main() {
	var q queuer = &queue{}
	q.Init()
	fmt.Println(q.ElemNum())
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue("abcdasdsad")
	fmt.Println(q)
	e, err := q.DeQueue()
	fmt.Println(e, err)
	fmt.Println(q)
}
