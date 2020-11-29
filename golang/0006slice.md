
### 0006 slice [code](demo/slice/slice_test.go)
- 底层实现
```
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
```
- 切片创建
> 由数组创建
```
arr := [...]int{300:3}
slice := arr[3:]
```
> 由内置make创建
```
a := make([]int,30)  // len = 30 cap = 30
b := make([]int,30,40) // len = 30 cap = 40
```
- 操作
> len cap append copy
```
b := make([]int, 30, 40) // len = 30 cap = 40
fmt.Println(cap(b))
fmt.Println(len(b))
b = append(b, 333)
```
- 与string 转化
```
str := "hello test"
a := []byte(str)
b := []rune(str)
```


---
*[👈 0000 golang](0000golang.md)*

[415 出品，必属精品](../note.md) 

tags `slice` `容器`



