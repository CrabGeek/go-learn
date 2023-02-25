package learn8

import (
	"fmt"
	"testing"
)

type Teacher struct {
	name   string
	age    int
	gender byte
}

func TestStruct(t *testing.T) {
	// 1. var 声明方式实例化结构体，初始化方式为: 对象.属性=值
	var t1 Teacher

	fmt.Println(t1)
	fmt.Printf("t1:%T, %v, %q \n", t1, t1, t1)

	t1.name = "Steven"
	t1.age = 34
	t1.gender = 1

	fmt.Println(t1)
	fmt.Println("==============================")

	t2 := Teacher{}
	t2.name = "David"
	t2.age = 30
	t2.gender = 1
	fmt.Println(t2)
	fmt.Println("==============================")

	t3 := Teacher{
		name:   "Josh",
		age:    28,
		gender: 1,
	}
	fmt.Println(t3)
	fmt.Println("==============================")

	t4 := Teacher{"Ruby", 30, 0}
	fmt.Println(t4)
	fmt.Println("==============================")
}
