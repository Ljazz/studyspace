# 两个数的和
def towSums(mynum, target):
    mydict = {}     # 建立一个字典，存储数据和下标的对应关系
    for i in range(len(mynum)):
        m = mynum[i]  # 定义m为当前待查询数字
        if target - m in mydict:    # 判定target-m是否已经在字典中
            return (mydict[target-m], i)    # 如果已经存在，则返回这两个数的下标
        else:
            mydict[m] = i   # 如果不存在则记录键值对


if __name__ == "__main__":
    num = [3, 4, 5, 7, 10]
    x, y = towSums(num, 11)
    print(num[x], num[y])
