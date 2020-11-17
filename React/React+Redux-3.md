# React+Redux的数据流管理

## 1. Flux架构

> Flux是一种应用架构，或者说是一种思想。

### 1.1 MVC和MVVM

1、MVC简介

MVC（Model View Controller），是Modle（模型）-View（视图）-Controller（控制器）的缩写。MVC不是框架，不是设计模式，也不是软件架构，而是一种架构模式。

框架、设计模式、软件架构、架构模式区别：

- 框架（Framework）：是一个系统的可重用设计，表现为一组抽象的可交互方法。就像若干类的构成，设计若干构件，以及构建之间相互依赖关系、责任分配和流程控制等。
- 设计模式（Design Pattern）：是一套被反复使用，多数人只晓得、经过发呢类的代码设计经验总结。其目的是为了代码的可重用性、让代码更容易被他人理解、保证代码的可靠性
- 软件架构（Software architecture）：是一系列相关的抽象模式，用于指导大型软件系统各个方面的设计。软件架构是一个系统的草图，软件体系结构是构建计算机软件实践的基础。
- 架构模式（风格）：一个架构模式描述软件系统里基本的结构组织或纲要。架构模式提供一些事先定义好的子系统，指定它们的责任，并给出把它们组织在一起的法则和指南。一个架构模式常常可以分解成很多个设计模式的联合使用。

### 1.2 Flux

Flux核心思想是利用单向数据流和逻辑单向流来应对MVC架构中出现状态混乱的问题

Flux由3部分组成

- Dispatcher(分发器)
  - 用于分发事件
- Store
  - 用于存储应用状态，同时响应事件并更新数据
- View
  - 表示视图层，订阅来自Store的数据，渲染到页面。

Flux的核心是单向数据流，运作方式是

> Action -> Dispatcher -> Store -> View

流程如下：

1. 创建Action（提供给Dispatcher）
2. 用户在View层交互去触发Action。
3. Dispatcher收到Action，要求Store进行相应的更新。
4. Store更新，通知View去更新
5. View收到通知，更新页面。

1、Dispatcher简介

Dispatcher是一个调度中心，管理所有数据流，所有事件通过它来分发。Dispatcher处理Action的分发，维护Store之间的依赖关系，负责处理View和Store之间建立Action的传递。

Dispatcher用于将Action分发给Store注册的回调函数，与普通的发布-订阅模式有两点不同

- 回调函数不是订阅到某一个特定的事件或频道，每个动作会分发给所有注册的回调函数
- 回调函数可以指定在其他回调函数之后调用。

Dispatcher通常是应用级单例，所以一个应用只需要一个dispatcher即可

```js
var Dispatcher = require('flux').Dispatcher;
var AppDispatcher = new Dispatcher();
```

Dispatcher包含的信息

- callbacks：DispatchToken和函数回调的字典
- isDispatching：展示当前Dispatcher是否处于Dispatch状态
- isHandled：Token检测一个函数是否被处理过
- isPending：Token检测一个函数是否被提交Dispatcher过
- lastID：最近依次被加入Dispatcher的函数体的唯一ID，即DispatchToken
- pendingPayload：需要传递给调用函数的数据

Dispatcher类的函数

(1)、register()函数：用于注册一个回调函数进入Dispatch，同时为这个callback生成DispatchToken，并加入字典。

```js
register(callback:(payload: TPayload)=>void):DispatchToken{
    var id=_prefix + this._ladtID++;
    this._callbacks[id] = callback;
    return id;
}
```

(2)、unregister()函数：可以通过DispatchToken将callback从字典中删除。

```js
unregister(id: DispatchToken): void {
    invariant(
      this._callbacks[id],
      'Dispatcher.unregister(...): '%s' does not map to a registered callback.',
      id
    );
    delete this._callbacks[id];
}
```

**注意**：invariant是一个用于描述错误的Node包，接收两个参数，第1个参数是条件，第2个参数是报错信息描述。

(3)、waitFor()函数：是一个等待函数，当执行到某些依赖的条件不满足时，就去等待它完成。

