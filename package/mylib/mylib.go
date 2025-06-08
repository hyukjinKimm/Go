package mylib

import "fmt"

func Hello() {
	fmt.Println("Hello from mylib!")
}

func secret() {
	fmt.Println("This is secret")
}

var Version = "1.0.0"

var secretKey = "abc123"

func init() {
	fmt.Println("myLib package 초기화")
}
