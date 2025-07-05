package main

import "fmt"

func main() {

	// slice1 과 slice2는 서로 다른 슬라이스지만 같은 배열을 참조한다.
	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4, 5) // slice1의 참조 배열에 4, 5 값을 추가한다

	// slice1의 변경이 slice2에 영향간다
	slice1[1] = 100
	// slice1에 데이터를 추가하면 slice2의 값이 변경된다.
	slice1 = append(slice1, 500)

	fmt.Println("slice1 :", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2 :", slice2, len(slice2), cap(slice2))

	/*
		결과 값
		slice1 : [0 100 0 500] 4 5
		slice2 : [0 100 0 500 5] 5 5


		분석
			slice1과 slice2가 참조하는 배열은 동일한 배열을 바라본다.
			Slice 내부의 len 과 cap은 다르다


		슬라이스 사용 방안
			1. 슬라이스를 다른 함수나 goroutine에 전달할 때는 항상 copy 사용 검토
			2. 슬라이스 공유가 의도된 것인지 명확히 문서화
			3. 코드 리뷰에서 슬라이스 사용 패턴 중점 체크
			4. 성능이 중요한 부분에서는 벤치마크 테스트로 copy 비용 검증
	*/
}
