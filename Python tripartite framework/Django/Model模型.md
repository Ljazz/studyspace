<!-- TOC -->

- [1. Model模型](#1-model模型)
  - [1.1. 、什么是Model模型](#11-什么是model模型)
  - [1.2. 、字段](#12-字段)
    - [1.2.1. 、字段类型](#121-字段类型)
    - [1.2.2. 、字段选项](#122-字段选项)
    - [1.2.3. 、关联关系字段](#123-关联关系字段)
      - [1、ForeignKey](#1foreignkey)
      - [2、ManyToManyField](#2manytomanyfield)
      - [3、OneToOneField](#3onetoonefield)
    - [1.2.4 跨文件模型](#124-跨文件模型)
    - [1.2.5 字段命名限制](#125-字段命名限制)
    - [1.2.6 自定的字段类型](#126-自定的字段类型)
      - [1、`Field.__init__()`方法接收以下参数](#1field__init__方法接收以下参数)
      - [2、字段解析](#2字段解析)
      - [自定义字段编写文档](#自定义字段编写文档)
      - [实用方法](#实用方法)
      - [通用建议](#通用建议)
  - [1.3 Meta选项](#13-meta选项)
    - [可用的选项](#可用的选项)
  - [1.4 模型属性](#14-模型属性)
  - [1.5 模型方法](#15-模型方法)
    - [1.5.1 重写之前定义的模型方法](#151-重写之前定义的模型方法)
    - [1.5.2 执行自定义SQL](#152-执行自定义sql)
  - [1.6 模型继承](#16-模型继承)
    - [1.6.1 抽象基类](#161-抽象基类)
    - [1.6.2 多表继承](#162-多表继承)
    - [1.6.3 代理模型](#163-代理模型)
    - [1.6.4 多表继承](#164-多表继承)
    - [1.6.5 不能用字段名“hiding”](#165-不能用字段名hiding)
  - [在包中管理模型](#在包中管理模型)

<!-- /TOC -->

# 1. Model模型

## 1.1. 、什么是Model模型

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

模型在使用的时候，需要将模型所在的`app`添加到`setting.py`文件中的`INSTALLED_APPS`中，然后通过命令`python manage.py makemigrations`和`python manage.py migrate`。可以将模型同步到数据库中。

## 1.2. 、字段

### 1.2.1. 、字段类型

模型中每个字段都应该是某个`Field`类的实例。这些字段类的作用有
- 字段类型用以指定数据库数据类型

<table>
<thead>
        <th>字段类型</th>
        <th>定义</th>
        <th>描述</th>
        <th>额外必填参数说明</th>
    </thead>
<tbody>
<tr>
<td width="10%">AutoField</td>
<td width="30%">class AutoField(**options)</td>
<td width="40%">一个IntegerField，它根据可用的ID自动递增。 您通常不需要直接使用它； 如果没有另外指定，主键字段将自动添加到模型中</td>
<td></td>
</tr>
<tr>
<td width="10%">BigAutoField</td>
<td width="30%">class BigAutoField(**options)</td>
<td width="40%">64位整数，非常类似于AutoField，不同之处在于，它可以容纳1到9223372036854775807之间的数字。</td>
<td></td>
</tr>
<tr>
<td width="10%">BigIntegerField</td>
<td width="30%">class BigIntegerField(**options)</td>
<td width="40%">64位整数，非常类似IntegerField，不同之处在于，它可以容纳-9223372036854775808至9223372036854775807之间的数字。</td>
<td></td>
</tr>
<tr>
<td width="10%">BinaryField</td>
<td width="30%">class BinaryField(max_length=None, **options)</td>
<td width="40%">一个用于存储原始二进制数据的字段，可以存储bytes、bytearray、memoryview类型数据。默认情况下，该字段的editable设置为False，在这种情况下，它不能包含在ModelForm中。</td>
<td>max_length：字段的最大长度（以字符为单位）</td>
</tr>
<tr>
<td width="10%">BooleanField</td>
<td width="30%">class BooleanField(**options)</td>
<td width="40%">布尔类型字段，当未定义default时，默认值为None。</td>
<td></td>
</tr>
<tr>
<td width="10%">CharField</td>
<td width="30%">class CharField(max_length=None, **options)</td>
<td width="40%">定义字符型字段</td>
<td>max_length：字段的最大长度（以字符为单位）</td>
</tr>
<tr>
<td width="10%">DateField</td>
<td width="30%">class DateField(auto_now=False, auto_now_add=False, **options)</td>
<td width="40%">日期字段，Python中由datetime.date实例表示。</td>
<td>
<li>auto_now：每次保存对象时自动设置为现在，只有在Model.save()时自动更新，其他方式(QuerySet.update())对其他字段进行更新时，该字段不会更新。</li>
<li>auto_now_add：首次创建对象时，将字段自动设置为现在。</li>
</td>
</tr>
<tr>
<td width="10%">DateTimeField</td>
<td width="30%">class DateTimeField(auto_now=False, auto_now_add=False, **options)</td>
<td width="40%">日期和时间字段，Python中由datetime.datetime实例表示。</td>
<td>
<li>auto_now：每次保存对象时自动设置为现在，只有在Model.save()时自动更新，其他方式(QuerySet.update())对其他字段进行更新时，该字段不会更新。</li>
<li>auto_now_add：首次创建对象时，将字段自动设置为现在。</li>
</td>
</tr>
<tr>
<td width="10%">DecimalField</td>
<td width="30%">class DecimalField(max_digits=None, decimal_places=None, **options</td>
<td width="40%">固定精度的十进制数字，在Python中由Decimal实例表示。它使用DecimalValidator。</td>
<td>
<li>max_digits：数字中允许的最大位数。请注意，此数字必须大于或等于decimal_places。</li>
<li>decimal_places：用数字存储的小数位数。</li>
</td>
</tr>
<tr>
<td width="10%">DurationField</td>
<td width="30%">class DurationField(**options)</td>
<td width="40%">用于存储时间段的字段-使用timedeltaPython中建模。在PostgreSQL上使用时，使用的数据类型是间隔，在Oracle上使用的数据类型是INTERVAL DAY (9) TO SECOND (6)</td>
<td></td>
</tr>
<tr>
<td width="10%">EmailField</td>
<td width="30%">class EmailField(max_length=254, **options)</td>
<td width="40%">邮件字段，本质上还是CharField，使用EmailValidator核查该值是否是有效的电子邮件地址。</td>
<td></td>
</tr>
<tr>
<td width="10%">FileField</td>
<td width="30%">class FileField(upload=None, max_length=100, **options</td>
<td width="40%">文件上传字段。不支持primary_key参数。</td>
<td>
<dl>
<dt>upload_to：此属性提供了一种设置上传目录和文件名的方法，并且可以通过两种方法进行设置。在这两种情况下，该值都将传递到Storage.save()方法。如果指定字符串值或路径，
则它可能包含strftime()格式，该格式将替换为文件上传的日期/时间。如果使用默认的FileSystemStorage，则字符串值将附加到MEDIA_ROOT路径中，以形成本地文件系统上将存储上载文件的位置。也可以使用其他存储形式。
upload_to也可以是可调用的，例如函数。这将被调用以获得上载路径，包括文件名。此可调用对象必须接收两个参数：</dt><dd>instance：定义FileField的模型实例。 更具体地说，这是附加当前文件的特定实例。</dd>
<dd>filename：最初提供给文件的文件名。 确定最终目标路径时，可以考虑也可以不考虑。</dd>
<dt>
<pre>
def user_directory_path(instance, filename):
    # file will be upload to MEDIA_ROOT/user_{id\filename}
    return 'user_{0}/{1}'.format(instance.user.id, filename)
class MyModel(models.Model):
    upload = models.FileField(upload_to=user_directory_path)
</pre>
</dt>
</dl>
</td>
</tr>
<tr>
<td width="10%">FilePathField</td>
<td width="30%">class FilePathField(path='', match=None, recursive=False, allow_files=True, allow_folders=False, max_length=100, **options)</td>
<td width="40%">一个CharField，其选择仅限于文件系统上某个目录中的文件名。</td>
<td>
<li>FilePathField.path：必选参数。目录的绝对文件系统路径，应从中选择此FilePathField。path也可以是可调用的，例如在运行时动态设置路径的函数。</li>
<li>
<pre>
import os
from django.conf import settings
from django.db import models

def images_path():
    return os.path.join(settings.LOCAL_FILE_DIR, 'images')

class MyModel(models.Model):
    file = models.FilePathField(path=images_path)
</pre>
</li>
<li>FilePathField.match：可选的。 FilePathField将用于过滤文件名的正则表达式（作为字符串）。 请注意，正则表达式将应用于基本文件名，而不是完整路径。 例如：“ foo。* \。txt $”，它将与名为foo23.txt的文件匹配，但与bar.txt或foo23.png的文件不匹配。</li>
<li>FilePathField.recursive：可选的。 正确或错误。 默认值为False。 指定是否应包含路径的所有子目录</li>
<li>FilePathField.allow_files：可选的。 正确或错误。 默认值为True。 指定是否应包含指定位置的文件。 这个或allow_folders必须为True。</li>
<li>FilePathField.allow_folders：可选的。 正确或错误。 默认值为False。 指定是否应包含指定位置的文件夹。 这个或allow_files必须为True。</li>
<li>匹配适用于基本文件名，而不是完整路径.</li>
</td>
</tr>
<tr>
<td width="10%">FloatField</td>
<td width="30%">class FloatField(**options)</td>
<td width="40%">Python中由float实例表示的浮点数。当本地化为False或TextInput时，此字段的默认表单小部件为NumberInput；否则为TextInput</td>
<td></td>
</tr>
<tr>
<td width="10%">ImageField</td>
<td width="30%">class ImageField(upload_to=None, height_field=None, width_field=None, max_length=100, **options)</td>
<td width="40%">从FileField继承所有属性和方法，但还验证上载的对象时有效的图像。除了FileField可用的特殊属性外，ImageField还具有height和width属性。</td>
<td>
<li>ImageField.height_field：每次保存模型实例时，模型字段的名称都会自动填充图像的高度</li>
<li>ImageField.width_field：每次保存模型实例时，模型字段的名称都会自动填充图像的宽度。</li>
</td>
</tr>
<tr>
<td width="10%">IntegerField</td>
<td width="30%">class IntegerField(**options)</td>
<td width="40%">一个整数。在Django支持的所有数据库中，-2147483648到2147483647之间的值都是安全的。它使用MinValueValidator和MaxValueValidator根据默认数据库支持的值来验证输入。当本地化为False时，此字段的默认表单小部件为NumberInput，否则为TextInput。</td>
<td></td>
</tr>
<tr>
<td width="10%">GenericIPAddressField</td>
<td width="30%">class GenericIPAddressField(protocol='both', unpack_ipv4=False, **options)</td>
<td width="40%">字符串格式的IPv4或IPv6地址（例如192.0.2.30或2a02：42fe :: 4）。 此字段的默认表单窗口小部件是TextInput。</td>
<td>
<li>GenericIPAddressField.protocol：将有效输入限制为指定的协议。 可接受的值为“ both”（默认），“ IPv4”或“ IPv6”。 匹配不区分大小写。</li>
<li>GenericIPAddressField.unpack_ipv4：解压缩IPv4映射的地址，例如:: FFFF：192.0.2.1。 如果启用此选项，则该地址将解压缩为192.0.2.1。 默认设置为禁用。 只能在协议设置为“ both”和“ both”时使用。</li>
</td>
</tr>
<tr>
<td width="10%">NullBooleanField</td>
<td width="30%">class NullBooleanField(**options)</td>
<td width="40%">像BooleanField一样为null = True。 使用该字段代替该字段，因为在将来的Django版本中可能会不推荐使用该字段。</td>
<td></td>
</tr>
<tr>
<td width="10%">PositiveIntegerField</td>
<td width="30%">class PositiveIntegerField(**options)</td>
<td width="40%">类似于IntegerField，但必须为正数或零（0）。 在Django支持的所有数据库中，0到2147483647之间的值都是安全的。 出于向后兼容的原因，可接受值0。</td>
<td></td>
</tr>
<tr>
<td width="10%">PositiveSmallIntegerField</td>
<td width="30%">class PositiveSmallIntegerField(**options)</td>
<td width="40%">类似于PositiveIntegerField，但只允许特定点（与数据库有关）下的值。 在Django支持的所有数据库中，0到32767的值都是安全的。</td>
<td></td>
</tr>
<tr>
<td width="10%">SlugField</td>
<td width="30%">class SlugField(max_length=50, **options)</td>
<td width="40%">Slug是一个报纸术语。 Slug是某事物的简短标签，仅包含字母，数字，下划线或连字符。 它们通常在URL中使用。像CharField一样，您可以指定max_length。 如果未指定max_length，则Django将使用默认长度50。 意味着将Field.db_index设置为True。</td>
<td>
<li>SlugField.allow_unicode：如果为True，则该字段除了ASCII字母外还接受Unicode字母。 默认为False。</li>
</td>
</tr>
<tr>
<td width="10%">SmallAutoField</td>
<td width="30%">class SmallAutoField(**options)</td>
<td width="40%">与AutoField相似，但仅允许值处于一定限制（与数据库有关）之下。 在Django支持的所有数据库中，从1到32767的值都是安全的</td>
<td></td>
</tr>
<tr>
<td width="10%">SmallIntegerField</td>
<td width="30%">class SmallIntegerField(**options)</td>
<td width="40%">类似于IntegerField，但仅允许在特定点（与数据库有关）下的值。 从-32768到32767的值在Django支持的所有数据库中都是安全的。</td>
<td></td>
</tr>
<tr>
<td width="10%">TextField</td>
<td width="30%">class TextField(**options)</td>
<td width="40%">大文本字段。 此字段的默认表单窗口小部件是Textarea。如果指定max_length属性，它将反映在自动生成的表单字段的Textarea小部件中。 但是，它不是在模型或数据库级别强制执行的。 为此使用CharField。</td>
<td></td>
</tr>
<tr>
<td width="10%">TimeField</td>
<td width="30%">class TimeField(auto_now=False, auto_now_add=False, **options)</td>
<td width="40%">时间，在Python中由datetime.time实例表示。 接受与DateField相同的自动填充选项。此字段的默认表单窗口小部件为TimeInput。</td>
<td></td>
</tr>
<tr>
<td width="10%">URLField</td>
<td width="30%">class URLField(max_length=200, **options)</td>
<td width="40%">URL的CharField，由URLValidator验证。此字段的默认表单窗口小部件是URLInput。像所有CharField子类一样，URLField采用可选的max_length参数。 如果未指定max_length，则使用默认值200。</td>
<td></td>
</tr>
<tr>
<td width="10%">UUIDField</td>
<td width="30%">class UUIDField(**options)</td>
<td width="40%">用于存储通用唯一标识符的字段。 使用Python的UUID类。 在PostgreSQL上使用时，它存储在uuid数据类型中，否则存储在char（32）中。通用的唯一标识符是primary_key的自动字段的很好替代。 数据库不会为您生成UUID，因此建议使用default</td>
<td>
<pre>
import uuid
from django.db import models

class MyUUIDModel(models.Model):
    id = models.UUIDField(primary_key=True, default=uuid.uuid4, editable=False)
</pre>
</td>
</tr>
</tbody>
</table> 

**FileField和FieldFile**

1、FileField

FileField.storage：一个存储对象，用于处理文件的存储和检索。此字段默认表单窗口小部件是ClearableFileInput。

在模型中使用`FileField`或`ImageField`步骤：
1. 在settings文件中，需要定义：MEDIA_ROOT作为Django存储上传文件目录的完整路径；MEDIA_URL作为该目录的基本公共URL，确保该目录能够被Web服务器的用户写入
2. 在模型中添加`FileField`或`ImageField`，并定义upload_to选项，以指定要用于上传文件的MEDIA_ROOT子目录。
3. 将存储在数据库中的所有文件都是该文件的路径（相对于MEDIA_ROOT）。在使用的时，例如，如果您的ImageField称为mug_shot，则可以使用{{object.mug_shot.url}}在模板中获取图像的绝对路径。

2、FieldFile

当访问模型上的FileField时，将提供FieldFile的实例作为访问基础文件的代理。
- FieldFile.name：文件名，包括从关联FileField的Storage根目录开始的相对路径。
- FieldFile.size：基础Storage.size（）方法的结果。
- FieldFile.url：一个只读属性，通过调用基础Storage类的url（）方法来访问文件的相对URL。
- FieldFile.open(mode='rb')：以指定模式打开或重新打开与此实例关联的文件。与标准的Python open（）方法不同，它不返回文件描述符。由于在访问基础文件时会隐式打开基础文件，因此可能不需要调用此方法，除非将指针重置为基础文件或更改模式。
- FieldFile.close()：行为类似于标准的Python file.close（）方法，并关闭与此实例关联的文件。
- FieldFile.save(name, content save=True)：此方法采用文件名和文件内容，并将它们传递到该字段的存储类，然后将存储的文件与model字段关联。如果要手动将文件数据与模型上的FileField实例相关联，则使用save（）方法来保留该文件数据。 
  - 接受两个必需的参数：
    - name是文件名
    - content是包含文件内容的对象
    - 可选的save参数控制在更改与该字段关联的文件之后是否保存模型实例。默认为True。
    - content参数应该是django.core.files.File的实例，而不是Python的内置文件对象
```python
from django.core.files import File
# open an existing file using python's built-in open()
f = open('/path/to/hello.world')
myfile = File(f)
# 或者可以像这样从Python字符串构造一个
from django.core.files.base import ContentFile
myFile = ContentFile("hello world")
```
- FieldFile.delete(save=True)：删除与此实例关联的文件，并清除该字段上的所有属性。 注意：如果在调用delete（）时碰巧打开了文件，此方法将关闭文件。
    - 可选的save参数控制在删除与该字段关联的文件之后是否保存模型实例。 默认为True。
    - 请注意，删除模型时，不会删除相关文件。 如果您需要清理孤立的文件，则需要自己处理

### 1.2.2. 、字段选项

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
    name = models.CharField(blank=True, choices=MedalType.choices, max_length=10)
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

9、db_column

用于此字段的数据库列的名称。 如果未指定，Django将使用该字段的名称。

10、db_index

如果为True，将为此字段创建数据库索引。

11、db_tablespace

如果此字段已建立索引，则用于该字段的索引的数据库表空间的名称。 缺省值是项目的DEFAULT_INDEX_TABLESPACE设置（如果已设置）或模型的db_tablespace（如果有）。 如果后端不支持索引的表空间，则忽略此选项。

12、editable

如果为False，则该字段将不会显示在管理员或任何其他ModelForm中。在模型验证期间也将跳过它们。默认为True。

13、error_messages

使用error_messages参数可以覆盖该字段将引发的默认消息。 传递字典，其中包含与您要覆盖的错误消息相匹配的键。错误消息键包括null，空白，无效，invalid_choice，unique和unique_for_date。 这些错误消息通常不会传播到表单。

14、validators

要为此字段运行的验证器列表。

### 1.2.3. 、关联关系字段

#### 1、ForeignKey

定义：class ForeignKey(to, on_delete, **options)

多对一的关系。需要两个位置参数
- 被关联的类
- on_delete选项。

如果要创建一个递归关系，一个与其自身有多对一关系的对象，则使用`models.ForeignKey('self', on_delete=models.CASCADE)

若需要在尚未定义的模型上创建关系，则可以使用模型的名称，而不是模型对象本身：
```python
from django.db import models

class Car(models.Model):
    manufacturer = models.ForeignKey(
        'Manufacturer',
        on_delete=models.CASCADE,
    )


class Manufacturer(models.Model):
    pass
```

当模型被子类化为具体模型并且与抽象模型的app_label不相关时，可以解析在抽象模型上以这种方式定义的关系

```python
from django.db import models

class AbstractCar(models.Model):
    manufacturer = models.ForeignKey('Manufacturer', on_delete=models.CASCADE)

    class Meta:
        abstract = True


class Manufacturer(models.Model):
    pass


class Car(AbstractCar):
    pass
```

若要引用在另一个应用程序中定义的模型，可以显式指定带有完整应用程序标签的模型。

```python
from django.db import models

class Car(models.Model):
    manufacturer = models.ForeignKey(
        'production.Manufacturer',
        on_delete=models.CASCADE
    )
```

**参数**

<font color='skyblue'>1、**ForeignKey.on_delete**</font>
- CASCADE

  级联删除。Django模拟SQL约束ON DELETE CASCADE的行为，并删除包含ForeignKey的对象。未在关联模型上调用Model.delete()，但会为所有已删除对象发送`pre_delete`和`post_delete`信号。

- PROTECT：

通过引发ProtectedError(django.db.IntegrityError的子类)来防止删除被引用的对象。

- RESTRICT：

通过引发RestrictedError（django.db.IntegrityError的子类）来防止删除引用的对象。 与PROTECT不同，如果引用的对象还引用了在同一操作中通过CASCADE关系删除的另一个对象，则允许删除引用的对象。

```ipython
In [1]: artist_one = Artist.objects.create(name='artist one')

In [2]: artist_two = Artist.objects.create(name='artist two')

In [3]: album_one = Album.objects.create(artist=artist_one)

In [4]: album_two = Album.objects.create(artist=artist_two)

In [5]: song_one = Song.objects.create(artist=artist_one, album=album_one)

In [6]: song_two = Song.objects.create(artist=artist_one, album=album_two)

In [7]: Artist.objects.all()
Out[7]: <QuerySet [<Artist: Artist object (1)>, <Artist: Artist object (2)>]>

In [8]: Album.objects.all()
Out[8]: <QuerySet [<Album: Album object (1)>, <Album: Album object (2)>]>

In [9]: Song.objects.all()
Out[9]: <QuerySet [<Song: Song object (1)>, <Song: Song object (2)>]>

In [10]: album_one.delete()
---------------------------------------------------------------------------
RestrictedError                           Traceback (most recent call last)
<ipython-input-18-80d037d130de> in <module>
----> 1 album_one.delete()

E:\Program Files\Python38\lib\site-packages\django\db\models\base.py in delete(self, using, keep_parents)
    941
    942         collector = Collector(using=using)
--> 943         collector.collect([self], keep_parents=keep_parents)
    944         return collector.delete()
    945

E:\Program Files\Python38\lib\site-packages\django\db\models\deletion.py in collect(self, objs, source, nullable, collect_related, source_attr, reverse_dependency, keep_parents, fail_on_restricted)
    334                             restricted_objects[key] += objs
    335                 if restricted_objects:
--> 336                     raise RestrictedError(
    337                         'Cannot delete some instances of model %r because '
    338                         'they are referenced through restricted foreign keys: '

RestrictedError: ("Cannot delete some instances of model 'Album' because they are referenced through restricted foreign keys: 'Song.album'.", <itertools.chain object at 0x000002099AC5E6D0>)

In [11]: artist_two.delete()
---------------------------------------------------------------------------
RestrictedError                           Traceback (most recent call last)
<ipython-input-33-59bc5290b4a3> in <module>
----> 1 artist_two.delete()

E:\Program Files\Python38\lib\site-packages\django\db\models\base.py in delete(self, using, keep_parents)
    941 
    942         collector = Collector(using=using)
--> 943         collector.collect([self], keep_parents=keep_parents)
    944         return collector.delete()
    945

E:\Program Files\Python38\lib\site-packages\django\db\models\deletion.py in collect(self, objs, source, nullable, collect_related, source_attr, reverse_dependency, keep_parents, fail_on_restricted)
    334                             restricted_objects[key] += objs
    335                 if restricted_objects:
--> 336                     raise RestrictedError(
    337                         'Cannot delete some instances of model %r because '
    338                         'they are referenced through restricted foreign keys: '

RestrictedError: ("Cannot delete some instances of model 'Artist' because they are referenced through restricted foreign keys: 'Song.album'.", <itertools.chain object at 0x000002099ACAC5E0>)       


In [12]: artist_one.delete()
Out[12]: (4, {'app_model.Song': 2, 'app_model.Album': 1, 'app_model.Artist': 1})

In [13]: Artist.objects.all()
Out[13]: <QuerySet [<Artist: Artist object (2)>]>

In [14]: Album.objects.all()
Out[14]: <QuerySet [<Album: Album object (2)>]>

In [15]: Song.objects.all()
Out[15]: <QuerySet []>
```

- SET_NULL：

将ForeignKey设置为null；仅当null为True时，才有可能。

- SET_DEFAULT:

将ForeignKey设置为其默认值； 必须为ForeignKey设置默认值。

- SET()：

将ForeignKey设置为传递给SET()的值，或者如果传递了可调用对象，则为调用它的结果。在大多数情况下，有必要传递一个callable，以避免在导入models.py时执行查询：

```python
from django.conf import settings
from django.contrib.auth import get_user_model
from django.db import models

def get_sentinel_user():
    return get_user_model().objects.get_or_create(username="deleted")[0]


class MyModel(models.Model):
    user = models.ForeignKey(
        settings.AUTH_USER_MODEL,
        on_delete=models.SET(get_sentinel_user)
    )
```

- DO_NOTHING：

不采取行动。 如果您的数据库后端强制执行参照完整性，则除非您手动将SQL ON DELETE约束添加到数据库字段，否则这将导致IntegrityError。

<font color='skyblue'>2、**ForeignKey.limit_choices_to**</font>

使用ModelForm或admin呈现此字段时，为该字段的可用选项设置限制（默认情况下，可以选择queryset中的所有对象）。可以使用字典、Q对象或可返回字典或Q对象的可调用对象。

```python
from django.db import models

staff_member = models.ForeignKey(
    "User",
    on_delete=models.CASCADE,
    limit_choices_to={'is_staff': True}
)
```

<font color="skyblue">3、**ForeignKey.related_name**</font>

用于相关对象到此对象的关系的名称。它也是related_query_name的默认值（用于目标模型的反向过滤器名称的名称）

<font color="skyblue">4、**ForeignKey.related_query_name**</font>

用于目标模型的反向过滤器名称的名称。如果设置，则默认为related_name或default_related_name的值，否则默认为模型的名称

<font color="skyblue">5、**ForeignKey.to_field**</font>

关系所涉及的相关对象上的字段。 默认情况下，Django使用相关对象的主键。 如果您引用其他字段，则该字段必须具有unique = True。

#### 2、ManyToManyField

定义：`class ManyToManyField(to, **options)`

多对多关系。需要一个位置参数：与模型相关的类，其作用与对ForeignKey完全相同，包括递归和惰性关系。

**参数**：

- ManyToManyField.symmetrical

  定义自身为多对多关系时，如下

```python
from django.db import models

class Person(models.Model):
    friends = models.ManyToManyField("self")
```
若不希望与自己的多对多关系具有对称性，可以将`ManyToManyField.symmetrical`的值设置为False。

- ManyToManyField.through
  
  自动生成一个表来管理多对多关系。若想要手动指定中间表，则可以使用through选项来指定表示要使用的中间表的Django模型。

- ManyToManyField.through_fields

  仅在明确给出一个自定义中间模型时使用。django将会正常地决定使用中间模型的那些字段来自动地建立多对多关系。

```python
from django.db import models

class Person(models.Model):
    name = models.CharField(max_length=40)

class Group(models.Model):
    name = models.CharField(max_length=128)
    members = models.ManyToManyField(
      Person,
      through="MemberShip",
      through_fields=('group', 'person'),
    )

class MemberShip(models.Model):
    group = models.ForeignKey(Group, on_delete=models.CASCADE)
    person = models.ForeignKey(Person, on_delete=models.CASCADE)
    inviter = models.ForeignKey(
      Person,
      on_delete=models.CASCADE,
      related_name='membership_invites'
    )
    invite_reason = models.CharField(max_length=64)
```
Membership 有*两个*外键``Person``（`person`和`inviter`），这使得联接关系显的歧义，django并不知道使用哪一个。这个例子中，你必须显式的指定Django应该使用哪个外键， 通过`through_fields`， 就想上例中一样。

`through_fields`接受一个2元组（`fiels1，fields2`）,其中``fields1``是多对多关系字段被定义的那个模型的外键名称（本例中的``group``）， `fields2`是目标模型的外键名称（本例中的`person`）

MemberShi当中介表中有不止一个外键指向任一（或均）参与多对多关系的模型时，*必须*指定`through_fields`。

#### 3、OneToOneField

定义：`class OneToOneField(to, on_delete, parent_link=False, **options)`

一对一的关系。类似于`unique=True`的ForeignKey，但是关系的“反向”则将直接返回的单个对象。

一般应用于“扩展”一个模型的主键。若没有为OneToOneField指定related_name参数，则django将适用当前模型的小写名称作为默认值。

```python
from django.db import models
from django.conf import settings

class MySpecialUser(models.Model):
    user = models.OneToOneField(
      settings.AUTH_USER_MODEL,
      on_delete=models.CASCADE
    )
    supervisor = models.OneToOneField(
      settings.AUTH_USER_MODEL,
      on_delete=models.CASCADE,
      related_name='superviso_of'
    )
```
如果相关表中的条目不存在，则在访问反向关系时会引发RelatedObjectDoesNotExist异常。 这是目标模型的Model.DoesNotExist异常的子类。

### 1.2.4 跨文件模型

关联另一个应用的模型也是可以的，但是在使用前需要导入被关联的模型，或者也可以使用`应用.模型名`。

```python
from django.db import models
from geography.models import ZipCode

class Restaurant(models.Model):
    zip_code = models.ForeignKey(
      ZipCode,
      on_delete=models.SET_NULL,
      blank=True,
      null=True
    )
```
或者
```python
from django.db import models

class Restaurant(models.Model):
    zip_code = models.ForeignKey(
      'geography.ZipCode',
      on_delete=models.SET_NULL,
      blank=True,
      null=True
    )
```

### 1.2.5 字段命名限制

Django对模型的字段名有一些限制
1. 一个字段的名称不能是Python保留字，因为会导致Python语法错误
2. 一个字段名称不能包含连续的多个下划线，原因在于django查询语法的工作方式
3. 字段名不能以下划线结尾

### 1.2.6 自定的字段类型

所有的Django模型字段都是`django.db.models.Field`的子类。对于所有的模型字段来说，Django记录的大部分信息是一样的。存储的行为由`Field`处理。

```python
from django.db import models

class HandField(models.Field):
    description = "A hand of cards (bridge style)"

    def __init__(self, *args, **kwargs):
        kwargs['max_length'] = 104
        super().__init__(*args, **kwargs)
```

#### 1、`Field.__init__()`方法接收以下参数
- verbose_name
- name
- primary_key
- max_length
- unique
- blank
- null
- db_index
- rel：用于关联字段（像ForeignKey）。仅用于进阶用途
- default
- editable
- serialize：若为False，字段传给Django的序列化器时不会被序列化。默认为True
- unique_for_date
- unique_for_month
- unique_for_year
- choices
- help_text
- db_column
- db_tablespace：仅为创建索引，如果后端支持tablespaces。一般情况下可以忽略此选项。
- auto_created：若字段是自动创建的；则为True，用于OneToOneField的模型继承。

#### 2、字段解析

`deconstruct()`方法在模型迁移过程中会告诉Django如何获取你的薪资u但的一个实例，并将其转换为序列化形式，特别是，传递什么参数给`__init__()`来创建它。

如果未在继承的字段之前添加任何选项，就不需要编写新的`deconstruct()`方法。若修改了传递给`__init__()`的参数（像HandField中一样），需要增补被传递的值。

`deconstruct()` 返回包含 4 个项目的元组
- 字段的属性名
- 字段类的完整导入路径
- 位置参数（以列表的形式）
- 关键字参数（以字典的形式）

注意，这与 为自定义类 的 `deconstruct()` 方法不同，它返回包含 3个项目的元组。并不需要担心前面两个参数值;基类`Field`已包含处理字段属性名和导入路径的代码。我们应该关注位置参数和关键字参数，这些是我们有可能改的东西。

如上例`HandField`类中，总是强制设置`max_length`在__init__()`方法中。基类`Field`中的**deconstruct()**方法会注意到这个值，并尝试在关键字参数中返回它。

```python
from django.db import models


class HandField(models.Model):
    def __init__(self, *args, **kwargs):
        kwargs['max_length'] = 120
        super().__init__(*args, **kwargs)

    def deconstruct(self):
        name, path, args, kwargs = super().deconstruct()
        del kwargs['max_length']
        return name, path, args, kwargs
```
若添加了一个新的关键字参数，需要在**deconstruct**中新增加代码，将其传入kwargs。

```python
from django.db import models


class CommaSepField(models.Model):
    """Implements comma-separated storage of lists"""
    def __init__(self, separator="", *args, **kwargs):
        self.separator = separator
        super().__init__(*args, **kwargs)

    def deconstruct(self):
        name, path, args, kwargs = super().deconstruct()
        # Only include kwarg if it's not the default
        if self.separator != ",":
            kwargs['separator'] = self.separator
        return name, path, args, kwargs
```
对于你的字段实例的任意配置，`deconstruct()` 必须返回能传递给 `__init__`的参数重构状态。

#### 自定义字段编写文档

除了为其提供对开发者很有用的 docstring 外，你也需要让后台管理员通过`django.contrib.admindocs`查看字段类型的简单介绍。为此，只需要在自定义字段的**description**属性提供描述性文本。

#### 实用方法

在创建Field子类后，可能会考虑重写一些标准方法，这取决于字段的行为。

1、自定义数据库类型

假设已创建一个PostgreSQL自定义字段，名为mytype，可以继承Field并实现`db_type()`方法

```python
from django.db import models

class MyTypeField(models.Field):
    def db_type(self, connection):
        return "mytype"
```

只要已经建立MyTypeField，就可以像实用其他Field类型一样在模型中实用它

```python
from django.db import models

class Person(models.Model):
    name = models.CharField(max_length=120)
    something_else = MyTypeField()
```

若想创建一个数据库来源不明确的应用时，则考虑各个数据库之间的不同之处。比如，date/time，在PostgreSQL中为`timestamp`，在MySQL中则为`datetime`。可以通过`db_type()`方法来检查`connection.settings_dict['ENGINE']`属性来进行操作。

```python
from django.db import models


class MyDateField(models.Field):
    def db_type(self, connection):
        if connection.settings_dict['ENGINE'] == 'django.db.backends.mysql':
            return 'datetime'
        else:
            return 'timestamp'
```

`db_type()`和`rel_db_type()`方法由Django框架在为应用构建`CREATE TABLE`语句时调用，即第一次创建数据表时。这些方法也在构建一个包含此模型字段的`WHERE`子句时调用，即在利用QuerySet方法（get、filter、和exclude）检出数据时或将此模型字段作为参数时。它们在其它时间不会被调用。

某些数据库列类型接会接收参数，例如`CHAR(25)`,参数25表示列的最大长度。类似用例中，该参数若在模型中指定比硬编码在`db_type()`方法中更灵活。

```python
# This is a silly example of hard-coded parameters.
class CharMaxlength25Fiedl(models.Field):
    def db_type(self, connection):
        return 'char(25)'

# In the Model
class MyModel(models.Model):
    my_field = CharMaxlength25Field()
```
更好的方式是在运行时指定参数值，即类实例化时。

```python
# This is a silly example of hard-coded parameters.
class BetterCharField(models.Field):
    def __init__(self, max_length, *args, **kwargs):
        self.max_length = max_length
        super().__init__(*args, **kwargs)

    def db_type(self, connection):
        return 'char({})'.format(self.max_length)

# In the Model
class MyModel(models.Model):
    my_field = BetterCharField(25)
```
若数据列要求配置复杂的SQL，从`db_type()`返回None。这会让Django创建SQL的代码跳过该字段。随后需要为该字段在正确的表中以某种方式创建列，这种方式允许你告诉django不处理此事。

`rel_db_type()`方法由字段调用，例如`ForeignKey`和`OneToOneField`，这些通过指向另一个字段来决定数据库列类型的字段。

```python
from django.db import models

class UnsignedAutoField(models.AutoField):
    def db_type(self, connection):
        return 'integer UNSIGNED AUTO_INCREMENT'

    def rel_db_type(self, connection):
        return 'integer UNSIGNED'
```

2、将值转为Python对象

若自定义`Field`处理的数据结构比字符串，日期或浮点数等更复杂，可能需要重写`from_db_value()`和`to_python()`。

若存在于字段子类中，则从数据库加载数据时，在所有情况下都将调用`from_db_value()`包括聚集调用和values()调用。

`to_python()`在反序列化和为表单应用`clean()`时调用

作为通用规则，`to_python`应该平滑的处理以下参数:
- 一个正确的类型
- 一个字符串
- None（若字段允许null=True）

```python
import re

from django.core.exceptions import ValidationError
from django.db import models
from django.utils.translation import gettext_lazy as _


def parse_hand(hand_string):
    """Takes a string of cards and splits into a full hand."""
    p1 = re.compile('.{26}')
    p2 = re.complie('..')
    args = [p2.findall(x) for x in p1.findall(hand_string)]
    if len(args) != 4:
    raise ValidationError(_("Invalid input for a Hand instance"))
    return Hand(*args)


class HandField(models.Field):
    def from_db_value(self, value, expression, connection):
        if value is None:
            return value
        return parse_hand(vlaue)
    
    def to_python(self, value):
        if isinstance(value, Hand):
            return value
        if value is None:
            return value
        return parse_hand(value)
```

3、将Python转为查询值

实用数据库需要双向转换，如果重写了`from_db_value()`方法，必须重写`get_prep_value()`将Python对象转回查询值。

```python
from django.db import models

class HandField(models.Field):
    def get_prop_value(self, value):
        return ''.join([''.join(l) for l in (value.north, value.east, value.south, value.west)])
```

4、将查询值转为数据库值

某些数据类型在数据库后端处理前要转为某种特定格式。`get_db_prep_value()`实现了这种转换。查询所使用的连接由`connection`参数指定。这允许你在需要时指定后台需求的转换逻辑。

例，Django中`BinaryField`
```python
def get_db_prep_value(self, value, connection, prepared=False):
    value = super().get_db_prep_value(value, connection, prepared)
    if value is not None:
        return connection.Database.Binary(value)
    return value
```
自定义字段需要与普通查询参数使用不同的转换规则，可以重写`get_db_prep_save()`

5、在保存前预处理数值

使用`pre_save()`。例如`DateTimeField`在`auto_now`或`auto_now_add`中利用此方法正确设置属性。

如果重写了此方法，必须在最后返回该属性的值。如果修改了值，那么也需要更新模型属性，这样持有该引用的模型总会看到正确的值。

6、为模型字段指定表单字段

为了自定义`ModelForm`使用的表单属性，必须重写`formfield()`。表单字段类能通过 form_class 和 choices_form_class 参数指定；如果字段指定了选项，则使用后者，反之前者。若未提供这些参数，将会使用 CharField 或 TypedChoiceField。

完整的 kwargs 被直接传递给表单字段的 `__init__()` 方法。一般的，你要做的全部工作就是为 form_class 参数配置一个合适的默认值，并在随后委托父类处理。这可能要求你编写一个自定义表单字段（甚至表单视图）。

```python
class HandField(models.Field):
    def formfield(self, **kwargs):
        # This is a fairly standard way to set up some defaults
        # while letting the caller override them.
        defaults = {'form_class': MyFormField}
        defaults.update(kwargs)
        return super().formfield(**defaults)
```

7、为序列化转换字段数据

自定义序列化器序列化值的流程，你要重写 value_to_string()。使用 value_to_string() 是在序列化之前获取字段值的最佳方法。举个例子，由于 HandField 使用字符串存储数据，我们能复用一些已有代码:

```python
class HandField(models.Field):
    # ...

    def value_to_string(self, obj):
        value = self.value_from_object(obj)
        return self.get_prep_value(value)
```

#### 通用建议

编写自定义字段是个棘手的，尤其是在 Python 类，数据库，序列化格式之间进行复杂转换的时候。下面有几个让事情更顺利的建议：

1. 借鉴已有的 Django 字段（位于 `django/db/models/fields/__init__.py`）。试着找到一个与你目标类似的字段，而不是从零开始创建。
2. 为字段类添加一个 `__str__()` 方法。在很多地方，字段代码的默认行为是对值调用 `str()`。所以 `__str()__` 方法会自动将 Python 对象转为字符串格式

## 1.3 Meta选项

使用内部Meta类来给模型赋予元数据

```python
from django.db import models

class Ox(models.Model):
    horn_length = models.IntegerField()

    class Meta:
        ordering = ["horn_length"]
        verbose_name_plural = "oxen"
```
### 可用的选项
| 选项 | 含义 |
| --- | --- |
| abstract | 如果abstract=True，则此模型将是抽象基类 |
| app\_label | 如果模型是在INSTALLED_APPS中的应用程序之外定义的，则它必须声明它属于那个应用程序：`app_label='myapp'`；若要使用`app_label.object_name`或`app_label.model_name`格式表示模型，则可以分别使用`model._meta.label`或`model._meta.label_lower`。 |
| base\_manager\_name | 管理器的属性名称，例如“对象”，用于模型的`_base_manager` |
| db\_table | 用于模型的数据库表的名称 |
| db\_tablespace | 用于该模型的数据库表空间的名称。默认设置是项目的DEFAULT_TABLESPACE设置（如果已设置）。如果后端不支持表空间，则忽略此选项。 |
| default\_manager\_name | 用于模型的`_default_manager`的管理器的名称 |
| default\_related\_name | 默认情况下，将从相关对象到该对象的关系使用的名称。默认值为<model_name>_set。此选项还设置`related_query_name`。由于字段的反向名称应该是唯一的，因此设置时应当谨慎。 |
| get\_latest\_by | 模型中的字段名称或字段名称列表，通常为DateField，DateTimeField或IntegerField。这指定了腰子模型管理器的`Latest()`和`earlyest()`方法中使用的默认字段。 |
| managed | 默认为True，这意味着Django将在迁移过程中或在迁移过程中创建适当的数据库表，并将其作为刷新管理命令的一部分删除。即，Django管理数据库表的生命周期。若为False，则不会对此模型执行数据库表创建，修改或删除操作。如果模型表示现有表或通过其他某种方式创建的数据库视图，则此功能很有用。当managed=False时，这是唯一的区别。模型处理的所有其他方面与正常情况完全相同。
| order\_with\_respect\_to | 使此对象相对于给定字段（通常为ForeignKey）可排序。者可用于使相关对象相对于父对象可排序。 |
| ordering | 对象的默认顺序，用于在获取对象列表时使用 |
| permissions | 创建此对象时可进入权限表的额外权限。将自动为每个模型创建添加，更改，删除和查看权限。 |
| default\_permissions | 默认为（“添加”，“更改”，“删除”，“查看”）。 您可以自定义此列表，例如，如果您的应用不需要任何默认权限，可以将其设置为空列表。 必须先在模型上指定它，然后才能通过迁移创建模型，以防止创建任何遗漏的权限。 |
| proxy | 如果proxy = True，则子类化另一个模型的模型将被视为代理模型。|
| required\_db\_features | 当前连接应具有的数据库功能列表，以便在迁移阶段考虑模型。 例如，如果将此列表设置为\['gis_enabled'\]，则该模型将仅在启用GIS的数据库上同步。 在使用多个数据库后端进行测试时，跳过某些模型也很有用。 避免模型之间的关系 |
| required\_db\_vendor | 该模型特定于的受支持数据库供应商的名称。当前内置的供应商名称是：sqlite，postgresql，mysql，oracle。如果该属性不为空并且当前连接供应商不匹配该模型，则该模型将不同步。 |
| select_on\_save | 确定Django是否将使用1.6之前的django.db.models.Model.save（）算法。旧算法使用SELECT来确定是否存在要更新的行。新算法直接尝试UPDATE。在极少数情况下，Django无法看到现有行的UPDATE。一个示例是PostgreSQL ON UPDATE触发器，该触发器返回NULL。在这种情况下，即使数据库中存在一行，新算法也将最终执行INSERT。通常不需要设置此属性。默认值为False。|
| indexes | 要在模型上定义的索引列表 |
| unique\_together | 字段名称集合必须是唯一的 |
| index\_together | 使用索引选项。 |
| constraints | 要在模型上定义的约束列表 |
| verbose\_name | 对象的人类可读名称，单数形式 |
| verbose\_name\_plural | 对象的复数名称 |

## 1.4 模型属性

模型当中最终的属性是 *Manager*。它是Django模型和数据库查询操作之间的接口，并且它被用作从数据库当中获取实例，如果没有指定自定义的 *Manager* ，默认名称是 **ojbects**.Manager只能通过模型类来访问，不能通过模型实例来访问。

## 1.5 模型方法

```python
from django.db import models


class Person(models.Model):
    first_name = models.CharField(max_length=50)
    last_name = models.CharField(max_length=50)
    birth_date = models.DateField()

    def baby_boomer_status(self):
        "Returns the perosn's baby-boomer status"
        import datetime
        if self.birth_date < datetime.date(1945, 8, 1):
            return "Pre-boomer"
        elif self.birth_date < datetime.date(1965, 1, 1):
            return "Baby boomer"
        else:
            return "Post-boomer"
    
    @property
    def full_name(self):
        "Returns the person's full name"
        return "%s %s"%(self.first_name, self.last_name)
```

**模型方法列表**
| 方法 | 描述 |
| --- | --- |
| 

### 1.5.1 重写之前定义的模型方法

### 1.5.2 执行自定义SQL

## 1.6 模型继承

### 1.6.1 抽象基类

### 1.6.2 多表继承

### 1.6.3 代理模型

### 1.6.4 多表继承

### 1.6.5 不能用字段名“hiding”

## 在包中管理模型
