# 进程

**进程**是正在运行的程序的实例，每一个程序可能会产生多个进程。

## fork()

在Unix/Linux系统中，`fork`函数被用于创建进程。这个函数比较特殊，对于普通的函数，被调用它一次，返回一次，但是调用`fork`一次，它返回两次。事实上，`fork`函数创建了新的进程，我们把它称为子进程，子进程几乎是当前进程（即父进程）的一个拷贝：它会复制父进程的代码段，堆栈段和数据段。

对于父进程，`fork`函数返回了子进程的进程号`pid`，对于子进程，`fork`函数则返回`0`，这也是`fork`函数返回两次的原因，根据返回值，我们可以判断进程是父进程还是子进程。

下面我们看一段C代码，它展示了fork的基本使用：

```c
#include<unistd.h>
#include<stdio.h>

int main(int argc, char const *argv[])
{
    int pid;
    pid = fork();   # 使用 fork 函数
    
    if (pid < 0){
        printf("Fail to create process\n");
    }
    else if (pid == 0){
        printf("I am child process (%d) and my parent is (%d)\n", getpid(), getppid());
    }
    else{
        printf("I (%d) just created a child process (%d)\n", getpid(), pid);
    }
    return 0;
}
```
其中，getpid用于获取当前进程号，getppid用于获取父进程号。

事实上，Python的os模块包含了普通的操作系统功能，该模块也提供了`fork`函数

```python
import os

pid = os.fork()

if pid < 0:
    print("Fail to create process")
elif pid == 0:
    print("I am child process (%s) and my parent is (%s).".format(os.getpid(), os.getppid()))
else:
    print("I (%s) jus created a child process (%s).".format(os.getpid(), pid))
```

运行上述代码

```text
I (12075) just created a child process (12076).
I am child process (12076) and my parent is (12075).
```

需要注意的是，虽然子进程复制了父进程的代码段和数据段等，但是一旦子进程开始运行，子进程和父进程就是相互独立的，它们之间不再共享任何数据。

## 多进程

Python提供了`multiprocessing`模块，利用它，可以编写跨平台的多进程程序，但需要注意的是`multiprocessing`在Windows和Linux平台的不一致性：一样的代码在Windows和Linux下运行的结果可能不同。因为Windows的进程模型和Linux不一样，Windows下没有fork。

```python
import os
from multiprocessing import Process

# 子进程要执行的代码
def child_proc(name):
    print("Run child process {} ({})...".format(name, os.getpid()))


if __name__ == "__main__":
    print("Parent process {}.".format(os.getpid()))
    p = Process(target=child_proc, args=("test", ))
    print("Process will start")
    p.start()
    p.join()
    print("Process end.")
```

上述代码中，我们从multiprocessing模块引入Process，Process是一个用于创建进程对象的类，其中，target指定了进程要执行的函数，args指定了参数。在创建了进程实例p之后，我们调用start方法开始执行该子进程，接着，我们又调用了join方法，该方法用于阻塞子进程以外的所有进程（这里指父进程），当子进程执行完毕后，父进程才会继续执行，它通常用于进程间的同步。

```text
Parent process 13566.
Process will start.
Run child process test (13567)...
Process end.
```

## multiprocessing与平台有关

```python
import random
import os
from multiprocessing import Process

num = random.randint(0, 100)

def show_num():
    print("pid: {}, num is {}".format(os.getpid(), num))


if __name__ == "__main__":
    print("pid: {}, num is {}".format(os.getpid(), num))
    p = Process(target=show_num)
    p.start()
    p.join()
```

```text
pid: 13669, num is 83.
pid: 13670, num is 83...
```

## 使用进程池创建多个进程

```python
import os, time
from multiprocessing import Pool

def foo(x):
    print("Run task {} (pid: {})...".format(x, os.getpid()))
    time.sleep(2)
    print("Task {} result is: {}".format(x, x*x))


if __name__ == "__main__":
    print("Parent process {}.".format(os.getpid()))
    p = Pool(4)     # 设置进程数
    for i in range(5):
        p.apply_async(foo, args=(i, ))  # 设置每个进程要执行的函数和参数
    print("Waiting for all subprocesses done...")
    p.close()
    p.join()
    print("All subprocesses done.")
```

