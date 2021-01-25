<!-- TOC -->

- [为什么要使用Django Formset](#为什么要使用django-formset)
- [Formset的分类](#formset的分类)
- [formset_factory的使用](#formset_factory的使用)
- [modelformset_factory的使用](#modelformset_factory的使用)
- [inlineformset_factory使用](#inlineformset_factory使用)
- [整个formset的验证](#整个formset的验证)
- [给Formset添加额外字段](#给formset添加额外字段)

<!-- /TOC -->

## 为什么要使用Django Formset

在不使用Formset情况下在同一页面上一键提交2张或多张表单的做法如下，在模板中给每个表单取不同的名字，如form1和form2。注意：form1和form2分别对应forms.py里的Form1()和Form2()。

```html
<form>
    {{ form1.as_p }}
    {{ form2.as_p }}
</form>
```
用户提交后，可以在视图中对用户提交的数据分别处理

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

## Formset的分类

Django针对不同的formset提供了3种方法：<font color='#ff2299'>formset_factory</font>，<font color='#ff2299'>modelformset_factory</font>和<font color='#ff2299'>inlineformset_factory。</font>

## formset_factory的使用

对于继承forms.Form的自定义表单，我们可以使用`formset_factory`。可以通过设置`extra`和`max_num`属性来确定想要展示的表单数量。注意：<font color='sky blue'>max_num优先级高级extra。</font>如下案例种，想要显示3个空表单（extra=3），但最后只会显示2个空表单，因为max_num=2

```python
from django import forms


class BookForm(forms.Form):
    name = forms.CharField(max_length=100)
    title = forms.CharField()
    pub_date = forms.DateField(required=False)

# forms.py - build a formset of books
from django.forms import formset_factory
from .forms import BookForm

# extra：想要显示空表单的数量
# max_num：表单显示最大数量，可选，默认1000

BookFormSet = formset_factory(BookForm, extra=3, max_num=2)
```

在视图文件views.py里，我们可以像使用form一样使用formset

```python
from .forms import BookFormSet
from django.shortcuts import render

def manage_books(request):
    if request.method == 'POST':
        formset = BookFormSet(requet.POST, request.FILES)
        if formset.is_valid():
            pass
    else:
        formset = BookFormSet()
    return render(request, 'manage_books.html', {'formset': formset})
```

模板中使用formset

```html
<form action="." method="POST">
    {{ formset }}
</form>
```

也可以如下方式使用

```python
<form method="post">
    {{ formset.management_form }}
    <table>
        {% for form in formset %}
        {{ form }}
        {% endfor %}
    </table>
</form>
```

## modelformset_factory的使用

Formset也可以直接由模型Model创建，这时需要使用<font color='red'>modelformset_factory</font>。可以指定需要显示的字段和表数量。

```python
from django.forms import modelformset_factory
from myapp.models import Author

AuthorFormSet = modelformset_factory(
    Author, fields=('name', 'title'), extra=3
)
```
上述这种方式并不推荐，因为对单个表单添加验证方法非常不方便。可以先创建自定义的ModelForm，添加单个表单验证，然后利用`modelformset_factory`创建formset。

```python
class AuthorForm(forms.ModelForm):
    class Meta:
        model = Author
        fields = ('name', 'title')
    
    def clean_name(self):
        pass
```
由ModelForm创建`formset`

```python
AuthorFormSet = modelformset_factory(Author, form=AuthorForm)
```

## inlineformset_factory使用

由recipe模型，Recipe与Ingredient是单对多的关系。一般的formset只允许一次性提交多个Recipe或多个Ingredient。但是希望同一页面上添加一个菜谱(Recipe)和多个原料（Ingredient），这时就可以使用inlineformset了。

```python
from django.db import models

class Recipe(models.Model):
    title = models.CharField(max_length=255)
    description = models.TextField()


class Ingredient(models.Model):
    recipe = models.ForeignKey(Recipe, on_delete=models.CASCADE, related_name='ingredient')
    name = models.CharField(max_length=255)
```

利用`inlineformset_factory`创建formset的方法如下所示。该方法的第一个参数和第二个参数都是模型，其中第一个参数必需是ForeignKey。

```python
# forms.py
from django.forms import ModelForm
from django.forms import inlineformset_factory
from .models import Recipe, Ingredient, Instruction


class RecipeForm(ModelForm):
    class Meta:
        model = Recipe
        fields = ('title', 'description')

IngredientFormSet = inlinefromset_factory(Recipe, Ingredient, fields=('name', ), extra=3, can_delelte=False, max_num=5)
```

views.py中使用formset创建和更详细recipe的代码如下。在对IngredientFormSet进行实例化的时候，必需指定recipe的实例。

```python
def recipe_update(request, pk):
    recipe = get_object_or_404(Recipe, pk=pk)
    if request.method == 'POST':
        form = RecipeForm(request.POST, instance=recipe)

        if form.is_valid():
            recipe = form.save()
            ingredient_formset = IngredientFormSet(request.POST, instance=recipe)

            if ingredient_formset.is_valid():
                ingredient_formset.save()
        return redirect('/recipe/')
    else:
        form = RecipeForm(instance=recipe)
        ingredient_formset = IngredientFormSet(instance=recipe)
    return render(request, 'recipe/recipe_update.html', {'form': form, 'ingredient_format': ingredient_forms})


def recipe_add(request):
    if request.method == 'POST':
        form = RecipeForm(request.POST)

        if form.is_valid():
            recipe = form.save()
            ingredient_formset = IngredientFormSet(request.POST, instance=recipe)

            if ingredient_formset.is_valid():
                ingredient_formset.save()
        return redirect('/recipe/')
    else:
        form = RecipeForm()
        ingredient_formset = IngredientFormSet()
    return render(request, 'recipe/recipe_add.html', {'form': form, 'ingredient_formset': ingredient_formset})
```
模板recipe/recipe_add.html代码

```html
<h1>Add Recipe</h1>
<form action="." method="post">
    {% csrf_token %}

    {{ form.as_p }}

    <fieldset>
        <legend>Recipe Ingredient</legend>
        {{ ingredient_formset.management_form }}
        {{ ingredinet_formset.non_form_errors }}
        {% for form in ingredient_formset %}
            {{ form.name.errors }}
            {{ form.name.label_tag }}
            {{ form.name }}
        {% endfor %}
    </fieldset>
    <input type="submit" value="Add recipe" class="submit" />
</form>
```

## 整个formset的验证

formst由多个表单组成，单个表单的验证可以通过自定义的clean方法来完成，然而有时我们需要对整个formset的数据进行验证。

案例：用户一次性提交多篇文章标题后，需要检查title是否已重复。先定义一个BaseFormSet，然后使用formset=BaseArticleFormSet添加formset的验证。

```python
from django.forms import BaseFormSet
from django.forms import formset_factory
from myapp.forms import ArticleForm


class BaseArticleFormSet(BaseFormSet):
    def clean(self):
        """Checks that no two articles have the same title."""
        if any(self.errors):
            return
        
        titles = []
        for form in self.forms:
            title = form.cleaned_data['title']
            if title in titles:
                raise forms.ValidationError("Articles in a set must have distinct titles.")
            titles.append(title)

ArticleFormSet = formset_factory(ArticleForm, formset=BaseArticleFormSet)
```

## 给Formset添加额外字段

在BaseFormSet里不仅可以天机formset的验证，而且可以添加额外的字段

```python
from django.forms import BaseFormSet
from django.forms import formset_factory
from myapp.forms import ArticleForm


class BaseArticleFormSet(BaseFormSet):
    def add_fields(self, form, index):
        super().add_fields(form, index)
        form.fields["my_field"] = forms.CharField()


ArticleFormSet = formset_factory(ArticleForm, formset=BaseArticleFormSet)
```
