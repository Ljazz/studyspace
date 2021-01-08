## 安装uwsgi

> pip install uwsgi

## uwsgi配置文件

```ini
[uwsgi]
# uwsgi使用配置我呢见启动
http = :10001
# 项目目录
chdir = /mnt/project/mysite
# 指定项目的application
module = site.wsgi
# 指定sock文件路径
socket = /mnt/project/mysite/uwsgi.sock
# 启动uwsgi的用户名和用户组
uid = root
gid = root
# 进程数
workers = 4
# 接受的最大请求数
max-requests = 1000
pidfile = /mnt/project/mysite/uwsgi.pid
# 启动主进程
master = true
# 自动移除 unix socket 和 pid 文件，当服务停止时
vacuum = true
# 序列化接受的内容
thunder-lock = true
# 启动线程
enable-threads = true
# 设置中断时间
harakiri = 100
# 设置缓冲
post-buffering = 65535
# 不设置会导致上传大文件失败
buffer-size = 65535
# 设置日志目录
daemonize = /mnt/logs/mysite/uwsgi.log
```

## 启动uwsgi

> uwsgi --ini site_uwsgi.ini

## 安装和设置Nginx

### 安装Nginx

> sudo apt install nginx

## 配置nginx对应的配置文件
```conf
#user  nobody;
#worker_processes  1;

#error_log  logs/error.log;
#error_log  logs/error.log  notice;
#error_log  logs/error.log  info;

#pid        logs/nginx.pid;


events {
    worker_connections  1024;
}


http {
    include       mime.types;
    default_type  application/octet-stream;

    #log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
    #                  '$status $body_bytes_sent "$http_referer" '
    #                  '"$http_user_agent" "$http_x_forwarded_for"';

    #access_log  logs/access.log  main;

    sendfile        on;
    #tcp_nopush     on;

    #keepalive_timeout  0;
    keepalive_timeout  65;
    # 此处值代表nginx 设置 16个 512k 的块进行缓存，总共大小为16*512k 解决问题2 //磁珠值代表每块大小 解决问题2
    proxy_buffers 64 512k;
    proxy_buffer_size 512k;
    #gzip  on;

    server {
        listen       10000;
        #server_name  localhost;

        # 编码设置
        charset     utf-8;
        #charset koi8-r;

        #access_log  logs/host.access.log  main;

        # 静态文件目录

        location /static {
            alias /home/project/mysite/static; # your Django project's static files - amend as required
        }
        location / {
             proxy_pass http://127.0.0.1:10001;
             proxy_set_header Host $host:$server_port;
             proxy_set_header X-Real-IP $remote_addr;
             proxy_set_header X-Real-PORT $remote_port;
             proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;

             #add_header Access-Control-Allow-Origin *;
             #proxy_pass http://127.0.0.1:10001;
             #proxy_set_header Host $host:$10001;
             #proxy_set_header X-Real-IP $remote_addr;
             #proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
             #proxy_set_header X-Forwarded-Host  $host:$10001;
        }
    }
}
```

### 启动命令

> nginx -c /mnt/project/mysite/mysite_nginx.conf