package learn2

import (
	"fmt"
	"strings"
	"testing"
)

func TestFunction2(t *testing.T) {
	var result string = stringToLower("AbcdefGHijklMNOPqrstUVWxyz", processCase)
	fmt.Println(result)
	result = stringToLower2("AbcdefGHijklMNOPqrstUVWxyz", processCase)
	fmt.Println(result)
}

func processCase(str string) string {
	var result string = ""

	for i, value := range str {
		if i%2 == 0 {
			result += strings.ToUpper(string(value))
		} else {
			result += strings.ToLower(string(value))
		}
	}

	return result
}

func stringToLower(str string, f func(string) string) string {
	fmt.Printf("%T \n", f)
	return f(str)
}

// 声明了一个函数类型，通过type关键字，caseFunc会形成一种新的类型
type caseFunc func(string) string

func stringToLower2(str string, f caseFunc) string {
	fmt.Printf("%T \n", f)
	return f(str)
}

type processFunc func(int) bool

func TestFunction2Second(t *testing.T) {
	var slice []int = []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice = ", slice)
	odd := filter(slice, isOdd)
	even := filter(slice, isEven)
	fmt.Println("奇数元素: ", odd)
	fmt.Println("偶数元素: ", even)
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func filter(slice []int, f processFunc) []int {
	var result []int

	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
