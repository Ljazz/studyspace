import time
import threading

gen = None  # 全局生成器，共long_io使用


def gen_coroutine(f):
    def wrapper(*args, **kwargs):
        global gen
        gen = f()
        gen.next()

    return wrapper()


def long_io():
    def fun():
        print("开始执行IO操作")
        global gen
        time.sleep(5)
        try:
            print("完成IO操作，并send结果唤醒挂起程序继续执行")
            gen.send("io result")
        except Exception as e:
            print(str(e))

    threading._start_new_thread(fun, ())


@gen_coroutine
def req_a():
    print("开始处理请求req_a")
    ret = yield long_io()
    print("ret: {}".format(ret))
    print("完成处理请求req_a")


def req_b():
    print("开始处理请求req_a")
    time.sleep(2)
    print("开始处理请求req_b")


def main():
    req_a()
    req_b()
    while 1:
        pass


if __name__ == "__main__":
    main()
