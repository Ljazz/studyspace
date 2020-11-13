# React的事件与表单

## 1 事件系统

React有自己的事件系统，定义的事件处理器会接收到合成事件(SyntheticEvent)的实例

React事件系统和浏览器事件系统相比主要做了两件事：事件代理和事件自动绑定。这两个特性也正是合成事件的实现机制。

React的事件书写方式与传统HTML事件监听器的书写基本相似。React事件书写采用驼峰方式。

示例：给一个按钮添加单击事件

```js
// this.handleClick表示当前组件中定义的事件
<button onClick={this.handleClick}>Click me</button>
```

### 1.1 合成事件的事件代理

跟传统的事件处理机制不同的是React将所有自定义事件都绑定到结构的最顶层。使用一个事件监听器watch所有事件，并且内部包含一个映射表，记录事件与组件事件处理函数之间的对应关系。当事件出发时，React会根据映射关系找到真正的事件处理函数并调用。当组件被安装或被卸载时，对应的函数会被自动添加到事件监听器的内部映射表或者从表中删除。

### 1.2 事件的自动绑定

在React中，所有事件会被自动绑定到组件实例，并且会对该引用进行缓存，从而实现CPU和内存性能上的优化。但是如果使用ES6 Class或无状态的函数式写法，默认不会绑定this。在调用方法时需要手动绑定this。

常见的绑定方式，以button标签添加一个单击事件为例

1、在构造函数中使用bind()绑定this

在构造函数constructor()内使用bind()绑定this，等之后调用这个方法时无须再次绑定。

```js
import React, {Component} from 'react';
import {render} from 'react-dom';

class Button extends Component{
    constructor(props){
        super(props);
        // 在构造函数内完成this绑定
        this.handleClick = this.handleClick.bind(this);
    }
    // 自定义单击事件
    handleClick(){
        console.log('Clicked');
    }
    render(){
        return(
            <button onClick={this.state.handleClick}>
            click me
            </button>
        )
    }
}
```

2、使用箭头函数绑定this

每次调用函数时去绑定this，会生成一个新的方法实例，对性能有一定的影响。

```js
import React, {Component} from 'react';
import {render} from 'react-dom';

class Button extends Component{
    // 自定义单击事件
    handleClick(){
        console.log('Clicked');
    }
    render(){
        return(
            <button onClick={()=>this.handleClick()}>   // 利用箭头函数绑定
            click me
            </button>
        )
    }
}
```

3、使用bind()方法绑定this

```js
import React, {Component} from 'react';
import {render} from 'react-dom';

class Button extends Component{
    // 自定义单击事件
    handleClick(){
        console.log('Clicked');
    }
    render(){
        return(
            <button onClick={this.handleClick.bind(this)}>   // 用bind()方法绑定
            click me
            </button>
        )
    }
}
```

4、使用属性初始化器语法绑定this

因为使用属性初始化器语法绑定this的方法在创建时使用了箭头函数，所以一创建就绑定了this，因此在调用的时候无序再次绑定。

```js
import React, {Component} from 'react';
import {render} from 'react-dom';

class Button extends Component{
    // 用箭头函数定义单击事件，直接绑定到this
    handleClick=()=>{
        console.log('Clicked');
    }
    render(){
        return(
            <button onClick={this.handleClick}>
            click me
            </button>
        )
    }
}
```

### 1.3 在React中使用原生事件

在需要使用浏览器原生事件时，可以通过nativeEvent属性去获取和使用。由于原生事件需要绑定在真实的DOM中，所以一般在componentDidMount()生命周期阶段进行绑定操作。

注意：在React中使用DOM原生事件记得必须在组件卸载时是手动移除，否则将可能出现内存泄漏问题。

使用原生事件示例：

```js
import  React, {Component} from 'react';
import {render} from 'react-dom';

class NativeEvent extends Component{
    // 真实DOM加载完成才能取绑定原生事件到节点
    componentDidMount(){
        this.refs.myDiv.addEventListener('click', handleClick, false)
    }
    // 单击事件
    handleClick = (e) => {
        console.log(e)
    }
    // 组件卸载时，需要移除
    componentWillUnmount(){
        this.refs.myDiv.removeEventListener('click')
    }
    render(){
        return(
            <div ref="myDiv">click me</div>
        );
    }
}
ReactDOM.render(
    <NativeEvent />,
    document.getElementById('root')
);
```

### 1.4 合成事件与原生事件混用

若页面有个模态框，当单击模态框周围区域，需要将模态框隐藏。由于无法将模态框组件中的事件绑定到body上，因此只能通过原生事件在弹窗挂载到DOM完成之后获取body，然后通过原生事件绑定到body取实现。

```js
import  React, {Component} from 'react';
import {render} from 'react-dom';

class Demo extends Component{
    constructor(props){
        super(props);
        this.state = {
            isModalshow: true    // 默认弹窗打开
        };
        // 将this绑定到单击事件
        this.handleClickModal = this.handleClickModal.bind(this);
    }
    componentDidMount(){
        
    }
}
```
