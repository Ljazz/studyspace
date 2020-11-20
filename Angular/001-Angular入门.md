# Angular

## 1. Angular简述

### 1.1 核心概念

Angular框架有7大核心概念，它们是Angular的重要组成部分

![Angular的七大核心概念](images/Angular的七大核心概念.jpg)

Angular以HTML作为模板语言并扩展HTML元素及属性，使得应用组件开发保持高度清晰、一致。

### 1.2 AngularJS应用构成元素

- 模型（Model）：AngularJS程序中用于展示到页面的数据，本质是一个JavaScript对象。
- 视图（View）：从用户角度来看，视图就是用户所看到的网页内容；从AngularJS应用的角度来说，视图则是AngularJS指令与表达式经过解析后的DOM元素。
- 控制器（Controller）：AngularJS应用中用于处理业务逻辑的JavaScript方法。
- 作用域（Scope）：可以把作用域理解为一个容器，在控制器中可以访问这个容器，然后往容器中放入一些模型数据，在视图中就可以通过表达式将数据展现给用户
- 指令（Directives）：扩展的HTML属性或标签，能够被AngularJS框架识别，根据不同的指令执行相应的动作。例如，ng-app指令，作为html元素的扩展属性，能够被AngularJS识别，从而启动Angular框架
- 表达式（Expressions）：用于单向页面输出信息
- 模板（Template）：AngularJS以HTML作为模板语言，AngularJS模板实际上就是HTML片段。

### 1.3 AngularJS表达式

1、表达式的定义式

---
> {{expression}}

---

2、表达式中的四则运算

AngularJS表达式支持加减乘除四则运算及字符串拼接运算。

```html
<!DOCTYPE html>
<html ng-app>
<head>
    <meta charset="UTF-8">
    <title>ch01_05</title>
    <script type="text/javascript" src="../angular-1.5.5-angular.js"></script>
</head>
<body>
    <div>1 + 2 = {{1 + 2}}</div>
    <div>5 - 3 = {{5 - 3}}</div>
    <div>5 * 3 = {{5 * 3}}</div>
    <div>10 / 2 = {{10 / 2}}</div>
    <div>"hello " + "world" = {{"hello " + "world"}}</div>
</body>
</html>
```

3、表达式中的逻辑运算

AngularJS表达式除了支持算术运算外，还支持逻辑运算。

```html
<!DOCTYPE html>
<html ng-app>
<head>
    <meta charset="UTF-8">
    <title>ch01_02</title>
    <script type="text/javascript" src="../angular.js"></script>
</head>
<body>
    <div> 1 || 0 = {{1 || 0}}</div>
    <div>true && false = {{true && false}}</div>
</body>
</html>
```

4、表达式与作用域

AngularJS表达式可以访问作用域中的数据，把数据输出到HTML页面

```html
<!DOCTYPE html>
<html ng-app>
<head>
    <meta charset="UTF-8">
    <title>ch01_03</title>
    <script type="text/javascript" src="../angular.js"></script>
</head>
<body ng-init="person={'name': 'jane'};arr=['angularjs', 'jquery', 'react']">
    <div>{{person.name}}</div>
    <div>{{arr[0]}}</div>
</body>
</html>
```

上述示例中，通过ng-init指令向作用域中增加一个person对象和arr数组。

## 2. 双向数据绑定

### 2.1 AngularJS双向数据绑定

数据绑定是AngularJS框架在视图（DOM元素）与作用域之间建立的数据同步机制。所谓“双向”，是指界面的操作能够实时同步到作用域中，作用域中的数据修改也能够实时回显到界面中。

作用域可以被视为一个容器，里面有一些基于key-value的数据。

如下图，有两个输入框，当输入框内容发生变化时，AngularJS框架就把表单内容同步到作用域中对应的变量中，而当改变作用域中的变量时，AngularJS又会把修改后的变量值同步到表单中，这就是AngularJS的双向数据绑定。

![AngularJS双向数据绑定图解](./images/AngularJS双向数据绑定图解.jpg)

### 2.2 ng-model指令

AngularJS内部为我们提供了一个内置指令ng-model用于建立数据绑定。该指令只能用在表单元素上。

---
> `<input type='text' name="uname" ng-model="uname" />`

---

上述代码中，在input输入框上添加ng-model指令后，AngularJS框架就会子对应的作用域中创建一个uname属性和该输入框进行绑定。

**示例**：原生JS通过document对象的getElementById()方法获取输入框对象，响应输入框keyup事件。

```html
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>原生JS操作DOM</title>
</head>
<body>
    <div>用户名：<input = type='text' name='uname' id='uname' /></div>
    <div><span id='info'></span></div>
    <script>
        var uname = document.getElementById('uname');
        var info = document.getElementById('info');
        uname.onkeyup = function(){
            info.innerHTML = uname.value;
        }
    </script>
</body>
</html>
```

