# Python入门介绍

## 安装Python

### Linux下源码安装Python（Ubuntu为例）

1、下载Python安装包

> `wget https://www.python.org/ftp/python/3.9.0/Python-3.9.0.tgz`

或者

> `wget https://www.python.org/ftp/python/3.9.0/Python-3.9.0.tar.xz`

2、解压压缩包

> tar -zxvf Python-3.9.0.tgz`

3、进入解压缩后的文件夹

> cd Python-3.9.0

4、在`/usr/local`目录下创建一个新的文件夹`python3.9`

> mkdir /usr/local/python3.9

5、编译安装

> ./configure --prefix=/usr/local/python3.9
>
> make && make install

6、建立Python链接

> ln -s /usr/local/python3.9/bin/python /usr/bin/python3.9
>
> ln -s /usr/local/python3.9/bin/pip /usr/bin/pip3.9

7、验证python和pip是否正常可用

```shell
# python3 -V
Python 3.6.1
# pip3 -V
pip 9.0.1 from /usr/local/python3/lib/python3.6/site-packages (python 3.6)
```

## virtualenv虚拟环境

1、安装`virtualenv`第三方包

> `pip install virtualenv`

2、搭建`virtualenv`环境

```shell
$ mkdir ~/work
$ cd ~/work
$ virtualenv venv # virtualenv 环境的
目录
```

3、启动`virtualenv`环境

```shell
$ source venv/bin/active
(venv)$
```

