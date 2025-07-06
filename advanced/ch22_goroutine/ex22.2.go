package main

import (
	"fmt"
	"sync"
)

/*
*
동시 작업을 위한 객체

	Add() : 작업 개수
	Done() : 작업 완료될 때 마다 호출
	Wait() : 모든 작업 완료까지 대기
*/
var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := 0; i < b; i++ {
		sum += i
	}
	fmt.Println("%d~ %d 까지 합계는 %d입니다.", a, b, sum)
	wg.Done() // 작업 완료됨을 알림
}

func main() {
	wg.Add(10) //동시 작업 수 10개
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000000)
	}
	wg.Wait() // 10개의 작업이 모두 완료 되기까지 대기
}
