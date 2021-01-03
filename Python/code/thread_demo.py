from threading import Thread, current_thread


def thread_test(name):
    print("thread {} is running...".format(current_thread().name))
    print("hello ", name)
    print("thread {} ended".format(current_thread().name))


if __name__ == '__main__':
    print("thread {} is running...".format(current_thread().name))
    print("hello world")
    t = Thread(target=thread_test, args=("test", ), name="TestThread")
    t.start()
    t.join()
    print("thread {} ended".format(current_thread().name))