**示例**：AngularJS数据绑定机制实现

```html
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>AngularJS数据绑定机制</title>
    <script type="text/javascript" src="../angular.js"></script>
</head>
<body>
    <div>用户名：<input type='text' ng-model='uname' /></div>
    <div>{{uname}}</div>
</body>
</html>
```

### 2.3 ng-bind指令

ng-bind指令是和数据绑定相关的另一个指令，其作用是实现作用域到视图的单向数据绑定，和表达式功能类似，可用于向界面中输出作用域中的数据。

AngularJS表达式`{{name}}`会受到网络影响。若遇到网络问题，就会导致AngularJS加载缓慢，浏览器就会直接将AngularJS表达式当作字符串渲染到页面中。ng-bind指令在AngularJS没有加载完毕的时候是不会解析执行的，只有AngularJS架子完毕才会执行。

### 2.4 数据绑定示例：价格计算器

```html
<!DOCTYPE html>
<html ng-app>
<head>
    <meta charset="utf-8">
    <title>价格计算器</title>
    <script type="text/javascript" src="../angular.js"></script>
</head>
<body ng-init="price=10;num=1">
    <div>单价：<input type="number" ng-model="price" /></div><br />
    <div>数量：<input type="number" ng-model="num" /></div><br />
    <div>总价：{{price * num}}</div>
</body>
</html>
```

## 3. AngularJS与MVC

### 3.1 MVC模式简介

MVC是一种软件架构模式，独立于任何一门语言。

MVC是Model（模型）、View（视图）、Controller（控制器）的首字母缩写，MVC核心思想是把数据的管理、业务逻辑控制和数据的展示分离开，使程序的逻辑性和可维护性更强。它们之间的关系如下图

![模型、视图、控制器关系图](./images/模型、视图、控制器关系图.jpg)

View（视图）为用户可操作的软件界面，用户通过视图和程序进行交互，在视图中会触发不同的事件，例如单击按钮、输入文字等，不同的事件能够触发控制器执行相应的业务逻辑处理。

Controller（控制器）主要用于相应用户请求，在控制器中可操作模型数据，进行业务逻辑处理，根据处理结果分发到不同的视图。

Model（模型）为程序中的模型数据，是控制器与视图之间传递信息的载体。

### 3.2 AngularJS中的MVC

AngularJS中的MVC分别指

- Model（模型）：作用域对象（例如$rootScope对象）中的属性
- View（视图）：DOM元素，从用户角度来看就是HTML页面，在View中可以通过AngularJS表达式访问模型数据
- Controller（控制器）：用户自定义的构造方法，作用域中的模型数据可以通过依赖注入的方式注入控制器中。

1、AngularJS控制器的定义

AngularJS控制器是一个构造方法。JavaScript方法可以作为对象模板实例化对象，当方法作为对象模板时，方法本身即为对象的构造方法。所以可以像定义一个普通方法一样定义一个控制器。

```html
<script type="text/javascript">
 function LoginController($scope, $log){
     $scope.name="admin";
     $scope.pword="123456";
 }
</script>
```

除了上述的方式外，还可以使用模块实例的controller()方法来声明一个控制器。该方法可接收两个参数，第一个参数为控制器名称，第二个参数为一个匿名方法，即控制器的构造方法，具体使用方法如下：

```html
<script type="text/javascript">
    var app = angular.module("app", []);
    app.controller("LoginController", function($scope, $log){
        $scope.name="admin";
        $scope.pword="123456";
    });
</script>
```

上述代码中，AngularJS框架在window对象中增加了一个全局的angular对象，可以调用angular对象的module方法返回一个模块实例，然后调用模块实例的controller()方法来声明一个控制器。在定义控制器时制定了两个参数，即$scope和$log：$scope时作用域对象，是控制器与视图之间传递信息的载体；$log为AngularJS框架内置的日志服务对象，用于向控制台中输入日志信息。当为控制器构造方法指定这两个参数后，表示控制器依赖于这两个对象，控制器实例化时会把这两个对象注入控制器中。

**注**：AngularJS1.3版本之后已不再支持全局控制器，第一种控制器定义方式只适用于AngularJS1.3之前的版本。

2、控制对象的实例化

控制器对象的实例化用的时AngularJS内置的ng-controller指令。ng-controller的使用方法和ng-app指令类似，也是作为标签的扩展性使用。

---
> `<div ng-controller="LoginController"></div>`

---

AngularJS框架遇到ng-controller指令时会根据ng-controller指令指定的控制器名称查找控制器构造方法，然后使用对应的构造方法实例化控制器对象，并将控制器依赖的对象注入控制器对象中。