```js
waitFor(ids:Array<DispatchToken>):void{
    // 判断当前是否处于Dispatching，不处于Dispatching状态就不执行当前函数
    invariant(
        this._isDispatching,
        'Dispatcher.waitFor(...): Must be invoked while dispatching.'
    );
    // 从DispatchToken的数组中进行遍历，如果遍历到的DispatchToken处于Pending状态， =就暂时跳过
    for (var ii = 0; ii < ids.length; ii++) {
        var id = ids[ii];
        if (this._isPending[id]) {
            invariant(
                this._isHandled[id],
                'Dispatcher.waitFor(...): Circular dependency detected while ' +
                'waiting for '%s'.',
                id
            );
            continue;
        }
        // 检查Token对应的callback是否存在
        invariant(
            this._callbacks[id],
            'Dispatcher.waitFor(...):'%s' does not map to a registered callback.',
            id
        );
        //调用对应Token的callback函数
        this._invokeCallback(id);
    }
}
```

(4)、dispatch()函数：Dispatcher用于分发payload的函数。首先判断当前Dispatcher是否已经处于Dispatching状态中了。如果是，就不去打断。然后通过_startDispatching更新状态。更新状态结束以后，将非Pending状态的callback通过_invokeCallback执行（Pending在这里的含义可以简单理解为还没准备好或者被卡住了）。所有任务执行完成以后，通过_stopDispatching恢复状态。

```js
dispatch(payload: TPayload): void {
  invariant(
    !this._isDispatching,
    'Dispatch.dispatch(...): Cannot dispatch in the middle of a dispatch.'
  );
  this._startDispatching(payload);
  try {
    for (var id in this._callbacks) {
      if (this._isPending[id]) {
        continue;
      }
      this._invokeCallback(id);
    }
  } finally {
    this._stopDispatching();
  }
}
```

(5)、_startDispatching()函数：该函数将所有注册的callback的状态都清空，并标记Dispatcher的状态进入Dispatching。

```js
_startDispatching(payload: TPayload): void {
  for (var id in this._callbacks) {
    this._isPending[id] = false;
    this._isHandled[id] = false;
  }
  this._pendingPayload = payload;
  this._isDispatching = true;
}
```

(6)、_stopDispatching()函数：删除传递给callback的参数_pendingPayload，让当前Dispatcher不再处于Dispatch状态。

```js
_stopDispatching(): void {
  delete this._pendingPayload;
  this._isDispatching = false;
}
```

Dispatch要点小结

- register(function callback)：注册回调函数，返回一个可被waitFor()使用的Token；
- unregister(string id)：通过Token移除回调函数；
- waitFor(array ids)：在指定的回调函数执行之后才执行当前回调，该方法只能在分发动作的回调函数中使用；
- dispatch(object payload)：给所有注册的回调函数分发一个payload；
- isDispatching()：boolean值，返回Dispatcher当前是否处在分发的状态。

2、Action简介

Action可以看作是一个交互动作，改变应用状态或View的更新，都需要通过触发Action来实现。Action执行的结果就是调用了Dispatcher来处理相应的事情。Action是所有交互的入口，改变应用的状态或者有View需要更新时就需要通过Action实现。

Action是一个JavaScript对象，用来描述一个行为，里面包含了相关的信息。

Action的写法

```js
{
  actionName: "create-post"
  data: {
    content: "new post"
  }
}
```

Action对象主要由两部分构成：type(类型)和payload(载荷)。type是一个字符串常量，用于表示这个动作的标记。所谓的动作(Action)就是用来封装传递数据的。

```js
import AppDispatcher from './AppDispatcher';

var Actions = {
  addTodo(text){
    AppDispatcher.dispatch({
      type: 'ADD_TODO',
      text,
    });
  },
  // 其他的Actions
};
```

3、Store和View简介

Store包含应用状态和逻辑，不同的Store管理不同的应用状态。Store负责保存数据和定义修改数据的逻辑，同时调用Dispatcher的register()方法将自己设为监听器。每当发起一个Action(动作)去触发Dispatcher，Store的监听器就会被调用，用于执行是否更新数据的操作

Store在Flux中的特性是，管理应用所有额数据；只对外暴露getter方法，用于获取Store的数据，而没有暴露setter方法，这意味着不饿能通过Store去修改数据。若要修改Store数据，必须通过Action动作去触发Dispatcher实现。

只要Store发生变更，它就会使用emit()方法通知View更新并展示新的数据。

### 1.3 Flux缺点

