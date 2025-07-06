package main

import "fmt"

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Println("int", t)
	case string:
		fmt.Println("string", t)
	case float64:
		fmt.Println("float64", t)
	case bool:
		fmt.Println("bool", t)
	default:
		fmt.Println("unknown type", t)
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hello")
	PrintVal(Student{10})
}
