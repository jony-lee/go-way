package main

import (
	"errors"
	"fmt"
)

type List struct {
	data   []int
	length int
	cap    int
}

type Lister interface {
	Init(int, int) error         //初始化线性表
	Len() int                    //返回元素个数
	Cap() int                    //返回线性表容量
	Clear()                      //清空线性表
	expansion()                  //对线性表进行扩容
	IsEmpty() bool               //判断线性表是否为空
	GetElem(int) (int, error)    //获取第i个地方的元素值
	LocateElem(int) (int, error) //查找元素，返回元素位置，否则返回错误
	Insert(int, int) error       //插入元素
	Delete(int) (int, error)     //删除元素
}

func (l *List) Init(length int, capacity int) error {
	if length > capacity {
		return errors.New("error: length > capacity")
	}
	l.data = make([]int, length, capacity)
	l.length = length
	l.cap = capacity
	return nil
}
func (l *List) IsEmpty() bool {
	if l.length == 0 {
		return true
	}
	return false
}

//具体清空方法是什么？
func (l *List) Clear() {
	l.data = []int{}
	l.length = 0
	l.cap = 0
	//if l.length == 0{
	//	return
	//}
	//for i:=0;i<l.length;i++{
	//	l.data[i] = 0
	//}
}

func (l *List) GetElem(i int) (int, error) {
	if l.length == 0 || i < 0 || i >= l.length {
		return 0, errors.New("index out of range")
	}
	return l.data[i], nil
}

func (l *List) LocateElem(e int) (int, error) {
	for i := 0; i < l.length; i++ {
		if l.data[i] == e {
			return i, nil
		}
	}
	err := errors.New("the element is not in this list")
	return 0, err
}
func (l *List) expansion() {
	//具体扩容策略参照切片的扩容策略，此处直接引用切片扩容
	l.data = append(l.data, 0)
	l.cap = cap(l.data)
}

func (l *List) Insert(i, e int) error {
	if i < 0 || i > l.length {
		return errors.New("index out of range")
	}
	l.expansion()
	l.length++
	for j := l.length - 1; j > i; j-- {
		l.data[j] = l.data[j-1]
	}
	l.data[i] = e
	return nil
}
func (l *List) Delete(i int) (int, error) {
	if i < 0 || i > l.length {
		return 0, errors.New("index out of range")
	}
	e := l.data[i]
	for ; i < l.length-1; i++ {
		l.data[i] = l.data[i+1]
	}
	l.length--
	return e, nil
}
func (l *List) Len() int {
	return l.length
}

func (l *List) Cap() int {
	return l.cap
}
func (l List) String() string {
	s := "["
	for i := 0; i < l.length; i++ {
		s += fmt.Sprintf("%d,", l.data[i])
	}
	s += "]\n"
	return s
}
func main() {
	var a List
	//var a = &List{}
	//_ = a.Init(1, 2)
	fmt.Println(a)
	fmt.Println(a.IsEmpty())
	if err := a.Insert(0, 2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
	fmt.Println(a.GetElem(0))
	fmt.Println(a.LocateElem(0))
	e, _ := a.Delete(0)
	fmt.Println(e, a)
	fmt.Println(a.Len())
}
