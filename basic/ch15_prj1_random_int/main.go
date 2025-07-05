package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	//result Random int (0~99 사이)
	result := rand.IntN(100)

	/*
		사용자 입력
		  	result < input :: down && 시행횟수++
			result > input :: up && 시행횟수++
			result == imput :: success!! : x`시행횟수
	*/
	tryNum := 0
	for {
		var input int
		_, err := fmt.Scanf("%d", &input)
		if err != nil {
			panic(err)
			stdin.ReadString('\n')
		}
		if input > result {
			tryNum++
			fmt.Println("Down!")
			continue
		} else if input < result {
			tryNum++
			fmt.Println("Up!")
			continue
		}
		fmt.Printf("success!!( try : %d)", tryNum)
		break
	}
}
