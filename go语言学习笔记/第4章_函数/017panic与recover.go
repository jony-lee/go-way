package 第四章_函数

import "fmt"

//recover能够捕获任意的panic
//考虑通过匿名函数保护recover
func testPanic(x, y int) {
	z := 0
	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		z = x / y
	}()
	fmt.Println("x / y =", z)
}

func main() {
	testPanic(5, 0)
}
