<!-- TOC -->

- [1 Map](#1-map)
  - [1.1 map定义](#11-map定义)
  - [1.2 map基本使用](#12-map基本使用)
  - [1.3 判断某个键是否存在](#13-判断某个键是否存在)
  - [1.4 map的遍历](#14-map的遍历)
  - [1.5 使用delete()函数删除键值对](#15-使用delete函数删除键值对)
  - [1.6 按照指定顺序遍历map](#16-按照指定顺序遍历map)
  - [1.7 元素为map类型的切片](#17-元素为map类型的切片)
  - [1.8 值为切片类型的map](#18-值为切片类型的map)
  - [1.9 map的长度](#19-map的长度)
  - [1.10 map是引用类型的](#110-map是引用类型的)

<!-- /TOC -->

# 1 Map

maps是一种无需的基于`key-value`的数据结构，Go中map是引用类型，必须初始化才能使用

## 1.1 map定义

Go中，`map`的定义语法如下：

```go
map[KeyType]ValueType
```
其中
- KeyType：表示键的类型
- ValueType：表示键对应值的类型

map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：

```go
make(map[KeyValue]ValueType, [cap])
```
其中cap表示map的容量，该参数虽然不是必须，但是我们应该在初始化map的时候就为其指定一个合适的容量。

使用map过程中需要注意的几点
- map是无序的，每次打印出来的map都可能会不一样，他不能通过index获取，而必须通过key获取
- map的长度是不固定的，也就是和slice一样，也是一种引用类型
- 内置的len函数同样适用于map，返回map拥有的key 的数量
- map的key可以是所有可比较的类型，如布尔型、整数型、浮点型、复杂型、字符串型...

## 1.2 map基本使用

map中的数据都是成对出现的，map的基本使用代码如下

```go
func main() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)
}
```
输出
```
map[小明:100 张三:90]
100
type of a:map[string]int
```

map也支持在声明的时候填充元素

```go
func main() {
    rating := map[string]float32 {
        "C":5,
        "Go":4.5,
        "Python":4.5,
        "C++":2
    }
    fmt.Println(rating)
}
```

## 1.3 判断某个键是否存在

Go中有个判断map中键是否存在的特殊写法，格式如下

```go
value, ok := map[key]
```

```go
func main() {
	countryCapitalMap := make(map[string]string)

	// map 插入 key-value 对，各个国家对应的首都
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
    countryCapitalMap["India"] = "New Delhi"
    
	// 查看元素在集合中是否存在
	captial, ok := countryCapitalMap["United States"]
	// 如果 ok 是 true， 则存在，否则不存在
	if ok {
		fmt.Println("Capital of United States is ", captial)
	} else {
		fmt.Println("Capital of United States is not paresent")
	}
}
```
运行结果
```text
Capital of United States is not paresent
```

## 1.4 map的遍历

Go中使用`for range`遍历map

```go
value, ok := map[key]
```

示例代码如下：
```go
func main() {
	countryCapitalMap := make(map[string]string)

	// map 插入 key-value 对，各个国家对应的首都
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	// 使用 key 输出 map 值
	for country := range countryCapitalMap {
		fmt.Println("Capital of ", country, "is ", countryCapitalMap[country])
	}
}
```
运行结果
```text
Capital of  Japan is  Tokyo
Capital of  India is  New Delhi
Capital of  France is  Paris
Capital of  Italy is  Rome
```

若我们只想遍历key的时候，可以按下面方法

```go
func main() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["娜扎"] = 60
	for k := range scoreMap {
		fmt.Println(k)
	}
}
```
<font color='red'>注意</font>：遍历map时的元素顺序添加键值对的顺序无关

## 1.5 使用delete()函数删除键值对

使用`delete()`内建函数从map中删除一组键值对。删除函数不返回任何值。`delete()`函数格式如下

```go
delete(map, key)
```
其中
- map：表示要删除键值对的map
- key：表示要删除的键值对的键

示例：
```go
package main

import (
	"fmt"
)

func main() {
	/* 创建 map */
	countryCapitalMap := map[string]string{"France": "Paris", "India": "New Delhi", "Italy": "Rome", "Japan": "Tokyo"}

	fmt.Println("原始 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}

	/* 删除元素 */
	delete(countryCapitalMap, "China")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后 map")

	/* 打印 map */
	for country := range countryCapitalMap {
		fmt.Println("Capital of", country, "is", countryCapitalMap[country])
	}
}
```
运行结果：
```text
原始 map
Capital of France is Paris
Capital of Italy is Rome
Capital of Japan is Tokyo
Capital of India is New Delhi
Entry for France is deleted
删除元素后 map
Capital of Italy is Rome
Capital of Japan is Tokyo
Capital of India is New Delhi
```
## 1.6 按照指定顺序遍历map

```go
func main() {
    rand.Seed(time.Now().UnixNano()) // 初始化随机数种子

    var scoreMap = make(map[string]int, 200)

    for i := 0; i < 100; i++ {
        key := fmt.Sprintf("stu%02d", i) // 生成stu开头的字符串
        value := rand.Intn(100) // 生成0~99的随机整数
        scoreMap[key] = value
    }
    // 取出map中的所有key存入切片keys
    var keys = make([]string, 0, 200)
    for key := range scoreMap {
        keys = append(keys, key)
    }
    // 对切片进行排序
    sort.Strings(keys)
    // 按排序后的key遍历
    for _, key := range keys {
        fmt.Println(key, scoreMap[key])
    }
}
```

## 1.7 元素为map类型的切片

```go
func main() {
    var mapSlice = make([]map[string]string, 3)
    for index, value := range mapSlice {
        fmt.Printf("index:%d value:%v\n", index, value)
    }
    fmt.Println("after init")
    // 对切片中的map元素进行初始化
    mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
    mapSlice[0]["address"] = "沙河"
    for index, value := range mapSlice {
        fmt.Printf("index:%d value:%v\n", index, value)
    }
}
```

## 1.8 值为切片类型的map

```go
func main() {
    var sliceMap = make(map[string]string, 3)
    fmt.Println(sliceMap)
    fmt.Println("after init")
    key := "中国"
    value, ok := sliceMap[key]
    if !ok {
        value = make([]string, 0, 2)
    }
    value = append(value, "北京", "上海")
    sliceMap[key] = value
    fmt.Println(sliceMap)
}
```

## 1.9 map的长度

使用len函数可以确定map的长度

> `len(map)`  // 可以得到map的长度

## 1.10 map是引用类型的

与切片相似，映射是引用类型。当将映射分配给一个新的变量时，我们都指向相同的内部数据结构。因此，一个的变化会反映另一个。

```go
package main

import (  
    "fmt"
)

func main() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("Original person salary", personSalary)
    newPersonSalary := personSalary
    newPersonSalary["mike"] = 18000
    fmt.Println("Person salary changed", personSalary)

}
```
运行结果：
```text
Original person salary map[steve:12000 jamie:15000 mike:9000]  
Person salary changed map[steve:12000 jamie:15000 mike:18000] 
```

> map不能使用==操作符进行比较。==只能用来检查map是否为空。否则会报错：invalid operation: map1 == map2 (map can only be comparedto nil)
