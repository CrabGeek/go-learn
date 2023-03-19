package learn11

import (
	"fmt"
	"testing"
)

func FuncA() {
	fmt.Println("这是funcA")
}

func FuncB() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println("恢复啦, 获取recover的返回值:", msg)
		}
	}()

	fmt.Println("这是funcB")
	for i := 0; i < 10; i++ {
		fmt.Println("i: ", i)

		if i == 5 {
			panic("funcB 恐慌啦")
		}
	}
}

func FuncC() {
	defer func() {
		fmt.Println("延迟执行函数")
		msg := recover()
		fmt.Println("获取recover的返回值: ", msg)
	}()

	fmt.Println("这是funcC")
	panic("funcC 恐慌了")
}

func TestRecover(t *testing.T) {
	FuncA()
	FuncB()
	FuncC()

	fmt.Println("main over")
}
