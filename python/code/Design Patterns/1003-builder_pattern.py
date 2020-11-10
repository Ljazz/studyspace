from abc import ABCMeta, abstractmethod


class Toy(metaclass=ABCMeta):
    """玩具"""
    def __init__(self, name):
        self._name = name
        self.__components = []

    def getName(self):
        return self._name

    def addComponent(self, component, count=1, unit="个"):
        self.__components.append([component, count, unit])
        print("{} 增加了 {} {}{}".format(self._name, count, unit, component))

    @abstractmethod
    def feature(self):
        pass


class Car(Toy):
    """小车"""
    def feature(self):
        print("我是 {}, 我可以快速奔跑...".format(self._name))


class Manor(Toy):
    """
    庄园
    """
    def feature(self):
        print("我是 {}, 我可供观赏， 也可用来游玩!".format(self._name))


class ToyBuilder:
    """
    玩具构建者
    """
    def buildCar(self):
        car = Car("迷你小车")
        print("正在构建 {}...".format(car.getName()))
        car.addComponent("轮子", 4)
        car.addComponent("车身", 1)
        car.addComponent("发动机", 1)
        car.addComponent("方向盘")
        return car

    def buildManor(self):
        manor = Manor("淘淘小庄园")
        print("正在构建 {}".format(manor.getName()))
        manor.addComponent("客厅", 1, "间")
        manor.addComponent("卧室", 2, "间")
        manor.addComponent("书房", 1, "间")
        manor.addComponent("厨房", 1, "间")
        manor.addComponent("花园", 1, "个")
        manor.addComponent("围墙", 1, "堵")
        return manor


def testBuilder():
    builder = ToyBuilder()
    car = builder.buildCar()
    car.feature()
    print()
    mannor = builder.buildManor()
    mannor.feature()


if __name__ == "__main__":
    testBuilder()
