# 0012golang中interface原理[code](demo/interface/interface_test.go)

golang中interface是一个抽象概念，可以在go中实现多态。interface类型的变量可以保存任何实现该接口的类型的值，并且我们可以通过对interface进行类型断言获取对应类型的具体值。

## interface的使用

1. 用于接收任何值

在代码中最常用的就是声明一个空接口类型的变量，来接收任何类型的值，这个时候的接口称为空接口，表示不包含任何功能。

```
package interface_test

import (
	"fmt"
	"testing"
)

func TestNilInterface(t *testing.T) {
	var i interface{}
	fmt.Println(i) //<nil>
}
```

2. 实现接口用于多态

golang中interface被声明为一个类型

`type Ali interface{}`

接口中可以声明方法，其他类型实现这个接口的所有方法，就称这个类型实现了该接口，并且使用这个接口类型的值接收实现了此接口的类型的值，可以实现多态。**特别注意，使用\*T和T实现的方法调用方式不一样，使用T实现的方法，\*T和T都能调用；但是使用\*T实现的是能使用\*T调用**

```
type Ali interface {
	walk()
}
type Alis struct{}

//只能使用*Alis调用
func (a *Alis) walk() {
	fmt.Println("walking")
}


// Alis和*Alis类型都能调用
//func (a Alis) walk() {
//	fmt.Println("walkingaaaa")
//}
func TestAlis(t *testing.T) {
	var i Ali
	b := Alis{}
	i = &b
    // i = b  编译器报错
	i.walk()
}
```

3.对interface进行类型断言

```
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
	checkType(a)    // String
	checkType(b)    // Int
}
```

4.接口值判等

满足一下条件的interface{}值之间可以判等

- 都为interface{}的零值
- 具有相同的具体值和动态类型

```
func InterfaceEqual(a, b interface{}) bool {
	return a == b
}

func TestInterfaceEqual(t *testing.T) {
	var a interface{}
	var b interface{}
	fmt.Println(InterfaceEqual(a, b))   // true
	var c interface{} = "abc"
	var d interface{} = "abc"
	fmt.Println(InterfaceEqual(c, d))   // true

}
```

## interface原理

### interface的数据结构

#### `eface和iface`

```
// 空的interface{}
type eface struct {
	_type *_type
	data  unsafe.Pointer
}

// 至少带有一个函数的interface{}
type iface struct {
	tab  *itab
	data unsafe.Pointer
}
```

- `eface`

`eface`是一个两个机器字长表示的结构，`_type`和`data`组成，第一个字`_type`指向实际的类型描述的指针，`data`表示的是数据指针。
```
type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      tflag
	align      uint8
	fieldAlign uint8
	kind       uint8
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte
	str       nameOff
	ptrToThis typeOff
}
```

- `iface`

```
type itab struct {
	inter *interfacetype
	_type *_type
	hash  uint32 // copy of _type.hash. Used for type switches.
	_     [4]byte
	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}
```

`iface`表示至少带有一个函数的interface，也是两个机器字长表示，第一个字`tab`指向一个`itab`结构，第二个字表示数据指针。


TODO：图片-----iface和eface的结构图--------



- `iface`和`eface`中的`data`

`data`实际上存的就是具体数据的指针，在interface中的设计如下：

1.实际类型是一个值，则interface会保存一份这个值的拷贝，interface会在堆上分配一块内存，data指向它。

2.实际类型是一个指针，则interface会保存这个指针的一份拷贝，指针长度刚好是data的长度，所以data中存储的就是这个指针的值，也就是和实际变量指向同一个变量。

TODO：图片-----`data`存储值和指针的区别图--------


#### `itab`,`_type`

在上面`iface`的结构中有`itab`的结构体代码，可以看到`itab`中也是有`_type`这个成员的。

`itab`表示的是interface和实际类型之间的转换信息，对于每个interface和实际类型直接只要存在引关系，go运行的时候就会为这对具体的`<Tnterface,Type>`生成`itab`信息：

1.inter 指向对应的 interface 的类型信息。

2.type 和 eface 中的一样，指向的是实际类型的描述信息 _type

3.fun 为函数列表，表示对于该特定的实际类型而言，interface 中所有函数的地址。


TODO：图片-------itab的结构信息图----------

`_type`表示的是类型信息，具体信息是在编译期间生成的：

- size 为该类型所占用的字节数量。
- kind 表示类型的种类，如 bool、int、float、string、struct、interface 等。
- str 表示类型的名字信息，它是一个 nameOff(int32) 类型，通过这个 nameOff，可以找到- 类型的名字字符串

---
*[👈 0000 golang](0000golang.md)*

[415 出品，必属精品](../note.md) 

tags `interface`