package language_test

import (
	"container/list"
	"fmt"
	"reflect"
	"testing"
)

func PrintValueTpye(a interface{}) {
	// 方法1：
	println(reflect.TypeOf(a).Name())
	// 方法2：
	fmt.Println(reflect.TypeOf(a))
	// 方法3：
	fmt.Printf(`%T`, a)
	fmt.Println()
	fmt.Println()
}

// TestType 声明实例 - 显示完整声明
func TestValue2(t *testing.T) {
	fmt.Println("变量短类型声明")
	a := 3
	PrintValueTpye(a)
	b, c := false, ""
	PrintValueTpye(b)
	PrintValueTpye(c)
}

// TestValue 声明实例 - 显示完整声明
func TestValue(t *testing.T) {
	fmt.Println("变量显示完整声明")

	var int0 int
	fmt.Println(int0)

	var int1, int2 int
	fmt.Println(int1, int2)

	var str0 string
	fmt.Println(str0)

	var list1 = list.New()
	fmt.Println(list1)

	var map0 map[int]int = map[int]int{1: 3}
	fmt.Println(map0)

	var str1 string = "str1"
	fmt.Println(str1)

	var bool0 bool
	fmt.Println(bool0)
}

// TestStringTrans 验证string转换
func TestStringTrans(t *testing.T) {
	fmt.Println("验证string转换")
	a := "testString"
	b := []byte(a)
	c := []rune(a)

	PrintValueTpye(b)
	PrintValueTpye(c)
}

// TestStringConst 验证string的不可修改特性
func TestStringConst(t *testing.T) {
	fmt.Println("验证string的不可修改特性")
	a := "test"
	b := a[0]   // 内容不可修改
	b = 'd'
	fmt.Println(a)
	c := a[1:]  // 内容不可修改
	d := a[2:3] // 内容不可修改
	fmt.Println(len(a))

	fmt.Println(a)
	fmt.Printf("%c\n", b)

	fmt.Println(c)
	fmt.Println(d)

	PrintValueTpye(a)
	PrintValueTpye(c)
	PrintValueTpye(d)
	b = 'a' // error
	//c[1] = 'c'	// error
	c = "ss"
	//d[0] = 'c'	// error
	fmt.Println(a)
	fmt.Printf("%c", b)
	fmt.Println(c)
	fmt.Println(d)

}

// TestConst 验证const
func TestConst(t *testing.T) {
	fmt.Println("验证const")
	const (
		a = iota // 0
		b        // 1
		c        // 2
		_        // 3
		e
		f = iota * 3              // 5*3
		g                         // 6*3
		h float32    = iota       // 7
		i            = iota * 5.2 // 8*5
		j                         // 9*5
	)
	fmt.Println(a, b, c, e, f, g, h, i, j)
	//0 1 2 12 15 18 7 40 45
	//说明 iota 的值是从0开始递增的，

	//高效使用
	const (
		_ = iota
		K = 1 << (iota * 10) // 2*10
		M                    // 2*20
		G                    // 2*30
	)
	fmt.Println("KMG", K, M, G)
}

func TestSwitch(t *testing.T) {
	a := 3
	switch a {
	case 3:
		fmt.Println(3)
		fallthrough // 无论如何输出吓一跳
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	}

	switch a {
	case 3, 5:
		fmt.Println("3,5")
	}

	// case 代替if-else
	switch {
	case true:
		fmt.Println(true)
	}
}


