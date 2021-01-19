1、进入数据库

```bash
mysql -uUsername -pPassword
```

2、创建数据库

```bash
create database database_name;
```

3、创建数据库用户

```bash
CREATE USER 'username'@'%' IDENTIFIED BY 'password';
```

4、以root用户登录到数据库进行授权

```bash
GRANT ALL ON database_name.* TO 'username'@'%';
flush privileges; # 刷新权限
```
