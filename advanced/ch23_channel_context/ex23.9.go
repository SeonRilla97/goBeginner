package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	ctx := context.WithValue(context.Background(), "number", 9) // ctx객체, 취소함수
	go square7(ctx)
	time.Sleep(5 * time.Second)

	wg.Wait()
}
func square7(ctx context.Context) {
	if v := ctx.Value("number"); v != nil {
		n := v.(int)
		fmt.Println(n * n)
	}
	wg.Done()
}
