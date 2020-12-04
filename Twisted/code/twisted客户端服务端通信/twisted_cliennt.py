import datetime
from twisted.internet import error, reactor
from twisted.internet.protocol import Protocol, ClientFactory


class SBProtocol(Protocol):
    def __init__(self):
        """
        Constructor
        """
        pass

    def dataReceived(self, data):
        """
        接收数据时调用
        :param data:
        :return:
        """
        print('{} - {}'.format(datetime.datetime.now().isoformat(), data.decode('utf8')))

    def connectionMade(self):
        """
        建立连接时调用
        :return:
        """
        print('{}  [connectionMade]  -  Connected to the server'.format(datetime.datetime.now().isoformat()))

    def connectionLost(self, reason=error.ConnectionDone):
        """
        连接断开时调用
        :param reason:
        :return:
        """
        print('{}  [connectionLost]  -  Disconnected to the server'.format(datetime.datetime.now().isoformat()))


class ProtocolFactory(ClientFactory):
    def __init__(self):
        self.protocol = None

    def startedConnecting(self, connector):
        """
        建立连接时调用
        :return:
        """
        print('{}  [startedConnecting]  -  Started to connect.'.format(datetime.datetime.now().isoformat()))

    def buildProtocol(self, addr):
        """
        在创建Protocols的回调
        :param addr:
        :return:
        """
        print('{}  [buildProtocol]  -  {}'.format(datetime.datetime.now().isoformat(), addr))
        self.protocol = SBProtocol()
        return self.protocol

    def clientConnectionLost(self, connector, reason):
        """
        连接断开时调用
        :param connector:
        :param reason:
        :return:
        """
        print('{}  [clientConnectionLost]  -  Reason: {}'.format(datetime.datetime.now().isoformat(), reason))
        connector.connect()

    def clientConnectionFailed(self, connector, reason):
        """
        连接Server失败时的回调
        :param connector:
        :param reason:
        :return:
        """
        print('{}  [clientConnectionFailed]  -  Reason: {}'.format(datetime.datetime.now().isoformat(), reason))
        connector.connect()


if __name__ == "__main__":
    factory = ProtocolFactory()
    reactor.connectTCP('127.0.0.1', 8888, factory)
    reactor.run()
