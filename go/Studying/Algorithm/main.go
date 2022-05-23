package Algorithm

import (
	"fmt"
	"math"
	"sort"
)

func BinarySearch(num []int, target int) int {
	if !sort.IntsAreSorted(num) {
		sort.Ints(num)
	}
	var (
		left  int = 0
		right int = len(num)
		mid   int = 0
	)
	for {
		mid = int(math.Floor(float64((left + right) / 2)))
		if target > num[mid] {
			left = mid + 1
		} else if target < num[mid] {
			right = mid - 1
		} else {
			break
		}
		if left > right {
			mid = -1
			break
		}
	}
	return mid
}

func Sum(n int) int {
	var total int = 0
	for i := 0; i <= n; i++ {
		total = total + i
	}
	return total
}

func SumV2(n int) int {
	total := ((1 + n) * n) / 2
	return total

}

// 单链表
type linkNode struct {
	Data     int64
	NextNode *linkNode
}

func LinkNodeDesc() {
	node := new(linkNode)
	node.Data = 2
	node2 := new(linkNode)
	node2.Data = 3
	node.NextNode = node2
	node3 := new(linkNode)
	node3.Data = 4
	node2.NextNode = node3
	newnode := node
	for {
		if newnode != nil {
			fmt.Printf("%v\t", newnode.Data)
			newnode = newnode.NextNode
			continue
		}
		break
	}
}
