import os
import time

try:
    import qrcode
except ModuleNotFoundError:
    raise ModuleNotFoundError('qrcode未安装，请在终端执行pip install qrcode')

try:
    import alipay
except ModuleNotFoundError:
    raise ModuleNotFoundError('alipay未安装，请在终端执行pip install python-alipay-sdk')

APPNET = 'https://openapi.alipaydev.com/gateway.do'
RETURN_URL = 'http://127.0.0.1:8000/api/ali-pay/pay/'
APPID = '**********'
APP_PRIVATE = """-----BEGIN RSA PRIVATE KEY-----
************
-----END RSA PRIVATE KEY-----"""
APP_PUBLIC = "************"
ALI_PUBLIC = """-----BEGIN PUBLIC KEY-----
************
-----END PUBLIC KEY-----"""


def create_ali():  # 创建ali对象
    ali = alipay.AliPay(
        appid=APPID,
        app_notify_url=None,    # 默认回调url
        app_private_key_string=APP_PRIVATE,
        alipay_public_key_string=ALI_PUBLIC,    # 支付宝的公钥，验证支付宝回传消息使用，不是自己的公钥
        sign_type='RSA2',   # RSA 或者 RSA2
        debug=True,  # 默认False
    )
    return ali


def create_paypage_url(trade_no, amount, subject):   # 生成网页支付地址
    order_string = create_ali().api_alipay_trade_page_pay(
        out_trade_no=trade_no,
        total_amount=amount,
        subject=subject,
        notify_url=None,    # 可选，不填则使用默认notify url
        return_url=RETURN_URL   # 可选，不填则使用默认notify url
    )
    pay_url = APPNET + '?' + order_string
    return pay_url


def check_pay_state(trade_no, timeout, qr_code=None):   # 检测支付状态
    ok = False
    while True:
        time.sleep(1)
        timeout = timeout - 1
        result = create_ali().api_alipay_trade_query(trade_no)
        print(f"等待支付还剩{timeout}秒")
        if result.get('trade_status') == 'TRADE_SUCCESS':
            print('支付成功')
            out_trade_no = result['out_trade_no']
            print(trade_no, out_trade_no)
            if qr_code:
                os.remove(qr_code)
            ok = True
            return ok
        if timeout == 0:
            return ok


def create_qr_code(trade_no, amount, subject, dir):     # 创建二维码链接及图像地址
    pass


def cancel_trade(trade_no, timeout):    # 取消订单
    pass


if __name__ == '__main__':
    trade_no = str(time.time())
    price = 0.01
    subject = "vip充值"
    qr_code = create_qr_code(trade_no, price, subject, '')
    print(qr_code)
    check = check_pay_state(trade_no, 60)
    sorted()
