package learn4

import (
	"fmt"
	"testing"
)

func TestFunctionParameter(t *testing.T) {
	a := 10
	fmt.Printf("1. 变量a的内存地址: %p, 值为：%v \n\n", &a, a)
	fmt.Printf("==========int 型变量a的内存地址: %p \n\n", &a)
	changeIntVal(a)
	fmt.Printf("2. changeIntVal函数调用之后: 变量a的内存地址: %p, 值为: %v \n\n", &a, a)
	changeIntPtr(&a)
	fmt.Printf("3. changeIntPtr函数调用之后: 变量a的内存地址: %p, 值为: %v \n\n", &a, a)
}

func changeIntVal(a int) {
	fmt.Printf("-----------changeIntVal函数内: 值参数a的内存地址: %p, 值为 %v \n", &a, a)
	a = 90
}

func changeIntPtr(a *int) {
	fmt.Printf("-----------changeIntPtr函数内: 值参数a的内存地址: %p, 值为 %v \n", &a, a)
	*a = 50
}
