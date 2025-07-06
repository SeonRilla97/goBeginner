package main

import (
	"fmt"
	"sync"
	"time"
)

// 메시지가 있으면 빼와서 실행하고 그렇지 않다면 1초 간격으로 다른 일을 수행한다.
func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square3(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}

func square3(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminate")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}
