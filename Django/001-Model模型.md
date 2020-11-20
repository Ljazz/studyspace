# Model模型

## 什么是Model模型

Model(模型)就是数据模型。模型不是数据本身，而是抽象的描述数据的构成和逻辑关系。每个Django model实际上就是一个类，继承了models.Model。每个Model包含属性，关系和方法。当定义好model后，Django提供的一个接口会自动在数据库生成相应的数据表（table）。一般来说，每一个模型都映射一张数据库表。

```python
from django.db import models

class Musician(models.Model):
    first_name = models.CharField(max_length=50)
    last_name = models.CharField(max_length=50)
    instrument = models.CharField(max_length=100)

class Album(models.Model):
    artist = models.ForeignKey(Musician, on_delete=models.CASCADE)
    name = models.CharField(max_length=100)
    release_date = models.DateField()
    num_stars = models.IntegerField()
```

## 字段选项

一些通用参数

1、null

如果设置为`True`，当该字段为空时，Django会将数据库中该字段设置为`NULL`。默认为`False`。

2、blank

如果设置为`True`，该字段允许为空。默认为`False`。

与`null`的区别:

- `null`选项仅仅是数据库层次的设置。
- `blank`涉及到表单验证方面。

3、choices

一系列二元组，用作此字段的选项。如果提供了二元组，默认表单小部件是一个选择框，而不是标准的文本字段，并将限制给出的选项。

```python
YEAR_IN_SCHOOL_CHOICES = [
    ('FR', 'Freshman'),
    ('SO', 'Sophomore'),
    ('JR', 'Junior'),
    ('SR', 'Senior'),
    ('GR', 'Graduate'),
]
```

二元组的第一个值会存储在数据库中，第二个值会在表单中显示。对于模型实例，要获取该字段二元组中相对应的第二个值，使用`get_FOO_display()`方法。

```python
from django.db import models

class Person(models.Model):
    SHIRT_SIZES = (
        ('S', 'Small'),
        ('M', 'Medium'),
        ('L', 'Large'),
    )
    name = models.CharField(max_length=60)
    shirt_size = models.CharField(max_length=1, choices=SHIRT_SIZES)
```

```ipython
In [1]: p = Person(name="Fred Flintstone", shirt_size="L")
In [2]: p.save()
In [3]: p.shirt_size
Out[3]: 'L'
In [4]: p.get_shirt_size_display()
Out[4]: 'Large'
```

使用枚举方式定义**choices**

```python
from django.db import models

class Runner(models.Model):
    MedalType = models.TextChoices('MedalType', 'GOLD SILVER BRONZE')
    name = models.CharField(blank=True, choices=MedalType.choices, max_length=w10)
```

4、default

该字段的默认值。可以是一个值或者可调用的对象，如果是个可调用对象，每次实例化模型时都会调用该对象

5、help_text

额外的“帮助”文本，随表单控件一同显示。

6、primary_key

如果设置为True，将该字段设置为该模型的主键.若模型中没有对任何一个字段设置`primary_key=True`选项。Django会自动添加一个`IntegerField`字段，并设置为主键。

7、unique

如果设置为True，这个字段的值必须在整个表中保持唯一。

8、verbose_name

为改字段添加备注名称

## 字段类型

## 关联关系

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
