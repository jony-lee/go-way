package main

import (
	"errors"
	"fmt"
)

func e1() {
	fmt.Println("e1")
	err := errors.New("start err1")
	defer fmt.Println(err)
	err = errors.New("defer1 error")
}

func e2() {
	fmt.Println("e2")
	err := errors.New("start err2")
	fmt.Println(&err)
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("defer2 error")
	fmt.Println(&err)
}

func e3() {
	fmt.Println("e3")
	err := errors.New("start err3")
	fmt.Println(&err, "start")
	defer func(err error) {
		fmt.Println(&err, "func")
		fmt.Println(err)
	}(err)
	err = errors.New("defer3 error")
	fmt.Println(&err, "end")
}
func e4() {
	fmt.Println("0")
	defer fmt.Println("1")
	fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("4")
	return
}

func main() {
	defer func() {}()
	//e1()
	e2()
	e3()
	//e4()
	for i := 0; i < 5; i++ {
		j := i
		fmt.Println(&j)
		defer fmt.Println(j, &j, 4)
	}
}
