package main

import (
	"fmt"
	"sync"
	"time"
)

/*
ex : 파일 접근 시 모든 고루틴이 서로 다른 파일에 접근하도록 함

	=> 아에 사용하는 데이터가 분리되는 경우 ( 영역을 분리한다)
*/
type job interface {
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index)
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

func main() {
	var jobList [10]job

	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}

	wg.Wait()
}
