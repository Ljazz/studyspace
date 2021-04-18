<!-- TOC -->

- [1. Redis介绍](#1-redis介绍)
  - [1.1 Redis支持的数据结构](#11-redis支持的数据结构)
  - [1.2 Redis的应用场景](#12-redis的应用场景)
  - [1.3 准备Redis的环境](#13-准备redis的环境)
- [2. redis相关库](#2-redis相关库)
  - [2.1 安装](#21-安装)
  - [2.1 go-redis库](#21-go-redis库)
    - [2.1.1 链接Redis](#211-链接redis)
    - [2.1.2 基本指令](#212-基本指令)
      - [Keys()：根据正则获取keys](#keys根据正则获取keys)
      - [Type(): 获取key对应值的类型](#type-获取key对应值的类型)
      - [Del(): 删除缓存项](#del-删除缓存项)
      - [Exists(): 检测缓存项是否存在](#exists-检测缓存项是否存在)
      - [Expire(), ExpireAt(): 设置有效期](#expire-expireat-设置有效期)
      - [TTL(), PTTL(): 获取有效期](#ttl-pttl-获取有效期)
      - [DBSize(): 查看当前数据库key的数量](#dbsize-查看当前数据库key的数量)
      - [FlushDB(): 清空当前数据库](#flushdb-清空当前数据库)
      - [FlushAll(): 清空所有数据库](#flushall-清空所有数据库)
    - [2.1.3 字符串(string)类型](#213-字符串string类型)
      - [Set():设置](#set设置)
      - [SetEX(): 设置并指定过期时间](#setex-设置并指定过期时间)
      - [SetNX()：设置并指定过期时间](#setnx设置并指定过期时间)
      - [Get()：获取](#get获取)
      - [GetRange()：字符串截取](#getrange字符串截取)
      - [Incr()：增加+1](#incr增加1)
      - [IncrBy()：按指定步长增加](#incrby按指定步长增加)
      - [Decr()：减少1](#decr减少1)
      - [DecrBy()：按只当的步长减少](#decrby按只当的步长减少)
      - [Append()：追加](#append追加)
      - [StrLen()：获取长度](#strlen获取长度)
    - [2.1.4 列表(list)类型](#214-列表list类型)
      - [LPush()：将元素压入链表](#lpush将元素压入链表)
      - [LInsert()：在某个位置插入新元素](#linsert在某个位置插入新元素)
      - [LSet()：设置某个元素的值](#lset设置某个元素的值)
      - [LLen()：获取链表元素个数](#llen获取链表元素个数)
      - [LIndex()：获取链表下表对应的元素](#lindex获取链表下表对应的元素)
      - [LRange()：获取某个选定范围的元素集](#lrange获取某个选定范围的元素集)
      - [LPop(): 从链表左侧弹出数据](#lpop-从链表左侧弹出数据)
      - [LRem()：根据值移除元素](#lrem根据值移除元素)
    - [2.1.5 集合(set)类型](#215-集合set类型)
      - [SAdd()：添加元素](#sadd添加元素)
      - [SPop()：随机获取一个元素](#spop随机获取一个元素)
      - [SRem()：删除集合中指定的值](#srem删除集合中指定的值)
      - [SSMembers()：获取所有成员](#ssmembers获取所有成员)
      - [SIsMember()：判断元素是否在集合中](#sismember判断元素是否在集合中)
      - [SCard()：获取集合元素个数](#scard获取集合元素个数)
      - [SUnion：并集；SDiff()：差集；SInter()：交集](#sunion并集sdiff差集sinter交集)
    - [2.1.6 有序集合(zset)类型](#216-有序集合zset类型)
      - [ZAdd()：添加元素](#zadd添加元素)
      - [ZIncrBy()：增加元素分值](#zincrby增加元素分值)
      - [ZRange()、ZRevRange()：获取根据score排序后的数据段](#zrangezrevrange获取根据score排序后的数据段)
      - [ZRangeByScore()、ZRevRangeByScore()：获取score过滤后排序的数据段](#zrangebyscorezrevrangebyscore获取score过滤后排序的数据段)
      - [ZCard()：获取元素个数](#zcard获取元素个数)
      - [ZCount()：获取区间内元素个数](#zcount获取区间内元素个数)
      - [ZScore()：获取元素的score](#zscore获取元素的score)
      - [ZRank()、ZRevRank()：获取某个元素在集合中的排名](#zrankzrevrank获取某个元素在集合中的排名)
      - [ZRem()：删除元素](#zrem删除元素)
      - [ZRemRangeByRank()：根据排名来删除](#zremrangebyrank根据排名来删除)
      - [ZRemRangeByScore()：根据分值区间来删除](#zremrangebyscore根据分值区间来删除)
    - [2.1.7 哈希(hash)类型](#217-哈希hash类型)
      - [HSet()：设置](#hset设置)
      - [HMSet()：批量设置](#hmset批量设置)
      - [HGet()：获取某个元素](#hget获取某个元素)
      - [HGetAll()：获取全部元素](#hgetall获取全部元素)
      - [HDel()：删除某个元素](#hdel删除某个元素)
      - [HExists()：判断元素是否存在](#hexists判断元素是否存在)
      - [HLen()：获取长度](#hlen获取长度)
  - [2.2 redigo库](#22-redigo库)
    - [2.2.1 链接Redis](#221-链接redis)
    - [2.2.2 String类型Set、Get操作](#222-string类型setget操作)
    - [2.2.3 String批量操作](#223-string批量操作)
    - [2.2.4 设置过期时间](#224-设置过期时间)
    - [2.2.5 List队列操作](#225-list队列操作)
    - [2.2.6 Hash表](#226-hash表)
    - [2.2.7 Redis连接池](#227-redis连接池)
  - [2.3 连接Redis哨兵模式](#23-连接redis哨兵模式)
  - [2.4 连接Redis集群](#24-连接redis集群)
  - [根据前缀获取Key](#根据前缀获取key)
  - [执行自定义命令](#执行自定义命令)
  - [Pipeline](#pipeline)
  - [事务](#事务)
  - [Watch](#watch)

<!-- /TOC -->

# 1. Redis介绍

Redis是一个开源的，遵守BSD协议的高性能的key-value数据库。

Redis与其他key-value缓存产品有以下三个特点

1. Redis支持数据的持久化，可以将内存中的数据保存在磁盘中，重启的时候可以再次加载进行使用。
2. Redis不仅仅支持简单的key-value类型的数据，同时还提供string、list（链表）、set（集合）、hash表等数据结构的存储。
3. Redis支持数据的备份，即master-slave模式的数据备份。

## 1.1 Redis支持的数据结构

Redis支持的数据结构有

1. 字符串（strings）
2. 哈希（hashes）
3. 列表（lists）
4. 集合（sets）
5. 带范围查询的排序集合（sorted sets）
6. 位图（bitmaps）
7. hyperloglogs
8. 带半径查询
9. 地理空间索引

## 1.2 Redis的应用场景

- 缓存系统，减轻主数据库的压力
- 计数场景，比如微博、抖音中的关注数和粉丝数
- 人们排行榜，需要排序的场景特别适合使用ZSET
- 利用LIST可以实现队列的功能

## 1.3 准备Redis的环境

利用docker启动一个redis

```bash
docker run -name redis -p 6379:6379 -d redis:5.0.7
docker run it --netword host --rm redis:5.0.7 redis-cli
```

# 2. redis相关库

## 2.1 安装

比较常用的Go语言redis client库：[redigo](https://github.com/gomodule/redigo)，[go-redis](https://github.com/go-redis/redis)也可以用来连接redis数据库并进行操作且`go-redis`支持连接哨兵及集群模式的Redis。

[go-redis文档](https://pkg.go.dev/github.com/go-redis/redis)

> `go get -u github.com/go-redis/redis`
> `go get github.com/garyburd/redigo/redis`

## 2.1 go-redis库

### 2.1.1 链接Redis

```go
package main

import (
    "context"
    "fmt"
    "time"

    "github.com/go-redis/redis/v8"
)

//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。 
//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    pong, err := rdb.Ping(ctx).Result()
    if err != nil {
        fmt.Printf("connect redis failed, %v", err)
    }
}
```

### 2.1.2 基本指令

#### Keys()：根据正则获取keys

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB: 0,
    })

    // 表示获取所有的key
    keys, err := rdb.Keys(ctx, "*").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(keys)
}
```

#### Type(): 获取key对应值的类型

`Type()`方法用户获取一个key对应值的类型

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/reids"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB: 0,
    })

    vType, err := rdb.Type(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(vType) // string
}
```

#### Del(): 删除缓存项

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB: 0,
    })

    n, err := rdb.Del(ctx, "key1", "key2").Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("成功删除了 %v 个\n", n)
}
```

#### Exists(): 检测缓存项是否存在

`Exists()`方法用于检测某个key是否存在

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB: 0,
    })
    
    n, err := rdb.Exist(ctx, "key1").Result()
    if err != nil {
        panic(err)
    }
    if n > 0 {
        fmt.Println("存在")
    } else {
        fmt.Println("不存在")
    }
}
```

`Exists()`方法可以传入多个key，返回的第一个结果表示存在的key的数量

#### Expire(), ExpireAt(): 设置有效期

需要在设置好了缓存项后，在设置有效期

`Expire()`方法是设置某个时间段(time.Duration)后过期，`ExpireAt()`方法是在某个时间点(time.Time)过期失效

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
    "time"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB: 0,
    })

    res, err := rdb.Expire(ctx, "key", time.Minute * 2).Result()
    if err != nil {
        panic(err)
    }
    if res {
        fmt.Println("设置成功")
    } else {
        fmt.Println("设置失败")
    }

    res, err = rdb.ExpireAt(ctx, "key2", time.Now()).Result()
    if err != nil {
        panic(err)
    }
    if res {
        fmt.Println("设置成功")
    } else {
        fmt.Println("设置失败")
    }
}
```

#### TTL(), PTTL(): 获取有效期

`TTL()`方法可以获取某个键的剩余有效期

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
    "time"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB: 0,
    })

    // 设置一分钟的有效期
    rdb.Expire(ctx, "key", time.Minute)

    // 获取剩余有效期，单位：秒
    ttl, err := rdb.TTL(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(ttl)

    // 获取剩余有效期，单位：秒
    ttl, err := rdb.TTL(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(ttl)
}
```

#### DBSize(): 查看当前数据库key的数量

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB: 0,
    })

    num, err := rdb.DBSize(ctx).Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("数据库有 %v 个缓存项\n", num)
}
```

#### FlushDB(): 清空当前数据库

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB: 0,
    })

    // 清空当前数据库，因为连接的是索引为0的数据库，所以清空的就是0号数据库
    res, err := rdb.FlushDB(ctx).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}
```

#### FlushAll(): 清空所有数据库

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB: 0,
    })

    // 清空当前数据库，因为连接的是索引为0的数据库，所以清空的就是0号数据库
    res, err := rdb.FlushAll(ctx).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}
```

### 2.1.3 字符串(string)类型

#### Set():设置

仅支持字符串（包含数字）操作，不支持内置数据编码功能。若要存储Go的非字符串类型，需要提前手动序列化，获取时再反序列化。

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
    "time"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    _, err := rdb.Ping(ctx).Result()
    if err != nil {
        fmt.Printf("连接redis出错，错误信息：%v", err)
        return
    }

    // Set方法的最后一个参数表示过期时间，0表示永不过期
    err = rdb.Set(ctx, "key1", "value1", 0).Err()
    if err != nil {
        panic(err)
    }

    // Key2将会在2分钟后过期失效
    err = rdb.Set(ctx, "key2", "value2", time.Minute * 2).Err()
    if err != nil {
        panic(err)
    }
}
```

#### SetEX(): 设置并指定过期时间

设置键的同时，设置过期时间

```go
package main

import (
    "context"
    "github.com/go-redis/redis"
    "time"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379".
        Password: "",
        DB:       0,
    })

    err := rdb.SetEX(ctx, "key", "value", time.Hour * 2).Err()
    if err != nil {
        panic(err)
    }
}
```

#### SetNX()：设置并指定过期时间

> SetNX()与SetEX()的区别是，SetNX()仅当key不存在的时候才设置，如果key已经存在则不做任何操作，而SetEX()方法不管该key是否已经存在，缓存中直接覆盖

调用SetNX()是否调用成功，可以调用Result()方法，返回的第一个值表示是否设置成功了。若返回false，说明缓存key已经存在了，此次操作没有错误且不起作用；若返回true，说明key之前不存在缓存中，操作成功。

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
    "time"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379".
        Password: "",
        DB:       0,
    })

    res, err := rdb.SetNX(ctx, "key", "value", time.Minute).Result()
    if err != nil {
        panic(err)
    }
    if res {
        fmt.Println("设置成功")
    } else {
        fmt.Println("key已经存在缓存中，设置失败")
    }
}
```

#### Get()：获取

如果获取的key在缓存中并不存在，`Get()`方法将返回`redis.Nil`

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379".
        Password: "",
        DB:       0,
    })

    err := rdb.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := rdb.Get(ctx, "Key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("key: %v\n", val)

    val2, err := rdb.Get(ctx, "key-not-exist").Result()
    if err == redis.Nil {
        fmt.Println("key不存在")
    } else  if != nil {
        panic(err)
    } else {
        fmt.Printf("值为：%v\n", val2)
    }
}
```

#### GetRange()：字符串截取

`GetRange()`方法可以用来截取字符串的部分内容，第二个参数是下标索引的开始位置，第三个参数是小标索引的结束位置

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379".
        Password: "",
        DB:       0,
    })

    err := rdb.Set(ctx, "key", "hello world", 0).Err()
    if err != nil {
        panic(err)
    }
    val, err := rdb.GetRange(ctx, "key", 1, 4).Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("key: %v\n", val)
}
```

> 如果key不存在，调用GetRange()也不会报错，只是返回空字符串

#### Incr()：增加+1

`Incr()`、`IncrBy()`两个是对数字进行增加的操作，`incr`是执行原子`+1`操作，incrBy是增加指定的数。

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379".
        Password: "",
        DB:       0,
    })

    val, err := rdb.Incr(ctx, "number").Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("key当前的值为：%v\n", val)
}
```

#### IncrBy()：按指定步长增加

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379".
        Password: "",
        DB:       0,
    })

    val, err := rdb.IncrBy(ctx, "number", 2).Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("key当前的值为：%v\n", val)
}
```

#### Decr()：减少1

`Decr`、`DecrBy`方法是对数字进行减的操作

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379".
        Password: "",
        DB:       0,
    })

    val, err := rdb.Decr(ctx, "number").Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("key当前的值为：%v\n", val)
}
```

#### DecrBy()：按只当的步长减少

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379".
        Password: "",
        DB:       0,
    })

    val, err := rdb.DecrBy(ctx, "number", 2).Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("key当前的值为：%v\n", val)
}
```

#### Append()：追加

`Append()`表示往字符串后面追加元素，返回值是字符串的总长度

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379".
        Password: "",
        DB:       0,
    })

    err := rdb.Set(ctx, "key", "hello", 0).Err()
    if err != nil {
        panic(err)
    }
    length, err := rdb.Append(ctx, "key", " world").Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("当前缓存key的长度为：%v\n",length)

    val, err := rdb.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("当前缓存key的值为: %v\n", val)
}
```

#### StrLen()：获取长度

`StrLen()`方法可以获取字符串的长度

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379".
        Password: "",
        DB:       0,
    })

    err := rdb.Set(ctx, "key", "hello world", 0).Err()
    if err != nil {
        panic(err)
    }
    length, err := rdb.StrLen(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("当前缓存key的长度为：%v\n",length)
}
```

### 2.1.4 列表(list)类型

#### LPush()：将元素压入链表

`LPush()`方法将数据从左侧压入链表

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    //返回值是当前列表元素的数量
    n, err := rdb.LPush(ctx, "list", 1, 2, 3).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(n)
}
```

> 也可以从右侧压入连接，对应的方法为`RPush()`

#### LInsert()：在某个位置插入新元素

位置的判断，是根据相对应的参考元素判断

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    // 在名为key的缓存项值为100的元素前面插入一个值，值为123
    err := rdb.LInsert(ctx, "key", "before", "100", 123).Err()
    if err != nil {
        panic(err)
    }
}
```

> 如果key列表中有多个指定的缓存项，LInsert只会在第一个缓存项前面插入待插入项，并不会在所有的指定缓存项前面都插入。

#### LSet()：设置某个元素的值

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    // 下表是从0开始
    err := rdb.LSet(ctx, "list", 1, 100).Err()
    if err != nil {
        panic(err)
    }
}
```

#### LLen()：获取链表元素个数

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    length, err := rdb.LLen(ctx, "list").Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("当前链表的长度为: %v\n", length)
}
```

#### LIndex()：获取链表下表对应的元素

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    val, err := rdb.LIndex(ctx, "list", 0).Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("下标为0的值为: %v\n", val)
}
```

#### LRange()：获取某个选定范围的元素集

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    vals, err := rdb.LRange(ctx, "list", 0, 2).Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("从下标0到下标2的值: %v\n", vals)
}
```

#### LPop(): 从链表左侧弹出数据

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    val, err := rdb.LPop(ctx, "list").Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("移除的元素为：%v\n", val)
}
```

#### LRem()：根据值移除元素

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    n, err := rdb.LRem(ctx, "list", 2, "100").Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("移除了：%v个\n", n)
}
```

### 2.1.5 集合(set)类型

Redis set对外提供的功能与list类似，特殊之处在set可以自动排重。set也提供了判断一个成员是否在set集合内的接口，而list不提供。

Redis Set是string类型的无序集合。底层是一个value为null的hash表，所以添加、删除、查找的复杂度都为O(1)

集合数据的特征

1. 元素不能重复，保持唯一性
2. 元素无序，不能使用索引（下标）操作

#### SAdd()：添加元素

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    rdb.SAdd(ctx, "team", "kobe", "jordan")
    rdb.SAdd(ctx, "team", "curry")
    rdb.SAdd(ctx, "team", "kobe")
}
```

#### SPop()：随机获取一个元素

无序性，是随机的

`SPop()`方法是从集合中随机取出元素，如果想要一次获取多个元素，可以使用`SPopN()`方法

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    val, err := rdb.SPop(ctx, "team").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(val)
}
```

#### SRem()：删除集合中指定的值

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    n, err := rdb.SRem(ctx, "team", "kobe", "v2").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(n)
}
```

#### SSMembers()：获取所有成员

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    vals, err := rdb.SSMembers(ctx, "team").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(vals)
}
```

#### SIsMember()：判断元素是否在集合中

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    exists, err := rdb.SIsMember(ctx, "team", "jordan").Result()
    if err != nil {
        panic(err)
    }
    if exists {
        fmt.Println("存在集合中")
    } else {
        fmt.Println("不存在集合中")
    }
}
```

#### SCard()：获取集合元素个数

获取集合中元素个数

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    total, err := rdb.SCard(ctx, "team").Result()
    if err != nil {
        panic(err)
    }
    fmt.Printf("集合总共有 %v 个元素\n", total)
}
```

#### SUnion：并集；SDiff()：差集；SInter()：交集

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    rdb.SAdd(ctx, "setA", "a", "b", "c", "d")
    rdb.SAdd(ctx, "setB", "a", "d", "e", "f")

    // 并集
    union, err := rdb.SUnion(ctx, "setA", "setB").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(union)

    // 差集
    diff, err := rdb.SDiff(ctx, "setA", "setB").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(diff)

    // 交集
    inter, err := rdb.Inter(ctx, "setA", "setB").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(inter)
}
```

### 2.1.6 有序集合(zset)类型

Redis有序集合和集合一样也是string类型元素的集合，且不允许重复

有序集合的每个元素都会关联一个double类型的分数，redis通过这个分数对集合中的成员进行从小到大的排序。

有序集合成员是唯一的，但是`分数(score)`是可以重复的。

集合通过哈希表实现，所以添加、删除、查找的复杂度都是O(1)。集合中最大的成员数为2^32 - 1

#### ZAdd()：添加元素

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    rdb.ZAdd(ctx, "zSet", &redis.Z{
        Socre:  0,
        Member: 1,
    })
    rdb.ZAdd(ctx, "zSet", &redis.Z{
        Socre:  0,
        Member: 2,
    })
    rdb.ZAdd(ctx, "zSet", &redis.Z{
        Socre:  0,
        Member: 3,
    })
    rdb.ZAdd(ctx, "zSet", &redis.Z{
        Socre:  0,
        Member: 4,
    })
    rdb.ZAdd(ctx, "zSet", &redis.Z{
        Socre:  0,
        Member: 5,
    })
    rdb.ZAdd(ctx, "zSet", &redis.Z{
        Socre:  0,
        Member: 6,
    })
}
```

#### ZIncrBy()：增加元素分值

分值可以为负数，表示递减

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    rdb.ZIncrBy(ctx, "zSet", float64(rand.Intn(100)), "1")
    rdb.ZIncrBy(ctx, "zSet", float64(rand.Intn(100)), "2")
    rdb.ZIncrBy(ctx, "zSet", float64(rand.Intn(100)), "3")
    rdb.ZIncrBy(ctx, "zSet", float64(rand.Intn(100)), "4")
    rdb.ZIncrBy(ctx, "zSet", float64(rand.Intn(100)), "5")
    rdb.ZIncrBy(ctx, "zSet", float64(rand.Intn(100)), "6")
}
```

#### ZRange()、ZRevRange()：获取根据score排序后的数据段

根据分值排序后的，升序和降序的列表获取

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    // 获取排行耪
    // 后去分值
    res, err := rdb.ZRevRange(ctx, "zSet", 0, 2).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}
```

#### ZRangeByScore()、ZRevRangeByScore()：获取score过滤后排序的数据段

根据过滤之后的列表

需要提供分值区间

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    res, err := rdb.ZRangeByScore(ctx, "zSet", &redis.ZRangeBy{
        Min: "40",
        Max: "85",
    }).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}
```

#### ZCard()：获取元素个数

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    count, err := rdb.ZCard(ctx, "zSet").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(count)
}
```

#### ZCount()：获取区间内元素个数

获取分值在[40, 85]的元素的数量

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    n, err := rdb.ZCount(ctx, "zSet", "40", "85").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(n)
}
```

#### ZScore()：获取元素的score

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    score, err := rdb.ZScore(ctx, "zSet", "5").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(score)
}
```

#### ZRank()、ZRevRank()：获取某个元素在集合中的排名

`ZRank()`方法是返回元素在集合中的升序排名清空，从0开始。`ZRevRank()`与之相反

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    res, err := rdb.ZRevRank(ctx, "zSet", "2").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}
```

#### ZRem()：删除元素

通过元素的值来删除元素

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    res, err := rdb.ZRem(ctx, "zSet", "2").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}
```

#### ZRemRangeByRank()：根据排名来删除

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    // 按照升序排序删除第一个和第二个元素
    res, err := rdb.ZRemRangeByRank(ctx, "zSet", 0, 1).Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}
```

#### ZRemRangeByScore()：根据分值区间来删除

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    // 删除score在[40, 70]之间的元素
    res, err := rdb.ZRemRangeByScore(ctx, "zSet", "40", "70").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}
```

### 2.1.7 哈希(hash)类型

Redis hash是一个string 类型的field（字段）和value（值）的映射表，hash特别适合用于存储对象。

Redis中每个hash可以存储2^32-1键值对。

#### HSet()：设置

```go
package main

import (
    "context"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })

    rdb.HSet(ctx, "user", "key1", "value1", "key2", "value2")
    rdb.HSet(ctx, "user", []string{"key3", "value3", "key4", "value4"})
    rdb.HSet(ctx, "user", map[string]interface{}{"key5": "value5", "key6": "value6"})
}
```

#### HMSet()：批量设置

```go
package main

import (
    "context"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
    rdb.HMSet(ctx, "user1", map[string]interface{}{"name":"kevin", "age": 27, "address":"北京"})
}
```

#### HGet()：获取某个元素

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
    address, err := rdb.HGet(ctx, "user", "address").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(address)
}
```

#### HGetAll()：获取全部元素

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
    user, err := rdb.HGet(ctx, "user").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(user)
}
```

#### HDel()：删除某个元素

`HDel()`支持一次删除多个元素

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
    res, err := rdb.HDel(ctx, "user", "name", "age").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}
```

#### HExists()：判断元素是否存在

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
    res, err := rdb.HExists(ctx, "user", "address").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}
```

#### HLen()：获取长度

```go
package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "",
        DB:       0,
    })
    res, err := rdb.HLen(ctx, "user").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println(res)
}
```

## 2.2 redigo库

### 2.2.1 链接Redis

```go
package main

import (
    "fmt"
    "github.com/garyburd/redigo/redis"
)

func main() {
    c, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        fmt.Println("conn redis failed, ", err)
        return
    }
    defer c.Close()
    fmt.Println("redis conn success")
}
```

### 2.2.2 String类型Set、Get操作

```go
package main

import (
    "fmt"

    "github.com/garyburd/redigo/redis"
)

func main() {
    c, err := redis.Dial("tcp", "localhost:6379", redis.DialPassword("password"))
    if err != nil {
        fmt.Println("conn redis failed, ", err)
        return
    }
    defer c.Close()
    fmt.Println("redis conn success")

    _, err = c.Do("Set", "abc", 100)
    if err != nil {
        fmt.Println("Set failed, ", err)
        return
    }

    r, err := redis.Int(c.Do("Get", "abc"))
    if err != nil {
        fmt.Println("get abc failed, ", err)
        return
    }
    fmt.Println(r)
}
```

### 2.2.3 String批量操作

```go
package main
import (
    "fmt"
    "github.com/garyburd/redigo/redis"
)
func main() {
    c, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        fmt.Println("conn redis failed,", err)
        return
    }
    defer c.Close()
    _, err = c.Do("MSet", "abc", 100, "efg", 300)
    if err != nil {
        fmt.Println(err)
        return
    }
    r, err := redis.Ints(c.Do("MGet", "abc", "efg"))
    if err != nil {
        fmt.Println("get abc failed,", err)
        return
    }
    for _, v := range r {
        fmt.Println(v)
    }
}
```

### 2.2.4 设置过期时间

```go
package main
import (
    "fmt"
    "github.com/garyburd/redigo/redis"
)
func main() {
    c, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        fmt.Println("conn redis failed,", err)
        return
    }
    defer c.Close()
    _, err = c.Do("expire", "abc", 10)
    if err != nil {
        fmt.Println(err)
        return
    }
}
```

### 2.2.5 List队列操作

```go
package main
import (
    "fmt"
    "github.com/garyburd/redigo/redis"
)
func main() {
    c, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        fmt.Println("conn redis failed,", err)
        return
    }
    defer c.Close()
    _, err = c.Do("lpush", "book_list", "abc", "ceg", 300)
    if err != nil {
        fmt.Println(err)
        return
    }
    r, err := redis.String(c.Do("lpop", "book_list"))
    if err != nil {
        fmt.Println("get abc failed,", err)
        return
    }
    fmt.Println(r)
}
```

### 2.2.6 Hash表

```go
package main
import (
    "fmt"
    "github.com/garyburd/redigo/redis"
)
func main() {
    c, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        fmt.Println("conn redis failed,", err)
        return
    }
    defer c.Close()
    _, err = c.Do("HSet", "books", "abc", 100)
    if err != nil {
        fmt.Println(err)
        return
    }
    r, err := redis.Int(c.Do("HGet", "books", "abc"))
    if err != nil {
        fmt.Println("get abc failed,", err)
        return
    }
    fmt.Println(r)
}
```

### 2.2.7 Redis连接池

```go
package main
import(
    "fmt"
    "github.com/garyburd/redigo/redis"
)
var pool *redis.Pool  //创建redis连接池
func init(){
    pool = &redis.Pool{     //实例化一个连接池
        MaxIdle:16,    //最初的连接数量
        // MaxActive:1000000,    //最大连接数量
        MaxActive:0,    //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
        IdleTimeout:300,    //连接关闭时间 300秒 （300秒不使用自动关闭）    
        Dial: func() (redis.Conn ,error){     //要连接的redis数据库
            return redis.Dial("tcp","localhost:6379")
        },
    }
}
func main(){
    c := pool.Get() //从连接池，取一个链接
    defer c.Close() //函数运行结束 ，把连接放回连接池
        _,err := c.Do("Set","abc",200)
        if err != nil {
            fmt.Println(err)
            return
        }
        r,err := redis.Int(c.Do("Get","abc"))
        if err != nil {
            fmt.Println("get abc faild :",err)
            return
        }
        fmt.Println(r)
        pool.Close() //关闭连接池
}
```

## 2.3 连接Redis哨兵模式

```go
func initClient()(err error) {
    rdb := redis.NewFailoverClient(&redis.FailoverOptions{
        Master: "master",
        SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
    })

    ctx = context.Background()
    _, err = rdb.Ping(ctx).Result()
    if err != nil {
        return err
    }
    return nil
}
```

## 2.4 连接Redis集群

```go
func initClient()(err error) {
    rdb := redis.NewClusterClient(&redis.ClusterOptions{
        Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"}
    })

    ctx = context.Background()
    _, err = rdb.Ping(ctx).Result()
    if err != nil {
        return err
    }
    return nil
}
```

## 根据前缀获取Key

> `vals, err := rdb.Keys(ctx, "prefix").Result()`

## 执行自定义命令

> `res, err := rdb.Do(ctx, "set", "key", "value").Result()`

## Pipeline

`Pipeline`主要是一种网络优化。本质上意味着客户端缓冲一堆命令并一次性将它们发送到服务器。这些命令不能保证在事务中执行。这样做的好处是节省了每个命令的网络往返时间(RTT)

```go
pipe := rdb.Pipeline()

incr := pipe.Incr("pipeline_counter")
pipe.Expire("pipeline_counter", time.Hour)

_, err := pipe.Exec()
fmt.Println(incr.Val(), err)
```

`Pipelined`

```go
var incr *redis.IntCmd
_, err := rdb.Pipelined(func(pipe redis.Pipeliner) error {
    incr = pipe.Incr("pipeline_counter")
    pipe.Expire("pipeline_counter", time.Hour)
    return nil
})
fmt.Println(incr.Val(), err)
```

## 事务

Redis是单线程的，因此单个命令始终是原子的，但是来自不同客户端的两个给定命令可以依次执行。但是，`Multi/exec`能够确保在`multi/exec`两个语句之间的命令之间没有其他客户端正在执行命令。

```go
pipe := rdb.TxPipeline()

incr := pipe.Incr("tx_pipeline_counter")
pipe.Expire("tx_pipeline_counter", time.Hour)

_, err := pipe.Exec()
fmt.Println(incr.Val(), err)
```

`TxPipelined`

```go
var incr *redis.IntCmd
_, err := rdb.TxPipelined(func(pipe redis.Pipeliner) error {
    incr = pipe.Incr("tx_pipeline_counter")
    pipe.Expire("tx_pipeline_counter", time.Hour)
    return nil
})
fmt.Println(incr.Val(), err)
```

## Watch

某些场景下，除了要使用`MULTI/EXEC`命令外，还需要配合使用`WATCH`命令。在用户使用`WATCH`命令监视某个键之后，直到该用户执行EXEC命令的这段时间里，如果有其他用户抢先对被监视的键进行了替换、更新、删除等操作，那么当用户尝试执行EXEC的时候，事务将失败并返回一个错误，用户可以根据这个错误选择重试事务或者放弃事务。

> `Watch(fn func(*Tx) error, keys ...string) error`

Watch方法接收一个函数和一个或多个key作为参数。

```go
// 监视watch_count的值，并在值不变的前提下将其+1
key := "watch_count"
err = client.Watch(func(tx *redis.Tx) error {
    n, err := tx.Get(key).Int()
    if err != nil && err != redis.Nil {
        return err
    }
    _, err = tx.Pipelined(func(pipe redis.Pipeliner) error {
        pipe.Set(key, n+1, 0)
        return nil
    })
    return err
}, key)

func transactionDemo() {
    var (
        maxRetries   = 1000
        routineCount = 10
    )
    ctx, cancel := context.WitchTimeout(context.Background(), 5 * time.Second)
    defer cancel()

    // Increment 使用GET和SET命令以事务方式递增key 的值
    increment := func(key string) error {
        // 事务函数
        txf := func(tx *redis.Tx) error {
            // 获得key的当前值或零值
            n, err := tx.Get(ctx, key).Int()
            if err != nil && err != redis.Nil {
                return err
            }

            // 实际的操作代码
            n++

            // 操作权在Watch的key没有发生变化的情况下提交
            _, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
                pipe.Set(ctx, key, n, 0)
                return nil
            })
            return err
        }

        // 最多重试
        for i := 0; i < maxRetries; i++ {
            err := rdb.Watch(ctx, txf, key)
            if err == nil {
                // 成功
                return nil
            }
            if err == redis.TxFailedErr {
                // 乐观锁丢失，重试
                continue
            }
            // 返回其他错误
            return err
        }
        return errors.New("increment reached maximum number of retries")
    }

    // 模拟 routineCount 个并发同时去修改 counter3的值
    var wg sync.WaitGroup
    wg.Add(routineCount)
    for i := 0; i < routineCount; i++ {
        go func() {
            defer wg.Done()
            if err := increment("counter3"); err != nil {
                fmt.Println("increment error: ", err)
            }
        }()
    }
    wg.Wait()

    n, err := rdb.Get(context.TODO(), "counter3").Int()
    fmt.Println("ended with", n, err)
}
```
