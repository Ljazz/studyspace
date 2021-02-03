实验环境
- CentOS8.2
- Docker Version：20.10.3

### 1 拉取postgreSQL镜像

```bash
sudo docker pull postgres
```

### 2 创建映射目录

```bash
sudo mkdir /mnt/docker/postgres/data
```

### 3 运行容器

```bash
docker run -it --name postgresql --restart always -e POSTGRES_PASSWORD='123456' -e ALLOW_IP_RANGE=0.0.0.0/0 -v /mnt/docker/postgres/data:/var/lib/postgresql -p 55432:5432 -d postgres
```
说明
- `–name`: 自定义容器名称
- `-e POSTGRES_PASSWORD`：数据库密码
- `-e ALLOW_IP_RANGE=0.0.0.0/0`：这个表示允许所有ip访问，如果不加，则非本机 ip 访问不了
- `-v`:进行映射,本地目录：容器内路径
- `-p`：映射端口,宿主机端口：容器端口

### 4 进入postgres容器

```bash
docker exec -it postgresql bash
```

### 5 切换当前用户，登录数据库

```bash
su postgres
psql -U postgres -W
```