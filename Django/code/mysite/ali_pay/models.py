from django.db import models


class Pay(models.Model):
    trade_no = models.CharField('订单号', max_length=100)  # 支付前生成的订单号
    user_id = models.IntegerField('用户id')
    amount = models.FloatField('支付金额')
    out_trade_no = models.CharField('退单订单号', max_length=100)  # 退单用的订单号
    status = models.IntegerField('支付状态', default=0)  # -1：支付失败 0：尚未支付 1：支付成功
    create_date = models.DateTimeField('创建日期', auto_now_add=True)
    update_date = models.DateTimeField('更新日期', auto_now=True)
