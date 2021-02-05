<!-- TOC -->

- [Ubuntu安装Docker](#ubuntu安装docker)
- [红帽RHEL安装docker容器](#红帽rhel安装docker容器)
- [CentOS8安装docker](#centos8安装docker)

<!-- /TOC -->

# Ubuntu安装Docker
1. 更新现有的包列表

```bash
$ sudo apt update
```

2. 使用`apt`安装一些允许通过HTTPS才能使用的软件包

```bash
$ sudo apt install apt-transport-https ca-certificates curl software-properties-common
```

3. 添加官方Docker存储库的GPG密钥

```bash
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
```

4. 添加Docker存储库到APT源

```bash
$ sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu bionic stable"
```

5. 更新包列表

```bash
$ sudo apt update
```

6. 安装`docker-ce`

```bash
$ sudo apt install docker-ce
```

7. 检查装好的Docker运行状态

```bash
$ sudo systemctl status docker
● docker.service - Docker Application Container Engine
   Loaded: loaded (/lib/systemd/system/docker.service; enabled; vendor preset: enabled)
   Active: active (running) since Thu 2021-01-14 15:47:21 CST; 10min ago
     Docs: https://docs.docker.com
 Main PID: 25313 (dockerd)
    Tasks: 8
   CGroup: /system.slice/docker.service
           └─25313 /usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock
```

# 红帽RHEL安装docker容器

1. 更新现有的yum包

```bash
$ sudo yum update
```

2. 执行docker安装脚本

```bash
$ curl -sSLhttps://get.docker.com/ | sh
```

3. 启动docker服务

```bash
$ sudo service docker start
```

# CentOS8安装docker

1. 下载docker-ce的repo

```bash
curl https://download.docker.com/linux/centos/docker-ce.repo -o /etc/yum.repos.d/docker-ce.repo
```

2. 安装依赖

```bash
yum install https://download.docker.com/linux/Fedora/30/x86_64/stable/Packages/containerd.io-1.2.6-3.3.fc30.x86_64.rpm
```

3. 安装docker-ce

```bash
yum install docker-ce
```

4. 启动docker

```bash
systemctl start docker
systemctl enable docker
```