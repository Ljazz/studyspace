listValue = [1, 4, 5, 2]
listRight = [3, 2, -1, 1]
head = 0    # 初始化头指针
num = 3     # 要插入的元素
next, last = head, head     # 初始化表示插入位置的下一个元素和上一个元素的指针

def output():
    # 定义链表输入函数
    next = head
    while next != -1:
        print(listValue[next])
        next = listRight[next]

output()


while listValue[next] <= num and next != -1:
    last = next
    next = listRight[next]
listValue.append(num)
listRight.append(listRight[last])
listRight[last] = len(listValue) - 1
print('*'*20)
output()