主要体现在增加了代码量，使用Flux会让项目带入大量的概念和文件；单元测试难以进行。

## 2. Redux状态管理工具

### 2.1 Redux简介

Redux是一个“可预测的状态容器”，而实质也是Flux里面“单向数据流”的思想，但它充分利用函数式的特性，让整个实现更加优雅纯粹，使用更简单。

### 2.2 Redux三大特性

1、单一数据源

整个应用的state都存在一个JavaScript对象——Store中，可以将Store理解为全局的一个变量，且全局只有一个Store。

2、state是只读的

改变state的唯一方法就是触发Action，Action是一个普通的JavaScript对象，用于描述发生的事件。Store中有一个dispatch，dispatch接收Action参数，然后通过Store.dispatch(Action)来改变Store中的state。

```js
store.dispatch({
  type:types.RECEIVE_PRODUCTS, // Action名称
  products  // products表示该Action携带的状态，然后将存储在Store中
})
```

3、使用纯函数执行修改

纯函数：一个函数的返回结果只依赖于它的参数，相同的输入，永远只会得到相同的输出，而且没有任何可观察的副作用。

纯函数特性：

- 可缓存性（Cacheable），总能根据输入来做缓存
- 可移植性/自文档化（Portable/Self-Documenting），完全自给自足，纯函数需要的所有内容都能轻易获得。
- 可测试性（Testable），纯函数在测试时只需要简单地输入一个输入值，然后断言输出即可。
- 引用透明性（Referential Transparency），如果一个表达式在程序中可以被它等价的值替换而不影响结果，那么就说这段代码是引用透明的。传函数总能根据相同的输入返回相同的输出，所以能保证总是返回同一个结果。
- 并行代码，可以并行运行任意传函数。传函数不需要访问共享的内存，而且根据其定义，纯函数也不会因副作用而进入竞争状态(Race Condition)

纯函数Reducer作用是为了描述Action如何改变状态树。Reducer接收之前的state和Action，并返回新的state。

Reducer示例

```js
// 下面是一个Reducer，它负责处理Action，返回新的state

const addedlds = (state = [], action)=>{
  switch(action.type){
    case ADD_TO_CART:
      return [...state, action.productId]
    default:
      return state
  }
}
```

### 2.3 Redux的组成

Redux由3部分组成：Action、Reducer和Store。

1、Action

用来表达动作，Reducer根据Action更新State。

2、reducer

用来修改状态，定义整个应用的状态如何更改。

3、store

Redux中全局只有一个Store，用于存储整个应用的状态。

4个常用的api

- getSate()方法用于获取state
- dispatch(action)方法用于执行一个Action
- subscribe(listener)用于注册回调、监听state变化
- replaceReducer(nextReducer)更新当前Store内的Reducer

store是通过Redux提供的`createStore()`方法来创建的。`createStore()`方法有3个参数

- reducer(Function)：接收两个参数，分别是当前的state和要执行的Action，并返回新的state
- [preloadedState](any)：可选参数。初始state状态，在同构应用中可以使用这个参数去初始化state，或者从之前保存的用户会话中恢复并传给他。如果使用combineReducers()创建Reducer，它必须是一个普通对象，与传入的keys保持同样的结构；否则，可以任意传入任何Reducer能理解和识别的内容。
- enhancer(Function)：可选参数。enhancer就是指store enhancer，增强Store的功能。它是一个高阶函数，其参数是创建Store的函数，允许通过复合函数改变Store接口。

```js
import {createStore} from 'redux';

const addToCart = (state=initialState.quantityById, action)=>{
  switch(action.type){
    case ADD_TO_CART:
      const{productId} = action
      return {
        ...state, [productId]:(state[productId]||0)+1
      }
    default:
      return state
  }
}

let store = createStore(addToCart, ['hello', 'redux'])
```

### 2.4 Redux搭配React使用

安装react-redux

> npm install --save react-redux

react-redux提供了两个重要对象：Provider和connect

1、Provider简介

使用react-redux时需要先在最顶层创建一个Provider组件，用于将所有的React组件包裹起来，从而使React的左右组件都成为Provider的子组件。然后将创建好的Store作为Provider的属性传递给Provider

Provider示例

