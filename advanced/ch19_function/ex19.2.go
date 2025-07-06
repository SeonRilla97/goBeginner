package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer fmt.Println("반드시 호출됩니다.")
	defer f.Close()
	defer fmt.Println("파일을 닫았습니다.")

	fmt.Println("파일에 Helloworld를 씁니다. ")
	fmt.Fprintf(f, "Hello World")
}
