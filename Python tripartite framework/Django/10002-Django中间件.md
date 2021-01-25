## 什么是中间件

<font color='orange'>中间件</font>是一个用来处理Django的请求和相应的框架级别的钩子。它是一个轻量、低级别的插件系统，用于在全局范围内改变Django的输入和输出。每个中间件组件都负责做一些特定的功能。

因为中间件影响的是全局，所以在使用时需要谨慎使用，使用不当会影响性能

中间件是帮助我们在视图函数执行之前和执行之后都可以做一些额外的操作，它本质上就是一个自定义类，类中定义了几个方法，Django框架会在请求的特定的时间去执行这些方法。

Django的`setting.py`中关于中间件的配置
```python
MIDDLEWARE = [
    'django.middleware.security.SecurityMiddleware',
    'django.contrib.sessions.middleware.SessionMiddleware',
    'django.middleware.common.CommonMiddleware',
    'django.middleware.csrf.CsrfViewMiddleware',
    'django.contrib.auth.middleware.AuthenticationMiddleware',
    'django.contrib.messages.middleware.MessageMiddleware',
    'django.middleware.clickjacking.XFrameOptionsMiddleware',
]
```

MIDDLEWARE配置项是一个列表，列表中是一个个字符串，这些字符串其实是一个个类，也就是一个个中间件。

## 自定义中间件

中间可以定义五个方法，分别是：
- process_request(self, request)
- process_view(self, request, view_func, view_args, view_kwargs)
- process_template_response(self, request, response)
- process_exception(slef, request, exception)
- process_response(self, request, response)

以上方法的返回值可以是None或一个HttpResponse对象，如果是None，则继续按照django定义的规则向后继续执行，如果是HttpResponse对象，则直接将该对象返回给用户。

**自定义一个中间件示例**

```python
from django.utils.deprecation import MiddlewareMixin


class MD1(MiddlewareMixin):
    def process_request(self, request):
        print("MD1里面的 process_request")
    
    def process_response(self, request, response):
        print('MD1里面的 process_response')
        return response
```

**process_request**

`process_request`有一个参数，就是request，这里的request和视图函数中的request是一样的。它的返回值可以是None或者HttpResponse对象。返回值是None的话，按正常流程继续走，交给下一个中间件处理，如果是HttpResponse对象，Django将不执行视图函数，而将相应对象返回给浏览器。

```python

```
