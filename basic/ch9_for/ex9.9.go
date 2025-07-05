package main

import "fmt"

/*
*
for

0. for 초기문; 조건문; 후처리 { 코드블록 }
1. 무한 루프 : for {}
2. break : 루프 종료, label 설정 시 다중 루프 종료 가능
3. continue : 이후 코드 실행않고 후처리문 실행, 다음 조건문 실행
*/
func main() {
	a := 1
	b := 1

OuterFor:
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor
			}
		}
	}
	fmt.Println("%d * %d = %d\n", a, b, a*b)
}
