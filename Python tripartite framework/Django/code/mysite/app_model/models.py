from django.db import models
from django.contrib.auth.models import User
from django.urls import reverse
from django.utils.timezone import now
from django.template.defaultfilters import slugify
import uuid
import os


class Article(models.Model):
    STATUS_CHOICES = (
        ('d', '草稿'),
        ('p', '发表'),
    )

    title = models.CharField('标题', max_length=200, unique=True)
    slug = models.SlugField('slug', max_length=60)
    body = models.TextField('正文')
    pub_date = models.DateTimeField('发布时间', default=now, null=True)
    create_date = models.DateTimeField('创建时间', auto_now_add=True)
    mod_date = models.DateTimeField('修改时间', auto_now=True)
    status = models.CharField('文章状态', max_length=1, choices=STATUS_CHOICES)
    views = models.PositiveIntegerField('浏览量', default=0)
    author = models.ForeignKey(User, verbose_name="作者", on_delete=models.CASCADE)
    tags = models.ManyToManyField('Tag', verbose_name="标签集合", blank=True)

    def __str__(self):
        return self.title

    def get_absolute_url(self):
        return reverse('blog:article_detail', args=[str(self.id)])

    def viewed(self):
        self.views += 1
        self.save(update_fields=['views'])

    def save(self, force_insert=False, force_update=False, using=None,
             update_fields=None):
        if not self.slug or not self.id:
            self.slug = slugify(self.title)
        super(Article, self).save()
        # do something

    class Meta:
        ordering = ['-pub_date']
        verbose_name = 'article'


def user_directory_path(instance, filename):
    ext = filename.split('.')[-1]
    filename = '{}.{}'.format(uuid.uuid4().hex[:10], ext)
    # return the whole path to the file
    return os.path.join(instance.user.id, 'avatar', filename)


class UserProfile(models.Model):
    user = models.OneToOneField(User, on_delete=models.CASCADE, related_name='profile')
    avatar = models.ImageField(upload_to=user_directory_path, verbose_name='头像')


class AuthorManager(models.Manager):
    def get_queryset(self):
        return super().get_queryset().filter(role='A')


class EditorManager(models.Manager):
    def get_queryset(self):
        return super().get_queryset().filter(role='E')


class Person(models.Model):
    first_name = models.CharField(max_length=50)
    last_name = models.CharField(max_length=50)
    role = models.CharField(max_length=1, choices=(('A', 'Author'), ('E', 'Editor')))
    objects = models.Manager()
    authors = AuthorManager()
    editors = EditorManager()
