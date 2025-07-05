package main

import (
	"fmt"
	"strings"
)

func ToUpper(str string) string {
	var builder strings.Builder

	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	fmt.Println(ToUpper("Hello World"))
}
