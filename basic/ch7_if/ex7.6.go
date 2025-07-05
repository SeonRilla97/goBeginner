package main

import "fmt"

/**
쇼트 서킷
	1. && 연산 : 좌변 false 시 우변 확인 없이 false 반환
	2. || 연산 : 좌변 true 시 우변 확인 없이 true

if 중첩은 많이 봐줘서 2번까지만 하는게 좋다

*/

func getMyAge() (int, bool) {
	return 33, true
}
func main() {
	// if 초기문; 조건문 {} <- if 조건 검사 전 초기문을 삽입할 수 있다.
	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You are young")
	} else {
		fmt.Println("Nice Age : ", age)
	}
}
