package ch11_struct

import "fmt"

/*
구조체

1. 구조체 선언 시 자원 효율성을 위해 8byte 이하 타입을 한곳으로 뭉친다. (메모리 패딩)
2. 구조체는 다른 구조체를 내장 타입이나, 필드 방식으로 포함할 수 있다.=

- 구조체는 모듈 내부의 모든 기능이 단일 목적에 충실하게 모여있게 한다 -> 응집도를 증가시킨다.
	= 개별 데이터에 집중하는것이 아닌, 구조체간의 관계와 상호작용에 대해 고민하게 되기 때문
*/

type User struct {
	Name string
	ID   string
	Age  int
}

// VIPUser는 User를 내장 타입으로 포함했다.
type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

// VIPUser2는 User를 필드로 포함했다.
type VIPUser2 struct {
	User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 23}

	// 구조체에 다른 구조체를 내장
	vip1 := VIPUser{user, 3, 250}
	fmt.Println(vip1.UserInfo.Name)

	// 다른 구조체를 필드로 내장
	vip2 := VIPUser2{user, 3, 250}
	fmt.Println(vip2.Name)

}
