package main

import (
	"errors"
	"fmt"
)

//对于链实现的栈，只需要讲单链表的头节点看作是栈顶就ok
//s依旧是栈顶，s的data维护链栈的长度信息
type stack struct {
	data interface{}
	top  *stack
}
type stacker interface {
	Init()
	Clear()
	IsEmpty() bool
	GetTop() (interface{}, error)
	Push(interface{})
	Pop() (interface{}, error)
	ElemNum() int
}

func (s *stack) Init() {
	s.data = 0
	s.top = nil

}

func (s *stack) Clear() {
	s.Init()
}

func (s *stack) IsEmpty() bool {
	if s.top == nil {
		return true
	}
	return false
}
func (s *stack) GetTop() (interface{}, error) {
	if s.top == nil {
		return nil, errors.New("no element in stack")
	}
	return s.top.data, nil
}

func (s *stack) Push(a interface{}) {
	//无需扩容，只需在头部插入节点
	temp := &stack{
		data: a,
		top:  s.top,
	}
	s.data = s.data.(int) + 1
	s.top = temp

}

func (s *stack) Pop() (interface{}, error) {
	if s.top == nil {
		return nil, errors.New("pop an empty stack")
	}
	e := s.top.data
	s.top = s.top.top
	s.data = s.data.(int) - 1
	return e, nil
}
func (s *stack) ElemNum() int {
	return s.data.(int)
}

func (s *stack) String() string {
	node := s.top
	length := s.data
	st := fmt.Sprintf("length : %d:[", length)
	for node != nil {
		st = st + fmt.Sprintf("%d->", node.data)
		node = node.top
	}
	st += "]\n"
	return st
}

func main() {
	var s stacker = &stack{}
	s.Init()
	fmt.Println(s.IsEmpty())
	fmt.Println(s.GetTop())
	s.Push('1')
	s.Push(1)
	s.Push(1)
	s.Push(1)
	s.Push(1)
	s.Push(1)
	fmt.Println(s.ElemNum())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
