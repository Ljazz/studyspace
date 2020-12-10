def insertion_sort_1(nums):
    for i in range(1, len(nums)):
        print('第{}次排序过程'.format(i+1))
        for j in range(i, 0, -1):
            if nums[j] < nums[j-1]:
                nums[j], nums[j-1] = nums[j-1], nums[j]
            print(nums)  # [1, 2, 3, 4, 5, 6, 7, 8]


def insertion_sort_2(nums):
    for i in range(1, len(nums)):
        print('第{}次排序过程'.format(i+1))
        for j in range(i):
            print(nums)
            if nums[i] > nums[j]:
                ins = nums[i]
                nums.pop(i)
                nums.insert(j, ins)
                break


def insertion_sort_3(nums):
    for i in range(len(nums)):
        print('第{}次排序过程'.format(i+1))
        for j in range(i):
            print(nums)
            if nums[i] < nums[j]:
                ins = nums[i]
                nums.pop(i)
                nums.insert(j, ins)
                break


if __name__ == "__main__":
    nums = [5, 3, 6, 4, 1, 2, 8, 7]
    print('原始数组...')
    print(nums)
    print('-'*30)
    insertion_sort_1(nums)
