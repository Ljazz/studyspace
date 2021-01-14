## 什么是模板(Template)？模板是如何工作的

Django的模板是静态的html文件，它只决定了一个页面的样式或外观。它需要视图View传递过来的变量(Variable)或内容对象(Context object)才能被渲染成一个完整的页面。好处是是实现了样式与业务逻辑的分离，便于前后端和后端Web开发人员各自完成自己的开发工作。

```python
# blog/urls.py
from django.urls import path
from . import views

urlpatterns = [
    path('blog/article/<int:id>/', views.article_detail, name="article_detail")
]

# blog/views.py
from django.shortcuts import render, get_object_or_404
from .models import Article

def article_detail(request, id):
    article = get_object_or_404(Article, pk=id)
    return render(request, 'blog/article_detail.html', {"article": article})
```

下面是模板文件'blog/article_detail.html'的代码。在模板文件里我们可以通过双括号`{{ article }}`显示变量或内容对象，还可以通过点号`(.)`用来直接访问变量的属性。

```html
{% block content %}
{{ article.title }}
{{ article.pub_date }}
{{ article.body }}
{% endblock %}
```

## 模板(Template)文件的正确位置

对于模板文件，建议放在<font color='orange'>app/templates/app</font>，这样做的原因是为了安全。在views.py文件里建议使用<font color='orange'>app/template_name.html</font>调用template，这样会防止与其它同名template的冲突。

## 模板过滤器Filter

| 过滤器 | 例子 |
| --- | --- |
| lower,upper | `{{ article.title | lower }}` 大小写 |
| length | `{{ name | length}}` 长度 |
| default | `{{ value | default: "0"}}` 默认值 |
| date | `{{ picture.date | date:"Y-m-j" }}` 日期格式 |
| dicsort | `{{ value | dicsort:"name" }}` 字典排序 |
| escape | `{{ title | escape}}` 转义 |
| filesizeformat | `{{ file | filesizeformat }}` 文件大小 |
| first, last | `{{ list | first }}` 首或尾 |
| floatformat | `{{ value | floatformat }}` 浮点格式 |
| get_digit | `{{ value | get_digit }}` 位数 |
| join | `{{ list | join: "," }}` 字符链接 |
| make_list | `{{ value | make_list }}` 转字符串 |
| pluralize | `{{ number | pluralize }}` 复数 |
| random | `{{ list | random }}` 随机 |
| slice | `{{ list | slice: ":2" }}` 切 |
| slugify | `{{ title | slugify }}` 转为slug |
| striptags | `{{ body | striptags }}` 去除tags |
| time | `{{ value | time:"H:i" }}` 时间格式 |
| timesince | `{{ pub_date | timesince: given_date }}` | 
| truncatechars | `{{ title | truncatechars: 10 }}`  |
| truncatewords | `{{ title | truncatewords: 2 }}`  |
| truncatechars_html | `{{ title | truncatechars_html: 2 }}` |
| urlencode | `{{ path | urlencode }}` URL转义 |
| wordcount | `{{ body | wordcount }}` 单词字数 |

## 模板标签Tags

在Django模板里，变量是放在双括号(`{{ }}`)里的，而代码是放在`{% tag_name %}`标签里的。

<table>
<thead>
<th>标签</th>
<th>例子</th>
</thead>
<tbody>
<tr>
<td>{% block %}</td>
<td>
<pre>
{% block content %}
    代码块
{% endblock %}
</pre>
</td>
</tr>
<tr>
<td>{% csrf_token %}</td>
<td>
<pre>
{% csrf_token %} 表单专用
</pre>
</td>
</tr>
<tr>
<td>{% for %}</td>
<td>
<pre>
&ltul>
{% for athlete in athlete_list %}
    &ltli>{{ athlete.name }}&lt/li>
{% endfor %}
&lt/ul>
</pre>
</td>
</tr>
<tr>
<td>{% for ... empty %}</td>
<td>
<pre>
&ltul>
{% for athlete in athlete_list %}
    &ltli>{{ athlete.name }}&lt/li>
{% empty %}
    &ltli>Sorry, no athlete&lt/li>
{% endfor %}
&lt/ul>
</pre>
</td>
</tr>
<tr>
<td>{% if %}</td>
<td>
<pre>
{% if title != "python" %}
    Not python.
{% endif %}
{% if "hello" in greetings %}
    Say hello
{% endif %}
</pre>
</td>
</tr>
<tr>
<td>{% url %}</td>
<td>
<pre>
{% url "blog:article_detail" article.id %}
</pre>
</td>
</tr>
<tr>
<td>{% now %}</td>
<td>
<pre>
{% now "jS F Y H:i" %}
</pre>
</td>
</tr>
<tr>
<td>{% with %}</td>
<td>
<pre>
{% with total=business.employees.count %}
    {{ total }} employee {{ total | pluralize}}
{% endwith %}
</pre>
</td>
</tr>
<tr>
<td>{% load %}</td>
<td>
<pre>
# load tags and filters
{% load sometags library %}
</pre>
</td>
</tr>
<tr>
<td>{% include %}</td>
<td>
<pre>
{% include "header.html" %}
</pre>
</td>
</tr>
<tr>
<td>{% extends %}</td>
<td>
<pre>
{% extends "base.html" %}
</pre>
</td>
</tr>
</tbody>
</table>

## 模板的继承

Django使用<font color='orange'>extends</font>标签实现模板的继承。

```html
# base.html
{% block sidebar %}
{% endblock %}

{% block content %}
{% endblock %}

{% block footer %}
{% endblock %}

# template.html
{% extends "base.html" %}
{% block content %}
    {{ some code }}
{% endblock %}
```

<font color='orange'>extends</font>标签支持相对路径。如当文件目录结构如下时
```text
dir1/
    template.html
    base2.html
    my/
        base3.html
base.html
```
模板template.html中以下继承都是可以的
```html
{% extends "./base2.html %}
{% extends "../base.html %}
{% extends "./my/base3.html %}
```

## 模板文件加载静态文件

在模板文件中加载静态文件（如css文件和js文件）步骤如下
1. 先在app下新建static文件夹，然后把需要静态文件放进去，推荐路径app/static/app/静态文件
2. 在settings.py增加静态文件设置STATIC_URL。`STATIC_URL = '/static/'`
3. 模板中使用<font color='orange'>{% load static %}</font>

```html
{% load static %}

<!DOCTYPE html>
<html lang="en">
<head>
<title>{% block title %} Django Web Applications {% endblock %} </title>
<link rel="stylesheet" href="{% static 'app/custom.css' %}">
</head>
```

若静态文件不属于任何app，而属于整个项目，则需要在项目目录下的settings.py中设置
```python
STATICFILES_DIRS = [
    BASE_DIR / "static",
]
```
