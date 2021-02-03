实验环境
- CentOS8.2
- Docker Version：20.10.3

### 1 下载镜像

在终端进行如下命令拉取Jenkins镜像

```bash
sudo docker pull jenkins/jenkins
```

### 2 创建映射目录

```bash
mkdir /mnt/docker/jenkins
```

### 3 运行容器

```bash
docker run --name jenkins -p 8080:8080 -p 50000:50000 -v /mnt/docker/jenkins:/var/jenkins_home -d jenkins/jenkins
```

### 4 查看容器

```bash
[root@iZ2ze9guj1jb9w5pu6hwvfZ jenkins]# docker ps -a
CONTAINER ID   IMAGE                 COMMAND                  CREATED          STATUS                      PORTS                                                                                              NAMES
d3398cfbb475   jenkins/jenkins   "/sbin/tini -- /usr/…"   22 seconds ago   Exited (1) 20 seconds ago                                                                                                      jenkins
```

##### 上述问题解决方案

1、查看原因

```bash
[root@iZ2ze9guj1jb9w5pu6hwvfZ jenkins]# docker logs jenkins 
touch: cannot touch '/var/jenkins_home/copy_reference_file.log': Permission denied
Can not write to /var/jenkins_home/copy_reference_file.log. Wrong volume permissions?
```

2、解决方法

```bash
[root@iZ2ze9guj1jb9w5pu6hwvfZ jenkins]# chown -R 1000:1000 /mnt/docker/jenkins/
```

### 5 重启容器

```bash
docker restart jenkins
```