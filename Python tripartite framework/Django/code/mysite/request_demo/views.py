from django.http import HttpResponse


def index(request):
    # return HttpResponse("hello world！")
    # return HttpResponse("请求路径：{}".format(request.path)),
    # return HttpResponse(f"请求路径：{request.path}")
    # return HttpResponse(
    #     f"请求方法：{request.method}<br /> \
    #     GET参数：{request.GET}<br /> \
    #     POST参数：{request.POST}<br /> \
    #     META信息：{request.META}<br /> \
    #     user信息：{request.user}<br /> \
    #     请求路径：{request.path}"
    # )
    values = request.META.items()
    html = [f'<tr><td>{k}</td><td>{v}</td></tr>' for k, v in values]
    html_str = '\n'.join(html)
    return HttpResponse(f"<table>{html_str}</table>")
