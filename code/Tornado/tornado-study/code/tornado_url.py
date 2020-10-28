"""
    File Name       : tornado_url.py
    Description     ：路由配置
    Author          : mxm
    Created on      : 2020/8/14
"""
import tornado.web
import tornado.ioloop
import tornado.httpserver
import tornado.options

from abc import ABC
from tornado.options import options, define
from tornado.web import url, RequestHandler

define("port", default=8000, type=int, help="run server on the given port")

class IndexHandler(RequestHandler, ABC):
    def get(self):
        python_url = self.reverse_url("python_url")
        self.write('<a href="{}">itcast</a>'.format(python_url))


class ItcastHandler(RequestHandler, ABC):
    def initialize(self, subject):
        self.subject = subject

    def get(self):
        self.write(self.subject)


if __name__ == '__main__':
    tornado.options.parse_command_line()
    app = tornado.web.Application([
        (r"/", IndexHandler),
        (r"/cpp", ItcastHandler, {"subject": "c++"}),
        url(r"/python", ItcastHandler, {"subject": "python"}, name="python_url")
    ], debug=True)
    http_server = tornado.httpserver.HTTPServer(app)
    http_server.listen(options.port)
    tornado.ioloop.IOLoop.current().start()
