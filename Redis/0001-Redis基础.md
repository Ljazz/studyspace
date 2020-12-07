# 1. 初识NoSQL


## 1.1 什么是NoSQL

NoSQL是Not Only SQL的缩写，也是众多非关系型数据库的统称。NoSQL和关系型数据库一样，也是用来存储数据的仓库。

1. NoSQL解决了高并发读/写问题
2. NoSQL解决了海量数据的高效存储和访问问题
3. NoSQL实现了高可用性及高扩展性

NoSQL具有如下特点：
- 容易扩展，方便使用，数据之间没有关系
- 数据模型非常灵活，无须提前为要存储的数据建立字段类型，随时可以存储自定义的数据格式。
- 适合大量数据、高性能的存储
- 具有高并发读/写、高可用性

## 1.2 NoSQL与传统关系型数据库的比较

| | NoSQL | 关系型数据库 |
| --- | --- | --- |
| 使用成本 | 使用简单，易搭建，大部分是开源软件，比较廉价 | 通常需要安装部署，开源比较少，使用成本高。
| 存储形式 | 具有丰富的存储形式，如key-value(键值对)形式、图形结构形式、文档形式、列簇形式等，可以存储各种类型的数据 | 采用关系型数据模型来组织的，是行列表结构，通过行与列的二元形式表示出来，数据之间具有很强的关联性。采用二维表结构的形式对数据进行持久存储。 |
| 查询速度 | 将数据存储在系统缓存中，不需要经过SQL层的解析，查询效率很高 | 将数据存储在系统的硬盘中，在查询时需经过SQL层的解析，然后读入内存，实现查询，因此查询效率较低 |
| 扩展性 | 去掉了传统关系型数据表与字段之间的关系，实现了真正意义上的扩展。采用键值对的形式存储数据，消除了数据之间的耦合性 | 由于采用关系型数据模型来存储数据，数据与数据之间的关联性较强，存在耦合性，不易扩展 |
| 是否支持ACID特性 | 一般不支持ACID特性，它实现了最终一致性 | 支持ACID特性，具有严格的数据一致性 |
| 是否支持SQL语句 | 不支持 | 支持 |

**注意**
1. ACID特性是指数据库事务的执行要素，包括原子性、一致性、隔离性、持久性。
2. SQL是结构化查询语言、数据操纵语言、数据定义语言。

## 1.3 NoSQL使用场景

1. 对于大数据量、高并发的存储系统及相关应用
2. 对于一些数据模型比较简单的相关应用
3. 对数据一致性要求不是很高的业务场景
4. 对于给定key来映射一些复杂值得环境
5. 存储用户信息，如大型电商系统得购物车、会话等。
6. 对于多数据源的数据存储。
7. 对易变化、热点高频信息、关键字等信息的存储。

## 1.4 NoSQL的数据模型

关系型数据库的数据模型由数据结构、数据操作及完整性约束条件组成。

NoSQL的4种数据模型如下
- 键值对数据模型
  - 键值对数据模型就是采用键值对形式将数据存储在一张哈希表种的一类数据库，这张哈希表具有一个特定的键和一个指向特定数据的指针。键值对存储中的值可以是任意类型的值，如数字、字符串，也可以是封装在对象中的新的键值对。
- 列数据模型
  - 列数据模型就是将数据按照列簇形式来存储的一类数据库，通常用于存储分布式系统的海量数据。它也有键，这些键指向多个列，由数据库的列簇来统一安排。
- 文档数据模型
  - 文档数据模型以文档形式进行存储，它是键值对数据模型的升级版，是版本化的文档。它可以使用模式来指定某个文档结构，通常采用特定格式来存储半结构化的文档，最常使用的存储格式是XML、JSON。每个文档都是自包含的数据单元，是一系列数据项的集合。
- 图数据模型
  - 图数据模型采用图结构形式存储数据，它是最复杂的NoSQL，常被用于存储一些社交网络的社交关系，适用于存储高度互联的数据。它是由多个节点和多条边组成，节点表示实体，边表示两个实体之间的关系。

键值对数据模型、列数据模型、文档数据模型统称为聚合模型，它们有一个共同特点：可以把一组相互关联的对象看作一个整体单元来操作，通常把这个单元称为一个聚合。

