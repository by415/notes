### 命令行参数 [code](demo/goflag/goflag.go)

- 通过os获取命令行参数

`os.Args` 为运行时，到main函数已经填充好的数组

- flag 库获取命令行参数

1、使用flag.String(), Bool(), Int()等函数注册flag，下例声明了一个整数flag，解析结果保存在*int指针ip里：

`var ip = flag.Int("flagname", 1234, "help message for flagname")`

2、也可以将flag绑定到一个变量，使用Var系列函数：

```var flagvar int
 func init() {
 	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname")
 }
```

3、你可以自定义一个用于flag的类型（满足Value接口）并将该类型用于flag解析，如下：

```go
flag.Var(&flagVal, "name", "help message for flagname")
```
对这种flag，默认值就是该变量的初始值。
在所有flag都注册之后，调用： `flag.Parse()`来解析命令行参数写入注册的flag里。

解析后，flag后面的参数可以从flag.Args()里获取或用flag.Arg(i)单独获取。这些参数的索引为从0到flag.NArg()-1。

命令行flag语法：
```
-flag
-flag=x
-flag x  // 只有非bool类型的flag可以
```


---
*[👈 0000 golang](0000golang.md)*

[415 出品，必属精品](../note.md)
 
tags `commandline` `命令行`