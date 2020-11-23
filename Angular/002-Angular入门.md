# 6. 路由与多视图

## 6.1 创建多视图应用

### 6.1.1 使用`$routeProvider`创建映射

为了建立URL到AngularJS视图的映射，我们需要用到AngularJS的一个内置Provider对象`$routeProvider`，该对象用于创建路由映射，提供了一个when(path, route)方法和otherwise(params)方法，能够帮助我们把控制器、视图模板、URL关联起来。

1、 when(path, route)方法接收两个参数，具体如下：

（1）path：string类型，路由路径（和`$location.path`相对应），如果`$location.path`路径后有多余的"/"或者缺少"/", 路由依然能够匹配，并且`$location.path`会依据路由定义删除多余的"/"或者增加"/"。除此之外，在path中还可以使用占位符，需要使用“:”隔开，例如/ShowOrders:Num.

（2）route：Object类型，用于配置映射信息。该对象具有以下属性：

- controller：{string|function}类型，用于指定控制器名称或控制器构造方法。
- controllerAs：string类型，通过控制器标识符名称引用控制器。
- template：{string|function}类型，该属性可为字符串类型，用于指定视图模板，也可以是一个方法，该方法必须返回HTML模板内容。
- templateUrl：string类型，作用和template属性相同，不同的是，它用于指定视图模板文件路径
- redirectTo：重定向的地址。
- resolve：Object类型，用于指定注入控制器中的内容。

（2）otherwise(params)，该方法接收一个string类型的参数，用于匹配路由中未定义的URL。

```js
var routeModule = angular.module('routeModule', ['ngRoute']);
routeModule.config(['$routeProvider',function ($routeProvider) { 
    $routeProvider.when('/addOrder',{
        templateUrl: 'templates/add-order.html',
        controller: 'AddOrderController'
    }).when('/showOrders', {
        templateUrl: 'templates/show-orders.html',
        controller: 'ShowOrdersController'
    }).otherwise({
        redirectTo: '/addOrder'
    });
 }]);
```

如上面的代码所示angular.module()方法会返回一个模块实例，可以调用模块实例的config()方法对路由进行配置。config()方法会在模块加载时被执行，主要用于对服务进行配置。在上面的代码中，我们在config()参数方法中注入了一个AngularJS内置`$routeProvider`对象，然后使用`$routeProvider`的when()方法定义了两个路由，URL分别为/addOrder和/showOrders，把它们分别映射到视图templates/add-order.html和templates/show-orders.html

AngularJS的路由模块作为一个单独的模块，模块名称为ngRoute，我们如果在自定义的模块中使用它，需要添加ngRoute模块依赖，代码如下：

```js
var routeModule = angular.module('routeModule', ['ngRoute']);
```

### 6.1.2 创建多视图

新增订单页面

```html
<div>
    <h3>新增订单页面</h3>
</div>
```

显示订单列表页面

```html
<div>
    <h3>显示订单列表页面</h3>
</div>
```

### 6.1.3 通过路由切换视图

```html
<html>
<head>
<meta http-equiv="content-type" content="text/html; charset=UTF-8">
<script src="http://apps.bdimg.com/libs/angular.js/1.4.6/angular.min.js" rel="external nofollow"  rel="external nofollow" ></script>
<script src="http://apps.bdimg.com/libs/angular-route/1.3.13/angular-route.js" rel="external nofollow"  rel="external nofollow" ></script>

<script type="text/javascript">
angular.module('ngRouteExample', ['ngRoute'])
.controller('HomeController', function ($scope) { $scope.$route = $route;})
.controller('AboutController', function ($scope) { $scope.$route = $route;})
.config(function ($routeProvider) {
    $routeProvider.
    when('/home', {
        templateUrl: './templates/home.html',
        controller: 'HomeController'
    }).
    when('/about', {
        templateUrl: './templates/about.html',
        controller: 'AboutController'
    }).
    otherwise({
        redirectTo: '/home'
    });
});
</script>

  
</head>

<body ng-app="ngRouteExample" class="ng-scope">
  <script type="text/ng-template" id="embedded.home.html">
      <h1> Home </h1>
  </script>

  <script type="text/ng-template" id="embedded.about.html">
      <h1> About </h1>
  </script>

  <div> 
    <div id="navigation">  
      <a href="#/home">Home</a>
      <a href="#/about">About</a>
    </div>
      
    <div ng-view="">
    </div>
  </div>
</body>
</html>
```

## 6.2 通过URL向控制器传递参数

```html
var routerModule = angular.module('routeModule', ['ngRooute']);

routeModule.config(['$routeProvider', function($routeProvider){
    $routeProvider.when('/showOrder/:orderId", {
        templateUrl: 'templates/order-details.html',
        controller:'ShowOrderController'
    })
}]);

routeModule.controller("ShowOrderController", function($scope, $routeParams){
    $scope.order_id = $routeParams.orderId;
});
```
