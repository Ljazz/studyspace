"""
    File Name       : httpserver.py
    Description     ：HTTPserver
    Author          : mxm
    Created on      : 2020/8/12
"""
from abc import ABC

import tornado.web
import tornado.ioloop
import tornado.httpserver


class IndexHandler(tornado.web.RequestHandler, ABC):
    """
    主路由视图类
    """
    def get(self):
        """
        http的get请求
        :return:
        """
        self.write("HTTPServer")


class Sum(tornado.web.RequestHandler, ABC):
    def get(self):
        sum_str = str(33 + 22)
        self.write(sum_str)


if __name__ == '__main__':
    app = tornado.web.Application([
        (r'/', IndexHandler),
        (r'/sum/', Sum)
    ])
    # ----------------------------
    # 重写app.listen(8000)
    # app.listen(8000)
    # ----------------------------
    http_server = tornado.httpserver.HTTPServer(app)
    http_server.listen(8000)
    tornado.ioloop.IOLoop.current().start()
