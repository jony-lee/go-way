package main

import "fmt"

//使用copy时，是将右边的元素对齐覆盖到左边的元素中
func sliceCopy() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[5:8]
	n := copy(s[2:], s1)
	fmt.Println(n, s)

	s2 := make([]int, 6)
	n = copy(s2, s)
	fmt.Println(n, s2)
}

func main() {
	sliceCopy()
}
