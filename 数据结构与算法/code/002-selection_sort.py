def selection_sort_1(nums):
    for i in range(len(nums)):
        print('第{}次排序过程'.format(i+1))
        for j in range(i+1, len(nums)):
            if nums[i] > nums[j]:
                nums[i], nums[j] = nums[j], nums[i]
            print(nums)
        print('-'*30)


def selection_sort_2(nums):
    res = []
    while nums:
        min = nums[0]
        for item in nums:
            if min > item:
                min = item
        nums.remove(min)
        res.append(min)
    print(res)


def selection_sort_3(nums):
    res = []
    while nums:
        min_num = min(nums)
        res.append(min_num)
        nums.remove(min_num)
    return []


if __name__ == "__main__":
    nums = [5, 3, 6, 4, 1, 2, 8, 7]
    print('原始数组...')
    print(nums)
    print('-'*30)
    selection_sort_2(nums)
