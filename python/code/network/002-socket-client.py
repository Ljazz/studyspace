import socket

s = socket.socket()
s.connect(('192.168.2.10', 1234))
data = s.recv(1024)
s.close()
print('Received', data)
