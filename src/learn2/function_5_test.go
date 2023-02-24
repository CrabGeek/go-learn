package learn2

import (
	"fmt"
	"testing"
)

func TestFlexiableFunction(t *testing.T) {
	sum, avg, count := getScore(90, 82.5, 73, 64.8)
	fmt.Printf("学院共有%d门成绩, 总成绩为: %.2f, 平均成绩为: %.2f\n", count, sum, avg)
	fmt.Println()
	scores := []float64{92, 72.5, 93, 74.5, 89, 87, 74}
	sum, avg, count = getScore(scores...)
	fmt.Printf("学员共有%d门成绩, 总成绩为: %.2f, 平均成绩为: %.2f\n", count, sum, avg)
}

func getScore(scores ...float64) (sum, avg float64, count int) {
	for _, value := range scores {
		sum += value
		count++
	}

	avg = sum / float64(count)
	return
}
