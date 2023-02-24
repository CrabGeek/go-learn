package main

import "fmt"

func main() {
	chinese := 90
	english := 80.9
	var i int8 = 1 << 2

	avg := (chinese + int(english)) / 2
	avg2 := (float64(chinese) + english) / 2
	fmt.Printf("%T, %d\n", avg, avg)
	fmt.Printf("%T, %f\n", avg2, avg2)
	fmt.Println(i)

}
