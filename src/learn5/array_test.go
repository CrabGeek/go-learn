package learn5

import (
	"fmt"
	"testing"
)

func TestArrayPhase1(t *testing.T) {
	a := [4]float64{67.7, 89.8, 21, 78}
	b := [...]int{2, 3, 5}

	fmt.Printf("数组a 的长度为%d, 数组b的长度为 %d\n", len(a), len(b))
}

func TestItercationArray(t *testing.T) {
	a := [4]float64{67.7, 89.8, 21, 78}
	b := [...]int{2, 3, 5}

	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], "\t")
	}

	fmt.Println()

	for _, value := range b {
		fmt.Print(value, "\t")
	}
}

func Test2DArray(t *testing.T) {
	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	fmt.Println(len(a))
	fmt.Println(len(a[0]))

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

func TestArrayIsValueType(t *testing.T) {
	a := [...]string{"USA", "CHINA", "INDIA", "GERMANY", "FRANCE"}
	b := a
	b[0] = "SINGAPORE"

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}
