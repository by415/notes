package interface_test

import (
	"fmt"
	"testing"
)

type Ali interface {
	walk()
	// fly()
	// jump()
}

type Alis struct{}

// 只能使用*Alis调用
// func (a *Alis) walk() {
// 	fmt.Println("walking")
// }

// func (a *Alis) fly() {
// 	fmt.Println("fly")
// }

// func (a *Alis) jump() {
// 	fmt.Println("jump")
// }

// Alis和*Alis类型都能调用
func (a Alis) walk() {
	fmt.Println("walkingaaaa")
}

// func (a Alis) fly() {
// 	fmt.Println("flyaaaa")
// }

// func (a Alis) jump() {
// 	fmt.Println("jumpaaa")
// }

func TestNilInterface(t *testing.T) {
	var i interface{}
	fmt.Println(i) //<nil>
}

func TestAlis(t *testing.T) {
	var i Ali
	b := Alis{}
	i = &b
	i.walk()
}

func checkType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Other")
	}
}

func TestInterfaceType(t *testing.T) {
	var a interface{} = "abcd"
	var b interface{} = 23
	checkType(a)
	checkType(b)
}

func InterfaceEqual(a, b interface{}) bool {
	return a == b
}

func TestInterfaceEqual(t *testing.T) {
	var a interface{}
	var b interface{}
	fmt.Println(InterfaceEqual(a, b))
	var c interface{} = "abc"
	var d interface{} = "abc"
	fmt.Println(InterfaceEqual(c, d))

}
