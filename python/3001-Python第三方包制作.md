# Python第三方包制作

## 文件结构

```tree
|--smart
|  |--static
|  |  |--icon.svg
|  |  |--confg.json
|  |--engine
|  |  |--__init__.py
|  |  |--core.py
|  |--__init__.py
|  |--__version__.py
|  |--api.py
|  |--utils.py
|--tests
|  |--__init__.py
|--LICENSE
|--README.md
|--setup.py
```

文件结构说明：

- smart：项目核心代码模块，提供对外接口的实现。内部可以包含子模块和静态文件。
- tests：包括测试用例。
- LICENSE：版本相关信息。
- README：提供项目基本描述和相关使用方法介绍。
- setup.py：项目安装配置文件。

## setup.py安装配置介绍

```python
from pathlib import Path
from setuptools import setup


BASE_DIR = Path(__file__).parent

# 自动读取version信息
about = {}
filename = Path.joinpath(BASE_DIR, 'smart', '__version__.py')
with open(str(filename), 'r', encoding='utf-8') as f:
    exec(f.read(), about)

# 自动读取readme
with open('README.md', 'r', encoding='utf-8') as f:
    readme = f.read()

setup(
    name=about["__title__"],                    # 包名称
    version=about["__version__"],               # 包版本
    description=about["__description__"],       # 包详细描述
    long_description=readme,                    # 长描述，通常是readme，打包到PiPy需要
    author=about["__author__"],                 # 作者名称
    author_email=about["__author_email__"],     # 作者邮箱
    url=about["__url__"],                       # 项目官网
    packages=[      # 需要导入的包

    ],
    data_files=[    # 需要的静态文件

    ],
    include_package_data=True,                  # 是否需要导入静态数据文件
    python_requires=">=3.0, !=3.0.*, !=3.1.*, !=3.2.*, !=3.3.*",    # Python依赖版本
    install_requires=[      # 第三方库依赖

    ],
    zip_safe=False,     # 此项需要，否则卸载时报Windows error
    classifiers=[       # 程序的所属分类列表
        'Development Status :: 5 - Production/Stable',
        'Intended Audience :: Developers',
        'Natural Language :: English',
        'Programming Language :: Python',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.4',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6',
        'Programming Language :: Python :: 3.7',
        'Programming Language :: Python :: 3.8',
        'Programming Language :: Python :: Implementation :: CPython',
        'Programming Language :: Python :: Implementation :: PyPy'
    ],
)
```

## 源码安装

编写好源码和配置好`setup.py`脚本后，就可以使用`python`将当前模块安装到第三方库中

> python setup.py install

## 打包代码

- 验证`setup.py`文件正确性

> python setup.py check

- 若上述没有出现任何错误或警告，就可以使用下面命令进行打包。

> python setup.py sdist

## 打包whl文件

```txt
# --wheel-dir：为打包存储的路径
# 空格后为需要打包的的工程路径
pip wheel --wheel-dir=D:\\work\\base_package\\dist D:\\work\\base_package
```

使用上述方法就可以完成打包，打包后的`.whl`文件可以通过下面命令进行安装

> pip install xxx.whl

## 上传代码到PyPi

- 注册pypi账号：https://pypi.org/

### 直接上传

使用`register`命令时最简单的上传方式，但是使用HTTP并未加密，有泄露密码的可能。

```cmd
# 注册包
python steup.py register

# 上传包
python setup.pysdist upload
```

### 使用`twine`上传

1. 安装`twine`。`pip install twine`
2. 使用`twine`注册并上传代码。

```cmd
# 注册包
twine register dist/smart.whl
# 上传包
twine upload dist/*
```
