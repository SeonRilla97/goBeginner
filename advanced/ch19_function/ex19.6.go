package main

import "fmt"

// Capture에 대해 공부
// -> 추가 : 클로저라는 개념이 쓰임
// -> 익명함수의 v 변수는 함수가 제거되기 전까지 GC에 의해 해제되지 않는다.
func captureLoop() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}
