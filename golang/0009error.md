### 错误处理
####错误vs异常
> 错误指的是可能出现问题的地方出现了问题，比如打开一个文件时失败，这种情况在人们的意料之中 ；而异常指的是不应该出现问题的地方出现了问题，比如引用了空指针，这种情况在人们的意料之外。可见，错误是业务过程的一部分，而异常不是。 
  
> Golang中引入error接口类型作为错误处理的标准模式，如果函数要返回错误，则返回值类型列表中肯定包含error。error处理过程类似于C语言中的错误码，可逐层返回，直到被处理。
 
> Golang中引入两个内置函数panic和recover来触发和终止异常处理流程，同时引入关键字defer来延迟执行defer后面的函数
>

> Golang错误和异常是可以互相转换的：
 
 #### 错误转异常
 > 比如程序逻辑上尝试请求某个URL，最多尝试三次，尝试三次的过程中请求失败是错误，尝试完第三次还不成功的话，失败就被提升为异常了。
 
 #### 异常转错误
>  异常转错误，比如panic触发的异常被recover恢复后，将返回值中error类型的变量进行赋值，以便上层函数继续走错误处理流程。
 
#### 应当使用异常的场景

- 空指针引用
- 下标越界
- 除数为0
- 不应该出现的分支，如switch中，分支已经确定，default应当pannic
- 输入不应该引起函数错误

对于异常，我们应该在合适的上层进行recover,并打印响应错误；使得部署后的程序不会终止

错误处理的正确姿势

- 失败的原因只有一个时，不使用error 使用bool
- 没有失败时，比如setValue，不使用error
- error应当放在函数最后一个值，bool值也是， 遵循 command,ok/err 原则
- 错误值统一定义，而非随意插入
- 错误逐层传递时，层层都加日志
```go

    /// 多一条日志，附带有每个错误的详细位置。如果日志显示不合理，日志检索，定位极不方便
    _,err = func(){
        log.Error("this is a error")
        return errors.New("this is a error")
    }()
    if err != nil {
        log.err("func return error",err)
    }

    
    /// 少一次log打印，同时如果还有更多上层时，一条日志就是一个完整的错误链条
    _,err = func(){
        return errors.New("this is a error")
    }()
    if err != nil {
        log.err("func return error",err)
    }   

    
```
- 在错误时，使用defer进行资源释放
```go
func deferDemo() error {
    err := createResource1()
    if err != nil {
        return ERR_CREATE_RESOURCE1_FAILED
    }
    defer func() {
        if err != nil {
            destroyResource1()
        }
    }()
    err = createResource2()
    if err != nil {
        return ERR_CREATE_RESOURCE2_FAILED
    }
    defer func() {
        if err != nil {
            destroyResource2()
        }
    }()

    err = createResource3()
    if err != nil {
        return ERR_CREATE_RESOURCE3_FAILED
    }
    defer func() {
        if err != nil {
            destroyResource3()
        }
    }()

    err = createResource4()
    if err != nil {
        return ERR_CREATE_RESOURCE4_FAILED
    }
    return nil
}
```

- 尝试几次可以避免失败，不要立即返回。如url请求

- 当上层不关心错误时，不建议返回error;如关闭操作，资源释放

- 发生错误时，不要忽略有价值的返回值；如 文件读写，Read 会返回错误信息及读到的字节n

***对于返回值应当有清晰的说明，以便他人使用***

#### 异常处理的正确姿势

- 在程序开发阶段，坚持速错

使用panic放大错误，及时处理，比如ide自动生成方法，会内置一个panic提示必须实现

- 在程序部署后，应当恢复异常，避免程序终止

1.在recover中打印详细堆栈信息，关键业务信息

2.将异常转换为错误，以便调用者让程序恢复到健康状态继续运行

3.对于不该出现的分支，使用 panic

4.对输入不应该有问题的函数，使用panic设计而非error

##
参考 [Golang错误和异常处理的正确姿势](https://www.jianshu.com/p/f30da01eea97)