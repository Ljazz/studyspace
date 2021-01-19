<!-- TOC -->

- [什么是表单？表单何时使用？](#什么是表单表单何时使用)
- [表单实例化](#表单实例化)
- [模板文件中使用form](#模板文件中使用form)
- [表单实际使用案例](#表单实际使用案例)
- [表单的验证](#表单的验证)
- [通用视图里使用表单](#通用视图里使用表单)
- [自定义表单输入的widget](#自定义表单输入的widget)
- [表单数据初始化](#表单数据初始化)
- [Formset的使用](#formset的使用)

<!-- /TOC -->

## 什么是表单？表单何时使用？

类似模型，Django表单也由各种字段组成。表单可以自定义(forms.Form)，也可以由模型Model创建(forms.ModelForm)。注意：在模型里面用verbose_name来描述一个字段，表单中使用label

```python
from django import forms
from .models import Contact

class ContactForm1(forms.Form):
    name = forms.CharField(label="Your Name", max_length=255)
    email = forms.EmailField(label="Email address")


class ContactForm2(forms.ModelForm):
    class Meta:
        model = Contact
        fields = ('name', 'email', )
```

Django常用做法就是在app文件夹下创建一个forms.py，专门存放app所定义的各种表单，这样方便集中管理表单。如果要使用上述的表单，可以在views.py中直接import引入即可。

## 表单实例化

下面方法可以实例化一个空表单，但里面没有任何数据，可以通过 {{ form }} 在模板中渲染

```python
form = ContactForm()
```

用户提交的数据可以通过以下的方法与表单结合，生成与数据结合过的表单(Bound forms)。Django只能对Bound forms进行验证。
```python
form = ContactForm(data=request.POST, files=request.FILES)
```

## 模板文件中使用form

模板文件中我们可以通过{{ form.as_p }}，{{ form.as_li }}和{{ form.as_table}}中渲染表单。若想详细控制每个field的格式，可以采用如下方式

```html
{% block content %}
<div class="form-wrapper">
<form method="post" action="" enctype="multipart/form-data">
{% csrf_token %}
{% for field in form %}
<div class="fieldWrapper">
{{ field.errors }}
{{ field.label_tag }} {{ field }}
{% if field.help_text %}
<p class="help">{{ field.help_text|safe }}</p>
{% endif %}
</div>
{% endfor %}
<div class="button-wrapper submit">
<input type="submit" value="Submit" />
</div>
</form>
</div>
{% endblock %}
```

## 表单实际使用案例

设计一个用于用户注册的表单。在app目录下创建forms.py文件，然后再文件中创建一个RegistrationForm。

```python
from django import forms
from django.contrib.auth.models import User


class RegistrationForm(forms.Form):
    username = forms.CharField(label="Username", max_length=40)
    email = forms.EmailField(label="Email")
    password1 = forms.CharField(label="Password", widget=forms.PasswordInput)
    password2 = forms.CharField(label="Password Confirmation", widget=forms.PasswordInput)
```
使用forms.py的好处是：
- 所有的表单在一个文件里，非常便于后期维护，比如增添或修改字段
- forms.py可通过clean方法自定义表单验证，非常便捷

再views.py中使用RegistrationForm
```python
from django.shortcuts import render, get_object_or_404
from django.contrib.auth.models import User
from .forms import RegistrationForm
from django.http import HttpResponseRedirect


def register(request):
    if request.method == "POST":
        form = RegistrationForm(request.POST)
        if form.is_valid():
            username = form.cleaned_data['username']
            email = form.cleaned_data['email']
            password = form.cleaned_data['password2']
            # 使用内置User自带的create_user方法创建用户，不需要使用save()
            user = User.objects.create_user(username=username, password=password, email=email)
            # 如果直接使用objects.create()方法后不需要使用save()
            return HttpResponseRedirect("/accounts/login/")
    else:
        form = RegistrationForm()
    return render(request, "users/registration.html", {"form": form})
```
模板<font color='orange'>registration.html</font>内容如下，若需要通过表单上传图片或文件，一定不要忘了给form加<font color="orange">enctype="multipart/form-data"</font>

```html
<form action="." method="POST">
{{ form.as_p }}
</form>
```

<font color="orange">registrationForm</font>工作原理
- 当用户通过<font color="orange">POST</font>方法提交表单，我们将提交的数据与<font color="orange">RegistrationForm</font>结合，然后验证表单<font color="orange">RegistrationForm</font>的数据是否有效。
- 若表单数据有效，先使用Django User模型自带的<font color="orange">create_user</font>方法创建<font color="orange">user</font>对象，再创建<font color="orange">user_profile</font>。用户通过一张表单提交数据，实际上我们分别存储再两张表里。
- 如果用户注册成功，我们通过<font color="orange">HttpResponseRedirect</font>方法转到登录页面
- 如果用户没有提交表单或不是通过POST方法提交表单，我们转到注册页面，生成一张空的<font color="orange">RegistrationForm</font>

## 表单的验证

每个forms类可以通过<font color="orange">clean</font>方法自定义表单验证。若是对某些字段进行验证，可以通过<font color="orange">clean_字段名</font>方式自定义表单验证。如果用户提交数据未通过验证，会返回<font color="orange">ValidationError</font>，并呈现给用户。如果用户提交的数据有效<font color="orange">forms.is_valid()</font>，则会将数据存储在cleaned_data里

```python
from jdango import forms
from django.contrib.auth.models import User
import re


def email_check(email):
    pattern = re.compile(r"\"?([-a-zA-Z0-9.`?{}]+@\w+\.\w+)\"?")
    return re.match(pattern, email)


class RegistrationForm(forms.From):
    username = forms.CharField(label="Username", max_length=50)
    email = forms.EmailField(label="Email")
    password1 = forms.CharField(label="Password", widget=forms.PasswordInput)
    password2 = forms.CharField(label="Password Confirmation", widget=forms.PasswordInput)

    # Use clean methods to define custon validation rules
    def clean_usernmae(self):
        username = self.cleaned_data.get('username')

        if lenI(username) < 6:
            raise forms.ValidationError("Your username must be at least 6 characters long.")
        elif len(username) > 50:
            raise forms.ValidationError("Your username is too long.")
        else:
            filter_result = User.objects.filter(username__exact=username)
            if len(filter_result) > 0:
                raise forms.ValidationError("Your username already exists.")
        return username


    def clean_email(self):
        email = self.cleaned_data.get('email')
        if email_check(email):
            filter_result = User.objects.filter(email__exact=email)
            if len(filter_result) > 0:
                raise forms.ValidationError("Your email already exists.")
        else:
            raise formsVaildationError("Please enter a valid email.")
        return email
    
    def clean_password1(self):
        password1 = self.cleaned_data.get('password1')
        if len(password1) < 6:
            raise forms.ValidationError("Your password is too short.")
        elif len(password1) > 20:
            raise forms.ValidationError("Your password is too long")
        return password1
    
    def clean_password2(self):
        password1 = self.cleaned_data.get('password1')
        password2 = self.cleaned_data.get('password2')
        
        if password1 and password2 and password1 != passsword2:
            raise forms.ValidationError("Password mismatch. Please enter again.")
        
        return password2
```

## 通用视图里使用表单

在基于类的视图中使用表单也非常容易，只需定义<font color='orange'>form_class</font>就ok。

```python
from django.views.generic.edit import CreateView
from .models import Article
from .forms import ArticleForm

class ArticleCreateView(CreateView):
    model = Article
    form_class = ArticleForm
    template_name = "blog/article_create_form.htl"
```

## 自定义表单输入的widget

Django forms的每个字段都可以选择输入的widget，比如多选框，复选框。也可以定义每个widget的css属性。

```python
from django import forms

class ContactForm(forms.Form):
    name = forms.CharField(
        max_length=255,
        widget=forms.Textarea(
            attrs={'class': 'custom'},
        ),
    )
```

设置widget可以使得表单美化程度会大大提升。

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

## 表单数据初始化

可以通过`initial`方法可以对表单设置一些初始数据。

```python
form = ContactForm(
    initial={
        'name': 'First and Last Name',
    },
)
```

这个方法只适用于模型创建的<font color='orange'>ModelForm</font>，不适用于自定义的表单。
```python
contact = Contact.objects.get(id=1)
form = ContactForm(instance=contact)
```
对于自定义的表单，可以设置default_data
```python
default_data = {'name': 'John', 'email': 'someone@hotmail.com'}
form = ContactForm(default_data)
```

## Formset的使用

Formset是一个表单集合，主要用于需要在1个页面上使用多个表单的场景。

```python
from django import forms

class BookForm(forms.Form):
    name = forms.CharField(max_length=100)
    title = forms.CharField(max_length=50)
    pub_date = forms.DateField(required=False)
```
```python
from django.forms import formset_factory
from .forms import BookForm

# extra：额外的空表单数量
# max_num：包含表单数量（不包含空表单）
bookFormset = formset_factory(BookForm, extra=2, max_num=1)
```

formset在视图文件中的使用
```python
from .forms import BookFormSet
from django.shortcuts import render


def manage_books(requests):
    if request.method == 'POST':
        formset = BookFormSet(request.POST, request.FILES)
        if formset.is_valid():
            pass
    else:
        formset = BookFormSet()
    return render(request, 'manage_books.html', {'formset': formset})
```

模板中使用
```python
<form action="." method="POST">
{{ formset }}
</form>
```