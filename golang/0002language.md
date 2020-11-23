### 0002 部分重点基础 [code](demo/language/base_test.go)

#### 变量
go语言提供自动内存管理，我们不用关心生命周期和位置。
编译器使用[栈逃逸技术](https:\/\/blog.csdn.net\/u010853261\/article\/details\/102846449)自动分配空间，可能在堆或者栈上
#### 变量声明

- 显示完整声明
> var varName dataType [ = value ]
- 短类型声明
> varName := value \
> := 声明只能在函数/方法内

----------
#### 常量  
同c语言类似，用名称绑定一块内存地址。存在放程序的只读段(.rodata section)
> 
iota 预声明符 用于声明常量/枚举. 
> const()内一组使用时，值递增；\
> 单独const x = iota 其值每次都是从0开始

----------
#### 字符串
- 是常量，可以通过[]下标访问值，但是不能修改
```
    a := "testString"
    b := a[3]
    a[1]='3' //error
```
- ***字符串转切片[]bytes()慎用，会复制内容。数据量大时会有性能问题***
```
    a := "test"
    b := []bytes(a)
```
- 底层实现是个二元数据结构如：
```
    type stringStruct struct {
        str unsafe.Pointer  // 指向字节数组的指针
        len int             // 长度 
    }
```
- 字符串切片返回仍是指向相同底层字符数组的**子串**，同样不能修改
```
    a := "test"
    b := a[0]       
    b = 'd'         // ok  b为字符可修改。

    c := a[1:]      // 
    c[0] = 'x'      // c类型为string，不能通过下标修改
```
- 字符串转换
> []bytes(str) 转换为字节数组\
> []rune(str) 转unicode 字数组

- 运算
```
    a := "hello"
    b := "world"
    c := a + b  // 拼接
    d := len(c) // 内置求长度
    //遍历
    for i: = 0; i < len(a); i++ {
        fmt.Println(a[i])
    }

    for i,v := range a {
        fmt.Println(i,v)
    }
    
```
---
### 条件判断
- if
```
    if x:= true;x == true {
    }else {
    }
```
- switch
> switch 后可跟个初始化语句
>
> switch 后边表达式可选，没有的话，case相当于if-else
>
> 通过 fallthough 强制执行下一个 case语句（不会判断下一个case语句）
> 
> default 可以放在任意位置 
> switch 和 .(type) 进行类型查询

- for
```
for init; condition; post {
}
for condition {     // == while
}
for _,x range container {
} 
```


*[👈 0000 golang](0000golang.md)*

[415 出品，必属精品](../note.md)
 
tags `string` `const` `iota`