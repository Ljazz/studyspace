<!-- TOC -->

- [为什么需要使用cookie和session](#为什么需要使用cookie和session)
- [什么是cookie，cookie的应用场景](#什么是cookiecookie的应用场景)
- [Django中的cookie](#django中的cookie)
- [什么是session及session的工作原理](#什么是session及session的工作原理)
- [Django中的session](#django中的session)

<!-- /TOC -->

## 为什么需要使用cookie和session

HTTP协议本身是“无状态”的，在一次请求和下一次请求之间没有任何状态保持，服务器无法识别来自同一用户的连续请求。有了cookie和session，服务器就可以利用它们记录客户端的访问状态了，这样用户就不用再每次访问不同页面都需要登录了。

## 什么是cookie，cookie的应用场景

cookie是一种数据存储技术，它是将一段文本保存再客户端（浏览器或本地电脑）的一种技术，并且可以长时间的保存。当用户首次通过客户端访问服务器时，web服务器会发送给客户端一小段信息。客户端浏览器会将这段信息以cookie形式保存在本地某个目录下的文件内。当客户端下次再发送请求时会自动将cookie也发送到服务器端，这样服务器通过查验cookie内容就知道该客户端早访问过了。

cookie的常见应用场景
- 判断用户是否已经登录
- 记录用户登录信息（比如用户名、上次登录时间）
- 记录用户搜索关键词

cookie的缺点再于其<font color='orange'>不可靠</font>和<font color='orange'>不安全</font>，原因如下
- 浏览器不一定会保存服务器发来的cookie，用户可以通过设置选择是否保存cookie
- cookie是有声明周期的（通过Expire设置），如果超过周期，cookie就会被清除
- HTTP数据通过明文发送，容易受到攻击，因此不能在cookie中存放敏感信息（比如信用卡号，密码等）
- cookie以文件形式存储在客户端，用户可以随意修改的

## Django中的cookie

**设置cookies(保存数据到客户端)**

<font color='red'>response.set_cookie(key, value, expires)</font>

参数说明
- key：cookie的名称
- value：保存的cookie的值
- expires：保存的时间，以秒为单位

例如：`response.set_cookie('username', 'John', 60*60*24)`

一般在Django的视图中先生成不含cookie的response，然后set_cookie，最后把response返回给客户端（浏览器）

案例：
```python
# 1. 不使用模板
response = HttpResponse("hello world")
response.set_cookie(key, value, expires)
return response

# 2. 使用模板
response = render(request, 'xxx.html', context)
response.set_cookie(key, value, expires)
return response

# 3. 重定向
response = HttpResponseRedirect('/login/')
response.set_cookie(key, value, expires)
return response
```

**获取cookies，获取用户发来请求中的cookies**

```python
request.COOKIES['username']
request.COOKIES.get('username')
```

**检查cookies是否已经存在**
```python
request.COOKIES.has_key('<cookie_name>')
```

**删除cookies**
```python
response.delete_cookie('username')
```

使用cookie验证用户是否已经登录
```python
from django.http import HttpResponseRedirect
from django.shortcuts import render
from django.contrib.auth.models import User
from .forms import LoginForm


# 如果登录成功，设置cookie
def login(request):
    if request.method == 'POST':
        form = LoginForm(request.POST)

        if form.is_valid():
            username = form.cleaned_data['username']
            password = form.cleaned_data['password']

            user = User.objects.filter(username__exact=username, password__exact=password)

            if user:
                response = HttpResponseRedirect('/index/')
                # 将username写入浏览器cookie，失效时间为360秒
                response.set_cookie('username', username, 360)
                return response
            else:
                return HttpResponseRedirect('/login/')
        else:
            form = LoginForm()
        return render(request, 'users/login.html', {'form': form})


# 通过cookie判断用户是否已经登录
def index(request):
    # 提取浏览器中的cookie，如果不为空，表示已登录账号
    username = request.COOKIES.get('username', '')
    if not username:
        return HttpResponseRedirect('/login/')
    return render(request, 'index.html', {'username': username})
```

## 什么是session及session的工作原理

session又名会话，其功能、应用程序和cookie类似，用来存储少量的数据或信息。但由于数据存储在服务器上，而不是客户端上，所以比cookie更安全。

**Session工作流程如下：**
- 客户端向服务器发送请求时，看本地是否有cookie文件。如果有，就在HTTP的请求头(Request Headers)中，包含一行cookie信息
- 服务器接收到请求后，根据cookie信息，得到sessionId，根据sessionId找到对应的session，用这个session就能判断出用户是否登录等等。

使用Session的好处在于，即使用户关闭了浏览器，session仍将保持到会话过期。

## Django中的session

**设置session的值**
```python
request.session['key'] = value
request.session.se_expiry(time) # 设置过期时间，0表示浏览器关闭则失效
```

**获取session的值**
```python
request.session.get('key', None)
```

**删除session的值**
```python
del request.session['key']
```

**判断是否在session中**
```python
'fav_color' in request.session
```

**获取所有session的key和value**
```python
request.session.keys()
request.session.values()
request.session.items()
```

**setting.py有关session的设置**
```text
1. SESSION_COOKIE_AGE = 60 * 30
2. SESSION_EXPIRE_AT_BROWSER_CLOSE = True
```

**示例**：通过使用session来判断用户是否已登录
```python
from django.shortcuts import render
from django.http import HttpResponseRedirect
from django.contrib.auth.models import User
from .forms import LoginForm


# 若登录成功，设置session
def login(request):
    if request.method == 'POST':
        form = LoginForm(request.POST)

        if form.is_valid():
            username = form.cleaned_data['username']
            password = form.cleaned_data['password']

            user = User.objects.filter(username__exact=username, password__exact=password)

            if user:
                # 将username写入session，存入服务器
                request.session['username'] = username
                return HttpResponseRedirect('/index/')
            else:
                return HttpResponseRedirect('/login/')
        else:
            form = LoginForm()
        return render(request, 'users/login.html', {'form': form})


# 通过session判断用户是否登录
def index(request):
    # 获取session中的username
    username = request.session.get('username', '')
    if not username:
        return HttpResponseRedirect('/login/')
    return render(request, 'index.html', {'username': username})

```

**示例2**：通过session控制不让用户连续评论2次

```python
form django.http import HttpResponse

def post_comment(request, new_comment):
    if request.session.get('has_commented', False):
        return HttpResponse("You've already commented.")
    c = comments.Comment(comment=new_comment)
    c.save()
    request.session['has_commented'] = True
    return HttpResponse("Thanks for your comment！")
```