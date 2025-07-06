package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)
	wg.Add(1)
	go square2(&wg, ch, quit)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	quit <- true // quit 채널을 통해 모든 작업이 완료됐음을 알린다.
	wg.Wait()
}

func square2(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case v := <-ch:
			fmt.Printf("Square: %d\n", v*v)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
}
