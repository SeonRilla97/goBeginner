package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
1. 표준 입출력이란? : 여러 장치로부터 입력/출력을 fmt 패키지로 통합해서 제공한다.

	- 입력 :  메모리의 표준 입력 스트림에 임시 저장된다.
		- 입력값을 끝까지 읽지 못하고 종료되면 -> 나머지 문자는 스트림에 남아서 추후 입력 시 먼저 읽는다
			=> 스트림을 사용하고, 끝나면 무조건 비워라

	표준 입력의 인자는 변수의  주소값이 들어간다
		=> 내부적으로 Scanf 인자로 전달받은 변수가 포인터 타입인지 확인한다.

*/

func main() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, a, b)
	}

}
