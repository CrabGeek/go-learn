package learn7

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	var country = map[string]string{
		"China":  "Beijing",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
		"France": "Paris",
		"Italy":  "Rome",
	}

	fmt.Println(country)

	rating := map[string]float64{"c": 5, "Go": 4.5, "Python": 4.5, "C++": 3}

	fmt.Println(rating)

	for k, v := range country {
		fmt.Println("国家 ", k, " 首都 ", v)
	}
}
