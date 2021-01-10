from django.db import models
from django.contrib.auth.models import AbstractUser


class UserProfile(AbstractUser):
    """
    用户
    """
    APIkey = models.CharField('APIkey', max_length=30, default='abcdefghijklmn')
    money = models.IntegerField('余额', default=10)

    class Meta:
        verbose_name = '用户'
        verbose_name_plural = verbose_name

    def __str__(self):
        return self.username


class Book(models.Model):
    """
    书籍
    """
    title = models.CharField('书名', max_length=30, default='')
    isbn = models.CharField('isbn', max_length=30, default='')
    author = models.CharField('作者', max_length=20, default='')
    publish = models.CharField('出版社', max_length=30, default='')
    rate = models.DecimalField('评分', max_digits=5, decimal_places=2, default=0)
    create_time = models.DateTimeField('创建时间', auto_now_add=True)

    class Meta:
        verbose_name = '书籍'
        verbose_name_plural = verbose_name

    def __str__(self):
        return self.title
