<!-- TOC -->

- [依赖管理](#依赖管理)
  - [1.1 为什么需要依赖管理](#11-为什么需要依赖管理)
  - [1.2 godep](#12-godep)
    - [安装](#安装)
    - [基本命令](#基本命令)
    - [使用godep](#使用godep)
    - [vender机制](#vender机制)
    - [godep开发流程](#godep开发流程)
  - [1.3 go module](#13-go-module)
    - [GO111MODULE](#go111module)
    - [go mod命令](#go-mod命令)
    - [go.mod](#gomod)
    - [go get](#go-get)
    - [管理依赖](#管理依赖)
    - [go mod edit](#go-mod-edit)
  - [1.4 项目中使用go mod](#14-项目中使用go-mod)
    - [现有项目](#现有项目)
    - [新项目](#新项目)

<!-- /TOC -->

# 依赖管理

## 1.1 为什么需要依赖管理

早期的版本，Go所依赖的所有的第三方库都放在GOPATH目录下面。造成的问题就是同一个库只能保存一个版本的代码。

## 1.2 godep

Go1.5开始引入`vendor`模式，若项目目录下有vendor目录，go工具链会优先使用`vendor`内的包进行编译、测试等。

`godep`是一个通过vender模式实现的Go第三方依赖管理工具，类似的还有`dep`。

### 安装

安装`godep`工具

> `go get github.com/tools.godep`

### 基本命令

`godep`安装完成之后，可以通过`godep`查看支持的命令。

| 命令 | 含义 |
| --- | --- |
| godep save | 将依赖项输出并复制到Godeps.json文件中 |
| godep go | 使用保存的依赖项运行go工具 |
| godep get | 下载并安装具有指定依赖项的包 |
| godep path | 打印依赖的GOPATH路径 |
| godep restore | 在GOPATH中拉取依赖的版本 |
| godep update | 更新选定的包或go版本 |
| godep diff | 显示当前和以前保存的依赖项集之间的差异 |
| godep version | 查看版本信息 |

### 使用godep

在项目目录下执行`godep save`命令，会在当前项目中创建`Godeps`和`vender`两个文件夹。

其中`Godeps`文件夹下有一个`Godeps.json`文件，里面记录的是项目所依赖的包信息。`vender`文件夹下是项目所依赖的报的源代码文件。

### vender机制

Go1.5开始支持，能够控制Go语言程序编译时依赖包搜索路径的优先级：

> 根项目目录下的`vender` --> `$GOAPTH/src`

### godep开发流程

1. 保证程序能够正常编译
2. 执行`godep save`保存当前项目的所有第三方依赖的版本信息和代码
3. 提交Godeps目录和vender目录到代码库
4. 若要更新依赖的版本，可以直接修改Godeps.json文件中的对应项

## 1.3 go module

`go module`是Go1.11后官方推出的版本管理工具，并且从Go1.13开始，`go module`将是GO默认的依赖管理工具。

### GO111MODULE

启用`go module`支持首先要做的就是设置环境变量`GO111MODULE`，有`off`、`on`、`auto`三个可选项，默认值为`auto`。

1. GO111MODULE=off，禁用模块支持，编译时会从GOPATH和vendor文件夹中查找包。
2. GO111MODULE=on，启用模块支持，编译时会忽略GOPATH和vendor文件夹，只根据 go.mod下载依赖。
3. GO111MODULE=auto，当项目在$GOPATH/src外且项目根目录有go.mod文件时，开启模块支持。

用`go module`管理依赖后会在项目根目录下生成两个文件`go.mod`和`go.sum`。

### go mod命令

```txt
go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit        编辑go.mod文件
go mod graph       打印模块依赖图
go mod init        初始化当前文件夹, 创建go.mod文件
go mod tidy        增加缺少的module，删除无用的module
go mod vendor      将依赖复制到vendor下
go mod verify      校验依赖
go mod why         解释为什么需要依赖
```

### go.mod

go.mod文件记录了项目所有的依赖信息，其结构大致如下：

```txt
module github.com/Q1mi/studygo/blogger

go 1.12

require (
    github.com/DeanThompson/ginpprof v0.0.0-20190408063150-3be636683586
    github.com/gin-gonic/gin v1.4.0
    github.com/go-sql-driver/mysql v1.4.1
    github.com/jmoiron/sqlx v1.2.0
    github.com/satori/go.uuid v1.2.0
    google.golang.org/appengine v1.6.1 // indirect
)
```

其中，

- module用来定义包名
- require用来定义依赖包及版本
- indirect表示间接引用

**依赖的版本**

`go mod`支持语义化版本号，比如`go get foo@v1.2.3`，也可以跟git的分支或tag，比如`go get foo@master`，当然也可以跟git提交哈希，比如`go get foo@e3702bed2`。关于依赖的版本支持以下几种格式：

```txt
gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7
gopkg.in/vmihailenco/msgpack.v2 v2.9.1
gopkg.in/yaml.v2 <=v2.2.1
github.com/tatsushid/go-fastping v0.0.0-20160109021039-d7bb493dee3e
latest
```

**replace**

在国内访问golang.org/x的各个包都需要翻墙，你可以在go.mod中使用replace替换成github上对应的库。

```txt
replace (
    golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
    golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
    golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
```

### go get

在项目中执行`go get`命令可以下载依赖包，并且还可以指定下载的版本。

1. 运行`go get -u`将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
2. 运行`go get -u=patch`将会升级到最新的修订版本
3. 运行`go get package@version`将会升级到指定的版本号version

如果下载所有依赖可以使用`go mod download`命令。

### 管理依赖

我们在代码中删除依赖代码后，相关的依赖库并不会在go.mod文件中自动移除。这种情况下我们可以使用`go mod tidy`命令更新go.mod中的依赖关系。

### go mod edit

**格式化**

> go mod edit -fmt

**添加依赖项**

> go mod edit -require=golang.org/x/text

**移除依赖项**

> go mod edit -droprequire=golang.org/x/text

## 1.4 项目中使用go mod

### 现有项目

1. 项目根目录下执行`go mod init`，生成`go.mod`文件
2. 执行`go get`，查找并记录当前项目的依赖，同时生成一个`go.sum`记录每个依赖库的版本和哈希值

### 新项目

1. 执行`go mod init 项目名`命令，在当前项目文件夹下创建一个`go.mod`文件
2. 手动编辑`go.mod`的require依赖项或执行`go get`自动发现、维护依赖。
