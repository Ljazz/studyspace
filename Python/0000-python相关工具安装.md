<!-- TOC -->

1. [Linux安装python](#linux安装python)

<!-- /TOC -->

# Linux安装python
1、下载python的源码包

使用`wget`从[Python官方下载页面](https://www.python.org/downloads/source/ )下载Python的源码包

```bash
VERSION=3.8.2
wget https://www.python.org/ftp/python/${VERSION}/Python-${VERSION}.tgz
```

2、解压缩源码包

```bash
tar -zxvf Python-${VERSION}.taz
```

3、进入源码包

```bash
cd Python-${VERSION}
```

4、编译安装

```bash
./configure --prefix=/usr/local/python3.8  # --prefix指定安装目录
make
make install
```

5、验证是否安装成功
```bash
python3.8 --version
```
6、安装pip

```bash
wget https://bootstrap.pypa.io/get-pip.py
python get-pip.py -i https://pypi.tuna.tsinghua.edu.cn/simple/
```

<font color='red'>有可能出现的错误</font>
**问题1**：

错误详细信息：
```text
zipimport.ZipImportError: can't decompress data; zlib not available
Makefile:1079: recipe for target 'install' failed
make: *** [install] Error 1
```
问题解决方案
```bash
# Ubuntu/Debian下需安装的依赖
sudo apt-get install -y make build-essential libssl-dev zlib1g-dev libbz2-dev libreadline-dev libsqlite3-dev wget curl llvm libncurses5-dev libncursesw5-dev xz-utils tk-dev libffi-dev liblzma-dev python-openssl

# Fedora/CentOS/RHEL(aws ec2)下需安装的依赖
sudo yum install zlib-devel bzip2 bzip2-devel readline-devel sqlite sqlite-devel openssl-devel xz xz-devel libffi-devel
```

**问题2**：

pip安装mysqlcient报错

```bash
# Centos 安装mysqlclient
# 先安装mysql-devel，后安装mysqlclient

yum install mysql-devel

pip3 install mysqlclient
```