package array_test

import (
	"fmt"
	"reflect"
	"testing"
)

// TestTypeArrDiff 数组长度不同类型不同
func TestTypeArrDiff(t *testing.T) {
	var a [3]int = [3]int{1,1,1}
	var b [4]int
	c:= [...]int {1,3,4,55:3}

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(c))

	//a = append(a,3)	// error  arr cant append
}

func TestArrayCompare(t *testing.T) {
	a1:= [...]int{10:3}
	a2:= [...]int{10:3}
	if a1 == a2{
		println("==")
	}
	a2[4] = 3

	if a1 != a2{
		println("!=")
	}

	//a3:= [...]int{4:3}
	// invalid operation: a1 != a3 (mismatched types [11]int and [5]int)
	//if a1 != a3{
	//	println("!=")
	//}
}

// TestArrayDefine 数组初始化方法
func TestArrayDefine(t *testing.T) {
	var a [3]int = [3]int{1,1,1}
	fmt.Println(a)
	var b [3]int
	fmt.Println(b)

	c:= [...]int {1,3,4}
	fmt.Println(c)

	d := [4]int{1:10,3:20}
	fmt.Println(d)
	e := [...]int {3,4,4,30:2}
	fmt.Println(e)
}
