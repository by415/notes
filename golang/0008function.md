### 函数 [code](demo/function/function_test.go)

> 函数 是一种类型\
> 支持多返回值\
> 支持闭包\
> 支持可变参数\
> 不支持默认值参数\
> 不支持函数重载\
> 不支持函数嵌套，但支持匿名函数嵌套
> 
>
> func funcName(parm-list)(result-list){\
> }
```
func A(){
}
func B()int {
    return 1
}
func C(a,b int) int {
    return a+b
}
func C(a,b int)(sum int) {
    sum = a+b
    return
}
func D(a,b int)(sum int) {  // 指定返回值
    add:=func(a,b int)(s int){
        return a+b
    }
    sum = add(a,b)
    return
}

func swap(a,b int)(int,int){
    return b,a
}
```
### 不定参数
> parm ...type
>
> 所有不定参数类型相同\
> 不定参数必须是最后一个参数\
> 不定参数相当于切片\
> 切片可以作为不定参数传递， 切片名后边 ...\
> 形参为不定参数和形参为切片不是同类型函数
>  

### 函数签名
> c++的函数指针

### 匿名函数
> 有名/匿名函数
```go
func (a,b int)(int){
    return a+b
}
b:= func (a,b int)(int){
    return a-b
}
```

### defer 
> 函数结束时执行。\
> 多个时，先注册后执行\
> 主动调用 os.Exit(1) defer会被跳过

---

###闭包
 由函数和引用环境组成。一般为匿名函数，引用局部变量，全局变量构成
> 闭包 = 函数+ 引用环境 \

***环境引入为直接引入;编译器检测到引用，会将对应变量分配在堆上***
> 如果函数返回的闭包引用了该函数的局部变量：\
> 多次调用该函数，返回的多个闭包所引用的外部变量是多个副本,因为每次调用都会为局部变量分配内存\
> 用一个闭包函数多次，如果闭包修改了引用变量，那么对外部变量有影响的，因为闭包函数共享外部引用
>
***引用闭包的好处是减少全局变量*** 

闭包是附加了数据的行为  闭包的使用，有理解成本在，不够直接。不推荐使用

结构是附加了行为的数据

---
*[👈 0000 golang](0000golang.md)*

[415 出品，必属精品](../note.md) 

tags `function` `函数` `方法`