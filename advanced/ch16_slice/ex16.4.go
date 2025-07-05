package main

func changeArr(arr2 [5]int) {
	arr2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	// Slice는 내부에 배열의 참조가 복사 되기에 동일한 배열을 바라본다.
	changeSlice(slice)

	// array는 새로운 배열을 복사하여 넘긴다 -> 서롣 다른 배열에 대해 작업한다.
	changeArr(arr)

}
