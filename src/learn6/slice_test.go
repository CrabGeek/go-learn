package learn6

import (
	"fmt"
	"testing"
)

func TestSlice1(t *testing.T) {
	var numbers = make([]int, 3, 5)
	fmt.Printf("%T\n", numbers)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)
}

func TestSliceChunck(t *testing.T) {
	// 创建切片
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	// 打印原始切片
	printSlice(numbers)

	fmt.Println("numbers == ", numbers)

	fmt.Println("numbers[1:4] == ", numbers[1:4])

	fmt.Println("numbers[:3] == ", numbers[:3])

	fmt.Println("numbers[4:] == ", numbers[4:])

	numbers2 := numbers[:2]
	printSlice(numbers2)

	numbers3 := numbers[2:5]
	printSlice(numbers3)
}

func printSlice(x []int) {
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(x), cap(x), x)
}
