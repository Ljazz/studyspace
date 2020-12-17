"""
输入：

industry_list = [
    {
        "pid": "女装",
        "id": "连衣裙"
    },
    {
        "id": "女装"
    },
    {
        "pid": "女装",
        "id": "半身裙"
    },
    {
        "pid": "女装",
        "id": "A字裙"
    },
    {
        "id": "数码"
    },
    {
        "pid": "数码",
        "id": "电脑配件"
    },
    {
        "pid": "电脑配件",
        "id": "内存"
    },
]
输出：

result = {
    "数码": {
        "电脑配件": {
            "内存": {}
        }
    },
    "女装": {
        "连衣裙": {},
        "半身裙": {},
        "A字裙": {}
    }
}
"""


def convert_format(array):
    data = {}
    for item in array:
        item.setdefault('pid', '')
        data.setdefault(item['id'], {})
        data.setdefault(item['pid'], {})
        data[item['pid']][item['id']] = data[item['id']]
    return data['']


if __name__ == "__main__":
    industry_list = [
        {
            "pid": "女装",
            "id": "连衣裙"
        },
        {
            "id": "女装"
        },
        {
            "pid": "女装",
            "id": "半身裙"
        },
        {
            "pid": "女装",
            "id": "A字裙"
        },
        {
            "id": "数码"
        },
        {
            "pid": "数码",
            "id": "电脑配件"
        },
        {
            "pid": "电脑配件",
            "id": "内存"
        },
    ]
    output = convert_format(industry_list)
    print(output)