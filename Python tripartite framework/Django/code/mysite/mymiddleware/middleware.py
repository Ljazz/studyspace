from django.utils.deprecation import MiddlewareMixin


class MD1(MiddlewareMixin):
    def process_request(self, request):
        print('MD1中的 process_request')


class MD2(MiddlewareMixin):
    def process_request(self, request):
        print('MD2中的 process_request')
