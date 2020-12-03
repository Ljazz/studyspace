<!-- TOC -->

- [路由](#路由)
  - [1. 前端路由简介](#1-前端路由简介)
  - [2. 前端路由的实现原理](#2-前端路由的实现原理)
    - [2.1 history API方式](#21-history-api方式)
    - [2.2 Hash方式](#22-hash方式)
  - [3. react-router路由配置](#3-react-router路由配置)
    - [3.1 react-router的安装](#31-react-router的安装)
    - [3.2 路由配置](#32-路由配置)
    - [3.3 默认路由](#33-默认路由)
    - [3.4 嵌套路由](#34-嵌套路由)
    - [3.5 重定向](#35-重定向)
  - [4. react-router下的history](#4-react-router下的history)
    - [4.1 browserHistory](#41-browserhistory)
    - [4.2 hashHistory](#42-hashhistory)
    - [4.3 createMemoryHistorybr](#43-creatememoryhistorybr)
  - [5. react-router路由切换](#5-react-router路由切换)
    - [5.1 Link标签](#51-link标签)
    - [5.2 history属性](#52-history属性)
    - [5.3 传参](#53-传参)
  - [6 进入和离开的Hook](#6-进入和离开的hook)
    - [6.1 onEnter简介](#61-onenter简介)
    - [6.2 onLeave简介](#62-onleave简介)

<!-- /TOC -->

# 路由

## 1. 前端路由简介

前端路由其实就是通过JavaScript来配置路由。单页应用基本也是前后端分离的，因此后端不提供路由。前端路由和后端路由实现原理是一样的，只是实现方式不同。

## 2. 前端路由的实现原理

前端路由的主要是实现方式

- Hash
- HTML5的history API

### 2.1 history API方式

Mozilla官方文档描述

```txt
// 只读属性。返回一个整数，该整数表示会话历史中元素的数目，包括当前加载的页。例如，
在一个新的选项卡加载的一个页面中，这个属性返回1
window.history.length
Returns the number of entries in the joint session history.
// 允许Web应用程序在历史导航上显式地设置默认滚动恢复行为。此属性可以是自动的（auto）
或者手动的（manual）
window.history.scrollRestoration[=value]
Returns the scroll restoration mode of the current entry in the session history.Can be set, to change the scroll restoration mode of the current entry in the session history.
// 返回一个表示历史堆栈顶部的状态值。这是一种可以不必等待popstate事件而查看
状态的方式。
window.history.state
Returns the current serialized state, deserialized into an object.
// 通过当前页面的相对位置从浏览器历史记录（会话记录）加载页面
window.history.go([delta])
Goes back or forward the specified number of steps in the joint session history.A zero delta will reload the current page.If the delta is out of range, does nothing.
// 往上一页，用户可单击浏览器左上角的返回按钮模拟此方法，等价于 history.go(-1)
window.history.back()
Goes back one step in the joint session history.If there is no previous page, does nothing.
// 往下一页，用户可单击浏览器左上角的前进按钮模拟此方法，等价于 history.go(1)
window.history.forward()
Goes forward one step in the joint session history.If there is no next page, does nothing.
// 按指定的名称和URL（如果提供该参数）将数据push进会话历史栈，数据被DOM进行不透明
处理；可以指定任何可以被序列化的javaScript对象

window.history.pushState(data,title[,url])
Pushes the given data onto the session history, with the given title, and, if provided and not null, the given URL.
// 按指定的数据，名称和URL（如果提供该参数），更新历史栈上最新的入口。这个数据被DOM 进行了不透明处理。可以指定任何可以被序列化的javaScript对象。
window.history.replaceState(data,title[,url])
Updates the current entry in the session history to have the given data, title, and, if provided and not null, URL.
```

history.pushState()和history.replaceState()，这是history新增的API，可以为浏览器提供历史栈。这个两个API都接收3个参数

- 状态对象(state object)：一个JavaScript对象，与用pushState()方法创建的新历史记录条目关联。无论何时用户导航到新创建的状态，popState事件都会被触发，并且事件对象的state属性都包含历史记录条目的状态对象的复制。
- 标题(title)：FireFox浏览器目前会忽略该参数，虽然以后可能会用上，考虑到未来可能会对该方法进行修改，传一个空字符串会比较安全。或者也可以传入一个简短的标题，标明将要进入的状态。
- 地址(URL)：新的历史记录条目的地址。浏览器不会在调用pushState()方法后加载该地址，但之后可能会视图加载，例如用户重启浏览器。新的URL不一定是绝对路径；如果是相对路径，它将以档期那URL为基准；传入的URL与当前URL应该是同源的，否则，pushState()会抛出异常。该参数是可选的；如果不指定，则为文档当前URL。

pushState()会增加一条新的历史记录，当使用了这个接口，浏览器的返回按钮能返回到上一次访问的URL，在浏览器的历史记录中也会多出一条记录。而replaceState()会替换当前历史记录，浏览器的返回按钮无法返回到上一次的记录，浏览器的历史记录中也将用当前的URL替换为上一次的URL作为浏览记录存储。值得注意的是，pushState()和replaceState()这两个接口不会引起页面刷新，这一点同HASH一样，这样就实现了也买你无刷新情况下改变URL。接下来只要监听浏览器URL改变的事件即可对页面展示内容进行切换，在history中的事件是popState。

注意：pushState()和replaceState()不支持跨域。

### 2.2 Hash方式

Hash（哈希）也称作锚点，指的是URL中#号以后的字符。基本身用于页面定位，它可以配合id的元素显示在可视区域内。同样地，它可以改变浏览器URL，同时做到不刷新也买你。但依旧可以通过hashChange事件去监听其变化，从而进行页面展示内容的切换。

hash方式实现路由示例

```js
import React from 'react';
import {render} from 'react-dom';   // 各个路由对应的组件

const Home = () => <div>Home</div>
const About = () => <div>About</div>
const Inbox = () => <div>Inbox</div>

const App = React.createClass({
    getInitialState(){
        return{
            route: window.location.hash.substr(1) // 获取浏览器Hash值，并存储在state中
        }
    },
    componentDidMount(){
        //利用监听Hash事件去改变路由值
        window.addEventListener('hashchange', ()=>{
            this.setState({
                route: window.location.hash.substr(1)
            })
        })
    },
    render(){
        let Child;
        // 根据state的route值来相应当前内容的动态展示
        switch(this.state.route){
            case '/about': Child = About; break;
            case '/inbox': Child = Inbox; break;
            default: Child=Home;
        }
        return(
            <div>
                <a href="#/about">About</a>
                // 单击a标签会改变浏览器地址栏URL，但不会刷新页面
                <a href="#/inbox">Inbox</a>
                <Child />
            </div>
        )
    }
})

React.render(<App />, document.body)
```

## 3. react-router路由配置

### 3.1 react-router的安装

react-router 3.x安装：

> npm install -g react-router

react-router4版本之后安装：

> npm install react-router-dom -g

### 3.2 路由配置

以Webpack作为项目的模块管理器为例，react-router3.x的配置示例：

```jsx
import React from 'react';
import ReactDom from 'react-dom';
import {Route, IndexRoute, hashHistory} from 'react-router';
import App from './app/containers/app/';
import Home from './app/containers/home/';
import User from './app/containers/user/';

ReactDOM.render(
    <Router history={hashHistory}>
        <Route path="/" component={App}>
        {/* 当url为/时渲染Home*/}
        <IndexRoute component={Home} />
        <Route path="/user" component={User} />
        </Route>
    </Router>,
    document.getElementById('root')
);
```

然后在App组件中，使用`this.props.children`属性配置路由的组件：

```js
const Inbox = React.createClass({
    render(){
        <div>
        {/*渲染这个child路由组件*/}
        {this.props.children}
        </div>
    }
})
```

上述代码中，Router标签代表路由器。路由器中的history表达式react-router路由的模式。

`<Route>`是react-router最重要的组件，职责是在其path属性与某个location匹配时呈现指定的视图，用于配置路由和组件的对应关系。

`<Route path="/list" component={List} />`

path表示路由的路径，component代表路由对应的页面组件。其中/代表的根路径。

### 3.3 默认路由

在react-router3.x中可以指定默认路由(IndexRoute)

```js
// v3
<Router>
<Route path="/" component={App}>
<IndexRoute component={Home} />
<Route path="about" component={About} />
<Route path="help" component={Help} />
</Route>
</Router>
```

react-router v4中是没有`<IndexRoute>`组件，而是通过`<Switch>`组件提供相似的功能

react-router 4.x中指定默认路由示例

```js
// v4
const App = () => {
    <Switch>
        <Route exact path="/" component={Home} />
        <Route path='/about' component={About} />
        <Route path='/help' component={Help} />
    </Switch>
}
```

### 3.4 嵌套路由

react-router可以实现路由的嵌套，嵌套路由被描述成一种树形结构。react-router会深度遍历整个路由配置来匹配给出的路径。

```js
// v3
<Route path='parent' component={Parent}>
<Route path='child1' component={Child1} />
<Route path='child2' component={Child2} />
</Route>
```

```js
// v4
<Route path='parent' component={Parent} />
const Parent = () => (
    <div>
        <Route path='child1' component={Child1} />
        <Route path='child2' component={Child2} />
    </div>
)
```

### 3.5 重定向

在react-router 3.x版本中，如果要从一个路径重定向到另一个路径，可以使用`<IndexRedirect>`.

```js
// v3
<Route path='/' component={App}>
<IndexRedirect to='/AnotherApp' />
<Route path='anotherApp' component={anotherApp} />
</Route>
```

react-router 4.x中，使用`<Redirect>`

```js
// v4
import {Route, Redirect} from 'react-router';
<Route exact path='/' render={()=>(
    loggedIn ?(
        <Redirect to='/homePage' />
    ):(
        <LoginPage />
    )
)} />
```

## 4. react-router下的history

react-router中有history属性，用于监听浏览器地址栏的改变，并解析这个URL转换为location对象，从而匹配react-router配置的路由去渲染对应的视图。

react-router中有3种形式

- browserHistory
- hashHistory
- createMemoryHistorybr

### 4.1 browserHistory

browserHistory是基于使用浏览器history API实现的，也是react-router应用推荐的路由模式。可以从react-router中引入使用

```js
import {browerHistory} from 'react-router';

render(
    <Router history={browserHistory} routes={routes} />,
    document.getElementById('app')
)
```

这种模式的优点更像“真实的”URL，形如/index/homePage。这种模式的缺点就是当用户在子路由刷新或向服务器直接请求子路由，则会显示找不到，给出404报错。这时候就需要服务器去配置并处理这些路由的访问。

Nginx，可以使用try_files指令

```txt
server {
    ...
    location / {
        try_files $uri /index.html;
    }
}
```

Apache服务器，根目录下创建一个.htaccess文件夹

```txt
RewriteBase /
RewriteRule ^index\.html$-[L]
RewriteCond %{REQUEST_FILENAME}!-f
RewriteCond %{REQUEST_FILENAME}!-d
RewriteRule ./index.html[L]
```

### 4.2 hashHistory

hashHistory是基于哈希（#）实现的。形如/#/index/homePage?_k=adsis。示例如下

```js
import {hashHistory} from 'react-router'

render(
    <Router history={hashHistory} routes={routes} />,
    document.getElementById('app')
)
```

### 4.3 createMemoryHistorybr

createMemoryHistory模式用于服务端渲染。它不会再地址栏被操作或读取，但会在内存中进行历史记录的存储。

它与其他两种模式不同的是，需要手动去创建

```js
const history = createMemoryHistory(location)
```

## 5. react-router路由切换

### 5.1 Link标签

Link是react-router中用于路由相互跳转的其中一种方法。其本质就是一个被处理过的`<a>`标签，它可以接收Router的状态

Link标签实现路由切换示例

```js
render(){
    return(
        <div>
            <ul>
                <li><Link to="/">点击跳转首页</Link></li>
                <li><Link to="/login">点击跳转登录页</Link></li>
            </ul>
        </div>
    )
}
```

Link可以直到那个Route的连接时激活状态，并可以为该链接添加actionClassName或activeStyle属性。这就使得当用户在Tab切换的时候，可以方便地设置激活时的样式展示。

示例

```js
<Link to='/home' activeStyle={{color:'red'}}>Home</Link>
<Link to='/about' activeStyle={{color:'yellow'}}>About</Link>

// 或
<Link to='/home' activeClassName='active'>Home</Link>
<Link to='/about' activeClassName='active'>About</Link>
```

**注意**：如果连接到跟路由“/”，要使用`<IndexLink>`。

### 5.2 history属性

react-router3.x中，路由的跳转一般这样处理：

- 从react-router导出browserHistory
- 使用browserHistory.push()等方法进行路由跳转

react-router3.x示例中的路由跳转示例：

```js
import browserHistory from 'react-router';

// ...

browerHistory.push('/user');
```

react-router4.x中，不提供browserHistory等方法，而是通过高阶组件withRouter去使用history的方法实现路由跳转。

react-router4.x示例中的路由跳转示例

```js
import React, {Component} from 'react';
import {withRouter} from 'react-router-dom';

class User extends Component{
    constructor(props){
        super(props);
    }
    linkTo(){
        this.props.history.push('/about');
    }
    render(){
        return()
    }
}
export default withRouter(User);
```

### 5.3 传参

应用跳转时可能会涉及参数传递，再react-router中传参也很简答，先通过Route的path进行配置，然后再Link的to属性中添加需要的参数，最后再跳转后的页面去获取。

示例1：react-router中，使用this.props.params.query获取参数。

```js
<Router history={history}>
    <Route path="user/:id" component={User} />  // 路由中配置 :id
</Router>

//...

<Link to={{pathname:'/user/{id}'}} activeClassName='active'>
Click to userPage
</Link>
```

示例2：react-router中，使用this.props.location.query获取参数

```js
<Router history={history}>
    <Route path="user" component={User} />  // 路由中配置 :id
</Router>

//...

<Link to={{pathname:'/User', query:{id:id} }} activeClassName='active'>
Click to userPage
</Link>
```

示例3：react-router中使用this.props.params.id获取参数

如果采用history跳转，那么再user页面中可以使用this.params.id获取参数。

```js
<Router history={hashHistory}>
<Route path='/user/:id' component={User}></Route>
</Router>
// ...

hashHistory.push('/user/888')
```

## 6 进入和离开的Hook

Route可以定义onEnter和onLeave两个Hook，这两个钩子会在路由跳转确认时触发一次。这对于权限验证和路由跳转前数据持久化保存有很大作用

### 6.1 onEnter简介

onEnter Hook会在即将进入路由时触发，会从最外层的父路由开始，直到最下层子路由结束。它可以接收3个参数：

```js
type EnterHook = (nextState: RouterState, replaceState:RedirectFunction, callback?:Function) => any;
```

第一个参数nextState表示它接收的下一个router state，第2个参数replaceState function用于触发URL的变化，第3个参数callback用于设置回调函数，以便于继续往下执行。

**注意**：在onEnter Hook中使用callback会让变换过程处于阻塞状态，直到callback被回到。如果不能快速回调，这可能会导致整个ui失去响应

### 6.2 onLeave简介

onLeave Hook会在即将离开路由时触发，从最下层的子路由开始，直到最外层的父路由结束。它时一个用户自定义的函数，用于在离开时被调用。