```js
import React from 'react';
import {render} from 'react-dom';
import {createStore} from 'redux';
import {Provider} from 'react-redux';
// 导入reducer
import reducer from './reducers';
// store创建
const store = createStore(
  reducer
)

render(
  <Provider store={store}> // Provider需要包裹在整个应用组件的外部
  <App /> // React组件
  </Provider>,
  document.getElementById('root')
)
```

2、connect简介

connect主要作用时连接react组件与Redux Store。当前组件可以通过props获取应用中的state和Actions。

connect接收4个参数：mapStateToProps()、mapDispatchToProps()、mergeProps()和options。

(1)、`[mapStateToProps(state,ownProps):stateProps]`：该函数允许开发者将Redux的Store中的数据作为props绑定到组件上，返回对象的所有key都为当前组件的props。此时该组件回监听store的变化。一旦Store发生变化，mapStateToProps()函数就会被调用，并返回一个纯对象，这个对象将与当前组件的props合并。该函数的第2个参数ownProps表示当前组件自身的props。只要组件接收到新的props，mapStateToProps()就会被调用，计算出一个新的stateProps提供给组件使用。

```js
const mapStateToProps = (state)=>{
  // store的第一个参数就是Redux的Store
  return{
    list: state.list // 从Redux的Store中获取list
  }
}
```

(2)、`[mapDispatchToProps(dispatch,[ownProps]):dispatchProps]`：它的功能更是将Action作为props绑定到当前组件。返回当前组件的actioncreator，并与dispatch绑定。如果定了第2个参数ownProps，该参数的值将会传递到该组件的props中，之后一旦组件收到新的props，mapDispatchToProps()也会被调用

(3)、`[mergeProps(stateProps, dispatchProps, ownProps):props]`：如果指定了这个函数，可以将mapStateToProps()和mapDispatchToProps()返回值和当前组件的props作为参数，返回一个完整的props。不管是要和ownProps合并之后才会被赋值给组件使用。如果不传递这个参数，connect将使用Object.assign代替

(4)、`[option]`：可选的额外配置项，一般不太使用。如果指定这个参数，可以去制定connector的行为。

mapSateToProps()和mapDispatchToProps()函数示例

```js
import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import * as cartActions from '../../actions';
import ItemComponent from './ItemComponent';
import CartComponent from './CartComponent';
const App = React.createClass({
    render() {
        const lists = this.props;
        const cartActions = this.props.cartActions;
        return(
            <div>
                <ItemComponent {...lists} />
                <CartComponent {...cartActions} />
            </div>
        )
    }
})
// connect包裹App组件
export default connect(state => ({
    lists: state.list
}), dispatch => ({
    cartActions: bindActionCreators(cartActions, dispatch)
}))(App)
```

上述例子中，connect和Provider组件相互配合将Actions和lists以props的形式传入组件App中。在mapStateToProps()中，选取整个Store树种的list分支作为当前组件的props，并命名为lists。然后在组件种使用this.props.lists。在mapDispatchToProps()中，使用Redux提供的工具方法将cartActions与dispatch绑定，最后在组件中使用this.props.cartActions.

react-redux的connect属于高阶组件，它允许向一个现有组件添加新功能，同时不改变其结构，属于装饰器模式。

decorator属性应用示例

```js
import React from 'react';
import {render} from 'react-dom';
import connect from './connect';

@connect
class App extends React.Component{
  render(){
    return(<div>....</div>);
  }
}
```

上述代码中，组件中通过connect将Store中的数据转换为组件可用的数据，并且生产Action的派发哈桑农户。react-redux利用connect将当前“木偶组件”进行包裹，并为该组件传递数据。

**木偶组件**(Dumb components)也叫UI组件，是与Redux没有直接联系的组件。该组件不知道Store或Action的存在，需要通过props传入组件，让组件得到且使用它们。

“木偶组件”，就是该组件能独立运作，不依赖于这个应用的Actions或Stores存在而存在。它不必存在state，也不能使用state，允许接收数据和数据改动，但只能通过props来处理。原则上它只负责展示，向一个“木偶”，收到外界props的控制。

木偶组件中的事件可以使用Actions方法的调用，Actions中的函数是经过bindActionCreators处理过的，会直接派发，从而改变Store数据，触发视图重渲染。重渲染的过程其实就是由于props传入新的数据通过比较后对DOM的更新。

