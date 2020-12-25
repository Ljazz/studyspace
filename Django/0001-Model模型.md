# Model模型

## 1、什么是Model模型

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

## 2、字段

### 2.1、字段类型

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
<li>GenericIPAddressField.unpack_ipv4：解压缩IPv4映射的地址，例如:: ffff：192.0.2.1。 如果启用此选项，则该地址将解压缩为192.0.2.1。 默认设置为禁用。 只能在协议设置为“ both”和“ both”时使用。</li>
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
myfile = ContentFile("hello world")
```
- FieldFile.delete(save=True)：删除与此实例关联的文件，并清除该字段上的所有属性。 注意：如果在调用delete（）时碰巧打开了文件，此方法将关闭文件。
    - 可选的save参数控制在删除与该字段关联的文件之后是否保存模型实例。 默认为True。
    - 请注意，删除模型时，不会删除相关文件。 如果您需要清理孤立的文件，则需要自己处理

### 2.2、字段选项

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

### 2.3、关联关系字段

1、**ForeignKey**

定义：class ForeignKey(to, on_delete, **options)

多对一的关系。需要两个位置参数
- 被关联的类
- on_delete选项。

如果要创建一个递归关系，一个与其自身有多对一关系的对象，则使用`models.ForeignKey('self', on_delete=models.CASCADE)



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


