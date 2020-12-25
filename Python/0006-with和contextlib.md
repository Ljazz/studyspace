# 术语

与上下文管理器和with语句有关的概念
- 上下文管理协议（Context Management Protocol） ：包含方法 enter() 和 exit()，支持该协议的对象要实现这两个方法。
- 上下文管理器Context Manager） ：支持上下文管理协议的对象，这种对象实现了 enter() 和 exit() 方法。上下文管理器定义执行 with 语句时要建立的运行时上下文，负责执行 with 语句块上下文中的进入与退出操作。通常使用 with 语句调用上下文管理器，也可以通过直接调用其方法来使用。
- 运行时上下文（runtime context）：由上下文管理器创建，通过上下文管理器的 enter() 和 exit() 方法实现，enter() 方法在语句体执行之前进入运行时上下文，exit() 在语句体执行完后从运行时上下文退出。with 语句支持运行时上下文这一概念。
- 上下文表达式（Context Expression）：with 语句中跟在关键字 with 之后的表达式，该表达式要返回一个上下文管理器对象。
- 语句体（with-body） ：with 语句包裹起来的代码块，在执行语句体之前会调用上下文管理器的 enter() 方法，执行完语句体之后会执行 exit() 方法。

# 基本语法和工作原理

with语句的语法格式如下

```text
with context_expression [as target(s)]:
    with-body
```

这里context_expression要返回一个上下文管理器对象，该对象并不赋值给as子句中的target(s)，如果指定了as子句的话，会将上下文管理器的`__enter__()`方法的返回值赋值给target(s)。

target(s)可以是单个变量，或者由"()"括起来的元组（不能是仅仅由","分隔的变量列表，必须加"()"）。

Python对一些内建对象进行改进，加入了对上下文管理器的支持，可以用于with语句中，比如可以自动关闭文件、线程锁的自动获取和释放等。假设要对一个文件进行操作，使用with语句可以由如下代码

```python
with open('somefileName', 'r') as f:
    for line in f.readlines():
        rint(line)
```

这里使用了 with 语句，不管在处理文件过程中是否发生异常，都能保证 with 语句执行完毕后已经关闭了打开的文件句柄。如果使用传统的 try/finally 范式，则要使用类似如下代码：

```python
somefile = open('somefileName', 'r')
try:
    for line in somefile.readlines():
        print line
        # ...more code
finally:
    somefile.close()
```

# contextlib

要让一个对象用于with语句，就必须实现上下文管理，而实现上下文管理是通过`__enter__`和`__exit__`这两个方法实现的：

```python
class Query(object):
    def __init__(self, name):
        self.name = name

    def __enter__(self):
        print('Begin')
        return self

    def __exit__(self, exc_type, exc_value, traceback):
        if exc_type:
            print('Error')
        else:
            print('End')

    def query(self):
        pass


with Query('Bob') as q:
    q.query()
```

但是这样有些麻烦，Python的内建模块contextlib能让我们的实现变得更简便

```python
from contextlib import contextmanager
 
class Query(object):
 
    def __init__(self, name):
        self.name = name
 
    def query(self):
        print('Query info about %s...' % self.name)
 
@contextmanager
def create_query(name):
    print('Begin')
    q = Query(name)
    yield q
    print('End')
 
with create_query('Bob') as q:
    q.query()
```

contextlib提供的装饰器@contextmanger，这个装饰器接受一个generator，在生成器中用yield语句将想要用于with语句的变量输出出去。

　　加入我们想要在每次运行代码的前后都运行特定的代码，我们也可以选用@contextmanger这个装饰器实现

```python
@contextmanager
def tag(name):
    print("<%s>" % name)
    yield#这个yield调用会执行with语句中的所有语句，因此with语句中的所有内容都将会被运行
    print("</%s>" % name)
 
with tag("h1"):
    print("hello")
    print("world")

"""
#这段代码的执行效果是：
<h1>
hello
world
</h1>
"""
```

# @closing

此外，如果一个对象没有实现运行时上下文，他就不能被应用到with语句当中，我们可以使用contextlib提供的@closing装饰器将其变为上下文对象

```python
from contextlib import closing
from urllib.request import urlopen


with closing(urlopen('https://www.python.org')) as page:
    for line in page:
        print(line)
```

closing其实也是一个经过@contextmanger装饰的genterator

```python
@contextmanager
def closing(thing):
    try:
        yield thing
    finally:
        thing.close()
```

它的作用就是将任意对象变为上下文对象，并支持with语句

　　奇怪的是，并不是所有的对象都能够通过closing（）方法变为上下文对象：（错误是Query对象没有实现close（）方法）

```text
class Query(object):
    def __init__(self,name):
        self.name=name
...
    with closing(Query('bob')) as p:
        print(p.name)
#错误信息：
bob#仍然能输出
Traceback (most recent call last):
  File "c:\Users\Administrator.SC-201605202132\.vscode\extensions\ms-python.python-2019.9.34911\pythonFiles\ptvsd_launcher.py", line 43, in <module>
    main(ptvsdArgs)
  File "c:\Users\Administrator.SC-201605202132\.vscode\extensions\ms-python.python-2019.9.34911\pythonFiles\lib\python\ptvsd\__main__.py", line 432, in main
    run()
  File "c:\Users\Administrator.SC-201605202132\.vscode\extensions\ms-python.python-2019.9.34911\pythonFiles\lib\python\ptvsd\__main__.py", line 316, in run_file
    runpy.run_path(target, run_name='__main__')
  File "c:\users\administrator.sc-201605202132\appdata\local\programs\python\python36\Lib\runpy.py", line 263, in run_path
    pkg_name=pkg_name, script_name=fname)
  File "c:\users\administrator.sc-201605202132\appdata\local\programs\python\python36\Lib\runpy.py", line 96, in _run_module_code
    mod_name, mod_spec, pkg_name, script_name)
  File "c:\users\administrator.sc-201605202132\appdata\local\programs\python\python36\Lib\runpy.py", line 85, in _run_code
    exec(code, run_globals)
  File "c:\Users\Administrator.SC-201605202132\Envs\sort\app\forTest.py", line 26, in <module>
    print(p.name)
  File "c:\users\administrator.sc-201605202132\appdata\local\programs\python\python36\Lib\contextlib.py", line 185, in __exit__
    self.thing.close()
AttributeError: 'Query' object has no attribute 'close'
PS C:\Users\Administrator.SC-201605202132\Envs\sort>
```