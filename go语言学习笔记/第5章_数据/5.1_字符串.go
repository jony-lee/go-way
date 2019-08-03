package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//字符串实际上是一个复合结构
//type stringStruct struct {
//	str unsafe.Pointer
//	len int
//}

//以切片语法返回子串时，其内部依旧指向原字节数组
func stringSlice() {
	s := "abcdefg"

	s1 := s[:3]
	s2 := s[1:4]
	s3 := s[2:]
	fmt.Println(&s, &s1, &s2, &s3)
	//s = "eeeeeee" //字符串是值类型，当重新赋值字符串时，原字符串地址和内容不变，开辟新的空间存储该字符串
	fmt.Println(s, s1, s2, s3)
	fmt.Println(&s, &s1, &s2, &s3)
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s)))
	fmt.Printf("%#v\n", (*reflect.StringHeader)(unsafe.Pointer(&s1)))
}

//字符串遍历
func stringIter() {
	s := "字符串"
	for i := 0; i < len(s); i++ { //按byte遍历
		fmt.Printf("%d:[%c]\n", i, s[i])
	}
	for i, c := range s { //rune:返回数组索引号，以及Unicode字符
		fmt.Printf("%d [%c]\n", i, c)
	}
}

func main() {
	stringSlice()
	stringIter()
}
