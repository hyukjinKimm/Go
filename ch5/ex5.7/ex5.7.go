package main

import "fmt"

func main() {
  var a int 
  var b int

  n, err := fmt.Scanln(&a, &b) // ① 값을 입력받습니다.
  if err != nil { // 에러 발생 시
    fmt.Println(n, err) // 에러를 출력합니다.
  } else {
    fmt.Println(n, a, b)
  }
  var s string
  fmt.Print("문자열 입력: ")
  fmt.Scanf("%s", &s)

  fmt.Println("입력한 문자열:", s)
}
