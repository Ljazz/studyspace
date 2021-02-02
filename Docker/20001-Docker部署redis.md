# 1. 拉取redis的镜像

```bash
sudo docker pull redis
```

# 2. 不配置数据目录

```bash
sudo docker run -itd --name redis -p 16379:6379 redis --requirepass "psd"
```

# 3. 配置数据目录

```bash
sudo docker run -itd --name redis -p 6379:6379 -v /mnt/docker/redis/data:/data --restart always redis --appendonly yes --requirepass "psd"
```

**参数说明**
- `-d`：以守护进程的方式启动容器
- `-p 6379:6379`：绑定宿主主机端口，16379宿主主机端口，6379容器的端口
- `--name redis`：容器名称
- `--restart always`：开启启动
- `--privileged=true`：提升容器内权限
- `requirepass`：设置登录密码
- `-v /root/docker/redis/data:data`：映射数据目录
- `--appendonly yes`：开启数据持久化
