package main

import "fmt"

// 가변인자 -> 내부적으로 슬라이스로 처리
func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums: %v\n", nums)
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3)) // == []int{1,2,3}
	fmt.Println(sum(10, 20))
	fmt.Println(sum())
}
