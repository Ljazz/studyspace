# 判断是否时闰年
def is_leap(year):
    """
    判断是否是闰年
    """
    if year % 4 == 0 and year % 100 != 0 or year % 400 == 0:
        return True
    return False


if __name__ == '__main__':
    year = input('请输入年份:')
    print(is_leap(int(year)))
