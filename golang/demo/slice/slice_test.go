package slice_test

import (
	"fmt"
	"reflect"
	"testing"
)

// 验证string转切片
func TestSlice2String(t *testing.T) {
	str := "hello test"
	a := []byte(str)
	b := []rune(str)
	fmt.Println(a, b)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))
}

// TestSliceInit 验证切片生成
func TestSliceInit(t *testing.T) {
	arr := [...]int{3, 100: 4}
	sl := arr[4:]
	fmt.Println(arr)
	fmt.Println(sl)

	a := make([]int, 30)     // len = 30 cap = 30
	b := make([]int, 30, 40) // len = 30 cap = 40
	fmt.Println(cap(a))
	fmt.Println(len(a))

	fmt.Println(cap(b))
	fmt.Println(len(b))
	b = append(b, 333)
	println(len(b))
}
