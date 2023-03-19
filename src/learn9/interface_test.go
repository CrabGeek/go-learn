package learn9

import (
	"fmt"
	"testing"
)

type Phone interface {
	call()
}

type AndroidPhone struct {
}

type IPhone struct {
}

func (a AndroidPhone) call() {
	fmt.Println("我是安卓手机, 可以打电话了")
}

func (i IPhone) call() {
	fmt.Println("我是苹果手机, 可以打电话了")
}

func TestInterface(t *testing.T) {
	var phone Phone
	phone = new(AndroidPhone)
	fmt.Printf("%T, %v, %p \n", phone, phone, &phone)
	phone.call()

	phone = new(IPhone)
	fmt.Printf("%T, %v, %p \n", phone, phone, &phone)
	phone.call()
}
