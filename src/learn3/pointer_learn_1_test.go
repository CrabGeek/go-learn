package learn3

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	a := 10
	fmt.Printf("变量地址: %x \n", &a)
}

func TestPointer2(t *testing.T) {
	var a int = 120
	var ip *int
	ip = &a

	fmt.Printf("a的类型是%T, 值是%v \n", a, a)

	fmt.Printf("&a 的类型是%T, 值是%v \n", &a, &a)

	fmt.Printf("ip 的类型是%T, 值是%v \n", ip, ip)

	fmt.Printf("*ip 变量的类型是%T, 值是%v \n", *ip, *ip)

	fmt.Printf("*&a 变量的类型是%T, 值是%v \n", *&a, *&a)

	fmt.Println(a, &a, *&a)

	fmt.Println(ip, &ip, *ip, *(&ip), &(*ip))
}

type Student struct {
	name    string
	age     int
	married bool
	gender  int8
}

func TestPointer3(t *testing.T) {
	var s1 = Student{"Steven", 35, true, 1}
	var s2 = Student{"Sunny", 20, false, 0}

	var a *Student = &s1
	var b *Student = &s2

	fmt.Printf("s1类型为%T, 值为%v \n", s1, s2)
	fmt.Printf("s2类型为%T, 值为%v \n", s2, s2)

	fmt.Printf("a类型为%T, 值为%v \n", a, a)
	fmt.Printf("b类型为%T, 值为%v \n", b, b)

	fmt.Printf("*a类型为%T, 值为%v \n", *a, *a)
	fmt.Printf("*b类型为%T, 值为%v \n", *b, *b)

	fmt.Println(s1.name, s1.age, s1.married, s1.gender)
	fmt.Println(a.name, a.age, a.married, a.gender)

	fmt.Println(s2.name, s2.age, s2.married, s2.gender)
	fmt.Println(b.name, b.age, b.married, b.gender)

	fmt.Println((*a).name, (*a).age, (*a).married, (*a).gender)
	fmt.Println((*b).name, (*b).age, (*b).married, (*b).gender)

	fmt.Printf("&a 类型为%T, 值为%v \n", &a, &a)
	fmt.Printf("&b 类型为%T, 值为%v \n", &b, &b)

	fmt.Println(&a.name, &a.age, &a.married, &a.gender)
	fmt.Println(&b.name, &b.age, &b.married, &b.gender)
}

func TestNilPointer(t *testing.T) {
	var ptr *int
	if ptr == nil {
		fmt.Println("ptr is nil pointer")
	} else {
		fmt.Println("ptr is not nil pointer")
	}
}

func TestPointerArray(t *testing.T) {
	const COUNT int = 3
	a := [COUNT]string{"abc", "ABC", "123"}
	var ptr [COUNT]*string

	fmt.Printf("%T, %v \n", ptr, ptr)

	fmt.Println(ptr[0])

	for i := 0; i < COUNT; i++ {
		ptr[i] = &a[i]
	}

	fmt.Printf("%T, %v \n", ptr, ptr)

	fmt.Println(ptr[0])

	for i := 0; i < COUNT; i++ {
		fmt.Printf("a[%d] = %s \n", i, *ptr[i])
	}
}

func TestPointerToPointer(t *testing.T) {
	var a int
	var ptr *int
	var pptr **int

	a = 1234
	ptr = &a

	pptr = &ptr

	fmt.Printf("变量 a = %d \n", a)
	fmt.Printf("指针变量 *ptr = %d \n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d \n", **pptr)
}
