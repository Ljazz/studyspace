"""
    File Name       : templates.py
    Description     ：模板demo
    Author          : mxm
    Created on      : 2020/8/16
"""
import os

import tornado.web
import tornado.ioloop
import tornado.httpserver
import tornado.options
from abc import ABC
from tornado.options import define, options
from tornado.web import RequestHandler

define("port", default=8000, type=int, help="run server on the given port.")


class IndexHandler(RequestHandler, ABC):
    def get(self):
        houses = [
            {
                "price": 398,
                "title": "宽窄巷子+160平大空间+文化保护区双地铁",
                "score": 5,
                "comments": 6,
                "position": "北京市丰台区六里桥地铁"
            },
            {
                "price": 398,
                "title": "宽窄巷子+160平大空间+文化保护区双地铁",
                "score": 5,
                "comments": 6,
                "position": "北京市丰台区六里桥地铁"
            },
            {
                "price": 398,
                "title": "宽窄巷子+160平大空间+文化保护区双地铁",
                "score": 5,
                "comments": 6,
                "position": "北京市丰台区六里桥地铁"
            },
            {
                "price": 398,
                "title": "宽窄巷子+160平大空间+文化保护区双地铁",
                "score": 5,
                "comments": 6,
                "position": "北京市丰台区六里桥地铁"
            },
            {
                "price": 398,
                "title": "宽窄巷子+160平大空间+文化保护区双地铁",
                "score": 5,
                "comments": 6,
                "position": "北京市丰台区六里桥地铁"
            }
        ]
        self.render("index.html", houses=houses)


if __name__ == '__main__':
    current_path = os.path.dirname(__file__)
    tornado.options.parse_command_line()
    app = tornado.web.Application(
        [
            (r'/', IndexHandler),
        ],
        template_path=os.path.join(current_path, 'template')
    )
    http_server = tornado.httpserver.HTTPServer(app)
    http_server.listen(options.port)
    tornado.ioloop.IOLoop.current().start()