## 1.5 NoSQL数据库的分类

### 1、NoSQL数据库大致可分为四大类

1、键值对存储数据库

主要采用键值对形式存储数据的一类数据库。

典型代表：Redis(由C/C++语言开发)、Memcached、Voldemort、Berkeley DB、Tokyo Cabinet/Tyrant等。当采用该类数据库存储数据时，需要定义数据结构（半结构化）才能进行存储。

2、面向列存储数据库

主要按照列存储数据的一类数据库

典型代表：HBase（Java开发）、Cassandra（Java开发）、Riak（Erlang、C及JavaScript组合开发）等。当采用该类数据库存储数据时，需要定义数据结构（半结构化）才能进行存储。

3、面向文档数据库

主要用于存储文档的一类数据库。文档也是它最小单元，同一张表中存储的文档属性可以是多样化的，数据可以采用XML、JSON、JSONB等多种格式存储。

典型代表：MongoDB（由C++开发）、CouchDB（Erlang开发）、RavenDB等。当采用该类数据库存储数据时，不需要定义数据结构（非结构化）就可以存储。

4、面向图形数据库

主要用于存储图片信息的一类数据库

典型代表：Neo4j（Java开发）、infoGrid、infinite Graph等。

### 2、各类NoSQL数据库的比较

![各类NoSQL数据库比较](./images/各类NoSQL数据库比较.jpg)


# 2. 认识Redis

## 2.1 Redis简介

### 1、什么是Redis

Redis是用C语言开发的一款开源的、高性能的键值对存储数据库，采用了BSD协议，为了适应不同场景下的存储需求，提供了多种键值数据类型。

Redis支持的键值数据类型有字符串、列表、有序集合、散列和集合等。

### 2、Redis特性

1. 支持多种计算机编程语言
2. 具有丰富的数据类型，如String、List、Set、Hash、Sorted Set等
3. 支持多种数据结构，如哈希、集合、位图(多用于活跃用户数等的统计)、HyperLogLog（超小内存唯一值计数，由于只有12KB，因而是有一定误差范围的）、GEO（地理信息定位）。
4. 读/写速度快，性能高。
5. 支持持久化。Redis的持久化也就是备份数据，每隔一段时间将内存种的数据保存在磁盘中，在重启的时候会再次加载到内存中，从而实现数据持久化。Redis的持久化方式是RDB和AOF。
6. 简单且功能强大。
7. 实现高可用主从复制，主节点做数据副本。
8. 实现分布式集群和高可用。Redis Cluster支持分布式，进而可以实现分布式集群；Redis Sentinel支持高可用。

### 3、Redis使用场景

1. 做缓存。
2. 做计数器应用。
3. 实现消息队列系统。
4. 做实时系统、消息系统
5. 实现排行榜应用
6. 做数据过期处理
7. 做大型社交网络。
8. 分布式集群架构中的session分离。

## 2.2 搭建Redis环境

Windows环境下载地址：https://github.com/MicrosoftArchive/redis/releases

Linux环境下载地址：http://www.redis.net.cn/download/

### 1、Windows环境下搭建

1、下载Redis安装包

![下载页面](./images/redis下载页面.png)

2、解压压缩包

![解压缩目录](./images/Redis解压缩目录.png)

**注意**：也可以下载`.msi`文件，直接双击打开进行安装。

文件介绍：
- redis-server.exe：服务端程序，提供 redis 服务
- redis-cli.exe: 客户端程序，通过它连接 redis 服务并进行操作
- redis-check-dump.exe：RDB 文件修复工具
- redis-check-aof.exe：AOF 文件修复工具
- redis-benchmark.exe：性能测试工具，用以模拟同时由 N 个客户端发送 M 个 SETs/GETs 查询 (类似于 Apache 的 ab 工具)
- redis.windows.conf： 配置文件，将 redis 作为普通软件使用的配置，命令行关闭则 redis 关闭
- redis.windows-service.conf：配置文件，将 redis 作为系统服务的配置

3、运行cmd，cd进入解压目录，执行

> redis-server.exe redis.windows.conf

![运行](./images/redis命令行运行.png)

4、安装redis到Windows服务

> redis-server --service-install redis.windows.conf

