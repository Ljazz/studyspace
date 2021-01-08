## re模块的主要方法

| 方法 | 含义 |
| --- | --- |
| re.compile | 编译一个正则表达式模式（pattern） |
| re.match | 从头开始匹配，使用group()方法可以获取第一个匹配值 |
| re.search | 用包含方式匹配，使用group()方法可以获取第一个匹配值 |
| re.findall | 用包含方式匹配，把所有匹配到的字符放到以列表中的元素返回多个匹配值 |
| re.sub | 匹配字符并替换 |
| re.split | 以匹配到的字符当作列表分隔符，返回列表 |

## re模块中方法详细介绍

1、 recompile方法

compile函数用于编译正则表达式，生成一个正则表达式(Pattern)对象，供match()和search()两个函数使用。

函数定义：`re.compile(pattern[, flags])`

参数说明：

- pattern：一个字符串形式的正则表达式
- flags：可选，表示匹配模式，具体参数如下
    - re.I：忽略大小写
    - re.L：表示特殊字符集\w, \W, \b, \B, \s, \S依赖于当前环境
    - re.M：多行模式
    - re.S：即为`.`并且包括换行在内的任意字符(`.`不包括换行符)
    - re.U：表示特殊字符集 \w, \W, \b, \B, \d, \D, \s, \S 依赖于 Unicode 字符属性数据库
    - re.X：为了增加可读性，忽略口康哥和`#`后面的注释

### re.match和re.search方法区别
- re.match方法从头开始匹配
- re.search方法可以从字符串中任意位置匹配
    