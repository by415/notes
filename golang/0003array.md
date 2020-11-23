### 0003 array [code](demo/array/array_test.go)
- 声明
> var arr [2]int    
>
> arr := [...]int{1,3,4,5}       // 长度由{}列表决定
>
> arr := [3]int{1:1, 2:4}        // 通过索引赋值，其他用默认值
>
> arr := [...]int{3:4, 9:3}      // 通过索引赋值，长度由最后一个索引确定

- 特点
> 创建完，长度固定，不可追加
>
> 是值类型，赋值或者作为参数都是值拷贝
> 长度是类型的组成部分，[10]int 和[20]int 不是一个类型
>
> 可根据数组创建切片 
>

---
*[👈 0000 golang](0000golang.md)*

[415 出品，必属精品](../note.md)
 
tags `array` `容器`


