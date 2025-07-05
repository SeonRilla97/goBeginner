package main

type account struct {
	balance   int
	firstName string
	lastName  string
}

// 포인터 메서드
func (a1 *account) withdrawpointer(amount int) {
	a1.balance -= amount
}

// 값 타입 메서드
func (a1 account) withdrawValue(amount int) {
	a1.balance += amount
}

// 변경된 값을 가진 신규 구조체 반환
func (a1 account) withdrawReturnValue(amount int) account {
	a1.balance += amount
	return a1
}

func main() {
	var mainA *account = &account{100, "seon", "jinyu"}

	mainA.withdrawpointer(30) // mainA 금액 70
	mainA.withdrawValue(20)   //mainA 금액 50

	var mainB account = mainA.withdrawReturnValue(20) //mainB 금액 50
	mainB.withdrawpointer(30)                         //mainB 금액 20

}
