package learn2

import (
	"fmt"
	"testing"
)

func TestClourse(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d \t", i)
		fmt.Println(add2(i))
	}
}

func add2(x int) int {
	var sum int = 0
	sum += x
	return sum
}

func TestClourse2(t *testing.T) {
	pos := adder()
	for i := 0; i < 10; i++ {
		fmt.Println("===========================")
		fmt.Println(pos(i))
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		fmt.Printf("sum1 = %d \t", sum)
		sum += x
		fmt.Printf("sum2 = %d \t", sum)
		return sum
	}
}