```text
Waiting for all subprocesses done
Run task 0 (pid: 1712)...
Run task 1 (pid: 704)...
Run task 2 (pid: 13596)...
Run task 3 (pid: 5500)...
Task 1 result is: 1Task 0 result is: 0

Run task 4 (pid: 1712)...
Task 2 result is: 4
Task 3 result is: 9
Task 4 result is: 16
All subprocesses done.

Process finished with exit code 0

```

## 进程间通信

进程间的通信可以通过管道（Pipe），队列（Queue）等多种方式实现。Python的multiprocessing模块封装了底层的实现机制。

```python
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
```

运行结果:
```text
write, 1
read, 1
write, 2
write, 3
write, 4
write_task end
read, 2
read, 3
read, 4
read_task end
DONE
```

# 线程

线程（thread）是进程（process）中一个实体，一个进程至少包含一个线程。

**进程和线程的区别主要有**：
- 进程之间是相互独立的，多进程中，同一个变量，各自有一份拷贝存在于每个进程中，但互不影响；而同一个进程的多个线程是内存共享的，所以变量都由所有线程线程共享；
- 由于进程间是独立的，因此一个进程的崩溃不会影响到其他进程；而线程是包含在进程之内的，线程的崩溃就会引发进程的崩溃，继而导致同一进程内其他线程也崩溃；

## 多线程

在Python中，进行多线程编程的模块有两个：thread和threading。其中，thread是低级模块，threading是高级模块，对thread进行了封装。

```python
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
```

运行结果

```text
thread MainThread is running...
hello world
thread TestThread is running...
hello  test
thread TestThread ended
thread MainThread ended
```

## 锁

由于同一个进程之间的线程是内存共享的，所以当多个线程对同一个变量进行修改的时候，就会得到意想不到的结果。

```python
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
```

运行结果
```text
thread MainThread is running...
thread Thread-1 is running...
thread Thread-1 ended.
thread Thread-2 is running...
thread Thread-2 ended.
thread Thread-3 is running...
thread Thread-3 ended.
thread Thread-4 is running...
thread Thread-4 ended.
thread Thread-5 is running...
thread Thread-5 ended.
global num: 5000
thread MainThread ended.
```

上述代码中，num的值是不确定的，每运行一遍，会发现结果变了。

原因`num += 1`不是一个原子操作，也就是说它在执行时被分成若干步：
- 计算`num += 1`，存入临时变量tmp中；
- 将tmp的值赋给num

由于线程时交替运行的，线程在执行时可能中断，就会导致其他线程读到一个脏值。

为了保证计算的准确性，我们需要给`num += 1`这个操作加上锁。当某个线程开始执行这个操作时，由于该线程获得了锁，因此其他线程不能同时执行该操作，只能等待，直到锁被释放，这样就可以避免修改的冲突。创建一个锁可以通过`threading.Lock()`来实现

```python
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
```

运行结果

```text
thread MainThread is running...
thread Thread-1 is running...
thread Thread-1 ended.
thread Thread-2 is running...
thread Thread-3 is running...
thread Thread-2 ended.
thread Thread-3 ended.
thread Thread-4 is running...
thread Thread-5 is running...
thread Thread-4 ended.
thread Thread-5 ended.
global num: 50000
thread MainThread ended
```

## GIL锁

GIL锁的存在导致Python不能有效地使用多线程实现多核任务，因为在同一时间，只能有一个线程在运行。

`GIL`全称`Global Interpreter Lock`，译为**全局解释锁**。早期Python为了支持多线程，引入了GIL锁，用于解决多线程之间数据共享和同步的问题。但是这种方式后来被发现是非常低效的，当试图去除GIL的时候，却有大量的库代码依赖GIL，由于各种历史原因，GIL锁就一直保留到现在。

# ThreadLocal

同一进程道的多个线程之间是内存共享的，这意味着，当一个线程对全局变量做了修改，将会影响到其他所有线程，这是很危险的。为了避免多个线程同时修改全局变量，需要对全局变量的修改加锁。

