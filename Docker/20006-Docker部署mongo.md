实验环境
- CentOS8.2
- Docker Version：20.10.3

### 1. 拉取mongo:3.2镜像

```bash
sudo docker pull mongo:3.2
```

### 2. 创建映射目录

```bash
sudo mkdir /mnt/docker/mongo/db
```

### 3. 运行容器

```bash
sudo docker run --name mongo -p 27017:27017 -v /mnt/docker/mongo/db:/data/db -d mongo
```
说明
- `-p 27017:27017`：将容器的27017端口映射到主机的27017端口上（主机端口:容器端口）
- `-v /mnt/docker/mongo/db:/data/db`：将主机中`/mnt/docker/mongo/db`目录挂载到容器的`/data/db`，作为mongo数据库存储目录

### 4. 进入容器

```bash
sudo docker exec -it mongo bash
```

### 5. mongodb的使用

1、 用户创建和数据库创建

使用`mongo`进入mongo

```mongodb
# 创建用户
# 进入 admin 数据库
use admin
# 创建管理员用户
db.createUser(
    {
        user: "admin",
        pwd: "123456",
        roles: [{role: "userAdminAnydatabase", db: "admin"}]
    }
)
# 创建可读写权限的用户，对于只当的study数据
db.createUser(
    {
        user: "study",
        pwd: "study",
        roles: [{role: "read", db: "study"}]
    }
)

# 数据库创建
use study
```

### 6. 开启远程连接

在`mongo`的容器中

```bash
# 更新源
apt update
# 安装vim
apt install vim
# 修改mongo配置文件
vim /etc/mongod.conf.orig
```

将文件`mongod.conf.orig`中`bindIp: 127.0.0.1`注释或者换成`bindIp: 0.0.0.0`