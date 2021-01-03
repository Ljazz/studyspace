from threading import Thread, current_thread, Lock

num = 0
lock = Lock()


def calc():
    global num
    print("thread {} is running...".format(current_thread().name))
    for _ in range(10000):
        lock.acquire()      # 获取锁
        num += 1
        lock.release()      # 释放锁
    print("thread {} ended.".format(current_thread().name))


if __name__ == "__main__":
    print("thread {} is running...".format(current_thread().name))
    
    threads = []
    for i in range(5):
        threads.append(Thread(target=calc))
        threads[i].start()
    for i in range(5):
        threads[i].join()
    
    print("global num: {}".format(num))
    print("thread {} ended".format(current_thread().name))
