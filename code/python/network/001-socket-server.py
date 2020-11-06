import socket

s = socket.socket()
s.bind(('192.168.2.10', 1234))
s.listen(5)
while 1:
    conn, address = s.accept()
    print('a new connection from', address)
    conn.send('who are you?'.encode())
    conn.close()
