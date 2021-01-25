from django.http import HttpResponseRedirect
from django.shortcuts import render
from django.contrib.auth.models import User
from .forms import LoginForm


# 如果登录成功，设置cookie
def login(request):
    if request.method == 'POST':
        form = LoginForm(request.POST)

        if form.is_valid():
            username = form.cleaned_data['username']
            password = form.cleaned_data['password']

            user = User.objects.filter(username__exact=username, password__exact=password)

            if user:
                response = HttpResponseRedirect('/index/')
                # 将username写入浏览器cookie，失效时间为360秒
                response.set_cookie('username', username, 360)
                return response
            else:
                return HttpResponseRedirect('/login/')
        else:
            form = LoginForm()
        return render(request, 'users/login.html', {'form': form})


# 通过cookie判断用户是否已经登录
def index(request):
    # 提取浏览器中的cookie，如果不为空，表示已登录账号
    username = request.COOKIES.get('username', '')
    if not username:
        return HttpResponseRedirect('/login/')
    return render(request, 'index.html', {'username': username})
