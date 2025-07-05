package ch12_pointer

import "fmt"

/**
Pointer

특정 타입 변수의 주소값을 가지는 변수

필수 상식
	1. 인스턴스는 메모리에 생성된 데이터의 실체
	2. 포인터를 이용하면 인스턴스를 가리키게 할 수 있다.
	3. 함수 호출 시 인수를 통해 인스턴스를 전달하고, 그 값을 변경할 수 있다.
	4. 쓸모 없어진 인스턴스는 GC에 의해 자동으로 삭제된다.

스택 메모리와 힙 메모리
	1.스택 메모리 : 함수 내부에서 사용되는 변수
	2.힙 메모리 : 함수 외부로 공개되는 변수

	- Go언어의 탈출 분석을 통해 변수를 스택에 담을지 힙에 담을지 결정한다.
	- GO의 스택 영역은 동적 메모리 풀이다.
*/

type Data struct {
	value int
	data  [200]int
}

func changeData(arg *Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data

	// 포인터 타입 변수 선언 방법
	var p *Data = &data
	var pp *Data = &Data{}

	fmt.Println(p.value)
	fmt.Println(pp.data[100])
	// data 인스턴스의 주소가 함수로 넘어간다
	changeData(&data)
}
