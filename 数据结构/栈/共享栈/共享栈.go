package main

import (
	"errors"
	"fmt"
)

type stack struct {
	data []interface{}
	top1 int
	top2 int
}
type stacker interface {
	Init(int)
	Clear()
	IsEmpty() bool
	GetTop1() (interface{}, error)
	Push1(interface{}) error
	Pop1() (interface{}, error)
	GetTop2() (interface{}, error)
	Push2(interface{}) error
	Pop2() (interface{}, error)
	ElemNum() int
}

func (s *stack) Init(n int) {
	s.data = make([]interface{}, n, n)
	s.top1 = -1
	s.top2 = n
}

func (s *stack) Clear() {
	s.top1 = -1
	s.top2 = len(s.data)
}

func (s *stack) IsEmpty() bool {
	if s.top1 < 0 && s.top2 >= len(s.data) {
		return true
	}
	return false
}
func (s *stack) GetTop1() (interface{}, error) {
	if s.top1 < 0 {
		return nil, errors.New("no element in stack")
	}
	return s.data[s.top1], nil
}
func (s *stack) GetTop2() (interface{}, error) {
	if s.top2 >= len(s.data) {
		return nil, errors.New("no element in stack")
	}
	return s.data[s.top2], nil
}
func (s *stack) Push1(a interface{}) error {
	if s.top1 >= s.top2 {
		return errors.New("stack overflow")
	}
	s.top1++
	s.data[s.top1] = a
	return nil
}
func (s *stack) Push2(a interface{}) error {
	if s.top2 <= s.top1 {
		return errors.New("stack overflow")
	}
	s.top2--
	s.data[s.top2] = a
	return nil
}

func (s *stack) Pop1() (interface{}, error) {
	if s.top1 < 0 {
		return nil, errors.New("pop an empty stack")
	}
	e := s.data[s.top1]
	s.top1--
	return e, nil
}
func (s *stack) Pop2() (interface{}, error) {
	if s.top2 >= len(s.data) {
		return nil, errors.New("pop an empty stack")
	}
	e := s.data[s.top2]
	s.top2++
	return e, nil
}
func (s *stack) ElemNum() int {
	return s.top1 + 1 + len(s.data) - s.top2
}
func main() {
	var s stacker = &stack{}
	s.Init(10)
	fmt.Println(s.IsEmpty())
	fmt.Println(s.GetTop1())
	fmt.Println(s.Push1('1'))
	fmt.Println(s.ElemNum())
}
