class Apple:
    def fun1(self):
        return 'normal'
    @staticmethod
    def fun2():
        return 'staticmethod'
    @classmethod
    def fun3(cls):
        return 'classmethod'

print(Apple.fun1())
print(Apple.fun2())
print(Apple.fun3())
print('-' * 20)

a = Apple()
print(a.fun1())
print(a.fun2())
print(a.fun3())