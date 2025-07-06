package main

import "fmt"

type Stringer2 interface {
	String() string
}

type Student2 struct {
	Age int
}

func (s *Student2) String() string {
	return fmt.Sprintf("Student{Age:%d }", s.Age)
}

func PrintAge(stringer Stringer2) {
	s := stringer.(*Student2)
	fmt.Printf("Studeknt{Age:%d }", s.Age)
}

func main() {
	s := &Student2{15}
	PrintAge(s)
}
