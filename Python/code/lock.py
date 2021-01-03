from threading import Thread, current_thread


num = 0


def calc():
    global num
    print("thread {} is running...".format(current_thread().name))
    for _ in range(1000):
        num += 1
    print("thread {} ended.".format(current_thread().name))


if __name__ == '__main__':
    print("thread {} is running...".format(current_thread().name))
    threads = []
    for i in range(5):
        threads.append(Thread(target=calc))
        threads[i].start()

    for i in range(5):
        threads[i].join()

    print('global num: {}'.format(num))
    print("thread {} ended.".format(current_thread().name))
