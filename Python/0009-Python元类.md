# Python元类metaclass

元类(metaclass)的作用是定制类的创建行为。

## 1、python中一切皆对象，包括类

```python
In [1]: class Foo: 
   ...:     def hello(self): 
   ...:         print("hello world!") 
   ...:

In [2]: foo = Foo()

In [3]: print(type(foo))
<class '__main__.Foo'>

In [4]: print(type(foo.hello))
<class 'method'>

In [5]: print(type(Foo))
<class 'type'>

In [6]: temp = Foo

In [7]: Foo.var = 11

In [8]: print(Foo)
<class '__main__.Foo'>
```

## 2、类的创建过程

就上述示例而言，元类的确定过程
- 确定`Foo`的父类是否有参数`metaclass`
- 确定`Foo`的父类的父类是否有参数`metaclass`
- 使用默认元类`type`

上述的示例相当于下列

```python
def hello(self):
   print('hello world')


Foo = type("Foo", (object,), {"hello": hello})
```

## 3、动态创建类

Python中的类可以动态创建，用的就是默认元类type。动态创建类的type函数原型为：

---
type(object_or_name, bases dict)

---

参数说明：
- object_or_name：类名
- bases：基类，多个基类可以使用元组
- dict：类的属性和方法

```python
def init(self, name):
    self.name = name


def hello(self):
    print("hell {}".format(self.name))


Foo = type("Foo", (object,), {"__init__": init, "hello": hello, 'age': 10})
foo = Foo("xiaohu")
foo.hello()       # 打印 hello xiaohu
print(Foo.age)    # 10
```

## 4、自定义元类

在自定义一个元类的时候，需要继承默认元类`type`，并重写其中的`__new__()`方法

```python
class Author(type):
   def __new__(cls, name, bases, dict):
      # 添加属性
      dict["author"] = "xxx"
      return super().__new__(cls, name, bases, dict)
```

对模块中所有函数的继承类参数中添加metaclass参数

```python
class Foo(object, metaclass=Author):
   pass


foo = Foo()
print(foo.author)
```

