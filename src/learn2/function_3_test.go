package learn2

import (
	"fmt"
	"math"
	"testing"
)

// 匿名函数
func TestFunction3(t *testing.T) {
	func(data int) {
		fmt.Println("Hello ", data)
	}(100)
}

func TestFunction3Second(t *testing.T) {
	func(data string) {
		fmt.Println(data)
	}("Hello World")
}

func TestFunction3Third(t *testing.T) {
	var arr []float64 = []float64{1, 9, 16, 25, 30}
	visit(arr, func(v float64) {
		v = math.Sqrt(v)
		fmt.Printf("%.2f\n", v)
	})

	fmt.Println("===========================")

	visit(arr, func(v float64) {
		v = math.Pow(v, 2)
		fmt.Printf("%.0f\n", v)
	})
}

func visit(list []float64, f func(float64)) {
	for _, value := range list {
		f(value)
	}
}
