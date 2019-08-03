package main

import (
	"fmt"
	"unsafe"
)

//内存布局
//结构体的内存总是在创建的时候一次性分配的，各个字段按命名顺序紧凑排列。
//结构体中基本宽度由所有字段中最长的基础类型宽度确定
//利用unsafe包中相关函数可以输出字段情况

type point struct {
	x, y int
}
type value struct {
	id   int
	name string
	data []byte
	next *value
	point
}

func structMem() {
	v := value{
		id:    1,
		name:  "test",
		data:  []byte{1, 2, 3, 4},
		point: point{x: 100, y: 200},
	}
	s := `
		v: %p ~ %x, size: %d, align: %d

		field	address		offset	size
		------+----------+--------+------
		id		%p			%d		%d
		name	%p			%d		%d
		data	%p			%d		%d
		next	%p			%d		%d
		x		%p			%d		%d
		y		%p			%d		%d

	`
	fmt.Printf(s,
		&v, uintptr(unsafe.Pointer(&v))+unsafe.Sizeof(v), unsafe.Sizeof(v), unsafe.Alignof(v),
		&v.id, unsafe.Offsetof(v.id), unsafe.Sizeof(v.id),
		&v.name, unsafe.Offsetof(v.name), unsafe.Sizeof(v.name),
		&v.data, unsafe.Offsetof(v.data), unsafe.Sizeof(v.data),
		&v.next, unsafe.Offsetof(v.next), unsafe.Sizeof(v.next),
		&v.x, unsafe.Offsetof(v.x), unsafe.Sizeof(v.x),
		&v.y, unsafe.Offsetof(v.y), unsafe.Sizeof(v.y))
}

func main() {
	structMem()
}
