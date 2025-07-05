package main

/*
*
슬라이싱

배열을 슬라이스로 변환하는 기법으로, 신규 배열을 생성하지 않고 기존 배열의 참조를 이용하여 슬라이스를 생성한다.
*/
func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:2]

	arr[1] = 100
	slice = append(slice, 100)
}
