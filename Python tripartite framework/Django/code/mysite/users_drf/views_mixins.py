"""
    File Name       : views_mixins.py
    Description     ：
    Author          : mxm
    Created on      : 2020/8/11
"""
from users_drf.Serializer import BookModelSerializer
from users_drf.models import UserProfile, Book
from rest_framework import mixins, generics
from rest_framework.response import Response


class BookMixinView1(mixins.ListModelMixin, generics.GenericAPIView):
    queryset = Book.objects.all()
    serializer_class = BookModelSerializer

    def get(self, request, *args, **kwargs):    # 如果这里不加get函数，代表默认不支持get访问这个api
        APIKey = self.request.query_params.get("apikey", 0)
        developer = UserProfile.objects.filter(APIkey=APIKey).first()
        if developer:
            balance = developer.money
            if balance > 0:
                isbn = self.request.query_params.get("isbn", 0)
                developer.money -= 1
                developer.save()
                self.queryset = Book.objects.filter(isbn=int(isbn))
                return self.list(request, *args, **kwargs)
            else:
                return Response("兄弟，又到了需要充钱的时候！好开心")
        else:
            return Response("查无此人！")


class BookMixinView2(generics.ListAPIView):
    queryset = Book.objects.all()
    serializer_class = BookModelSerializer

    def get(self, request, *args, **kwargs):
        APIKey = self.request.query_params.get("apikey", 0)
        developer = UserProfile.objects.filter(APIkey=APIKey).first()
        if developer:
            balance = developer.money
            if balance > 0:
                isbn = self.request.query_params.get('isbn', 0)
                developer.money -= 1
                developer.save()
                self.queryset = Book.objects.filter(isbn=int(isbn))
                return self.list(request, *args, **kwargs)
            else:
                return Response("大兄弟，又到了该充钱的时刻了，是不是很惊喜！")
        else:
            return Response("查无此人")
