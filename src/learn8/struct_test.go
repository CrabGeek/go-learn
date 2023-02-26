package learn8

import (
	"fmt"
	"math"
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

type Emp struct {
	name   string
	age    int8
	gender byte
}

func TestFunc2(t *testing.T) {
	emp1 := new(Emp)
	fmt.Printf("emp1: %T, %v, %p \n", emp1, emp1, emp1)
	(*emp1).name = "David"
	(*emp1).age = 30
	(*emp1).gender = 1

	emp1.name = "David2"
	emp1.age = 31
	emp1.gender = 1
	fmt.Println(emp1)
	fmt.Println("------------------")
	syntacticSuger()
}

func syntacticSuger() {
	arr := [4]int{10, 20, 30, 40}
	arr2 := &arr
	fmt.Println((*arr2)[len(arr)-1])
	fmt.Println(arr2[0])

	arr3 := []int{100, 200, 300, 400}
	arr4 := &arr3
	fmt.Println((*arr4)[len(arr)-1])
}

type Human struct {
	name   string
	age    int8
	gender byte
}

func TestStructIsValueType(t *testing.T) {
	h1 := Human{"Steven", 35, 1}
	fmt.Printf("h1: %T, %v, %p \n", h1, h1, &h1)
	fmt.Println("--------------------------")

	h2 := h1
	h2.name = "David"
	h2.age = 30
	fmt.Printf("h2 修改后: %T, %v, %p \n", h2, h2, &h2)
	fmt.Printf("h1: %T, %v, %p \n", h1, h1, &h1)

	fmt.Println("--------------------------")

	changeName(h1)

	fmt.Printf("h1: %T, %v, %p \n", h1, h1, &h1)
}

func changeName(h Human) {
	h.name = "Daniel"
	h.age = 13
	fmt.Printf("函数体内修改后: %T, %v, %p \n", h, h, &h)
}

func TestNoNameStruct(t *testing.T) {
	// 匿名函数
	res := func(a, b float64) float64 {
		return math.Pow(a, b)
	}(2, 3)

	fmt.Println(res)

	// 匿名结构体
	addr := struct {
		province, city string
	}{"陕西省", "西安市"}

	fmt.Println(addr)

	cat := struct {
		name, color string
		age         int8
	}{
		name:  "绒毛",
		color: "黑白",
		age:   1,
	}

	fmt.Println(cat)
}
