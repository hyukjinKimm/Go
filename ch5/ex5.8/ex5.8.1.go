package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Print("정수 2개 입력: ")
	fmt.Scanf("%d %d\n", &a, &b)

	fmt.Println("입력한 값:", a, b)

	var s string
	fmt.Print("문자열 입력: ")
	fmt.Scanf("%s", &s)

	fmt.Println("입력한 문자열:", s)
}
