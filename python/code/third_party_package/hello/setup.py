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
