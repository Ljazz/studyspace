def merge_sort(arr):
    import math
    if len(arr) < 2:
        return arr
    middle = math.floor(len(arr) / 2)
    left, right = arr[: middle], arr[middle:]
    return merge(merge_sort(left), merge_sort(right))


def merge(left, right):
    result = []
    while left and right:
        if left[0] <= right[0]:
            result.append(left.pop(0))
        else:
            result.append(right.pop(0))
    while left:
        result.append(left.pop(0))
    while right:
        result.append(right.pop(0))
    return result


if __name__ == "__main__":
    nums = [9, 5, 3, 6, 4, 1, 2, 8, 7]
    print('原始数组...')
    print(nums)
    print('-'*30)
    print(merge_sort(nums))
