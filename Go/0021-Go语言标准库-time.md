# time包

time包提供了时间的显示和测量用的函数。日历的计算采用的是公历

## 时间类型

`time.Time`类型表示时间。我们可以通过`time.Now()`函数获取当前的时间对象，然后获取时间对象的年月日时分秒等信息。

```go
func timeDemo() {
    now := time.Now()  // 获取当前时间
    fmt.Printf("current time:%v\n", now)
    
    year := now.Year()   // 年
    month := now.Month() // 月
    day := now.Day()     // 日
    hour := now.Hour()   // 小时
    minute := now.Minute() // 分钟
    second := now.Second() // 秒
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
```

## 时间戳

时间戳是自1970年1月1日(08:00:00GMT)至当前时间的总毫秒。它也被称为Unix时间戳(UnixTimestamp)

基于时间对象获取时间戳的示例代码

```go
func timestampDemo() {
    now := time.Now()  // 获取当前时间
    timestamp1 := now.Unix()    // 时间戳
    timestamp2 := now.UnixNano()    // 纳秒时间戳
    fmt.Printf("current timestamp1:%v\n", timestamp1)
    fmt.Printf("current timestamp2:%v\n", timestamp2)
}
```

使用`time.Unix()`函数可以将时间戳转为时间格式

```go
func timestampDemo2(timestamp int64) {
    timeObj := time.Unix(timestamp, 0)  // 将时间戳转为时间格式
    fmt.Println(timeObj)
    year := timeObj.Year()  // 年
    month := timeObj.Month() // 月
    day := timeObj.Day()    // 日
    hour := timeObj.Hour()  // 小时
    minute := timeObj.Minute() // 分钟
    second := timeObj.Second() // 秒
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
```

## 时间间隔

`time.Duration`是`time`包定义的一个类型，它代表两个时间之间经过的时间，以纳秒为单位。`time.Duration`表示一段时间间隔，可表示的最长时间段大约290年。

time包中定义的时间间隔类型的常量

```go
const (
    Nanosecond  Duration = 1
    Microsecond          = 1000 * Nanosecond
    Millisecond          = 1000 * Microsecond
    Second               = 1000 * Millisecond
    Minute               = 60 * Second
    Hour                 = 60 * Minute
)
```

## 时间操作

### Add

Add方法主要解决时间+时间间隔的需求

```go
func (t Time) Add(d Duration) Time
```

求一个小时之后的时间

```go
func main() {
    now := time.Now()
    later := now.Add(time.Hour) // 当前时间加1小时后的时间
    fmt.Println(later)
}
```

### Sub

求两个时间之间的差值

```go
func (t Time) Sub(u Time) Duration
```

返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，将返回最大值/最小值。要获取时间点t-d(d为Duration)，可以使用t.Add(-d)。

### Equal

```go
func (t Time) Equal(u Time) bool
```

判断两个时间是否相同，会考虑时区的影响，因此不同失去标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息

### Before

```go
func (t Time) Before(u Time) bool
```
如果t代表的时间点在u之前，返回真；否则返回假

### After

```go
func (t Time) After(u Time) bool
```
如果t代表的时间点在u之后，返回真；否则返回假。

## 定时器

使用`time.Tick(时间间隔)`来设置定时器，定时器的本质上是一个通道(channel)。

