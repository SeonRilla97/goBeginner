package main

import "fmt"

/*
*
Slicew 정의 및 Append 방법

Slice의 내부 구조는 다음과 같다.
 1. 배열 포인터
 2. 현재 사용 용량(길이)
 3. 배열의 길이

Slice를 함수의 인자등으로 복사하면 len과 capacity는 값복사지만, 배열 포인터는 메모리 주소를 복사한다.

	=> 이로인해 예상치못한 버그가 발생할 수 있다.
*/
func main() {
	var slice []int

	for i := 1; i < 10; i++ {
		slice = append(slice, i)
	}

	slice = append(slice, 11, 12, 13, 14, 15)

	fmt.Println(slice)
}
