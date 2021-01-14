<!-- TOC -->

- [什么是Model模型？](#什么是model模型)
- [Django Model中字段（Field）的可选项和必选项](#django-model中字段field的可选项和必选项)
- [常见的Django Model META类选项](#常见的django-model-meta类选项)

<!-- /TOC -->

## 什么是Model模型？

Model（模型）简言而之即数据模型。模型不是数据本身（比如数据库里的数据），而是抽象的描述数据的构成和逻辑关系。每个Django model实际上是个类，继承了`models.Model`。每个Model应该包括属性、关系（即单对单、单对多、多对多）和方法。当定义好一个Model模型后，利用Django的接口会自动帮你在数据库中生成相应的数据表。

例如：属于出版社的实际案例
```python
# models.py
from django.db import models


class Publisher(models.Model):
    name = models.CharField(max_length=30)
    address = models.CharField(max_length=30)

    def __str__(self):
        return self.name


class Book(models.Model):
    name = models.CharField(max_length=30)
    description = models.TextField(blank=True, default="")
    publisher = ForeignKey(Publisher, on_delete=models.CASCADE)
    add_date = models.DateField()

    def __str__(self):
        return self.name
```

定义Django模型Model时，需要注意：
- 这个Field是否有必选项
- 这个Field是否必需（blank = True or False），是否可以为空(null = True or False)

## Django Model中字段（Field）的可选项和必选项

**橘黄色为必填项**

1. **CharField()    字符字段**
    - <font color='orange'>max_length = xxx or None</font>
    - 若不是必填项，可设置blank = True和default=''
    - 若用于username等字段，想使其唯一，可设置unique = True
    - 若有choice选项，可设置choice = XXX_CHOICES
2. **TextField()  文本字段**
    - max_length = xxx
    - 若不是必填项，可设置blank = True和default = ''
3. **DateField/DateTimeField()    日期与时间字段**
    - 一般建立设置默认日期default date，也可以使用`auto_now_add=True`
    - 对于DateField：`default=date.today`-先要`from datetime import date`
    - 对于DateTimeField：`default=timezone.now`-先要`from django.utils import timezone`
    - 对于上一次修改的日期可设置:`auto_now=True`
4. **EmailField() 邮件字段**
    - 若不是必填项，可设置blank=True和default=''
    - 一般Email用于用户名应是唯一的，应设置unique=True
5. **IntegerField()，SlugField(), URLField(), BooleanField()**
    - 可设置blank=True or null=True
    - 对于BooleanField一般设置default=True or False
6. **FileField(upload_to=None, max_length=100)-文件字段**
    - <font color='orange'>upload_to="/some folder/"</font>
    - max_length = xxxx
7. **ImageField(upload_to=None, height_field=None, width_field=None, max_length=100)**
    - <font color='orange'>upload_to="/some folder/"</font>
    - 其他选项是可选的
8. **ForeignKey(to, on_delete, **options)-单对多关系**
    - <font color='orange'>to必须直线其他模型</font>
    - 必须指定<font color="orange">on_delete</font>
      - CASCADE：删除级联
      - SET_NULL：设置为null，必须设置`null=True`
      - SET()：设置一个值
      - SET_DEFAULT：设置默认值
      - DO_NOTHING：不做任何操作
9.  **ManyToManyField(to, **options)-多对多关系**
    - to 必需指向其他模型，比如 User or 'self' .
    - 设置 "symmetrical = False " 若 多对多关系不是对称的
    - 设置 "through = 'intermediary model' " 如果需要建立中间模型来搜集更多信息
    - 可以设置 "related_name = xxx" 便于反向查询。

<font color='red'>注意：</font>一旦设置了<font color="orange">related name</font>，你将不能再通过<font color="orange">_set</font>方法来反向查询

## 常见的Django Model META类选项

```python
# models.py
from django.db import models

class Meta:
    # 按Priority降序, order_date升序排列.
    get_latest_by = ['-priority', 'order_date']
    # 自定义数据库里表格的名字
    db_table = 'music_album'
    # 按什么排序
    ordering = ['pub_date']
    # 定义APP的标签
    app_label = 'myapp'
    # 声明此类是否为抽象
    abstract = True
    # 添加授权
    permissions = (("can_deliver_pizzas", "Can deliver pizzas"),)
```
