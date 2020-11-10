"""
    File Name       : upload_file.py
    Description     ：上传文件
    Author          : mxm
    Created on      : 2020/8/15
"""
import tornado.web
import tornado.ioloop
import tornado.httpserver
import tornado.options
from abc import ABC
from tornado.options import options, define
from tornado.web import RequestHandler

define("port", default=8000, type=int, help="run sever on the given port.")


class IndexHandler(RequestHandler, ABC):
    def get(self):
        self.write("文件上传测试")


class UploadHandler(RequestHandler, ABC):
    def post(self):
        files = self.request.files
        img_files = files.get('img')
        if img_files:
            img_file = img_files[0]["body"]
            file = open('img.jpg', 'wb')
            file.write(img_file)
            file.close()
        self.write("upload file!")


if __name__ == '__main__':
    tornado.options.parse_command_line()
    app = tornado.web.Application([
        (r"/", IndexHandler),
        (r"/upload", UploadHandler),
    ])
    http_server = tornado.httpserver.HTTPServer(app)
    http_server.listen(options.port)
    tornado.ioloop.IOLoop.current().start()
