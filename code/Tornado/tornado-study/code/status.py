"""
    File Name       : status.py
    Description     ：tornado 状态码
    Author          : mxm
    Created on      : 2020/8/15
"""
import json

import tornado.web
import tornado.ioloop
import tornado.options
import tornado.httpserver
from tornado.options import options, define
from tornado.web import RequestHandler

define("port", default=8000, type=int, help="run server on the given port.")


class Err404Handler(RequestHandler):
    """对应/err/404"""
    def get(self):
        self.write("404")
        self.set_status(404) # 标准状态码，不用设置reason


class Err210Handler(RequestHandler):
    """对应/err/210"""
    def get(self):
        self.write("210")
        self.set_status(210, "非标准的状态码") # 非标准状态码，设置reason


class Err211Handler(RequestHandler):
    """对应/err/211"""
    def get(self):
        self.write("211")
        self.set_status(211) # 非标准状态码，不设置reason


class IndexHandler(RequestHandler):
    def prepare(self):
        if self.request.headers.get("Content-Type").startswith("application/json"):
            self.json_dict = json.loads(self.request.body)
        else:
            self.json_dict = None

    def post(self):
        if self.json_dict:
            for key, value in self.json_dict.items():
                self.write("<h3>{}</h3><p>{}</p>".format(key, value))

    def put(self):
        if self.json_dict:
            for key, value in self.json_dict.items():
                self.write("<h3>{}</h3><p>{}</p>".format(key, value))


if __name__ == '__main__':
    tornado.options.parse_command_line()
    app = tornado.web.Application([
        (r'/err/404', Err404Handler),
        (r'/err/210', Err210Handler),
        (r'/err/211', Err211Handler),
        (r'/', IndexHandler),
    ])
    http_server = tornado.httpserver.HTTPServer(app)
    http_server.listen(options.port)
    tornado.ioloop.IOLoop.current().start()