### go 语言陷阱 [code](demo/perfect/perfect_test.go)                                                                                                                                    

#### 多值赋值
```go
var x,y int = 2,3
var x,y = 2,"424"
x,y := f()
i,v := range S {}
x,y = y,x
```

#### range 局部变量复用
```go
	/// 闭包，变量存在竞争问题 
	// wg := sync.WaitGroup{}
	//for i, v := range si {
	//	wg.Add(1)
	//	go func() {
	//		fmt.Println(i, v)
	//		wg.Done()
	//	}()
	// wg.Wait()
	//}
    // 输出如下
	//2 3
	//2 3
	//8 9
	//8 9
	//8 9
	//8 9
	//8 9
	//8 9
	//8 9
```
```go

	/// 合理作为参数传递
	for i, v := range si {
		wg.Add(1)
		go func(i,v int) {
			fmt.Println(i, v)
			wg.Done()
		}(i,v)
	}
```
#### defer 陷阱

- 影响返回值
```go
func f()(r int)

函数 有名返回值 r 是被分配在栈上的。defer 作为函数的闭包函数\，拥有对r的引用操作。 在defer中操作r会影响实际返回值的

具体参考code


```
- 影响性能


#### 数组
无论何时都是明确值传递

> 数组直接赋值
>
> 作为函数参数
>
> 内嵌到结构体

#### 切片
```go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```
> 相当于c的数组，指针与值之间切换
>
> 如下两个声明，结果是不一样的 可以查看code中 TestSlice
>
> 	var a []int
>
> 	b := make([]int, 0)

多个切片引用同一个底层数组引发的混乱问题 参考code
> 通过切片[:]方式产生新切片引用 同一个底层数组，起始地址不同，cap相同，len不同
>
> a[0]=33 方式 修改，会影响所有引用该数据的内容
>
> append 多个数据触发空间分配，会重新拷贝该节点的内存。不影响其他切片
> 

#### 值，指针和引用

go 只有一种参数传递规则，值传递。如果传递的是指针或者复合结构，其指向依旧是同一个地址，所以会改变内容

复合结构包含 chan,map,slice,interface

```go

type hchan struct {
	qcount   uint           // total data in the queue
	dataqsiz uint           // size of the circular queue
	buf      unsafe.Pointer // points to an array of dataqsiz elements
	elemsize uint16
	closed   uint32
	elemtype *_type // element type
	sendx    uint   // send index
	recvx    uint   // receive index
	recvq    waitq  // list of recv waiters
	sendq    waitq  // list of send waiters

	// lock protects all fields in hchan, as well as several
	// fields in sudogs blocked on this channel.
	//
	// Do not change another G's status while holding this lock
	// (in particular, do not ready a G), as this can deadlock
	// with stack shrinking.
	lock mutex
}

// A header for a Go map.
type hmap struct {
	// Note: the format of the hmap is also encoded in cmd/compile/internal/gc/reflect.go.
	// Make sure this stays in sync with the compiler's definition.
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}

...

```

#### 函数名的意义

函数名与匿名函数字面量的值有三层含义

> 类型信息，表明其数据类型是函数类型
>
> 函数名代表函数的执行代码的起始位置
>
> 可以通过函数名进行函数调用， func_name(func_param_list).底层四层含义  ///  相当于C堆栈过程
>
>       准备好参数
>       修改PC的值，跳转到函数代码起始位置并执行
>       复制值到函数的返回值栈区
>       跳过RET返回到函数的下一条指令处继续执行
>

#### 引用语义

针对闭包，其可以引用和修改外部变量，相当于C++ 的引用了。
```go
	a := 3
	fmt.Printf("%p\n", &a)
	fmt.Println(a)

	func() () {
		fmt.Printf("%p\n", &a)
		a = 4
	}()
	fmt.Println(a)

输出
0xc000094298
3
0xc000094298
4

```
#### 习惯用法
##### 干净与强迫症
- 编译器不能通过未使用的局部变量和标签
- import 未使用的包 不能编译通过
- 所有控制结构，函数，方法定义 { 必须放在行尾
- go fmt 工具格式化代码，使得所有代码风格保持统一

##### comma,ok 表达式
- 获取map值
- 读取chan值
- 类型断言

##### 简写模式
- import() 多个包
- var() 声明多个变量

#### 函数和方法设计

对于复杂函数，方法的设计，可以使用如 GenerateInt(),generateInt()
两个同名函数的方法处理。优点在于

1.GenerateInt() 方法用于包外调用，并且保持简洁。

2.generateInt将详细设计做隔离。很好的一种分层处理

#### 多值返回值
多返回值函数，如果包含error，bool 类型返回值，则将error/bool放在最后一个返回值，作为一种编程风格。

没有对错之说

Go标准库就是遵循这样的设计

#### 对于接口的明确实现

期望结构实现某接口，将未实现当成错误暴漏在编译器,以下 2选1
```go
var (
	_               = MyTestStruct.test  /// MyTestStruct.test undefined (type MyTestStruct has no method test)
	_ InterfaceTest = MyTestStruct{} 	///cannot use MyTestStruct literal (type MyTestStruct) as type InterfaceTest in assignment:	MyTestStruct does not implement InterfaceTest (missing test method)
)
```

```go
type InterfaceTest interface {
	test()
}
type MyTestStruct struct {}
func (m MyTestStruct) test() {
}
var (
	_               = MyTestStruct.test  /// MyTestStruct.test undefined (type MyTestStruct has no method test)
	_ InterfaceTest = MyTestStruct{} 	///cannot use MyTestStruct literal (type MyTestStruct) as type InterfaceTest in assignment:	MyTestStruct does not implement InterfaceTest (missing test method)
)
```

---


*[👈 0000 golang](0000golang.md)*

[415 出品，必属精品](../note.md) 

tags `` `` 