package learn2

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var a int = 5

	if a == 0 {
		fmt.Println("Hello World")
	} else if a == 1 {
		fmt.Println("Bye~")
	} else {
		fmt.Println("lalala")
	}

}
