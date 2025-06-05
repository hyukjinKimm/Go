package main

import "fmt"

func main() {
	if x := 7; x%2 == 0 {
		fmt.Println("짝수 입니다.")
	} else {
		fmt.Println("홀수 입니다.")
	}
}
