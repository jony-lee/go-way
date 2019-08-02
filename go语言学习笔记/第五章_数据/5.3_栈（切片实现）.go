package main

func sliceStack(){
	stack := make([]int, 0, 5)//初始化栈
	push := func(x int) error {//栈push操作
		n := len(stack)
		if n == cap(stack) {
			return errors.New("Stack is full")
		}
		stack = stack[:n+1]
		stack[n] = x
		return nil
	}

	pop := func() (int, error) {//栈pop操作
		n := len(stack)
		if n == 0 {
			return 0, errors.New("Stack is empty")
		}
		x := stack[n-1]
		stack = stack[:n-1]
		return x, nil
	}
	//入栈测试
	for i := 0; i < 7; i++ {
		fmt.Printf("pop:%d,%v,%v\n", i, push(i), stack)

	}

	//出栈测试
	for i := 0; i < 7; i++ {
		x, err := pop()
		fmt.Printf("pop:%d,%v,%v\n", x, err, stack)
	}
}

func main() {
	sliceStack()
}
