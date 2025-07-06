package main

import "fmt"

func main() {
	//맵 생성
	m := make(map[string]string)
	// 삽입
	m["A"] = "ABC"
	m["B"] = "BBC"
	m["C"] = "CCC"
	m["D"] = "DDD"
	m["E"] = "EEE"
	// 변경
	m["B"] = "BBF"

	//순회
	for k, v := range m {
		fmt.Println(k, v)
	}

	//삭제
	delete(m, "D")

}
