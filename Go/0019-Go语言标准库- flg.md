<!-- TOC -->

- [1 os.Args](#1-osargs)
- [2 flag包基本使用](#2-flag包基本使用)
  - [2.1 导入flag包](#21-导入flag包)
  - [2.2 flag参数类型](#22-flag参数类型)
  - [2.3 定义命令行flag参数](#23-定义命令行flag参数)
  - [2.4 flag.Parse()](#24-flagparse)
  - [2.5 flag其他函数](#25-flag其他函数)

<!-- /TOC -->

Go语言内置的`flag`包实现了命令行参数的解析，`flag`包使得开发命令工具更为简单。

# 1 os.Args

如果只是简单的想要获取命令行参数，可以像下面的代码示例使用`os.Args`来获取命令行参数。

```go
package main

import (
    "fmt"
    "os"
)

// os.Args demo
func main() {
    // os.Args是一个[]string
    if len(os.Args) > 0 {
        for index, arg := range os.Args {
            fmt.Println("args[%d]=%v\n", index, arg)
        }
    }
}
```

将上面的代码执行go build -o "args_demo"编译之后，执行：

```
$ ./args_demo a b c d
args[0]=./args_demo
args[1]=a
args[2]=b
args[3]=c
args[4]=d
```

# 2 flag包基本使用

## 2.1 导入flag包

```go
import flag
```

## 2.2 flag参数类型

flag包支持的命令行参数类型有`bool`、`int`、`int64`、`uint`、`uint64`、`float`、`float64`、`string`、`duration`。

| flag参数 | 有效值 |
| --- | --- |
| 字符串flag | 合法字符串 |
| 整数flag | 	1234、0664、0x1234等类型，也可以是负数。 |
| 浮点数flag | 合法浮点数 |
| bool类型flag | 1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False。 |
| 时间段flag | 任何合法的时间段字符串。如”300ms”、”-1.5h”、”2h45m”。
合法的单位有”ns”、”us” /“µs”、”ms”、”s”、”m”、”h”。 |

## 2.3 定义命令行flag参数

<font color='red' size="4px">flag.Type()</font>

`flag.Type(flag名, 默认值, 帮助信息)*Type`例如我们要定义姓名、年龄、婚否三个命令行参数，我们可以定义如下

```go
name := flag.String("name", "颤三", "姓名")
age := flag.Int("age", 18, "年龄")
married := flag.Bool("married", false, "婚否")
delay := flag.Duration("d", 0, "时间间隔")
```
需要注意的是，此时`name`、`age`、`married`、`delay`均为对应类型的指针。

<font color='red' size="4px">flag.TypeVar()</font>

基本格式如下：`flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)`例如我们定义姓名、年龄、婚否三个命令行参数，可以按如下方式定义：

```go
var name string
var age int
var married bool
var delay time.Duration
flag.StringVar(&name, "name", "张三", "姓名")
flag.IntVar(&age, "age", 18, "年龄")
flag.BoolVar(&married, "married", false, "婚否")
flag.DurationVar(&delay, "d", 0, "时间间隔")
```

## 2.4 flag.Parse()

通过上述两种方式定义好命令行flag参数后，需要通过调用`flag.Parse()`来对命令行参数进行解析。

支持的命令行参数格式有以下几种：
- `-flag xxx`（使用空格，一个`-`符号）
- `--flag xxx`（使用空格，两个`-`符号）
- `-flag=xxx`（使用等号，一个`-`符号）
- `--flag=xxx`（使用等号，两个`-`符号）

其中，布尔类型的参数必须使用等号的方式指定

Flag解析在第一个非flag参数（单个`-`不是flag参数）之前停止，或者在终止符“-”之后停止。

## 2.5 flag其他函数

```go
flag.Args()     // 返回命令行参数后的其他参数，以[]string类型
flag.NArg()     // 返回命令行参数后的其他参数个数
flag.NFlag()    // 返回使用的命令行参数格式 
```

<font color="red">示例</font>

```go
func main() {
    // 定义命令行参数方式1
    var name string
    var age int
    var married bool
    var delay time.Duration
    flag.StringVar(&name, "name", "张三", "姓名")
    flag.IntVar(&age, "age", 18, "年龄")
    flag.BoolVar(&married, "married", false, "婚否")
    flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

    // 解析命令行参数
    flag.Parse()
    fmt.Println(name, age, married, delay)
    // 返回命令行参数后的其他参数
    fmt.Println(flag.Args())
    // 返回命令行参数后的其他参数个数
    fmt.Println(flag.NArg())
    // 返回使用的命令行参数个数
    fmt.Println(flag.NFlag())
}
```

<font color='red'>使用</font>

命令行参数使用提示：

```txt
$ ./flag_demo -help
Usage of ./flag_demo:
  -age int
        年龄 (default 18)
  -d duration
        时间间隔
  -married
        婚否
  -name string
        姓名 (default "张三")
```

正常使用命令行flag参数：

```txt
$ ./flag_demo -name 沙河娜扎 --age 28 -married=false -d=1h30m
沙河娜扎 28 false 1h30m0s
[]
0
4
```

使用非flag命令行参数：

```txt
$ ./flag_demo a b c
张三 18 false 0s
[a b c]
3
0
```