from django.shortcuts import render
from django.http import HttpResponseRedirect
from django.contrib.auth.models import User
from .forms import LoginForm


# 若登录成功，设置session
def login(request):
    if request.method == 'POST':
        form = LoginForm(request.POST)

        if form.is_valid():
            username = form.cleaned_data['username']
            password = form.cleaned_data['password']

            user = User.objects.filter(username__exact=username, password__exact=password)

            if user:
                # 将username写入session，存入服务器
                request.session['username'] = username
                return HttpResponseRedirect('/index/')
            else:
                return HttpResponseRedirect('/login/')
        else:
            form = LoginForm()
        return render(request, 'users/login.html', {'form': form})


# 通过session判断用户是否登录
def index(request):
    # 获取session中的username
    username = request.session.get('username', '')
    if not username:
        return HttpResponseRedirect('/login/')
    return render(request, 'index.html', {'username': username})
