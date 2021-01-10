"""
    File Name       : set_header.py
    Description     ：请求体头设置
    Author          : mxm
    Created on      : 2020/8/15
"""
import json

import tornado.web
import tornado.ioloop
import tornado.options
import tornado.httpserver
from tornado.options import define, options
from tornado.web import RequestHandler

define("port", default=8000, type=int, help="run server on the given port.")


class IndexHandler(RequestHandler):
    def set_default_headers(self):
        print("执行了set_default_headers()")
        # 设置get与post方式的默认响应体格式为json
        self.set_header("Content-Type", "application/json; charset=UTF-8")
        # 设置一个名为itcast、值为python的header
        self.set_header("itcast", "python")

    def get(self):
        print("执行了get")
        stu = {
            "name": "zhangsan",
            "age": 24,
            "gender": 1,
        }
        sut_json = json.dumps(stu)
        self.write(sut_json)
        self.set_header("itcast", "i love python") # 重写了header

    def post(self):
        print("执行了post")
        stu = {
            "name": "zhangsan",
            "age": 24,
            "gender": 1,
        }
        sut_json = json.dumps(stu)
        self.write(sut_json)


if __name__ == '__main__':
    tornado.options.parse_command_line()
    app = tornado.web.Application([
        (r"/", IndexHandler),
    ])
    http_server = tornado.httpserver.HTTPServer(app)
    http_server.listen(options.port)
    tornado.ioloop.IOLoop.current().start()

