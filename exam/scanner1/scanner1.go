package main

import "fmt"

func main() {
	var a string
	var b int

	fmt.Print("당신의 이름을 입력하세요:")
	fmt.Scanln(&a)
	fmt.Print("당신의 나이를 입력하세요:")
	fmt.Scanln(&b)
	fmt.Println("당신의 이름은", a, "이고, 나이는 ", b, "살 입니다.")

}
