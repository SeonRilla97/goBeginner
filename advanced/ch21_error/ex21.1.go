package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	defer file.Close()

	rd := bufio.NewReader(file)
	line, _ := rd.ReadString('\n')
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	// _는 이미 선언된 변수, := 연산자는 선언 및 할당 => 따라서 이미 선언된 _엔 알맞지 않은 연산자
	_, err = fmt.Println(file, line)
	return err
}

const filename string = "data.txt"

func main() {
	line, err := ReadFile(filename)

	if err != nil {
		err = WriteFile(filename, "This is WriteFile ")
		if err != nil {
			fmt.Println("파일 생성에 실패했습니다.", err)
			return
		}
	}
	line, err = ReadFile(filename)
	if err != nil {
		fmt.Println("파일 읽기에 실패했습니다")
		return
	}

	fmt.Println("파일내용" + line)
}
