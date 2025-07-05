package ch10_array

import "fmt"

/*
*
Array

동일타입 데이터로 이루어진 연속적인 데이터
 1. 내장함수 len() 으로 배열의 길이 파악 가능
 2. range 를 통해 순회
*/
func main() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	for _, arr := range a {
		for _, v := range arr {
			fmt.Println(v, " ")
		}
	}
}
