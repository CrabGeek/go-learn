package learn10

import (
	"fmt"
	"testing"
)

func FuncA() {
	fmt.Println("FuncA")
}

func FuncB() {
	panic("FuncB")
}

func FuncC() {
	fmt.Println("FuncC")
}

func TestPanic(t *testing.T) {
	FuncA()
	FuncB()
	FuncC()
}
