package main

import (
	"errors"
	"fmt"
)

type List struct {
	data interface{} //数据
	//head *List	//链表头指针,头节点没有数据
	next *List //指向下一个元素
}

type Lister interface {
	Init()                               //初始化线性表
	Len() int                            //返回元素个数
	Clear()                              //清空线性表
	IsEmpty() bool                       //判断线性表是否为空
	GetElem(int) (interface{}, error)    //获取第i个节点的元素值
	LocateElem(interface{}) (int, error) //查找元素，返回元素的节点位置，否则返回错误
	Insert(int, interface{}) error       //插入元素
	Delete(int) (interface{}, error)     //删除元素
}

func (l *List) Init() {
	l.next = nil
	l.data = 0
}
func (l *List) IsEmpty() bool {
	if l.next == nil {
		return true
	}
	return false
}

//具体清空方法是什么？
func (l *List) Clear() {
	l.Init()
}

//获取某位置的元素值，首先判断列表是不是空的，再判断查找索引是否在区间内
//node指向0位置，向下遍历i次指到对应位置
//返回对应位置数据
func (l *List) GetElem(i int) (interface{}, error) {
	if l.data == 0 {
		return 0, errors.New("empty list")
	}
	if i < 0 || i >= l.data.(int) {
		return 0, errors.New("index out of range")
	}

	node := l.next
	for i > 0 {
		node = node.next
		i--
	}
	return node.data, nil
}

//先判断列表是否为空，然后让node指向第0个元素位置，判断node的值，直到产生等号，或者node为nil
func (l *List) LocateElem(e interface{}) (int, error) {
	if l.data == 0 {
		return 0, errors.New("empty list")
	}
	var i int
	node := l.next
	for node != nil {
		if node.data == e {
			return i, nil
		}
		i++
		node = node.next
	}
	return 0, errors.New("the element is not in this list")
}

//先判断插入的位置是否合理，让node指向第0个元素位置，如果node为空直接插入返回
//让node指向插入元素位置的前驱结点，如果没有则指向l
//构建新元素结构体，让其next指向node.next
//然后让node.next指向新元素结构体
//元素个数+1
func (l *List) Insert(i int, e interface{}) error {
	if i < 0 || i > l.data.(int) {
		return errors.New("index out of range")
	}
	node := l
	for i > 0 {
		node = node.next
		i--
	}
	temp := &List{
		data: e,
		next: node.next,
	}
	node.next = temp
	l.data = l.data.(int) + 1
	return nil
	//old version
	//node:=l.next
	//if node == nil{
	//	l.next = &List{
	//		data:e,
	//		next:nil,
	//	}
	//	l.data = l.data.(int) + 1
	//	return nil
	//}
	//for i>0{
	//	node = node.next
	//	i--
	//}
	//temp:=&List{
	//	data:e,
	//	next:node.next,
	//}
	//node.next = temp
	//l.data = l.data.(int) + 1
	//return nil
}

//先判断list是否为空，空则返回错误，再判断删除节点是否在list范围内，不在则返回错误
//令node指向待删除元素的前驱，获取待删除元素的值，判断待删除元素是否有后继
//有后继则让前驱指向后继，否则指向nil
//元素个数-1
func (l *List) Delete(i int) (interface{}, error) {
	if l.data == 0 {
		return 0, errors.New("empty list")
	}
	if i < 0 || i >= l.data.(int) {
		return 0, errors.New("index out of range")
	}
	node := l //待删除元素的前驱
	for i > 0 {
		node = node.next
		i--
	}
	e := node.next.data
	if node.next.next != nil {
		node.next = node.next.next
	} else {
		node.next = nil
	}

	l.data = l.data.(int) - 1
	return e, nil
}

//直接返回元素个数
func (l *List) Len() int {
	return l.data.(int)
}

func (l List) String() string {
	node := l.next
	length := l.data
	s := fmt.Sprintf("length:%d [", length)
	for node != nil {
		s += fmt.Sprintf("%d->", node.data)
		node = node.next
	}
	s += "]\n"
	return s
}
func main() {
	var a Lister = &List{}
	//var a = &List{}
	a.Init()
	fmt.Println(a)
	fmt.Println(a.IsEmpty())
	if err := a.Insert(0, 2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
	fmt.Println(a.GetElem(0))
	fmt.Println(a.LocateElem(0))
	fmt.Println(a)
	e, _ := a.Delete(0)
	fmt.Println(e, a)
	fmt.Println(a.Len())
}
