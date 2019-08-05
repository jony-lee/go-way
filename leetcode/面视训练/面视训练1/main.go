package main

import "fmt"

//func toLowerCase(str string) string{
//	var sList = []int32{}
//	dt := 'A'-'a'
//	for _,ch := range str{
//		if ch>='A' && ch <='Z'{
//			ch -= dt
//		}
//		sList = append(sList, ch)
//	}
//	return string(sList)
//}

//func clumsy(N int) int {
//	if N<=4{
//		switch N {
//		case 1:return 1
//		case 2:return 2
//		case 3:return 6
//		case 4:return 7
//		}
//	}
//	rst := 0		//返回值
//	tmp := N		//阶段计算值
//	for i,opt:=N-1,0;i>0;i--{
//		opt = opt % 4
//		switch opt{
//		case 0:
//			tmp = tmp * i
//		case 1:
//			tmp = tmp / i
//		case 2:
//			tmp = tmp + i
//		case 3:rst += tmp;tmp = -i;
//		}
//
//		if i == 1{
//			switch opt {
//			case 0:rst += tmp
//			case 1:rst += tmp
//			case 2:rst += tmp
//			case 3:rst -= 1
//			}
//		}
//		opt++
//
//	}
//	return rst
//}

func addNegabinary(arr1 []int, arr2 []int) []int {
	var long, short, ret []int
	if len(arr1) >= len(arr2) {
		long = arr1
		short = arr2
	} else {
		long = arr2
		short = arr1
	}
	for i, isOdd := len(short)-1, -1; i >= 0; i-- {

	}

}

func main() {
	//a:= "HEllo"
	//fmt.Println(toLowerCase(a))

	//N:= 5
	//fmt.Println(clumsy(N))

	m := []int{1, 1, 1, 1, 1}
	n := []int{1, 0, 1}
	fmt.Println(addNegabinary(m, n))
}
