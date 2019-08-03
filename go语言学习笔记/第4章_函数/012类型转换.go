package 第四章_函数

func main() {
	x := 100.1
	p := (*float64)(&x)
	println(p)
}
