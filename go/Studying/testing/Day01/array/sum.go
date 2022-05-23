package array

func Sum_slice(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func Sum(numbers int) int {
	sum := 0
	for i := 0; i <= numbers; i++ {
		sum += i
	}
	return sum
}

// 不安全做法,考虑到切片有容量的概念
func SumAll(numToSum ...[]int) []int {
	lenNum := len(numToSum)
	sums := make([]int, lenNum)
	for i, number := range numToSum {
		sums[i] = Sum_slice(number)
	}
	return sums
}

// 优化版V2
func SumAllTails(numToSum ...[]int) []int {
	var sums []int
	for _, number := range numToSum {
		sums = append(sums, Sum_slice(number))
	}
	return sums
}

// 优化版V3
func SumAllV3(numToSum ...[]int) []int {
	var sums []int
	for _, number := range numToSum {
		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum_slice(number))
		}
	}
	return sums
}
