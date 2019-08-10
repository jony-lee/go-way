package main

import (
	"fmt"
	"unsafe"
)

//切片slice是go中三种引用类型之一（slice，map，结构体）
//切片内部通过指针引用底层数组
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

//引用类型需使用make函数或显示初始化语句
func sliceInit() {
	s1 := make([]int, 0, 0)    //指定len、cap，底层数组初始化为0
	s2 := make([]int, 3)       //省略cap，和len相等
	s3 := []int{10, 20, 5: 30} //按初始化元素分配底层数组，并设置len、cap
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))
	//Q. 其中s3的初始化和数组的初始化是一样的，那s3到底是数组呢，还是切片呢？
	//A.s3实际上是切片，具体参考下面的sliceAndArray函数，初始化数组的方式是[...]int{}
}

//定义和初始化的区别
func sliceDefine() {
	var a []int  //仅定义，但未进行初始化
	b := []int{} //定义且进行了初始化
	fmt.Println(a)
	fmt.Println(a == nil, b == nil)
}

func slicePoint() {
	s := []int{0, 1, 2, 3, 4}
	p := &s
	p0 := &s[0]
	p1 := &s[1]
	fmt.Println(p, p0, p1)
	(*p)[1] += 100
	*p1 += 100
	fmt.Println(s)
}

//slice可以获取元素地址，但不能像数组那样直接用指针访问元素内容
func sliceAndArray() {
	s := []int{0, 1, 2, 3}    //切片初始化
	a := [...]int{0, 1, 2, 3} //数组初始化
	pSlice := &s
	pArray := &a
	//fmt.Println(pSlice[0])		//无效操作，p[0] (type *[]int does not support indexing)
	fmt.Println((*pSlice)[0])
	fmt.Println(pArray[0]) //能够直接通过指针访问元素内容
}

//交错数组，即每一个二维数组的子数组长度不同，实际上这个二位数组不是连续存储的，各个子数组散列在内存的其他位置
//某一子数组发生扩容，其他数组不会发生变化。
//数组是固定大小的，不能扩容，只有切片能扩容，举个特别好玩的例子，就是寄居蟹换壳，寄居蟹相当于slice，寄居蟹的壳相当于数组
func arrayJagged() {
	x := [][]int{
		{1, 2},
		{10, 20, 30},
		{100},
	}
	fmt.Println("子数组地址及容量（扩容前）", x[2])
	for _, i := range x {
		fmt.Println(&i[0], cap(i))
	}
	x[2] = append(x[2], 200, 300, 400, 500)
	fmt.Println("子数组地址及容量（扩容后）", x[2])
	for _, i := range x {
		fmt.Println(&i[0], cap(i))
	}
}
func main() {
	sliceInit()
	//sliceDefine()
	arrayJagged()
}
