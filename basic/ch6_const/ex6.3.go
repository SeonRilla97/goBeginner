package main

import "fmt"

/*
상수

사용목적
1. 코드의 값에 의미를 부여하기 위함 (iota와 혼용)
2. 변하지 않는 값
3. 타입 미정 시 사용할 때 결정됨 -> 서버의 구성 설정 변수로 사용하면 좋을듯??
*/
const PI float64 = 3.14

const (
	Pig int = iota
	Cow
	Chicken
)

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("Pig")
	} else if animal == Cow {
		fmt.Println("Cow")
	} else if animal == Chicken {
		fmt.Println("Chicken")
	} else {
		fmt.Println("...")
	}
}

func main() {
	fmt.Println(Pig)
	fmt.Println(Cow)
	fmt.Println(Chicken)
}
