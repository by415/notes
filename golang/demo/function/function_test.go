package function_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

// 命名返回值，sum不用局部声明，之后直接return
func D(a, b int) (sum int) {
	add := func(a, b int) int {
		return a + b
	}
	sum = add(a, b)
	return
}

func swap(a, b int) (int, int) {
	return b, a
}

func TestFunction(t *testing.T) {
	fmt.Println(D(3, 4))
	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}

func P0(a [5]int) {
	for b := range a {
		fmt.Println(b)
	}
}
func P1(a ...int) {
	for b := range a {
		fmt.Println(b)
	}
}
func P2(a []int) {
	for b := range a {
		fmt.Println(b)
	}
}

//func P2(a []float32) {
//	for b := range a {
//		fmt.Println(b)
//	}
//}
// TestParmFunction 不定参数
func TestParmFunction(t *testing.T) {
	slice := []int{11, 22, 33}
	P1(slice...)

	arr := [...]int{1, 2, 3, 4, 5}
	/// 数组不能作为参数传递给 不定参数
	//P1(arr)
	//P1(&arr)

	P2(slice)
	P0(arr)
	fmt.Println(reflect.TypeOf(P0))
	fmt.Println(reflect.TypeOf(P1))
	fmt.Println(reflect.TypeOf(P2))
}

func Add(a, b int) int {
	return a + b
}

type Op func(int, int) int

func RunOp(f Op, a, b int) {
	fmt.Println(f(a, b))
}

// TestFuncPointer 验证函数指针
func TestFuncPointer(t *testing.T) {
	RunOp(Add, 1, 3)
}

func Wrap(op string) func(int, int) int {
	switch op {
	case "add":
		return func(a, b int) int {
			return a + b
		}
	case "sub":
		return func(a int, b int) int {
			return a - b
		}
	}
	return nil
}

func TestFuncLam(t *testing.T) {
	defer fmt.Println("end first")
	defer fmt.Println("end second")
	fmt.Println(Wrap("add")(3, 4))
	fmt.Println(Wrap("sub")(111, 4))
	os.Exit(0)
}

func closures(a int) func(int) int {
	return func(i int) int {
		println(&a, a)
		a = a + i
		return a
	}
}

func closures2() func(int) int {
	a := -100 /// 转存在 堆上
	return func(i int) int {
		println(&a, a)
		a = a + i
		return a
	}
}

func TestClosures(t *testing.T) {
	f1 := closures(2) /// 转存在 堆上
	fmt.Println(f1(2))
	fmt.Println(f1(4))
	fmt.Println(f1(111))

	f2 := closures(2) /// 转存在 堆上
	fmt.Println(f2(2))
	fmt.Println(f2(4))
	fmt.Println(f2(111))
	fmt.Println()
	fmt.Println()

	f3 := closures2()
	fmt.Println(f3(2))
	fmt.Println(f3(4))
	fmt.Println(f3(111))

}