```html
<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>价格计算器</title>
    <script src="../angular.js"></script>
</head>
<body>
    <div ng-app='myApp' ng-controller="myCtrl">
        <div>单价：<input type="number" ng-model="price"></div><br>
        <div>数量：<input type="number" ng-model="num"></div><br>
        <div>总价：{{totalPrice()}}</div>
    </div>
    <script>
        function CalcController($scope){
            $scope.price = 10;
            $scope.num = 1;
            $scope.totalPrice = function(){
                return $scope.price * $scope.num;
            }
        }
        var app = angular.module('myApp', []);
        app.controller("myCtrl", CalcController)
    </script>
</body>
</html>
```

### 3.4 控制器的作用域范围

使用ng-controller指令实例化控制器时会产生一个新的作用域对象，而且同一个页面中可以使用多个ng-controller指令来实例化不同的控制器对象。但是需要注意的是，每个控制器对应的作用域对象只能与ng-controller指令所在标签的开始标签与结束标签之间的DOM元素建立数据绑定。

```html
<!doctype html>
<html ng-app="app">
<head>
    <meta charset="UTF-8">
    <title>ch03_02</title>
    <script type="text/javascript" src="../angular.js">
    </script>
</head>
<body>
    <div ng-controller="UserController" style="border:#ccc solid 1px;">
        用户名：<input type="text" ng-model="name" placeholder="用户名"/><br>
        密码：<input type="password" ng-model="pword" placeholder="密码"/><br>
        <button>提交</button><br>
        <p>您输入的用户名：{{name}}</p>
        <p>您输入的密码：{{pword}}</p>
    </div>
    <br/>
    <div ng-controller="InfoContoller" style="border:#ccc solid 1px;">
        个人爱好：<input type="text" ng-model="love" placeholder="个人爱好"/>
        <p>您输入的个人爱好：{{love}}</p>
    </div>
    <script>
        function UserController($scope,$log) {
            $scope.name="admin";
            $scope.pword="123456";
            $log.info("UserController->name:" +$scope.name);
            $log.info("UserController->pword:" + $scope.pword);
        }
        function InfoContoller($scope,$log) {
            $scope.love="足球";
            $log.info("InfoContoller->name:" + $scope.name);
            $log.info("InfoContoller->pword:" + $scope.pword);
            $log.info("InfoContoller->love:" + $scope.love);
        }
        var app = angular.module("app",[]);
        app.controller("UserController",UserController);
        app.controller("InfoContoller",InfoContoller);
    </script>
</body>
</body>
</html>
```

### 3.5 控制器中处理DOM事件

AngularJS应用中的DOM事件处理可以在控制器中完成。AngularJS框架提供了一系列的事件指令，这些指令是在原生的JavaScript事件名称前增加“ng-”前缀，例如：ng-click、ng-keyup等。

```html
<!doctype html>
<html ng-app="app">
<head>
    <meta charset="UTF-8">
    <title>ch03_03</title>
    <script type="text/javascript" src="../angular.js">
    </script>
</head>
<body>
<div class="container" ng-controller="LoginController">
<div class="login-title text-center">
        <h1><small>登录</small></h1>
</div>
<div class="form">
    <div class="form-group">
        <div class="col-xs-12">
            <div class="input-group">
                <span class="input-group-addon">
                <span class="glyphicon glyphicon-user"></span>
                </span>
                <input type="text" ng-model="uname" class="form-control"
                            placeholder="用户名">
            </div>
        </div>
    </div><br/><br/>
    <div class="form-group">
        <div class="col-xs-12">
            <div class="input-group">
                <span class="input-group-addon">
                <span class="glyphicon glyphicon-lock"></span>
                </span>
                <input type="text" ng-model="pword" class="form-control"
                            placeholder="密码">
            </div>
        </div>
    </div><br/><br/>
    <div class="form-group form-actions">
        <div class="col-xs-4 col-xs-offset-4 ">
            <button ng-click="login()" type="submit" class="btn btn-sm btn-info">
                <span class="glyphicon glyphicon-log-in"> </span>
                    登录
            </button>
            <button type="reset" class="btn btn-sm btn-info">
                <span class="glyphicon glyphicon-off"> </span>
                重置
            </button>
        </div>
    </div>
</div>
</div>
<script>
    function LoginController($scope) {
        $scope.login = function(){
            if($scope.uname == "admin"
                && $scope.pword == "admin"){
                alert("登录成功！");
            } else {
                alert("用户名密码错误，请检查！");
            }
        }
    }
    var app = angular.module("app",[]);
    app.controller("LoginController",LoginController);
</script>
</body>
</html>
```

## 4. 应用模块化
