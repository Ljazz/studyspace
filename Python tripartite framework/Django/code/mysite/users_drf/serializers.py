"""
    File Name       : serializers.py
    Description     ：
    Author          : mxm
    Created on      : 2020/8/9
"""
from rest_framework import serializers
from users_drf.models import UserProfile, Book


class BookSerializer(serializers.Serializer):
    title = serializers.CharField(required=True, max_length=100)
    isbn = serializers.CharField(required=True, max_length=100)
    author = serializers.CharField(required=True, max_length=100)
    publish = serializers.CharField(required=True, max_length=100)
    rate = serializers.DecimalField(max_digits=5, decimal_places=2, default=0)
