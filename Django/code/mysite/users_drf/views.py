from users_drf.serializers import BookSerializer
from rest_framework.views import APIView
from rest_framework.response import Response
from users_drf.models import UserProfile, Book


class BookAPIView1(APIView):
    """
    使用serializer
    """
    def get(self, request, format=None):
        APIKey = self.request.query_params.get('apikey', 0)
        developer = UserProfile.objects.filter(APIkey=APIKey).first()
        if developer:
            balance = developer.money
            if balance > 0:
                isbn = self.request.query_params.get("isbn", 0)
                books = Book.objects.filter(isbn=int(isbn))
                books_serializer = BookSerializer(books, many=True)
                developer.money -= 1
                developer.save()
                return Response(books_serializer.data)
            else:
                return Response("兄弟，又到了需要充钱的时候！")
        else:
            return Response('查无此人')


class BookAPIView2(APIView):
    """
    使用ModelSerializer
    """
    def get(self, request, foramt=None):
        APIKey = self.request.query_params.get('apikey', 0)
        developer = UserProfile.objects.filter(APIkey=APIKey).first()
        if developer:
            balance = developer.money
            if balance > 0:
                isbn = self.request.query_params.get("isbn", 0)
                books = Book.objects.filter(isbn=int(isbn))
                books_serializer = BookSerializer(books, many=True)
                developer.money -= 1
                developer.save()
                return Response(books_serializer.data)
            else:
                return Response("兄弟，又到了需要充钱的时候！")
        else:
            return Response('查无此人')
