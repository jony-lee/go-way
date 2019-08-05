package main

import "fmt"

type Node struct {
	data  int
	Left  *Node
	Right *Node
}

func (*Node) String() string {

}

var i = -1

func ListToTree(arr []int) *Node {
	i += 1
	if i >= len(arr) {
		return nil
	}
	var t Node
	if arr[i] != 0 {
		t = Node{arr[i], nil, nil}
		t.Left = ListToTree(arr)
		t.Right = ListToTree(arr)
	} else {
		return nil
	}
	return &t
}
func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(ListToTree(a))
}
