# numpy

## Numpy是什么？

Numpy是Python科学计算的基础包，是一个开源的扩展库，用来支持大数据量的高维数组和矩阵运算

Numpy提供了矩阵数据类型、矢量处理、高精度运算等，它专为严格的数字处理而产生。

Numpy是一个数学库，主要用于数组计算，包含：
- 一个强大的N维数组对象ndarray；
- 广播功能函数
- 线性代数、傅里叶变换、随机数生成等功能

Numpy提供了两种基本的对象：ndarray（N-dimensional Array Object）和ufunc（Universal Function Object）。ndarray是存储单一数据类型的多维数组，而ufunc是能够对数组进行处理的函数。

## Numpy ndarray数组的创建

在使用numpy之前需要将它导入

---
`import numpy as np`

---

ndarray生成的方式
- 从已有的数据中创建
- 利用random创建
- 创建特定形状的多维数组
- 利用range、linspace函数生成等。

### 从已有数据中创建数组

直接对Python的基础数据类型（如列表、元组等）进行转换来生成ndarray

1）将列表转换成ndarray

```
In [1]: import numpy as np

In [2]: ls1 = [10, 20, 30, 40, 0]

In [3]: nd1 = np.array(ls1)

In [4]: print(nd1)
[10 20 30 40  0]

In [5]: print(type(nd1))
<class 'numpy.ndarray'>
```

2)嵌套列表可以转换成多维ndarray

```
In [1]: import numpy as np

In [2]: ls2 = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]

In [3]: nd2 = np.array(ls2)

In [4]: print(nd2)
[[1 2 3]
 [4 5 6]
 [7 8 9]]

In [5]: print(type(nd2))
<class 'numpy.ndarray'>
```

### 利用random模块生成数组

numpy.random模块常用函数

| 函数 | 描述 |
| --- | --- |
| numpy.random.random | 生成0到1之间的随机数 |
| numpy.random.uniform | 生成均匀分布的随机数 |
| numpy.random.randn | 生成标准正态的随机数 |
| numpy.random.randint | 生成随机的整数 |
| numpy.random.normal | 生成正态分布 |
| numpy.random.shuffle | 随机打乱顺序 |
| numpy.random.seed | 设置随机数种子 |
| random_sample | 生成随机的浮点数 |

```

```

## 创建特定形状的多维数组

参数
