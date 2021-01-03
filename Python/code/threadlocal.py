from threading import Thread, current_thread


def echo(num):
    print('函数echo：', current_thread().name, num)


def calc():
    print("thread {} is running...".format(current_thread().name))
    local_num = 0
    for _ in range(10000):
        local_num += 1
    echo(local_num)
    print("thread {} ended.".format(current_thread().name))


if __name__ == "__main__":
    print("thread {} is running...".format(current_thread().name))
    
    threads = []
    for i in range(5):
        threads.append(Thread(target=calc))
        threads[i].start()
    for i in range(5):
        threads[i].join()

    print("thread {} is ended.".format(current_thread().name))