如果不适用react-redux的Provider和connect这种机制，也可以直接使用Redux。但是本质是一样的，还是在React组件最外层将其包裹，作为props属性内层递进传递，但是不提倡这种做法

直接使用Redux示例

```js
class App extends Component{
  componentWillMount(){
    store.subscribe(state=>this.state(state))
  }
  render(){
    return(
      <BooksContainer state={this.state} onClick={()=>store.dispatch(actions.addToCart())} />
    );
  }
}
```

## 3 middleware中间件

### 3.1 什么是middleware

middleware本质上就是通过插件的形式，将原本Action->Reducer的流程改为Action->middleware1->middleware2...->Reducer。这正是Redux中间件middleware最优秀的特性，可以被链式组合和自由拔插。

1、手动记录

Redux的一个好处就是能让state的变化过程变得可预知和可追踪。每个Action发起后都会被计算并保存下来。

**需求**：需要记录应用中每一个Action发起的信息和state的状态。这样就可以通过查阅日志的方式找出哪个Action导致了state不正确。

最简单的做法是：在每次调用store.dispatch(action)前后手动记录。

记录日志示例：

```js
let action = actionTodo('Use Redux');
console.log('dispatching', action);
store.dispatch(action);
console.log('next state', store.getState());
```

2、封装dispatch()

有于手动记录代码是重复性额，因此可以将以上代码封装成一个函数。

记录日志的函数示例：

```js
function dispatchAndLog(store, action){
  console.log('dispatching', action);
  store.dispatch(action);
  console.log('next state', store.getSate());
}
```

3、替换dispatch()

由于Redux的Store是一个包含一些方法的普通对象，因此可以直接替换Redux的Store方法

```js
const next = store.dispatch;
// 获取Redux中的dispatch
store.dispatch = function dispatchAndLog(action){ // 重新覆盖原有的方法
  console.log('dispatching', action);
  let result = next(action);
  console.log('next state', store.getState());
  return result;
}
```

4、添加多个middleware

日志记录和报错记录示例：

```js
function patchStoreToAddLogging(store){
  const next = store.dispatch;
  store.dispatch = function.dispatchAndLog(action){
    console.log('dispatching', action);
    let result = next(action);
    console.log('next state', store.getState());
    return result;
  }
}

function patchStoreToAddCrashReporting(store){
  const next = store.dispatch;
  store.dispatch = function.dispatchAndReportErrors(action){
    try{
      return next(action);
    }catch(err){
      console.error('捕获一个异常！', err);
      Raven.captureException(err, {
        extra:{
          action,
          state: store.getState()
        }
      })
      throw err
    }
  }
}
```

(2)将store.dispatch的引用当作参数传递到另一个函数中，而不是直接改变它的值。

```js
function logger(store){
  const next = store.dispatch;
  //之前的做法
  // store.dispatch = function dispatchAndLog(action){}
  return function dispatchAndLog(action){
    console.log('dispatching', action);
    let result = next(action);
    console.log('next state', store.getState());
    return result;
  }
}
```

上面代码中的next()就是dispatch()，但是这个dispatch()函数每次执行时会保留上一个middleware传递的dispatch()函数的引用。接下来写一个方法将多个middleware连起来，该方法的主要作用是将上一次返回的函数赋值给store.dispatch

```js
function applyMiddlewareByMonkeypatching(store, middlewares){
  middlewares = middlewares.slice()
  middlewares.reverse()
  // 在每一个middleware中变换dispatch方法
  middlewares.forEach(middleware=>store.dispatch=middleware(store))
}
```

5、柯里化middleware

```js
function logger(store){
  return function wrapDispatchToAddLogging(next){
    return function dispatchAndLog(action){
      console.log('dispatching', action);
      let result = next(action);
      conole.log('next state', store.getState());
      return result;
    }
  }
}

```

ES6的箭头函数来表达函数柯里化(currying)：

```js
const logger = store => next => action =>{
  console.log('dispatching', action);
  let result = next(action);
  console.log('next state', store.getState());
  return result;
}
```

最终版：将logger和crashReporter引用到Redux Store：

```js
import {createStore, combineReducers, applyMiddleware} from 'redux';

let todoApp = combineReducers(reducers);
let store = createStore(
  todoApp, // applyMiddleware()告诉createStore()如何处理中间件
  applyMiddleware(logger, crashReporter)
)
```