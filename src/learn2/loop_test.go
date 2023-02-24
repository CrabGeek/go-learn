package learn2

import (
	"fmt"
	"testing"
)

func TestLoop(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
