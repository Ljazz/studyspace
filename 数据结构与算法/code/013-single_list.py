"""
单链表
"""

listValue = [1, 5, 6, 2, 4, 3]
listPointer = [3, 2, -1, 5, 1, 4]

head = 0    # head是指向链表第一个元素的指针，需要自己定义
next = head     # 给next赋值
while next != -1:   # next是指向下一个元素的指针，不等于-1代表后面还有元素
    print(listValue[next])  # 输出下一个元素中存储的值
    next = listPointer[next]    # 把指针变为下一个元素中存储的指针
