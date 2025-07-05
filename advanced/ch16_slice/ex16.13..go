package main

import (
	"fmt"
	"sort"
)

type student struct {
	Name string
	Age  int
}

/*
*
구조체 슬라이스의 정렬 구현

  - 슬라이스 타입에 대해 : Len, Less, Swap 메서드를 구현한다
*/
type Students []student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age > s[j].Age } // 내림차순 정렬
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	s := Students{
		{"a", 31},
		{"b", 21},
		{"c", 32},
		{"d", 40},
		{"e", 15},
	}

	sort.Sort(s)
	fmt.Println(s)
}
