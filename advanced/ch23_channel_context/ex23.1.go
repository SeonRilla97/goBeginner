package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch) // 채널을 닫아서 데이터를 다 보냈음 알린다.
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { //채널에서 데이터를 계속 기다린다.
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}
