package learn2

import (
	"fmt"
	"testing"
)

var a1 int = 7
var b1 int = 9

func TestFunction1(t *testing.T) {
	a1, b1, c1 := 10, 20, 0
	fmt.Printf("TestFunction1() 函数中 a1 = %d\n", a1)
	fmt.Printf("TestFunction1() 函数中 b1 = %d\n", b1)
	fmt.Printf("TestFunction1() 函数中 c1 = %d\n", c1)
	c1 = sum(a1, b1)
	fmt.Printf("TestFunction1() 函数中 c1 = %d\n", c1)
}

func sum(a1, b1 int) (c1 int) {
	a1++
	b1 += 2
	c1 = a1 + b1
	fmt.Printf("sum() 函数中 a1 = %d\n", a1)
	fmt.Printf("sum() 函数中 b1 = %d\n", b1)
	fmt.Printf("sum() 函数中 c1 = %d\n", c1)
	return c1
}
