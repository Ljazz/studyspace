"""
    File Name       : options.py
    Description     ：
    Author          : mxm
    Created on      : 2020/8/12
"""
from abc import ABC

import tornado.web
import tornado.ioloop
import tornado.httpserver
import tornado.options

tornado.options.define('port', default=8000, type=int, help="run server on the given port.")
tornado.options.define('itcast', default=[], type=str, multiple=True, help="itcast subjects.")


class IndexHandler(tornado.web.RequestHandler, ABC):
    """主路由处理类"""
    def get(self, *args, **kwargs):
        self.write("tornado.options")


if __name__ == '__main__':
    tornado.options.parse_command_line()
    print(tornado.options.options.itcast)
    app = tornado.web.Application([
        (r'/', IndexHandler),
    ])
    http_server = tornado.httpserver.HTTPServer(app)
    http_server.listen(tornado.options.options.port)
    tornado.ioloop.IOLoop.current().start()
