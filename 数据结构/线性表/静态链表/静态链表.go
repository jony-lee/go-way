package main

//TODO
const MaxListSize int = 1000

type List struct {
	data interface{}
	cur  int
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
	arr := [MaxListSize]List{}
	for i := 0; i < len(arr)-1; i++ {
		arr[i].cur = i + 1
	}
	arr[MaxListSize-1].cur = 0
}
