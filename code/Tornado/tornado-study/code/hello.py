from abc import ABC

import tornado.web
import tornado.ioloop


class IndexHandler(tornado.web.RequestHandler, ABC):
    """主路由处理视图类"""
    def post(self):
        self.write("hell Itcast!")


if __name__ == '__main__':
    app = tornado.web.Application([
        (r"/", IndexHandler),
    ])
    app.listen(8000)
    tornado.ioloop.IOLoop.current().start()