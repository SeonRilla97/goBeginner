package main

import "fmt"

/*
*
슬라이스를 복사할 때 사용하는 방법
*/
func main() {
	slice := []int{1, 2, 3, 4, 5}

	// 슬라이스 복사
	slice4 := make([]int, len(slice))
	cnt3 := copy(slice4, slice)
	fmt.Println(cnt3)

	//슬라이스의 요소 삭제
	newSlice := append(slice[:2], slice[3:]...) // ...의미 :: slice[3:] 을 개별적으로 풀어서 전달해라
	fmt.Println(newSlice)

	//슬라이스 요소 추가
	addSlice := append(slice, 0)
	copy(slice[3+1:], slice[3:]) //slice[3] 에 값 추가
	slice[3] = 100
	fmt.Println(addSlice)

}
