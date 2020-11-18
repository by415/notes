# notes
共享的markdown笔记

## 仓库管理规范：

1. 结构
```
notes
├── docker      
├── golang                  //笔记子模块
│   ├── 1.routine.md        //子模块笔记，笔记中通过[demo](demo/test_routine)
│   └── demo                //每一个子模块目录中都有个demo目录
│   │   └── test_routine    //对应某一个笔记的demo
│   │       └── main.go
│   └── images              //笔记引用的静态资源
└── k8s
```

2.提交规范

```
1.每个人首先git clone远程仓库到本地
2.主分支设置为master，不允许直接push，但是允许merge
3.个人不定期的同步master分支的笔记，自己在本地使用t/jilixin，t/ms等以t/开头的临时分支，不定期推送到自己的分支，通过merge的方式放进master
```

3.commit规范

```
[子模块] 简单描述
空行
详细描述

eg：
[docker] docker原理之cgroup

1.介绍了cgroup的工作原理
2.简单介绍了cgroup源码
```
