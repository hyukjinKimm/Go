package main

import "fmt"

func main() {
	fmt.Print("음식 이름을 입력해주세요:")
	var a string
	fmt.Scanln(&a)

	fmt.Print("음식 가격을 입력해주세요:")
	var b int
	fmt.Scanln(&b)

	fmt.Print("음식 수량을 입력해주세요:")
	var c int
	fmt.Scanln(&c)

	fmt.Print(a, c, "개를 주문하셨습니다. 총 가격은 ", b*c, "원입니다.")

}
