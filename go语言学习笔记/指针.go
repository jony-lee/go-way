package main

func main() {
	x := 10
	p := &x
	//p = 0xc00002df80
	println(p, *p)
}
