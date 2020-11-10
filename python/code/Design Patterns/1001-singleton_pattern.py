# 单例简单示例
class Singleton1:
    _instance = None

    def __new__(cls, *args, **kwargs):
        if not cls._instance:
            cls._instance = super().__new__(cls)
        return cls._instance

    def getInstance(self):
        print('id ==> {}'.format(id(self)))
        return self._instance


# 单例实现的方式1：重写__new__方法
class Singleton2:
    _instance = None
    _isFirstInit = False

    def __new__(cls, *args, **kwargs):
        if not cls._instance:
            cls._instance = super().__new__(cls)
        return cls._instance

    def __init__(self, name):
        if not self._isFirstInit:
            self._name = name
            print('First created {}'.format(self._name))
            Singleton2._isFirstInit = True
        else:
            print('Not First created {}'.format(name))

    def getInstance(self):
        print('唯一真神: {}'.format(self._name))


# 单例实现的方式2-自定义metaclass
class Singleton3(type):
    """
    单例实现方式 - metaclass
    """
    def __init__(cls, what, bases=None, dict=None):
        super().__init__(what, bases, dict)
        cls._instance = None    # 初始化全局变量cls._instance为None

    def __call__(cls, *args, **kwargs):
        # 控制对象的创建过程，如果cls._instance为None，则创建，否则直接返回
        if cls._instance is None:
            cls._instance = super().__call__(*args, **kwargs)
        return cls._instance


class CustomClass(metaclass=Singleton3):
    """
    用户自定义的类
    """
    def __init__(self, name):
        self.__name = name

    def getName(self):
        return self.__name


def SingletonDecorator(cls, *args, **kwargs):
    """
    定义单例装饰器
    """
    instance = {}

    def wrapperSingleton(*args, **kwargs):
        if cls not in instance:
            instance[cls] = cls(*args, **kwargs)
        return instance[cls]

    return wrapperSingleton


@SingletonDecorator
class Singleton4:
    """
    使用单例装饰器修饰一个类
    """
    def __init__(self, name):
        self.__name = name

    def getName(self):
        return self.__name


if __name__ == '__main__':
    s1 = Singleton1()
    s2 = Singleton1()
    print('*' * 20)
    print(s1.getInstance())
    print(s2.getInstance())
    print('*' * 20)
    print(id(s1.getInstance()), id(s2.getInstance()))
    print(Singleton1.__dict__)
    print('#' * 20)

    s11 = Singleton2('s11')
    s22 = Singleton2('s22')
    s11.getInstance()
    s22.getInstance()
    print(id(s11), id(s22))
    print(id(s11) == id(s22))
    print('#' * 20)

    tony = CustomClass('Tony')
    karry = CustomClass('Karry')
    print(tony.getName(), karry.getName())
    print('id(tony): ', id(tony), 'id(karry): ', id(karry))
    print('tony == karry :', tony == karry)
    print('#' * 20)

    tony = Singleton4('Tony')
    karry = Singleton4('Karry')
    print(tony.getName(), karry.getName())
    print('id(tony): ', id(tony), 'id(karry): ', id(karry))
    print('tony == karry :', tony == karry)
    print('#' * 20)
