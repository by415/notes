package main

import (
	"fmt"
	"practice/test_go/init"
	"time"
)

func funcIf() {
	a := 1
	if a == 2 {
		fmt.Println("xxx")
	} else if a == 1 {
		fmt.Printf("1")
	} else {
		fmt.Println("2")
	}
}

func funcFor() {
	//第一种形式
	i := 0
	for i < 3 {
		fmt.Println("i")
		i++
	}
	//第二种
	for j := 0; j < 3; j++ {
		fmt.Println("j")
	}
	//第三种
	for {
		fmt.Println("a")
		break
	}

	//范围for 遍历数组/切片/字符串/映射 信道读取
	m := map[int]int{
		3: 3,
		2: 2,
		1: 1,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}

	str := "hello world"
	for pos, ch := range str {
		fmt.Printf("%d %c\n", pos, ch)
	}
}

func funcSwitch(n int) {
	switch { //if-elseif的优化版本
	case 0 <= n && n < 10:
		fmt.Println(1)
	case 10 <= n && n < 100:
		fmt.Println(2)
	case 100 <= n && n < 1000:
		fmt.Println(3)
	}

	switch n { //普通switch
	case 1, 2, 3:
		fmt.Println(1)
	case 4, 5, 6:
		fmt.Println(2)
	case 7, 8, 9:
		fmt.Println(3)
	}

	//判断类型
	var t interface{}
	a := '0'
	t = a
	switch t := t.(type) {
	default:
		fmt.Println("default")
		fmt.Printf("%T", t)
	case bool:
		fmt.Printf("%T", t)
	case int:
		fmt.Printf("%T", t)
	}
}

func funcFunc() (n int) {
	n = 100
	n++
	return
}

func funcDefer() { //用来做结束释放，类似C++对象RAII
	defer fmt.Println("end")
	defer fmt.Println("end2")
	fmt.Println("start")
}

func funcNew() {
	// new不会初始化数据结构，只会置零。
	// go内置垃圾回收，没有局部变量作用域概念，new的作用似乎和普通变量取地址一样
}

//自定义类型
type MyFile struct {
	name string
	path string
	num  int
}

//复合文字构造 初始化列表？
func funcConstruct() {
	f := MyFile{"filename", "filepath", 10} //必须一一对应
	fmt.Println(f.name, f.path)

	f2 := MyFile{num: 10, path: "xxx"} //这样可以随意自由组合
	fmt.Println(f2.name, f2.path)
}

//make 仅适用于slice map channel
//make内置复杂结构的初始化函数
func funcMake() {
	m := make([]int, 100)
	m[0] = 100
	for i := range m {
		fmt.Printf("%d ", m[i])
	}
}

//数组，仅用于固定长度时，一半固定不变的数据结构，如配置
//数组赋值是全内存拷贝，所以涉及拷贝请使用切片
func funcArray() {
	a := [10]int{1, 2, 3}
	for n := range a {
		fmt.Printf("%d ", a[n])
	}
}

//切片 切片即动态数组，类似C++的vector，有起始地址，容量，长度三元组
func funcSlice2(buf []int) {
	buf = append(buf, 1)
}
func funcSlice() {
	s := make([]int, 10)
	for i := range s {
		fmt.Printf("%d ", s[i])
	}
	fmt.Printf("\n")
	fmt.Println(len(s), cap(s))
	funcSlice2(s)
	fmt.Println(len(s), cap(s))
	s = append(s, 1)
	s = append(s, 1)
	fmt.Println(len(s), cap(s))
}

//map kv结构数据
func funcMap() {
	var m = map[string]int{
		"hello": 1,
		"world": 2,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println(m["xxx"]) //这个操作不会插入
	m["xxx"] = 0
	fmt.Println(len(m))
	v, ok := m["xxx"]
	fmt.Println(v, ok)
	delete(m, "xxx") //删除某项
	fmt.Println(len(m))
	fmt.Println(m)
}

//内建函数append来扩容切片
func funcAppend() {
	x := []int{1, 2, 3}
	fmt.Println(len(x), cap(x)) // 3 3
	x = append(x, 4)
	fmt.Println(len(x), cap(x)) // 4 6 扩容2倍逻辑
	y := []int{4, 5, 6}
	x = append(x, y...)
	fmt.Println(len(x), cap(x)) // 7 12 扩容2倍逻辑
	fmt.Println(x)
}

func funcConst() {
	const (
		a = iota
		b
		c
		_
		d
	)
	fmt.Println(a, b, c, d) // 0 1 2 4

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
	)
	fmt.Println(KB, MB, GB)
}

var name string

//文件的初始化函数,引用了该包就会调用
func init() {
	name = "hello world"
	xxx.Fun()
}

//接口
type Sequence []int

func (s Sequence) Len() int {
	return len(s)
}

// func (s Sequence) Append(a int) {
// 	s = append(s, a)
// }

func (s *Sequence) Append(a int) {
	*s = append(*s, a)
}

func funcInterface() {
	var s Sequence
	fmt.Println(s.Len())
	s.Append(1)
	fmt.Println(s.Len())
}

func funcTypeAssert() {
	var value interface{}
	value = "hello"
	str := value.(string)
	fmt.Println(str)

	i, ok := value.(int)
	if ok {
		fmt.Println("ok")
	}
	fmt.Println(i)
}

//goroutinue
func funcGo() {
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("go")
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("end")
}

//channel 协程间通信
func funcChan() {
	c := make(chan int)
	go func() {
		fmt.Println("sync1")
		c <- 100
	}()
	a := <-c
	fmt.Println("sync2", a)

	c2 := make(chan int, 100)

	go func() {
		fmt.Println("produce:")
		for i := 0; i < 100; i++ {
			c2 <- i
		}
	}()
	for i := 0; i < 100; i++ {
		a2 := <-c2
		fmt.Println("consume:", a2)
	}
}

func main() {
	//funcIf()
	//funcFor()
	//funcSwitch(4)
	//fmt.Println(funcFunc())
	//funcDefer()
	//funcConstruct()
	//funcMake()
	//funcArray()
	//funcSlice()
	//funcMap()
	//funcAppend()
	//funcConst()
	//funcInterface()
	//funcTypeAssert()
	//funcGo()
	funcChan()
}
