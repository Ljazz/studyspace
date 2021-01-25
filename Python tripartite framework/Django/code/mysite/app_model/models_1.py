from django.db import models
from django.urls import reverse


# 自定义Manager方法
class HighRatingManager(models.Manager):
    def get_queryset(self):
        # return super(HighRatingManager, self).get_queryset().filter(rating='1')
        return super().get_queryset().filter(rating='1')


class Product(models.Model):
    # CHOICES选项
    RATING_CHOICES = (
        ("1", "very Good"),
        ("2", "Good"),
        ("3", "Bad"),
    )

    # 数据表字段
    name = models.CharField('名称', max_length=30)
    rating = models.CharField(max_length=1, choices=RATING_CHOICES)

    # MANAGERS方法
    objects = models.Manager()
    high_rating_products = HighRatingManager()

    # META类选项
    class Meta:
        verbose_name = 'product'
        verbose_name_plural = 'products'

    # __str__方法
    def __str__(self):
        return self.name

    # 重写save方法
    def save(self, *args, **kwargs):
        # do something
        super().save(*args, **kwargs)
        # do something else

    # 自定义绝对路径
    def get_absolute_url(self):
        return reverse('product_details', kwargs={'pk': self.id})

    # 定义其它方法
    def do_something(self):
        pass
