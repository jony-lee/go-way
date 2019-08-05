package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
【题目】


【思路】
//TODO



*/

//【代码】
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	s := []string{}
	for {
		s = append(s, strconv.Itoa(l.Val))
		if l.Next == nil {
			break
		}
		l = l.Next
	}
	return strings.Join(s, "->")
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := head
	root := &ListNode{
		Val:  0,
		Next: head,
	}
	delNode := root
	for node.Next != nil {
		node = node.Next
		if n > 1 {
			n--
			continue
		}
		delNode = delNode.Next
	}

	delNode.Next = delNode.Next.Next
	return root.Next
}
func isMatch(s string, p string) bool {
	if len(s) == 0 && (len(p) == 0 || p == "*") { //字符串为空，匹配为空或*可成功
		return true
	}
	if len(s) != 0 && len(p) == 0 { //字符串不为空，匹配为空则失败
		return false
	}
	j := 0
	for i := 0; i < len(s); i++ {
		if s[i] == p[j] || p[j] == '?' {
			j++
		}
		if j == len(p)-1 && p[j] == '*' {
			return true
		}
		if p[j] == '*' && p[j+1] != s[i+1] {
			j++
			continue
		}
	}
}

//【主函数】
func main() {
	//a:= &ListNode{
	//	1,
	//	&ListNode{
	//		2,
	//		&ListNode{
	//			3,
	//			&ListNode{
	//				4,
	//				&ListNode{
	//					5,
	//					nil,
	//				},
	//			},
	//		},
	//	},
	//}
	//fmt.Println(removeNthFromEnd(a,2))

	s := "asd"
	p := "asd"
	fmt.Println(isMatch(s, p))
}

/*
【总结】：






*/
