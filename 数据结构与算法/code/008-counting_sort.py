def countingSort(arr, maxValue):
    bucktLen = maxValue + 1
    bucket = [0] * bucktLen
    sortedIndex = 0
    arrLen = len(arr)
    for i in range(arrLen):
        if not bucket[arr[i]]:
            bucket[arr[i]] = 0
        bucket[arr[i]] += 1
    for j in range(bucktLen):
        while bucket[j] > 0:
            arr[sortedIndex] = j
            sortedIndex += 1
            bucket[j] -= 1
    return arr


if __name__ == "__main__":
    nums = [9, 5, 3, 6, 4, 1, 2, 8, 7]
    print('原始数组...')
    print(nums)
    print('-'*30)
    print(countingSort(nums, max(nums)))
