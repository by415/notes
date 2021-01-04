# golang值传递和指针传递[code](demo/valuepoint/valuepoint_test.go)



golang中传递参数是分为两种情况，即**值传递**、**指针传递**，比较容易理解值传递就是在传递的时候拷贝一份参数的副本，指针传递其实也是拷贝了一份这个指针的副本，但是原指针和副本指针都指向了同一块内存地址，所以操作副本也会改变原值。

但是在golang中使用一些类型变量传递的时候，发现传递的是值却也能再操作副本后改变原值，`slice、map`，也称之为引用类型

### slice

slice和数组的区别：

```
Arrays, after declared of some size, cannot be resized, whereas slices can be resized. Slices are reference-types while arrays are value-type.
```
slice的结构在源码中为：
```
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```
可以看到slice其实是一个结构体，内部有一个指针，指向了底层数组，所以我们在值传递的时候，传递了一份指针的拷贝，但是指向的是同一份底层数组。

**值得注意**

既然slice是引用类型，那么关于slice的len和cap在传递中会不会发生改变呢？比如我传递一个slice，使用append进行元素添加，这个时候原来的slice的len会不会发生变化？还有cap呢？

```
var a = make([]int, 3, 6)

func change(b []int) {
	fmt.Printf("b 的地址：%p\n", &b)
	fmt.Printf("b 的len：%d\n", len(b))
	fmt.Printf("b 的cap：%d\n", cap(b))
	fmt.Println("b的值为：", b)
	b[0] = 2
	b = append(b, 4)
	b[0] = 10
	b[1] = 3
	b[2] = 30
	fmt.Printf("修改后的 b 的地址：%p\n", &b)
	fmt.Printf("修改后的 b 的len： %d\n", len(b))
	fmt.Printf("修改后的 b 的cap： %d\n", cap(b))
	fmt.Println("修改后的b的值为：", b)
}
	fmt.Printf("传参之前a 的地址：%p\n", &a)
	fmt.Printf("传参之前 a 的len：%d\n", len(a))
	fmt.Printf("传参之前 a 的cap：%d\n", cap(a))
	fmt.Println("传参之前 a的值为：", a)
	change(a)
	fmt.Printf("传参之后a 的地址：%p\n", &a)
	fmt.Printf("传参之后 a 的len：%d\n", len(a))
	fmt.Printf("传参之后 a 的cap：%d\n", cap(a))
	fmt.Println("传参之后 a的值为：", a)
```
这里声明了一个长度为3，容量为6的slice，然后使用参数传递，修改元素，输出结果为：
```
传参之前a 的地址：0xd3bdc0
传参之前 a 的len：3
传参之前 a 的cap：6
传参之前 a的值为： [0 0 0]
b 的地址：0xc0000044c0
b 的len：3
b 的cap：6
b的值为： [0 0 0]
修改后的 b 的地址：0xc0000044c0
修改后的 b 的len： 4
修改后的 b 的cap： 6
修改后的b的值为： [10 3 30 4]
传参之后a 的地址：0xd3bdc0
传参之后 a 的len：3
传参之后 a 的cap：6
传参之后 a的值为： [10 3 30]
```
结论1：slice的len和cap不会随着参数传递的修改而改变，所以说一个slice是唯一的，即它指向的数组下标是唯一的，如果参数传递中修改了在它下标范围内的元素，则原切片元素改变。