除了对全局变量的修改进行加锁，也可以使用线程自己的局部变量，因为局部变量只有线程自己能看到，对同一进程的其他线程是不可访问的。

```python
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

```
运行结果

```text
thread MainThread is running...
thread Thread-1 is running...
函数echo： Thread-1 10000
thread Thread-2 is running...
thread Thread-1 ended.
函数echo： Thread-2 10000
thread Thread-2 ended.
thread Thread-3 is running...
thread Thread-5 is running...
函数echo： Thread-3 10000
thread Thread-3 ended.
函数echo： Thread-5 10000
thread Thread-5 ended.
thread Thread-4 is running...
函数echo： Thread-4 10000
thread Thread-4 ended.
thread MainThread is ended.
```

上述代码中这种线程使用自己的局部变量的方法虽然可以避免多线程对同一变量的访问冲突，但还是会出现一些问题。在实际的开发中，我们会调用很多函数，每个函数又有很多局部变量，这时每个函数都这么传参数是不可取的。

为了解决上述问题，一个比较容易实现的做法就是创建一个全局字典，以线程的ID作为key，线程的局部数据作value，这样就可以消除函数传参的问题。

```python
from threading import Thread, current_thread

global_dict = {}


def echo():
    num = global_dict[current_thread()]     # 线程根据自己的ID获取数据
    print(current_thread().name, num)


def calc():
    print('thread %s is running...' % current_thread().name)
    global_dict[current_thread()] = 0
    for _ in range(10000):
        global_dict[current_thread()] += 1
    echo()
    print('thread %s ended.' % current_thread().name)


if __name__ == "__main__":
    print('thread %s is running...' % current_thread().name)

    threads = []
    for i in range(5):
        threads.append(Thread(target=calc))
        threads[i].start()
    for i in range(5):
        threads[i].join()

    print('thread %s ended.' % current_thread().name)

```

运行结果
```text
thread MainThread is running...
thread Thread-1 is running...
函数echo： Thread-1 10000
thread Thread-2 is running...
thread Thread-1 ended.

D:\Git_projects\studyspace\Python\code>
D:\Git_projects\studyspace\Python\code>
D:\Git_projects\studyspace\Python\code>python thread_global_dict.py
thread MainThread is running...
thread Thread-1 is running...
Thread-1 10000
thread Thread-2 is running...
thread Thread-1 ended.
Thread-2 10000
thread Thread-4 is running...
thread Thread-2 ended.
thread Thread-3 is running...
thread Thread-5 is running...
Thread-3 10000
thread Thread-3 ended.
Thread-5 10000
thread Thread-5 ended.
Thread-4 10000
thread Thread-4 ended.
thread MainThread ended.
```

上述的方式也并不是最完美的解决方法，因为global_dict是个全局变量，所有线程都可以对它进行修改。


Python提供了ThreadLocal对象，它可以真正做到线程之间的数据隔离，而且不用查找dict

```python
from threading import Thread, current_thread, local

global_data = local()


def echo():
    num = global_data.num
    print(current_thread().name, num)


def calc():
    print("thread {} is running...".format(current_thread().name))

    global_data.num = 0
    for _ in range(10000):
        global_data.num += 1
    echo()

    print("thread {} ended.".format(current_thread().name))


if __name__ == "__main__":
    print("thread {} is running...".format(current_thread().name))

    threads = []
    for i in range(5):
        threads.append(Thread(target=calc))
        threads[i].start()
    for i in range(5):
        threads[i].join()

    print("thread {} ended.".format(current_thread().name))

```

运行结果

```text
thread MainThread is running...
thread Thread-1 is running...
Thread-1 10000
thread Thread-2 is running...
thread Thread-1 ended.
thread Thread-5 is running...
thread Thread-4 is running...
Thread-2 10000
thread Thread-3 is running...
Thread-5 10000
thread Thread-5 ended.
thread Thread-2 ended.
Thread-3 10000
Thread-4 10000
thread Thread-4 ended.
thread Thread-3 ended.
thread MainThread ended.
```