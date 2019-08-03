package main

import (
	"fmt"
	"reflect"
)

//字段标签并不是注释，而是用来对字段进行描述的元数据，尽管它不属于数据成员，但却是类型的组成部分。
//可以利用反射获取标签信息。
//TODO 有关反射的内容当前不处理，后续章节会有详细介绍，到时候再学习
type user1 struct {
	name string `昵称`
	sex  byte   `性别`
}

func structTag() {
	u := user1{"Tom", 1}
	v := reflect.ValueOf(u)
	t := v.Type()
	for i, n := 0, t.NumField(); i < n; i++ {
		fmt.Printf("%s: %v\n", t.Field(i).Tag, v.Field(i))
	}
}
func main() {
	structTag()
}
