from abc import ABCMeta, abstractmethod


class ReceiveParcel(metaclass=ABCMeta):
    """接收包裹抽象类"""
    def __init__(self, name):
        self.__name = name

    def getName(self):
        return self.__name

    @abstractmethod
    def receive(self):
        pass


class TonyReception(ReceiveParcel):
    """Tony接收"""
    def __init__(self, name, phoneNum):
        super().__init__(name)
        self.__phoneNum = phoneNum

    def getPhoneNum(self):
        return self.__phoneNum

    def receive(self, parcelContent):
        print("货物主人：{}，手机号：{}".format(self.getName(), self.getPhoneNum()))
        print("接收到一个包裹，包裹内容：{}".format(parcelContent))


class WendyReception(ReceiveParcel):
    """Wendy代收"""
    def __init__(self, name, receiver):
        super().__init__(name)
        self.__receiver = receiver

    def receive(self, parcelContent):
        print("我是{}的朋友，我来帮他代收快递！".format(self.__receiver.getName()))
        if self.__receiver is not None:
            self.__receiver.receive(parcelContent)
        print("代收人：{}".format(self.getName()))


def testReceiveParcel():
    tony = TonyReception("Tony", "18512345678")
    print("Tony接收：")
    tony.receive("雪地靴")
    print()

    print("Wendy代收：")
    wendy = WendyReception("Wendy", tony)
    wendy.receive("雪地靴")


if __name__ == "__main__":
    testReceiveParcel()
