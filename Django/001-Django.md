# Django小知识点

## 1. Windows中搭建开发环境

### 安装Python

安装时可以直接勾选`添加python到环境变量中`，如果漏掉这一步，可以在安装完毕之后自行添加到环境变量中

安装完成并配置好环境变量后，打开命令行，输入`python -V or python --version`，如果打印出python的版本号，说明安装成功

```cmd
C:\Users\dusai> python -V
Python 3.7.0
```

### 配置虚拟环境

虚拟环境(virtualenv，或venv)是Python多版本管理的利器，可以使每个项目环境与其他项目独立开来，保持环境的干净，解决包冲突问题。

新建文件夹`django_project`，并进入文件夹

```cmd
E:\>cd django_project
E:\django_project>
```

输入配置venv的命令，其中`env`为虚拟环境的放置目录：

```cmd
E:\django_project> python -m venv env
```

创建完成后，输入`env\Scripts\activate.bat`，即可进入虚拟环境

```cmd
E:\django_project> env\Scripts\activate.bat
(env) E:\django_project>
```

盘符前面有(env)标识符说明进入venv成功

## 2. Django路由中的path函数

path是Django的路由函数：

- 第一个参数：路由，访问路径
- 第二个参数：视图，处理方法
- namespace：可以确保反查到唯一的url，即使不同的app使用了相同的url

> 如果在path方法中使用了`namespace`参数，在对应`app`的`urls.py`文件中需要添加`app_name='app'`，否则在运行时会报错。

## 3. 模板中静态文件载入

- 使用Django3 ==> {% load static %}
- 使用Django2 ==> {% load staticfiles %}

## 4. markdown

1、安装markdown
> Markdown是一种轻量级的标记语言，它允许人们“使用易读易写的纯文本格式编写文档，然后转换成有效的或者HTML文档。

markdwon的安装：`pip install markdown -i https://pypi.tuna.tsinghua.edu.cn/simple`

`markdown.markdown`语法接收两个参数

- 第一个参数是需要渲染的正文内容
- 第二个参数载入了常用的语法扩展
  - `markdown.extensions.extra`中包括了缩写、表格等扩展
  - `markdown.extensions.codehilite`是代码高亮扩展

Django处于安全的考虑，会将输出的HTML代码进行转义，**这可能使得要渲染的正文内容无法正常显示**。管道符`|`是Django中过滤器的写法，而`|safe`就相当于给要渲染的正文内容贴上了一个不需要转义的标签。

## 5. Django的Form

`Form`实例可以绑定到数据，也可以不绑定数据。

- 如果绑定到数据，就能够验证该数据并将表单呈现为HTML并显示数据
- 如果它未绑定，则无法进行验证，但仍然可以将空白表单呈现为HTML
- `Form`对象的主要任务是验证数据。使用绑定后的`Form`实例，调用`is_valid()`方法验证并返回指定数据是否有效的布尔值。
- [The Forms API](https://docs.djangoproject.com/zh-hans/2.1/ref/forms/api/)

`Form`不仅负责验证数据，还可以"清洗"数据，将其标准化为一致的格式，这个特性使得它允许以各种方式输入特定字段的数据，并且始终产生一直的输出。一旦`Form`使用数据创建了一个实例并对其进行了验证，就可以通过`cleaned_data`属性访问清洗之后的数据。

## 6. CSRF攻击

CSRF攻击可以理解为：攻击者盗用了你的身份，以你的名义发送恶意请求。

- 用户登录了博客网站A，浏览器记录下这次会话，并保持了登录状态；
- 用户在没有退出登录的情况下，又非常不小心的打开了邪恶的攻击网站B
- 攻击网站B在页面中植入了恶意代码，悄无声息的向博客网站A发送了删除文章的请求，此时浏览器误以为是用户在操作，从而顺利的执行了删除操作。

由于浏览器的同源策略，CSRF攻击者并不能得到你的登录数据实际内容，但是可以欺骗浏览器，让恶意请求附上正确的登录数据。不要小看CSRF攻击的威力：倘若是你的银行账户具有此类安全漏洞，黑客就可以神不知鬼不觉转走你的所有存款。

### **CSRF令牌**

Django中提交表单必须加`csrf_token`，这个就是CSRF令牌，它防范CSRF攻击的流程如下：

- 当用户访问 django 站点时，django 反馈给用户的表单中有一个隐含字段 `csrf_token`，这个值是在服务器端随机生成的，每次都不一样；
- 在后端处理 POST 请求前，django 会校验请求的 cookie 里的`csrf_token`和表单里的`csrf_token`是否一致。一致则请求合法，否则这个请求可能是来自于 CSRF攻击，返回 403 服务器禁止访问。

## 7. Django的session

- `authenticate()`方法验证用户名和密码是否匹配，如果是，返回用户数据
- `login()`方法实现用户登录，将用户数据保存在session

### 什么是session

session在网络应用中，称为“会话控制”，它存储特定用户会话所需的属性及配置信息。

当用户在Web页之间跳转时，存储在Session对象中的变量将不会丢失，而是在整个用户会话中一直存在。

Session最常见的用法就是存储用户的登录数据

## 8. Paginator类

可以利用`from django.core.paginator import Paginator`引入`Paginator`类，

## 9.
