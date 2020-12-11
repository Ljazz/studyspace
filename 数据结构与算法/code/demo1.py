def selctionSort(a):
    sort_list = []
    while len(a):
        # sort_list.append(min(a))
        # a.remove(min(a))
        min_s = a[0]
        for s in a:
            if s < min_s:
                min_s = s
        sort_list.append(min_s)
        a.remove(min_s)
    print(sort_list)


def insertionSort(arr):
    for i in range(len(arr)):
        for j in range(i):
            if arr[i] < arr[j]:
                arr[i], arr[j] = arr[j], arr[i]
    print(arr)


def bubbleSort(arr):
    for i in range(len(arr)-1):
        for j in range(len(arr) - 1 - i):
            if arr[j] > arr[j+1]:
                arr[j], arr[j+1] = arr[j+1], arr[j]
    print(arr)


if __name__ == "__main__":
    a = ['b', 'e', 'a', 'd', 'f', 'c']

    # 选择排序
    # selctionSort(a)

    # 插入排序
    insertionSort(a)

    # 冒泡排序
    # bubbleSort(a)
