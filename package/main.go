package main

import (
	"fmt"

	"example.com/package/mylib"
)

func main() {
	fmt.Println("main 시작")

	// mylib 패키지의 공개 함수 호출
	mylib.Hello()

	// 공개 변수 출력
	fmt.Println("버전:", mylib.Version)

	// 비공개 함수, 변수는 외부 접근 불가
	// mylib.secret() // 컴파일 에러
	// fmt.Println(mylib.secretKey) // 컴파일 에러
}
