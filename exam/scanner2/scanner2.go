package main

import "fmt"

func main() {
	var a int
	fmt.Print("하나의 정수를 입력하세요:")
	fmt.Scanln(&a)

	if a%2 == 0 {
		fmt.Print("입력한 숫자 ", a, "는 짝수입니다.")
	} else {
		fmt.Print("입력한 숫자 ", a, "는 홀수입니다.")
	}

}
