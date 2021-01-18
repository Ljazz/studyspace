from django.db import models
from django.contrib.auth.models import User
from datetime import date


class Publisher(models.Model):
    name = models.CharField(max_length=30)
    address = models.CharField(max_length=60)

    def __str__(self):
        return self.name


class Book(models.Model):
    name = models.CharField(max_length=30)
    description = models.TextField(blank=True, default='')
    publisher = models.ForeignKey(Publisher, on_delete=models.CASCADE)
    add_date = models.DateField()

    def __str__(self):
        return self.name


class Restaurant(models.Model):
    name = models.TextField()
    address = models.TextField(blank=True, default='')
    telephone = models.TextField(blank=True, default='')
    url = models.URLField(blank=True, null=True)
    user = models.ForeignKey(User, default=1, on_delete=models.CASCADE)
    date = models.DateField(default=date.today)

    def __str__(self):
        return self.name


class Dish(models.Model):
    name = models.TextField()
    description = models.TextField(blank=True, default='')
    price = models.DecimalField('USD amount', max_digits=8, decimal_places=2, blank=True, null=True)
    user = models.ForeignKey(User, default=1, on_delete=models.CASCADE)
    date = models.DateField(default=date.today)
    image = models.ImageField(upload_to="myrestaurants", blank=True, null=True)
    restaurant = models.ForeignKey(Restaurant, null=True, related_name='dishes', on_delete=models.CASCADE)

    def __str__(self):
        return self.name


class Review(models.Model):
    RATING_CHOICES = ((1, 'one'), (2, 'two'), (3, 'three'), (4, 'four'), (5, 'five'))
    rating = models.PositiveIntegerField('Rating', blank=False, default=3, choices=RATING_CHOICES)
    comment = models.TextField(blank=True, null=True)
    user = models.ForeignKey(User, default=1, on_delete=models.CASCADE)
    date = models.DateField(default=date.today)

    class Meta:
        abstract = True


class RestaurantReview(Review):
    restaurant = models.ForeignKey(Restaurant, on_delete=models.CASCADE)

    def __str__(self):
        return "{} review".format(self.restaurant.name)

    class Meta:
        # 按Priority降序，order_date升序
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
        permissions = (("can_deliver_pizzas", "Can deliver pizzas"))
