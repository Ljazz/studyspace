def bubble_sort_1(nums):
    for i in range(len(nums)-1):
        print('第{}次排序过程'.format(i+1))
        for j in range(0, len(nums)-1-i):
            if nums[j] > nums[j+1]:
                nums[j], nums[j+1] = nums[j+1], nums[j]
            print(nums)


if __name__ == "__main__":
    nums = [5, 3, 6, 4, 1, 2, 8, 7]
    print('原始数组...')
    print(nums)
    print('-'*30)
    bubble_sort_1(nums)
