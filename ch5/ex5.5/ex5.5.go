package main

import "fmt"

func main() {
  var a int // ① 값을 저장할 변수
  var b int

  n, err := fmt.Scan(&a, &b) // ② 입력 두 개 받기
  if err != nil { // ③ 에러가 발생하면 에러 코드 출력
     fmt.Println(n, err) 
  } else { // ④ 정상 입력되면 입력값 출력
     fmt.Println(n, a, b)
  }
  	var s string
	fmt.Print("문자열 입력: ")
	fmt.Scanf("%s", &s)

	fmt.Println("입력한 문자열:", s)
}