package main

import "fmt"

func main() {

	for true {
		var name string
		fmt.Print("이름을 입력하세요 (종료를 입력하면 종료:)")
		fmt.Scanln(&name)

		if name == "종료" {
			fmt.Println("프로그램을 종료합니다.")
			break
		}

		var age int
		fmt.Print("나이를 입력하세요: ")
		fmt.Scanln(&age)
		fmt.Println("입력한 이름: ", name, "나이: ", age)

	}
}
