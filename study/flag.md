<!--
 * @Description: 
 * @Autor: 小明～
 * @Date: 2021-09-15 17:15:18
 * @LastEditors: 小明～
 * @LastEditTime: 2021-09-15 17:32:34
-->
## flag包学习

> flag包主要用于解析命令行参数

```
var port = ""
flag.StringVar(&port, "port", "", "启动端口") //使用命令行参数
flag.StringVar(&port, "port", "9000", "启动端口") //使用默认参数
flag.Parse()   //从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行

使用
go run main.go -port 8080
```

#### 根据需要接受类型的不同flag包还提供了其他与StringVar一样功能的函数
```
flag.BoolVar
flag.IntVar
flag.Int64Var
等等...具体见flag包
```
[flag包](https://studygolang.com/pkgdoc)
<!-- 姓名|技能|排行
--|:--|:--
刘备|哭|大哥
关羽|打|二哥
张飞|骂|三弟 -->