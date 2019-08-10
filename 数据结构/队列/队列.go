package main

import (
	"errors"
	"fmt"
)

type queue struct {
	data  []interface{}
	front int
	rear  int
}
type queuer interface {
	Init(int)
	Clear()
	IsEmpty() bool
	GetHead() (interface{}, error)
	EnQueue(interface{}) error
	DeQueue() (interface{}, error)
	ElemNum() int
}

func (q *queue) Init(n int) {
	q.data = make([]interface{}, n, n)
}
func (q *queue) Clear() {
	q.front = 0
	q.rear = 0
}
func (q *queue) IsEmpty() bool {
	return q.front == q.rear
}
func (q *queue) GetHead() (interface{}, error) {
	if q.front == q.rear {
		return nil, errors.New("empty queue")
	}
	return q.data[q.front], nil
}

func (q *queue) EnQueue(e interface{}) error {
	if (q.rear+1)%len(q.data) == q.front {
		return errors.New("queue overflow")
	}
	q.data[q.rear] = e
	q.rear = (q.rear + 1) % len(q.data)
	return nil
}
func (q *queue) DeQueue() (interface{}, error) {
	if q.front == q.rear {
		return nil, errors.New("dequeue an empty queue")
	}
	var e = q.data[q.front]
	q.front = (q.front + 1) % len(q.data)
	return e, nil
}
func (q *queue) ElemNum() int {
	return (q.rear - q.front + len(q.data)) % len(q.data)
}

func (q queue) String() string {
	var s string
	s = fmt.Sprintf(`
		queue status
-------------------------------
Queue	|	%v
EleNum	|	%d
Length	|	%d
`, q.data, q.ElemNum(), len(q.data))
	return s
}

func main() {
	var q queuer = &queue{}
	q.Init(5)
	fmt.Println(q.ElemNum())
	_ = q.EnQueue(1)
	_ = q.EnQueue(2)
	_ = q.EnQueue("abcdasdsad")
	fmt.Println(q)

	e, err := q.DeQueue()
	fmt.Println(e, err)
	fmt.Println(q)
}
