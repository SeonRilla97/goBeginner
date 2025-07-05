package main

import "fmt"

/*
	변수 선언 : 메모리 주소와 타입이 정의된다
	Go언어는 강타입 언어다 -> 암시적 변환이 없기 때문에 연산시 형변환이 필수다
	변수에 값을 할당 안할 시 기본값으로 들어간다 ( 정수 : 0 / 실수 0.0 / String : "" / 그 외 nil
	실수 -> 정수 형변환 시 "버림" 한다
*/

func main() {
	a := 3
	var b float64 = 3.5

	var c int = int(b)
	d := float64(a * c)

	var e int64 = 7
	f := int64(d) * e

	var g int = int(b * 3)
	var h int = int(b) * 3

	fmt.Println(g, f, h)
}
