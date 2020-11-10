"""
    File Name       : uri.py
    Description     ：
    Author          : mxm
    Created on      : 2020/8/15
"""
import json

import tornado.web
import tornado.ioloop
import tornado.options
import tornado.httpserver

from abc import ABC
from tornado.options import options, define
from tornado.web import RequestHandler

define("port", default=8000, type=int, help="run server on the given port.")


class IndexHandler(RequestHandler, ABC):
    def get(self):
        # self.write("hello world! 1")
        # self.write("hello world! 2")
        # self.write("hello world! 3")
        # self.write("hello world! 4")
        # self.write("hello world! 5")

        stu = {
            "name": "zhangsan",
            "age": 24,
            "gender": 1,
        }
        stu_json = json.dumps(stu)
        self.write(stu_json)
        self.set_header("Content-Type", "application/json; charset=UTF-8")

        # self.write(stu)



class SubjectCityHandler(RequestHandler, ABC):
    def get(self, subject, city):
        self.write(("Subject: {}<br /> City: {}".format(subject, city)))


class SubjectDateHandler(RequestHandler, ABC):
    def get(self, subject, date):
        self.write("Subject: {}<br/> Date: {}".format(subject, date))


if __name__ == '__main__':
    tornado.options.parse_command_line()
    app = tornado.web.Application([
        (r"/", IndexHandler),
        (r"/sub-city/(.+)/([a-z]+)", SubjectCityHandler),   # 无名方式
        (r"/sub-date/(?P<subject>.+)/(?P<date>\d+)", SubjectDateHandler),   # 命名方式
    ])
    http_server = tornado.httpserver.HTTPServer(app)
    http_server.listen(options.port)
    tornado.ioloop.IOLoop.current().start()
