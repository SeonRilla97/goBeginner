package main

import "fmt"

type opFunc2 func(int, int) int

func getOperator2(op string) opFunc2 {

	if op == "+" {
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}

	} else {
		return nil
	}
}

func main() {
	var operator opFunc2
	operator = getOperator2("*")

	var result = operator(3, 4)
	fmt.Println(result)
}
