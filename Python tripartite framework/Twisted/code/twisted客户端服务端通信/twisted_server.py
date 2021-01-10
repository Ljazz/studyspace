from twisted.internet import reactor, task
from twisted.internet.protocol import Factory, Protocol


class EchoProtocol(Protocol):
    def __init__(self):
        self.task = task.LoopingCall(self.sendData)

    def connectionMade(self):
        self.task.start(60 * 1)

    def sendData(self):
        self.transport.write('hello client!'.encode('utf8'))


class EchoFactory(Factory):
    protocol = EchoProtocol


if __name__ == "__main__":
    factory = EchoFactory()
    reactor.listenTCP(8888, factory)
    reactor.run()
