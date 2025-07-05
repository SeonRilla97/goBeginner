package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("난 %d살 %s야", s.Age, s.Name)
}

func main() {
	student := Student{
		"철수", 12,
	}
	var stringer Stringer
	stringer = student

	fmt.Printf("%s\n", stringer.String())
}
