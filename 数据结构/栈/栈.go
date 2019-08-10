package main

import (
	"errors"
	"fmt"
)

type stack struct {
	data []interface{}
	top  int
}
type stacker interface {
	Init(int)
	Clear()
	IsEmpty() bool
	GetTop() (interface{}, error)
	Push(interface{}) error
	Pop() (interface{}, error)
	ElemNum() int
}

func (s *stack) Init(n int) {
	s.data = make([]interface{}, n, n)
	s.top = -1
}

func (s *stack) Clear() {
	s.top = -1
}

func (s *stack) IsEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}
func (s *stack) GetTop() (interface{}, error) {
	if s.top < 0 {
		return nil, errors.New("no element in stack")
	}
	return s.data[s.top], nil
}

func (s *stack) Push(a interface{}) error {
	if s.top >= len(s.data)-1 {
		//return errors.New("stack overflow")
		//利用切片自动扩容
		s.data = append(s.data, 0)
	}
	s.top++
	s.data[s.top] = a
	return nil
}

func (s *stack) Pop() (interface{}, error) {
	if s.top < 0 {
		return nil, errors.New("pop an empty stack")
	}
	e := s.data[s.top]
	s.top--
	return e, nil
}
func (s *stack) ElemNum() int {
	return s.top + 1
}
func main() {
	var s stacker = &stack{}
	s.Init(2)
	fmt.Println(s.IsEmpty())
	fmt.Println(s.GetTop())
	fmt.Println(s.Push('1'))
	fmt.Println(s.Push('1'))
	fmt.Println(s.Push('1'))
	fmt.Println(s.Push('1'))
	fmt.Println(s.Push('1'))
	fmt.Println(s.Push('1'))
	fmt.Println(s.Push('1'))
	fmt.Println(s.Push('1'))
	fmt.Println(s.Push('1'))
	fmt.Println(s.ElemNum())
}
