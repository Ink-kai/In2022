package main

import (
	_ "container/ring"
	"fmt"
	algorithm "pro22/Studying/Algorithm"
	"sort"
)

func main() {
	var arr = []int{1, 23, 4, 43, 5, 45, 645, 67, 57, 56, 756, 3, 54, 423, 36, 4523}
	sort.Ints(arr)
	fmt.Printf("-----------------------二分查找非递归-----------------------------\n")
	fmt.Printf("现在数组：%v\n数字位置：%d", arr, algorithm.BinarySearch(arr, 43))
	var sumNum int = 123456
	fmt.Printf("\n----------------------------求和--------------------------------\n")
	fmt.Printf("%d 求和:%d", sumNum, algorithm.Sum(sumNum))
	fmt.Printf("\n--------------------------求和-V2版-----------------------------\n")
	fmt.Printf("%d 求和:%d", sumNum, algorithm.SumV2(sumNum))
	fmt.Printf("\n----------------------------链表--------------------------------\n")
	algorithm.LinkNodeDesc()
	fmt.Printf("\n----------------------------循环链表--------------------------------\n")
	r := &algorithm.Ring{Value: -1}
	r.Link(&algorithm.Ring{Value: 1})
	r.Link(&algorithm.Ring{Value: 2})
	r.Link(&algorithm.Ring{Value: 3})
	r.Link(&algorithm.Ring{Value: 4})
	temp := r.Unlink(3)
	// node := r
	// for {
	// 	fmt.Printf("%v\n", node.Value)
	// 	node = node.NextRing()
	// 	if node == r {
	// 		break
	// 	}
	// }
	node := temp
	for {
		fmt.Printf("%v\n", node.Value)
		node = node.NextRing()
		if node == temp {
			break
		}
	}
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s := preByteConcat(10, letterBytes)
	fmt.Println(s)
}
func preByteConcat(n int, str string) string {
	buffer := make([]byte, 0, n*len(str))
	for i := 0; i < n; i++ {
		buffer = append(buffer, str...)
	}
	return string(buffer)
}
