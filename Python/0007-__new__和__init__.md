# `__new__`和`__init__`的区别于联系

`__new__`和`__init__`的区别主要表现在；`__new__`负责对象的创建而`__init__`负责对象的初始化。
- `__new__`：创建对象时调用，会返回当前对象的一个实例
- `__init__`：创建完对象后调用，对当前对象一些实例初始化，无返回值

1. 类中，若`__new__`和`__init__`同时存在，hi优先调用`__new__`


```python
class ClsTest(object):
    def __init__(self):
        print('init')
    def __new__(cls, *args, **kwargs):
        print('new')
ClsTest()
```

输出：

> new

2. 若`__new__`返回一个对象的实例，会隐式调用`__init__`

```python
class ClsTest(object):
    def __init__(self):
        print('init')
    def __new__(cls, *args, **kwargs):
        print("new %s"%cls)
        return object.__new__(cls, *args, **kwargs)
ClsTest()
```

输出：

```text
new <class '__main__.ClsTest'
init
```

3. `__new__`方法会返回所构造的对象，`__init__`则不会。`__init__`无返回值。

```python
class ClsTest(object):
    def __init__(self, cls):
        cls.x = 2
        print('init')
        return cls
ClsTest()
```

输出：

```text
init
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
TypeError: __init__() should return None, not 'ClsTest'
```

4. 若`__new__`没有正确返回当前类cls的实例，那`__init__`是不会被调用的，即使是父类的实例也不行

```python
class ClsTest1(object):
    pass

class ClsTest2(ClsTest1):
    def __init__(self):
        print("init")
    def __new__(cls, *args, **kwargs):
        print('new %s'%cls)
        return object.__new__(ClsTest1, *args, **kwargs)
b = ClsTest2()
print(type(b))
```

输出

```text
new <class '__main__.ClsTest2'>
<class '__main__.ClsTest1'>
```

注意：
1. 继承自`object`的新式类才有`__new__`
2. `__new__`至少要有一个参数cls，代表要实例化的类，此参数在实例化时由Python解释器自动提供，`__new__`必须要有返回值，返回实例化出来的实例，可以return父类`__new__`出来的实例，或者直接时object的`__new__`出来的实例。
3. `__init__`有一个参数，就是这个`__new__`返回的实例，`__init__`在`__new__`的基础上可以完成一些其他初始化的动作，`__init__`不需要返回值
4. 如果`__new__`返回一个对象的实例，会隐式调用`__init__`
