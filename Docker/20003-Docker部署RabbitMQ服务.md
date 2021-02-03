实验环境
- CentOS8.2
- Docker Version：20.10.3

### 1. 拉取rabbitmq:3.8.9版本镜像

```bash
sudo docker pull rabbitmq:3.8.9
```

### 2. 创建需要映射的目录

```bash
mkdir -p /mnt/docker/rabbitmq/lib /mnt/docker/rabbitmq/log
```

### 3. 运行容器(端口映射、配置用户名密码)

```bash
sudo docker run -itd --name rabbitmq -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=admin -p 15672:15672 -p 5672:5672 rabbitmq:3.8.9
```

### 4. 运行容器(端口映射、持久化路径、配置用户名密码)

```bash
sudo docker run -itd --name rabbitmq -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=admin -p 15672:15672 -p 5672:5672 -v /mnt/docker/rabbitmq/lib:/var/lib/rabbitmq -v /mnt/docker/rabbitmq/log:/var/log/rabbitmq rabbitmq:3.8.9
```

### 5. 开启控制台展示

```bash
# 进入rabbitmq容器内部
sudo docker exec -it rabbitmq bash
# 开启控制台插件
rabbitmq-plugins enable rabbitmq_management
```
