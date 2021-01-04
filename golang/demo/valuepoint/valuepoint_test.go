package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestValue(t *testing.T) {

}

func reference(st []int) {
	st[0] = 1
	fmt.Println(st)

}
func TestReference(t *testing.T) {
	s := make([]int, 10)
	s[0] = 0
	reference(s)
}

// 64位操作系统下的指针是8个字节
// 所有一个slice类型占24个字节
func TestSizeofSlice(t *testing.T) {
	fmt.Println(unsafe.Sizeof([]string{})) //24
	a := 8
	b := &a
	fmt.Println(unsafe.Sizeof(b)) //8
}

//
func TestSizeofMap(t *testing.T) {
	fmt.Println(unsafe.Sizeof(map[int]int{}))
}

func change(a *int) {
	fmt.Println("传递后的地址：", &a)
	*a = 10
}

func TestPoint(t *testing.T) {
	a := 5
	change(&a)
	fmt.Println("原地址：", &a, a)
}
