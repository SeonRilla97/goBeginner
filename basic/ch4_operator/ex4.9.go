package ch4_operator

import (
	"fmt"
	"math"
	"math/big"
)

/**
연산자

정수의 언더플로우, 오버플로우를 항상 유의하여 타입을 지정해야한다.
실수의 정밀도를 올리기 위해선
	1. math.Nextafter()을 사용한다
	2. math.Float()을 사용한다.

*/

func main() {
	// nextafter 예제
	var x float64 = 0.1
	var y float64 = 0.2
	var z float64 = 0.3

	fmt.Printf("%0.18f + %0.18f : %v\n", z, x+y, equal(x+y, z))

	//math.float 예제
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)
	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d))

}

func equal(a, b float64) bool {
	/*
	   부동소수점은 연속적이지 않고 계단식으로 표현된다.
	   a에서 b 방향으로 한 계단 이동했을 때 b에 도달하면,
	   a와 b는 부동소수점 정밀도 범위 내에서 같다고 판단한다.
	*/
	return math.Nextafter(a, b) == b
}
