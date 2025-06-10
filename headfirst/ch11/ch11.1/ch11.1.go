package mypkg

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethosWithParameters(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}

func (m MyType) MethosWithParameters(f float64) {
	fmt.Println("MethodWithParameter called with", f)
}

func (m MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}

func (m MyType) MethodNotInInterface() {
	fmt.Println("MethodNotInInterface called,")
}

func main() {

}
