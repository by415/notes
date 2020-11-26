package perfect

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"unsafe"
)

//2 3
//2 3
//8 9
//8 9
//8 9
//8 9
//8 9
//8 9
//8 9
func TestRange(t *testing.T) {
	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//for i, v := range si {
	//	fmt.Println(i, v)
	//}

	wg := sync.WaitGroup{}

	/// 闭包，变量存在竞争问题
	for i, v := range si {
		wg.Add(1)
		go func() {
			fmt.Println(i, v)
			wg.Done()
		}()
	}

	//for i := 0; i < 8; i++ {
	//	wg.Add(1)
	//	go func() {
	//		fmt.Println(si[i])
	//		wg.Done()
	//	}()
	//}
	//2 3
	//2 3
	//8 9
	//8 9
	//8 9
	//8 9
	//8 9
	//8 9
	//8 9

	/// 合理
	//for i, v := range si {
	//	wg.Add(1)
	//	go func(i,v int) {
	//		fmt.Println(i, v)
	//		wg.Done()
	//	}(i,v)
	//}
	//1 2
	//0 1
	//8 9
	//6 7
	//4 5
	//7 8
	//5 6
	//3 4
	//2 3
	wg.Wait()
}

func deferT() (r int) {
	r = 4
	defer func() {
		r = 22
	}()
	return 3
}

func deferTT() (r int) {
	t := 4
	defer func() {
		t = 33 // 不会影响返回值
		r = 44 // 影响返回值
	}()
	return t
}

func deferTTT() (r int) {
	t := 4
	defer func(r int) {
		r = r + 4 // defer 修改形参，不影响返回值
		return
	}(t)
	return t

}

func TestDefer(t *testing.T) {
	fmt.Println(deferT())
	fmt.Println(deferTT())
	fmt.Println(deferTTT())
}
func f(a [4]int) {
	a[2] = 2222
	fmt.Println(a)
}

// 数组任何时候都是只拷贝，不存在任何异常问题
func TestArray(t *testing.T) {
	arr := [4]int{1, 3, 4, 5}
	arrb := arr
	arrb[2] = 33
	fmt.Printf("%p,%v\n", &arr, arr)
	fmt.Printf("%p,%v\n", &arrb, arrb)

	f(arr)

	fmt.Printf("%p,%v\n", &arr, arr)
	fmt.Printf("%p,%v\n", &arrb, arrb)

	c := struct {
		s [4]int
	}{
		s: arr,
	}

	d := c
	c.s[2] = 1111
	d.s[2] = 4444

	fmt.Println(arr)
	fmt.Println(c)
	fmt.Println(d)
}

func TestSlice(t *testing.T) {
	var a []int
	b := make([]int, 0)

	//a is nil
	//b is not nil
	// 区别在于 b 的 array unsafe.Pointer 是有分配空间的

	//type slice struct {
	//	array unsafe.Pointer
	//	len   int
	//	cap   int
	//}
	if a == nil {
		fmt.Println("a is nil")
	} else {
		fmt.Println("a is not nil")
	}
	if b == nil {
		fmt.Println("b is nil")
	} else {
		fmt.Println("b is not nil")
	}

	sa := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	sb := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	fmt.Println(len(a), cap(a), sa.Data)
	fmt.Println(len(b), cap(b), sb.Data)
	fmt.Println(sa)
	fmt.Println(sb)

}

func TestSliceShareData2(t *testing.T) {
	a := []int{1, 2, 3, 4, 6, 7, 8}
	b := a[0:4]
	fmt.Println(a)
	fmt.Println(b)
	b = append(b, 2, 2, 2)

	fmt.Println(a)
	fmt.Println(b)

	//[1 2 3 4 6 7 8]
	//[1 2 3 4]

	//[1 2 3 4 2 2 2]
	//[1 2 3 4 2 2 2]
}

func TestSliceShareData(t *testing.T) {
	a := []int{1, 2, 3, 4, 6, 7, 8}
	sa := (*reflect.SliceHeader)(unsafe.Pointer(&a))

	/// 地址不同，指向的结构是同一个底层数组，偏移 = 起始下标*8
	b := a[0:4]
	sb := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	c := a[1:4]
	sc := (*reflect.SliceHeader)(unsafe.Pointer(&c))

	fmt.Println(sa)
	fmt.Printf("%p\n\n", (&a))
	fmt.Println(sb)
	fmt.Printf("%p\n\n", (&b))
	fmt.Println(sc)
	fmt.Printf("%p\n\n", (&c))

	b[2] = 22 /// 两个都改变
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	b = append(b, 3, 4, 4, 5, 5, 6, 6, 6, 67, 7, 7, 7, 7, 7, 7, 7)
	fmt.Println()
	fmt.Println()
	fmt.Println(sa)
	fmt.Printf("%p\n\n", (&a))
	fmt.Println(sb)
	fmt.Printf("%p\n\n", (&b))
	fmt.Println(sc)
	fmt.Printf("%p\n\n", (&c))

	a[2] = 444 /// a,c两个都改变 ,b已经重新分配空间，不影响
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func TestReference(t *testing.T) {
	a := 3
	fmt.Printf("%p\n", &a)
	fmt.Println(a)

	func() () {
		fmt.Printf("%p\n", &a)
		a = 4
	}()
	fmt.Println(a)
}

func GenerateInt() {
	generateInt()
}
func generateInt() {
}
func TestFunction(t *testing.T) {
	GenerateInt()
}

type InterfaceTest interface {
	test()
}
type MyTestStruct struct {
}

//func (m MyTestStruct) test() {
//	panic("implement me")
//}

var (
	//_               = MyTestStruct.test /// MyTestStruct.test undefined (type MyTestStruct has no method test)
//_ InterfaceTest = MyTestStruct{}    ///cannot use MyTestStruct literal (type MyTestStruct) as type InterfaceTest in assignment:	MyTestStruct does not implement InterfaceTest (missing test method)
)

func TestInterface(t *testing.T) {
}