- 启动服务：`redis-server --service-start`
- 停止服务：`redis-server --service-stop`

### 2、Linux环境下搭建

```
root@iZpr21qe3dl1aaZ:~# cd /usr/local     # 进入/usr/localmul
root@iZpr21qe3dl1aaZ:~# mkdir redis       # 创建redis目录
root@iZpr21qe3dl1aaZ:~# cd redis          # 进入redis目录
root@iZpr21qe3dl1aaZ:~# wget http://download.redis.io/releases/redis-4.0.9.tar.gz
root@iZpr21qe3dl1aaZ:~# tar -zxvf redis-4.0.9.tar.gz    # 解压压缩包
root@iZpr21qe3dl1aaZ:~# cd redis-4.0.9      # 进入目录
root@iZpr21qe3dl1aaZ:~# make                # 进行编译
root@iZpr21qe3dl1aaZ:~# make install        # 进行安装
```

## 2.3 Redis客户端

### 1、命令行客户端

Redis的命令行客户端`redis-cli`(Redis Command Line Interface)是Redis自带的基于命令行的客户端，主要用于与服务器端进行交互。

### 2、可视化客户端

Redis可视化客户端也称远程客户端，可以连接远程Redis数据库进行操作。

两款可视化工具：
- [Redis Desktop Manager(RDM)](https://redisdesktop.com/download)
- [TreeSoft数据库管理系统TreeDMS](http://www.treesoft.cn/dms.html)


# 3. Redis数据类型

> 目前redis数据库支持5种数据类型，分别是String(字符串)、Hash(哈希)、List(列表)、Set(集合)及Sorted Set(有序集合)

## 3.1 字符串（String）命令

字符串类型是Redis种最基本的数据类型，它是二进制安全的，任何形式的字符串都可以存储，包括二进制数据、序列化后的数据、JSON化的对象，甚至是一张经Base64编码后的图片。String类型的键最大能存储512MB的数据。

Redis字符串类型的相关命令用于管理Redis的字符串。

### 3.1.1、设置键值对

#### 1、SET命令：设置键值对

**命令格式**：`SET key value [EX seconds [PX milliseconds] [NX|XX]`

SET命令将字符串值value设置到key种。若key存在，将会覆盖原来的旧值，并且是忽略类型的。针对某个带有时间的key来说，当SET命令执行成功时，这个key上的生存时间会被清除。

SET命令的可选参数
- EX seconds：用于设置key的过期时间为多少秒(seconds)。其中，`SET key value EX seconds`等价于`SETEX key seconds value`。
- PX milliseconds：用于设置key的过期时间为多少毫秒(milliseconds)。其中，`SET key value PX milliseconds`相当于`PSETEX key milliseconds value`。
- NX：表示当key不存在时，才对key进行设置操作。其中，`SET key value NX`等价于`SETNX key value`.
- XX：表示当key存在时，才对key进行设置操作。

**返回值**：如果SET命令设置成功，则会返回OK。如果设置了NX或XX，但因为条件不足而设置失败，则会返回空批量回复(NULL Bulk Reply)。

```
127.0.0.1:0>SET stuName-1 '刘河飞'
"OK"
127.0.0.1:0>SET stuID-1 20180001
"OK"
127.0.0.1:0>SET age-1 22
"OK"
127.0.0.1:0>SET sex-1 '男'
"OK"
127.0.0.1:0>SET height-1 171
"OK"
127.0.0.1:0>SET weight-1 75
"OK"
127.0.0.1:0>SET className-1 '软件工程1班'
```

#### 2、MSET命令：设置多个键值对

**命令格式**：`MSET key value [key value...]`

使用MSET命令同时设置多个键值对。若某个key已经存在，那么MSET命令会用新值覆盖旧值。MSET命令是一个原子性操作，所有给定key都会在同一时间内被设置更新，不存在在某些key被更新了而另一些key没有被更新的情况。

**返回值**：总是返回OK，因为MSET命令不可能设置失败。

```
127.0.0.1:0>MSET stuName-2 '赵雨梦' stuID-2 20181762 age-2 24 sex-2 '女' height-2 175 weight-2 73 birthday-2 1994-04-23 className-2 '网络工程1班'
"OK"
127.0.0.1:0>MSET stuName-3 '宋飞' stuID-3 20180023 age-3 23 sex-3 '男' height-3 168 weight-3 67 birthday-3 1995-08-18 className-3 '网络工程1班'
"OK"
```

#### 3、SETNX命令：设置不存在的键值对

**命令格式**：`SETNX key value`

SETNX是`set if not exists`的缩写。如果key不存在，则设置值，当且仅当key不存在时。如果key已经存在，则SETNX什么也不做。

**返回值**：SETNX命令设置成功过返回1，设置失败返回0

```
127.0.0.1:0>SETNX collegeName '计算机学院'
"1"
127.0.0.1:0>SETNX collegeName '计算机工程学院'
"0"
```

#### 4、MSETNX命令：设置多个不存在的键值对

**命令格式**：`MSETNX key value [key value...]`

使用MSETNX命令同时设置多个键值对，当且仅当所有给定key都不存在时设置。若有一个给定的key已经存在，那么MSETNX命令也会拒绝执行所有给定key的设置操作。MSETNX命令是原子性的，它可以用来设置多个不同key表示不同字段的唯一性逻辑对象，所有字段要么全部被设置，要么全部设置失败。

**返回值**：所有key设置成功返回1；如果所有给定key都设置失败，返回0

```
127.0.0.1:0>MSETNX Chinese-teacher '郭涛' Math-teacher '杨艳' English-teacher '吴芳'
"1"
127.0.0.1:0>MSETNX Chinese-teacher '陈城' Math-teacher '杨小艳'
"0"
```

### 3.1.2、获取键值对

#### 1、GET命令：获取键值对的值

**命令格式**：`GET key`

使用GET命令获取key中设置的字符串值。如果key中存储的值不是字符串类型的，则会返回一个错误，因为GET命令只能用于处理字符串的值；当key不存在时，返回null。

**返回值**：当key存在时，返回key所对应的值；如果key不存在，返回null；如果key不是字符串类型的，返回错误。

```
127.0.0.1:0>GET stuID-1
"20180001"
127.0.0.1:0>GET stuName-1
"刘河飞"
127.0.0.1:0>GET age-1
"22"
127.0.0.1:0>GET sex-1
"男"
127.0.0.1:0>GET height-1
"171"
127.0.0.1:0>GET weight-1
"75"
127.0.0.1:0>GET birthday-1
null
127.0.0.1:0>GET className-1
"软件工程1班"
```

#### 2、MGET命令：获取多个键值对的值

**命令格式**：`MGET key [key ...]`

使用MGET命令同时返回多个给定key的值，key之间使用空格隔开。如果在给定的key中有不存在的key，那么这个key返回的值为null。

**返回值**：一个包含所有给定key的值的列表

```
127.0.0.1:0>MGET stuName-1 stuID-1 age-1 sex-1 height-1 weight-1 birthday-1 className-1
 1)  "刘河飞"
 2)  "20180001"
 3)  "22"
 4)  "男"
 5)  "171"
 6)  "75"
 7)  null
 8)  "软件工程1班"
```

#### 3、GETRANGE命令：获取键的子字符串值

**命令格式**：`GETRANGE key start end`

使用GETRANGE命令来获取key中字符串值从start开始到end结束的子字符串，小标从0开始（字符串截取）。start和end参数是整数，可以取负值。当取负值时，表示从字符串最后开始计数，-1表示最后一个字符，-2表示倒数第二个字符，依此类推。

**返回值**：返回截取的子字符串。

```
127.0.0.1:0>SET motto "this a content that is test what is a 'getrange' command"
"OK"
127.0.0.1:0>GETRANGE motto 0 100
"this a content that is test what is a 'getrange' command"
127.0.0.1:0>GETRANGE motto -8 -1
" command"
127.0.0.1:0>GETRANGE motto 0 -3
"this a content that is test what is a 'getrange' comma"
```

### 3.1.3 键值对的偏移量

#### 1、SETBIT命令：设置键的偏移量

**命令格式**：`SETBIT key offset value`

使用SETBIT命令对可以所存储的字符串值设置或清除指定偏移量上的位(bit)。vlaue参数值决定了位的设置或清除，value值取0或1。当key不存在时，自动生成一个新的字符串值。这个字符串时动态的，可以扩展，以确保将value保存到指定的偏移量上。当这个字符串扩展时，使用0来填充空白位置。offset参数必须时大于或等于0，并且小于2^32(bit映射被限制在512MB之内)的正整数。默认情况下，bit初始化为0.

**返回值**：返回指定偏移量原来存储的位。

```
127.0.0.1:0>SETBIT stuName-1 6 1
"0"
127.0.0.1:0>SETBIT stuName-1 7 0
"1"
127.0.0.1:0>SETBIT collegeName 100 0
"1"
```


### 2、GETBIT命令：获取键的偏移量值

**命令格式**：`GETBIT key offset`

对可以所存储的字符串值，使用GETBIT命令来获取指定偏移量上的位(bit)。当offset的值超过了字符串的最大长度，或者key不存在时，返回0。

**返回值**：返回字符串值指定偏移量上的位(bit)。

```
47.93.11.106:0>GETBIT stuName-1 6
"1"
47.93.11.106:0>GETBIT suName-1 7
"0"
47.93.11.106:0>GETBIT collegeName 100
"0"
```

### 3.1.4 设置键的生存时间

#### 1、SETEX命令：为键设置生存时间（秒）

**命令格式**：`SETEX key seconds value`

使用SETEX命令将value值设置到key中，并设置key的生存时间为多少秒(seconds)。如果key已经存在，则SETEX命令将覆写旧值。

等价于`SET key value EXPIRE key seconds`

SETEX命令是一个原子性命令，它设置value与设置生成时间是在同一时间完成的。

**返回值**：设置成功时，返回OK；当seconds参数不合法时，返回错误。

```
47.93.11.106:0>SETEX schoolName 100 '清华大学'
"OK"
47.93.11.106:0>GET schoolName
"清华大学"
47.93.11.106:0>TTL schoolName
"79"
47.93.11.106:0>GET schoolName
"清华大学"
47.93.11.106:0>TTL schoolName
"63"
47.93.11.106:0>TTL schoolName
"12"
47.93.11.106:0>TTL schoolName
"-2"
47.93.11.106:0>GET schoolName
null
```

#### 2、PSETEX命令：为键设置生存时间（毫秒）

**命令格式**：`PSETEX key milliseconds value`

使用PSETEX命令设置key的生存时间，以毫秒为单位。设置成功时返回OK。

```
47.93.11.106:0>PSETEX school_address 30000 '北京'  # 设置学校地址为北京，生存时间为30000毫秒
"OK"
47.93.11.106:0>GET school_address   # 获取学校地址
"北京"
47.93.11.106:0>PTTL school_address  # 查看学校地址剩余多少生存时间（毫秒）
"7169"
47.93.11.106:0>GET school_address
null
```

### 3.1.5 键值对的值操作

#### 1、SETRANGE命令：替换键的值

**命令格式**：`SETRANGE key offset value`

使用SETRANGE命令从指定的位置(offset)开始将key的值替换为新的字符串。若key不存在，就当空白字符串处理。若给定key原始存储的字符串长度比偏移量小，那么源字符串和偏移量之间的空白用零字节（Zerobytes，"\x00"）来填充。

**返回值**：返回执行SETRANGE命令之后的字符串长度

```
47.93.11.106:0>SET mykey "hello world"
"OK"
47.93.11.106:0>GET mykey
"hello world" 
47.93.11.106:0>SETRANGE mykey 5 Redis
"11"
47.93.11.106:0>GET mykey
"helloRedisd"
```

#### 2、GETSET命令：为键设置新值

**命令格式**：`GETSET key value`

使用GETSET命令将给定key的值设置为value，并返回key的旧值。当key存在但不是字符串类型时，将会返回错误。

**返回值**：返回给定key的旧值。如果key不存在，则返回null；如果key存在但不是字符串类型的，则返回错误。

```
47.93.11.106:0>EXISTS motto
"0"
47.93.11.106:0>GETSET motto "没有存款，就是拼的理由"
null
47.93.11.106:0>GET motto
"没有存款，就是拼的理由"
47.93.11.106:0>GETSET motto "拼个春夏秋冬，赢个无悔人生"
"没有存款，就是拼的理由"
47.93.11.106:0>GET motto
"拼个春夏秋冬，赢个无悔人生"
```

#### APPEND命令：为键追加值

**命令格式**：`APPEND key value`

如果key存在且是字符串类型，则将value追加到key旧值的末尾。如果key不存在，将key设置值为value。

**返回值**：返回追加value之后，key中字符串的长度。

```
47.93.11.106:0>GET motto
"拼个春夏秋冬，赢个无悔人生"
47.93.11.106:0>APPEND motto "，努力过，拼搏过..."
"66"
47.93.11.106:0>GET motto
"拼个春夏秋冬，赢个无悔人生，努力过，拼搏过..."
```

### 3.1.6 键值对的计算

#### 1、BITCOUNT命令：计算比特位数量

**命令格式**：`BITCOUNT key [start] [end]`

使用BITCOUNT命令计算在给定的字符串中被设置为1的比特位数量。有两个参数：start和end。如果不设置这两个参数，则表示对整个字符串进行计数；如果指定了这两个参数，是在这个范围内计数。若key不存在，会被当作空字符串处理，计数结果为0。

**返回值**：执行BITCOUNT命令之后，返回被设置为1的位的数量。

```
47.93.11.106:0>SET stuName '赵云'
"OK"
47.93.11.106:0>BITCOUNT stuName
"26"
47.93.11.106:0>BITCOUNT stuName 0 10
"26"
47.93.11.106:0>BITCOUNT stuName 0 13
"26"
47.93.11.106:0>BITCOUNT stuName 0 3
"18"
```

#### 2、BITOP命令：对键进行位元运算

**命令格式**：`BITOP operation destkey key [key...]`

使用BITOP命令对一个或多个保存二进制位的字符串key进行位元运算，并将运算结果保存到destkey中。operation表示位元运算操作符，可以是AND、OR、NOT、XOR4种操作种的任意一种。

- `BITOP AND destkey key [key...]`：表示对一个或多个key求逻辑并，并将结果保存到destkey中。
- `BITOP OR destkey key [key...]`：表示对一个或多个key求逻辑或，并将结果保存到destkey中。
- `BITOP NOT destkey key`：表示对给定的key求逻辑非，并将结果保存到destkey中。
- `BITOP XOR destkey key [key...]`：表示对一个或多个key求逻辑异或，并将结果保存到destkey中。

除了NOT操作外，其余的操作都是可以接收一个或多个key作为参数。当使用BITOP命令来进行不同长度的字符串的位元运算时，较短的字符串所缺少的部分会被看作0。空key也被看作包含0的字符串序列。

**返回值**：返回保存到destkey中的字符串的长度，这个长度和输入key中最长的字符串的长度相等。

```
47.93.11.106:0>SETBIT age 0 1
"0"
47.93.11.106:0>SETBIT age1 3 1
"0"
47.93.11.106:0>BITOP AND result age age1
"1"
47.93.11.106:0>GET result
"\x00"
47.93.11.106:0>BITOP OR result age age1
"1"
47.93.11.106:0>GET result
"\x90"
47.93.11.106:0>BITOP NOT result age age1
"ERR BITOP NOT must be called with a single source key."
47.93.11.106:0>BITOP NOT result age
"1"
47.93.11.106:0>GET result
"\x7F"
47.93.11.106:0>BITOP XOR result age age1
"1"
47.93.11.106:0>GET result
"\x90"
47.93.11.106:0>GETBIT age 0
"1"
47.93.11.106:0>BITCOUNT age
"1"
47.93.11.106:0>BITCOUNT age1
"1"
```

#### 3、STRLEN命令：统计键的值的字符长度

**命令格式**：`STRLEN key`

使用命令STRLEN统计key的值的字符长度。当key存储的不是字符串时，返回一个错误，当key不存在时，返回0。

```
47.93.11.106:0>GET motto
"拼个春夏秋冬，赢个无悔人生，努力过，拼搏过..."
47.93.11.106:0>STRLEN motto
"66"
```

### 3.1.7 键值对的值增量

#### 1、DECR命令：让键的值减1

**命令格式**：`DECR key`



## 3.2 哈希（Hash）命令


## 3.3 列表（List）命令

## 3.4 集合（Set）命令

## 3.5 有序集合（Sorted Set）


# 4. Redis必备命令

# 5. Redis数据库