```
var a = make([]int, 3, 3)

func change(b []int) {
	fmt.Printf("b 的地址：%p\n", &b)
	fmt.Printf("b 的len：%d\n", len(b))
	fmt.Printf("b 的cap：%d\n", cap(b))
	fmt.Println("b的值为：", b)
	b[0] = 2
	b = append(b, 4)
	// fmt.Printf("c 的地址：%p\n", &c)
	// fmt.Printf("c 的len： %d\n", len(c))
	// fmt.Printf("c 的cap： %d\n", cap(c))
	// fmt.Println("c的值为：", c)
	b[0] = 10
	b[1] = 3
	b[2] = 30
	fmt.Printf("修改后的 b 的地址：%p\n", &b)
	fmt.Printf("修改后的 b 的len： %d\n", len(b))
	fmt.Printf("修改后的 b 的cap： %d\n", cap(b))
	fmt.Println("修改后的b的值为：", b)

	// fmt.Printf("append之后b 的地址：%p\n", &b)
	// fmt.Printf("append之后 b 的len：%d\n", len(b))
	// fmt.Printf("append之后 b 的cap：%d\n", cap(b))
	// fmt.Println("append之后 a的值为：", a, "append之后b的值为：", b)

}

func main() {
	fmt.Printf("传参之前a 的地址：%p\n", &a)
	fmt.Printf("传参之前 a 的len：%d\n", len(a))
	fmt.Printf("传参之前 a 的cap：%d\n", cap(a))
	fmt.Println("传参之前 a的值为：", a)
	change(a)
	fmt.Printf("传参之后a 的地址：%p\n", &a)
	fmt.Printf("传参之后 a 的len：%d\n", len(a))
	fmt.Printf("传参之后 a 的cap：%d\n", cap(a))
	fmt.Println("传参之后 a的值为：", a)
}
```
这里声明了一个len和cap都为3的slice，然后使用参数传递，修改元素，输出结果为：
```
传参之前a 的地址：0x6ebdc0
传参之前 a 的len：3
传参之前 a 的cap：3
传参之前 a的值为： [0 0 0]
b 的地址：0xc0000044c0
b 的len：3
b 的cap：3
b的值为： [0 0 0]
修改后的 b 的地址：0xc0000044c0
修改后的 b 的len： 4
修改后的 b 的cap： 6
修改后的b的值为： [10 3 30 4]
传参之后a 的地址：0x6ebdc0
传参之后 a 的len：3
传参之后 a 的cap：3
传参之后 a的值为： [2 0 0]
```
结论2：当slice为参数传递时，使用append函数扩容了之后，会创建一个新的slice底层数组，修改这个新的切片元素，原切片元素不会发生改变。但是在扩容前的修改依旧会影响原值。


### map

map的实现依赖的是哈希算法，数组和哈希是两种最常见数据结构，数组是用来表示元素的序列，哈希则表示键值之间的映射关系。哈希我们这里简单的解释下：

```
通过一些哈希算法，让键均匀的分布在一个有序的数组上，然后数组下面挂各式各样的节点，将值存在节点中，这就是一个简单的哈希过程。
```
golang的map实现也是基于此，这里不做详细展开，这里只说明为何map为引用类型。

go15.6源码中的map结构为：

```
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

// mapextra holds fields that are not present on all maps.
type mapextra struct {
	// If both key and elem do not contain pointers and are inline, then we mark bucket
	// type as containing no pointers. This avoids scanning such maps.
	// However, bmap.overflow is a pointer. In order to keep overflow buckets
	// alive, we store pointers to all overflow buckets in hmap.extra.overflow and hmap.extra.oldoverflow.
	// overflow and oldoverflow are only used if key and elem do not contain pointers.
	// overflow contains overflow buckets for hmap.buckets.
	// oldoverflow contains overflow buckets for hmap.oldbuckets.
	// The indirection allows to store a pointer to the slice in hiter.
	overflow    *[]*bmap
	oldoverflow *[]*bmap

	// nextOverflow holds a pointer to a free overflow bucket.
	nextOverflow *bmap
}

// A bucket for a Go map.
type bmap struct {
	// tophash generally contains the top byte of the hash value
	// for each key in this bucket. If tophash[0] < minTopHash,
	// tophash[0] is a bucket evacuation state instead.
	tophash [bucketCnt]uint8
	// Followed by bucketCnt keys and then bucketCnt elems.
	// NOTE: packing all the keys together and then all the elems together makes the
	// code a bit more complicated than alternating key/elem/key/elem/... but it allows
	// us to eliminate padding which would be needed for, e.g., map[int64]int8.
	// Followed by an overflow pointer.
}
```

buckets 是一个指针，最终它指向的是bmap结构体,但是编译期间会动态的增加这个结构体为：
```
type bmap struct {
    topbits  [8]uint8
    keys     [8]keytype
    values   [8]valuetype
    pad      uintptr
    overflow uintptr
}
```
bmap就是平时我们所说的桶，golang的map中桶最多装8个key，具体哈希算法我们这里不做展开。

通过map的结构可以看出，传递map的时候拷贝的副本中的指针也原map始指向同一块内存地址的，因此map也是一种引用类型。

### 根据需求确认传参类型

- 如果参数是个大的结构体，则应该使用指针传递避免大拷贝影响性能，同时推荐结构体尽量使用指针传递的方式。
- map、slice是引用类型，使用值传递即可。
- 其他场景推荐尽量使用值传递，因为值传递开销小于指针传递，因为Go使用逃逸分析来确定变量是否可以安全地分配到函数的栈帧上，这可能比在堆上分配变量开销小的多。通过值传递可以简化Go中的逃逸分析，并为变量提供更好的分配机会