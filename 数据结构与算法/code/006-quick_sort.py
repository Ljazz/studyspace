def quick_sort(arr, left=None, right=None):
    left = 0 if not isinstance(left, (int, float)) else left
    right = len(arr)-1 if not isinstance(right, (int, float)) else right
    if left < right:
        partitionIndex = partition(arr, left, right)
        quick_sort(arr, left, partitionIndex-1)
        quick_sort(arr, partitionIndex+1, right)
    return arr


def partition(arr, left, right):
    pivot = left
    index = pivot + 1
    i = index
    while i <= right:
        if arr[i] < arr[pivot]:
            swap(arr, i, index)
            index += 1
        i += 1
    swap(arr, pivot, index-1)
    return index-1


def swap(arr, i, j):
    arr[i], arr[j] = arr[j], arr[i]


if __name__ == "__main__":
    nums = [9, 5, 3, 6, 4, 1, 2, 8, 7]
    print('原始数组...')
    print(nums)
    print('-'*30)
    print(quick_sort(nums))
