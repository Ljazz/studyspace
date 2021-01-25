<!-- TOC -->

- [自定义字段属性和错误信息](#自定义字段属性和错误信息)
- [自定义表单输入的widget和css属性](#自定义表单输入的widget和css属性)
- [表单数据初始化和实例化](#表单数据初始化和实例化)
- [自定义表单验证validators](#自定义表单验证validators)
- [一个页面同时提交2张或多张表单](#一个页面同时提交2张或多张表单)

<!-- /TOC -->

## 自定义字段属性和错误信息

对于每个字段你可以设置其是否为必需，最大长度和最小长度。还可以针对每个属性自定义错误信息。

```python
from django import forms


class LoginForm(forms.Form):
    username = forms.CharField(
        required=True,
        max_length=20,
        min_length=6,
        error_messages={
            'required': '用户名不能为空',
            'max_length': '用户名长度不得超过20个字符',
            'min_length': '用户名长度不得少于6个字符',
        }
    )
    password = forms.CharField(
        required=True,
        max_length=20,
        min_length=6,
        error_messages={
            'required': '密码不能为空',
            'max_length': '密码长度不得超过20个字符',
            'min_length': '密码长度不得少于6个字符',
        }
    )
```

对于继承ModelForm类的表单，可以在<font color='orange'>Meta</font>选项下<font color='orange'>widget</font>中来自定义错误信息。

```python
from django.forms import ModelForm, Textarea
from myapp.models import Author


class AuthorForm(ModelForm):
    class Meta:
        model = Author
        fields = ('name', 'title', 'birth_date')
        widgets = {
            'name': Textarea(attrs={'cols': 80, 'rows': 20})
        }
        labels = {
            'name': 'Author',
        }
        help_texts = {
            'name': 'Some useful help text.',
        }
        error_messages = {
            'name': {
                'max_length': "This writer's name is too long.",
            },
        }
```

## 自定义表单输入的widget和css属性

Django forms的每个字段都可以选择喜欢的输入widget，比如多选，复选框。还可以定义每个widget的css属性。如果不指定，Django会默认使用widget。

```python
from django import forms

class ContactForm(forms.Form):
    name = forms.CharField(
        max_length=255,
        widget=forms.Textarea(
            attrs=('class': 'custom')
        )
    )
```

设置widget可以使表单的美化程度大大提升，方便用户选择输入。例如下面对年份使用<font color='orange'>SelectDateWidget</font>，复选框可以用<font color='orange'>CheckboxSelectMultiple</font>，单选框可以使用<font color='orange'>RadioSelect</font>和<font color='orange'>Select</font>。常见文本输入可以用<font color='orange'>TextInput</font>和<font color='orange'>TextArea</font>

```python
from django import forms

BIRTH_YEAR_CHOICES = ('1980', '1981', '1982')
COLORS_CHOICES = (
    ('blue', 'Blue'),
    ('green', 'Green'),
    ('black', 'Black'),
)

class SimpleForm(forms.Form):
    birth_year = forms.DateField(widget=forms.SelectDateWidget(years=BIRTH_YEAR_CHOICES))
    favorite_colors = forms.MultipleChoiceField(
        required=False,
        widget=forms.CheckboxSelectMultiple,
        choices=COLORS_CHOICES,
    )
```

## 表单数据初始化和实例化

在有些场景中我们需要对表单设置一些初始数据，这一问题可以通`initial`方法解决。

```python
form = ContactForm(
    initial={
        'name': 'First and Last Name',
    },
)
```
其编辑修改类的场景中，我们还要给表单提供现有对象的数据，而不是渲染一张空表单。这个过程叫做<font color='red'>表单与数据的结合</font>

对于由继承<font color='orange'>ModelForm</font>类的表单，可以按照下面的方法进行示例化

```python
contact = Contact.objects.get(id=1)
form = ContactForm(instance=contact)
```

对于自定义的表单，可以设置<font color='orange'>default_data</font>。对于用户提交的数据，括号里可以使用request.POST。

```python
default_data = {'name': 'Jhon', 'email': 'someone@hotmail.com'}
form = ContactForm(default_data)
```

## 自定义表单验证validators

对于表单验证除了通过clean方法自定义外，还可以选择自定义validators。

```python
from django.core.exceptions import ValidationError
import re


def mobile_validate(value):
    mobile_re = re.compile(r'^(13[0-9]|15[012356789]|17[678]|18[0-9]|14[57])[0-9]{8}$')
    if not mobile_re.match(value):
        raise validationError('手机号格式错误')


class UserInfo(forms.Form):
    email = forms.EmailField(error_messages={'required': '邮箱不能为空'},)
    mobile = forms.CharField(validators=[mobile_validate,], error_messages={'required': '手机不能为空'}, widget=froms.TextInput(attrs={'class': 'form-control', 'placeholder': '手机号码'}))
```

## 一个页面同时提交2张或多张表单

很多场景中我们希望用户在同一页面上点击一个按钮同时提交2张或多张表单，这时我们可以在模板中给每个表单取不同对的名字。

```html
<form>
    {{ form1.as_p }}
    {{ form2.as_p }}
</form>
```

用户点击提交之后，我们可以在视图里对用户提交的数据分别进行处理

```python
if request.method == 'POST':
    form1 = Form1(request.POST, prefix='form1')
    form2 = Form2(request.POST, prefix='form2')

    if form1.is_valid() or form2.is_valid():
        pass
    else:
        form1 = Form1(prefix='form1')
        form2 = Form2(prefix='form2')
```
