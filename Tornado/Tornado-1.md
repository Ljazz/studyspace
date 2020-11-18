# Torando基础

## 1. Torando框架基础

### 1.1 Tornado简介

Torando框架的特点：

- 非阻塞式服务器：非阻塞式的最大特点式不管程序是否执行成功，都会立即向客户端返回结果，这样做的好处式提高用户体验
- 运行速度快：因为Tornado框架支持多线程操作，所以运行速度和响应速度极快
- 支持并发操作：能够在不影响用户体验得前提下，同时打开上前甚至更多的链接。
- 支持WebSocket链接：可以跟大多数操作系统实现无缝连接

Torando组成部分:

- tornado.web：创建web应用程序的Web框架
- HTTPServer和AsyncHTTPClient：处理HTTP请求，实现HTTP服务器与异步客户端功能
- IOLoop和IOStream：里面保存了实现异步网络功能的类库，这是Torando实现高并发的基础
- tornado.gen：一个基于Generator(生成器)实现的异步开发接口库，可以使用同步方式编写异步处理代码

### 1.2 安装Torando

> pip install tornado

## 2. 示例：hello world
