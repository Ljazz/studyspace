def shell_sort(arr):
    import math
    gap = 1
    while gap < len(arr) / 3:
        gap = gap * 3 + 1
    while gap > 0:
        for i in range(gap, len(arr)):
            temp = arr[i]
            j = i - gap
            while j >= 0 and arr[j] > temp:
                arr[j+gap] = arr[j]
                j -= gap
            arr[j+gap] = temp
        gap = math.floor(gap / 3)
    print(arr)


if __name__ == "__main__":
    nums = [9, 5, 3, 6, 4, 1, 2, 8, 7]
    print('原始数组...')
    print(nums)
    print('-'*30)
    shell_sort(nums)
