package ch8_switch

import "fmt"

/*
Switch

Case 종료 시 Break가 없어도 종료된다.
  - 만약 다음 case까지 확인이 필요한 경우 fallthrough를 사용한다
*/
type ColorType int

const (
	Red ColorType = iota
	Green
	Blue
)

func main() {

	switch color := ColorType(0); color {
	case Red:
		fmt.Println("Red")
		fallthrough // case Green 까지 실행
	case Green:
		fmt.Println("Green")
	case Blue:
		fmt.Println("Blue")
		return
	default:
		fmt.Println("Unknow")
	}
}
