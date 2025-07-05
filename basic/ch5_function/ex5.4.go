package main

import "fmt"

/**
함수

1. 함수에 전달된 인수는 "복사" 된다
2. 멀티 반환 할 수 있다.
3. 변수명을 지정하여 반환할 수 있다.
*/

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)

	d, success := Divide(9, 0)
	fmt.Println(d, success)
}

func Divide(a int, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}
