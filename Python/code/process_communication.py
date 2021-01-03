import time
from multiprocessing import Process, Queue


# 向队列中写入数据
def write_task(q):
    try:
        n = 1
        while n < 5:
            print("write, {}".format(n))
            q.put(n)
            time.sleep(1)
            n += 1
    except BaseException:
        print("write_task error")
    finally:
        print("write_task end")


# 从队列读取数据
def read_task(q):
    try:
        n = 1
        while n < 5:
            print("read, {}".format(q.get()))
            time.sleep(5)
            n += 1
    except BaseException:
        print("read_task error")
    finally:
        print("read_task end")


if __name__ == '__main__':
    q = Queue()     # 父进程创建Queue， 并传给各个子进程
    pw = Process(target=write_task, args=(q, ))
    pr = Process(target=read_task, args=(q, ))
    pw.start()  # 启动子进程pw， 写入
    pr.start()  # 启动子进程pr， 读取
    pw.join()   # 等待pw结束
    pr.join()   # 等待pr结束
    print("DONE")
