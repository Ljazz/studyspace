"""
    File Name       : apiview.py
    Description     ：
    Author          : mxm
    Created on      : 2020/8/16
"""
import tornado.web
import tornado.ioloop
import tornado.httpserver
import tornado.options

from abc import ABC
from tornado.options import options, define
from tornado.web import RequestHandler

class BaseHandler(RequestHandler, ABC):
    pass