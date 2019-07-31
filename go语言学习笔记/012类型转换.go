package main

func main() {
	x := 100.1
	p := (*float64)(&x)
	println(p)
}
