<!-- TOC -->

- [什么是缓存Cache](#什么是缓存cache)
- [为什么要使用缓存Cache](#为什么要使用缓存cache)
- [缓存Cache的应用场景](#缓存cache的应用场景)
- [Django缓存设置](#django缓存设置)
- [Django代码中使用Cache](#django代码中使用cache)
- [Django中使用Cache的高级技巧](#django中使用cache的高级技巧)

<!-- /TOC -->

## 什么是缓存Cache

<font color='orange'>缓存</font>是一类可以更快的读取数据的介质统称，也指其它可以加快数据读取的存储方式。一般用来存储临时数据，常用介质是读取速度很快的内存。一般来说从数据库多次把所需要的数据提取出来，要比从内存或者硬盘等一次读出来付出的成本更大。

## 为什么要使用缓存Cache

Django中，当用户请求到达视图后，视图会先从数据库提取数据放到模板中进行动态渲染，渲染后的结果就是用户看到的网页。如果用户每次请求都从数据库提取数据并渲染，将极大降低性能，不仅服务器压力大，而且客户端也无法及时获得响应。如果将渲染后的结果放到速度更快的缓存中，每次有请求过来，先检查缓存中是否有对应的资源，如果有，直接从缓存中取出来返回响应，节省取数据和渲染的时间，不仅提高系统性能，还能提高用户体验。

**案例**：每当我们访问博客首页时，下面视图都会从数据库中提取文章列表，并渲染到模板中。但是在大多数情况下，博客更新的不会很频繁，所以文章列表是不变的。所以用户在一定时间内多次访问首页时都从数据库重新读取同样的数据是一种很大的浪费。

```python
from django.shortcuts import render

def index(request):
    # 读取数据库并渲染到网页
    article_list = Article.objects.all()
    return render(request, 'index.html', {'article_list': article_list})
```

使用缓存Cache可以帮助我们解决上述的问题。当用户首次访问博客首页时，我们从数据库中提取文章列表，并将其存储到缓存中。当用户在单位时间内再次访问首页时，Django先检查缓存是否过期，在检查缓存里文章列表资源是否存在，如果存在，直接从缓存读取数据，并渲染模板。

```python
from django.shortcuts import render
from django.views.decorators.cache import cache_page


@cache_page(60 * 15)    # 秒数， 这里是15分钟
def index(request):
    article_list = Article.objects.all()
    return render(request, 'index.html', {'article_list': article_list})
```

## 缓存Cache的应用场景

缓存主要适用于对页面实时性要求不高的页面。存放在缓存的数据，通常是频繁访问的，而不会经常修改的数据。常用的场景
- 博客文章。假设用户一天更新一篇文章，那么可以为博客设置1天的缓存，一天后会刷新。
- 购物网站。商品的描述信息几乎不会变化，而商品的购买数量需要根据用户实时更新。我们可以只选择缓存商品描述信息
- 缓存网页片段。比如网页导航菜单和脚部

## Django缓存设置

Django提供多种缓存方式，若要使用缓存，需要在`settings.py`文件中进行设置，然后应用。根据缓存介质的不同，需要设置不同的缓存后台Backend。

<font color='sky blue'>Memcached缓存</font>

Memcached是基于内存的缓存，Django原生支持的最快最有效的缓存系统。对于大多数场景，推荐使用Memcached，数据缓存在服务器端。使用前需要通过pip安装memcached的插件`python-memcached`和`pylibmc`，可以同时支持多个服务器上面的memcached。

```python
# localhost
CACHES = {
    'default': {
        'BACKEND': 'django.core.cache.backends.memcached.MemcachedCache',
        'LOCATION': '127.0.0.1:11211',
    }
}

# unix socket
CACHES = {
    'default': {
        'BACKEND': 'django.core.cache.backends.memcached.MemcachedCache',
        'LOCATION': 'unix:/tmp/memcached.sock'
    }
}

CACHES = {
    'default': {
        'BACKEND': 'django.core.cache.backends.memcached.MemcachedCache',
        'LOCATION': [
            '172.19.26.240:11211',
            '172.19.26.242:11211',
        ],
        # 也可以给缓存机器加权重，权重高的承担很多的请求
        # 'LOCATION': [
        #     ('172.19.26.240:11211', 5),
        #     ('172.19.26.242:11211', 1),
        # ]
    }
}
```

数据库缓存

```python
caches = {
    'DEFAULT': {
        'BACKEND': 'django.core.cache.backends.db.DatabaseCache',
        'LOCATION': 'my_cache_table',
    }
}
```

文件系统缓存

```python
caches = {
    'DEFAULT': {
        'BACKEND': 'django.core.cache.backends.filebased.FileBasedCache',
        'LOCATION': '/var/tmp/django_cache',    # 文件夹路径
        # 'LOCATION': 'c:\foo\bar',    # windows下文件夹路径
    }
}
```

本地内存缓存

```python
CACHES = {
    'default': {
        'BACKEND': 'django.core.cache.backends.locmem.LocMemCache',
        'LOCATION': 'unique-snowflake',
    }
}
```

## Django代码中使用Cache

当做好Cahce的设置后，在代码中有三种方式使用Cache
- 在视图View中使用
- 在路由URLConf中使用
- 在模板中使用

<font color='sky blue'>在视图View中使用cache</font>

```python
from django.views.decorators.cache import cache_page

@cache_page(60 * 15)
def my_view(request):
    pass
```

<font color='sky blue'>在路由URLConf中使用cache</font>

```python
from django.views.decorators.cache import cache_page

urlpatterns = [
    path('foo/<int:code>/', cache_page(60 * 15)(my_view))
]
```

<font color='sky blue'>在模板中使用cache</font>

```html
{% load cache %}
{% cache 500 sidebar request.user.username %}
    .. sidebar for logged in user ...
{% endcache %}
```

## Django中使用Cache的高级技巧

<font color='sky blue'>使用cache_control</font>

用户通常会面对两种缓存
- Ta自己的浏览器缓存（私有缓存）
- Ta的提供者缓存（公共缓存）

公共缓存有多个用户使用，而受其他人的控制。这就产生了你不想遇到的敏感数据的问题。比u人说你的银行账号被存储在公众缓存中。因此，Web应用程序需要以某种方式告诉缓存那些数据是私有的，那些是公共的。

解决方案是表示出某个页面缓存应当是私有的。Django中是使用cache_control视图装饰器完成此项工作

```python
from django.views.decorators.cache import cache_control

@cache_control(private=True)
def my_view(request):
    pass
```

该装饰器负责在后台发送相应的HTTP头部

还有一些其它犯法可以控制缓存参数。例如HTTP允许应用程序执行以下操作：
- 定义页面可以被缓存的最大时间
- 指定某个缓存是否总是检查较新版本，仅当无更新时才传递所缓存内容

Django中，可以使用cache_control视图装饰器中指定这些缓存参数。如下案例中，cache_control告诉缓存对每次访问都重新验证并在最长3600秒内保持所缓存版本。

```python
from django.views.decorators.cache import cache_control

@cache_control(must_revalidate=True, max_age=3600)
def my_view(request):
    pass
```

在cache_control()中，任何合法的Cache-Control HTTP指令都是有效的。下面时完整列表：
- public=True
- private=True
- no_cache=True
- no_transform=True
- must_revalidate=True
- proxy_revalidate=True
- max_age=num_seconds
- s_maxage=num_seconds

<font color='sky blue'>使用vary_on_headers</font>

缺省情况下，Django的缓存系统使用所请求的路径来创建其缓存键。着意味着不同用户请求同样路径都会得到同样的缓存版本，不考虑用户端user-agent，cookie和语言配置的不同，除非你使用Vary头部通知缓存机制需要考虑请求头里的cookie和语言的不同。

要在Django中完成上述工作，可以使用vary_on_headers视图装饰器，

```python
from django.views.decorators.vary import vary_on_headers

@vary_on_headers('User-Agent', 'Cookie')
def my_view(request):
    pass
```

<font color='sky blue'>使用never_cache禁用缓存</font>

如果i想用头部完全禁掉缓存，可以使用`django.views.decorators.cache.never_cache`装饰器。如果不再使用中使用缓存，服务器端肯定不会缓存的，然后用户的客户端和浏览器还是会缓存一些数据，这时你可以使用never_cache禁用掉客户端的缓存。

```python
from django.views.decorators.cache import never_cache

@never_cache
def my_view(request):
    pass
```