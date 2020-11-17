<!-- TOC -->

- [React的事件与表单](#react的事件与表单)
  - [1 事件系统](#1-事件系统)
    - [1.1 合成事件的事件代理](#11-合成事件的事件代理)
    - [1.2 事件的自动绑定](#12-事件的自动绑定)
    - [1.3 在React中使用原生事件](#13-在react中使用原生事件)
    - [1.4 合成事件与原生事件混用](#14-合成事件与原生事件混用)
  - [2 表单（Forms）](#2-表单forms)
    - [2.1 受控组件](#21-受控组件)
    - [2.2 非受控组件](#22-非受控组件)
    - [2.3 受控组件和非受控组件对比](#23-受控组件和非受控组件对比)
    - [2.4 表单组件得几个重要属性](#24-表单组件得几个重要属性)
  - [3 React的样式处理](#3-react的样式处理)
    - [3.1 基本样式设置](#31-基本样式设置)

<!-- /TOC -->

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
        document.body.addEventListener('click', e=>{
            //判断是否在模态框内单击
            if(e.target && e.target.matches('div.modal'))
                return;
            // 单击模态框之外关闭
            thi.setState({
                isModalShow: false
            })
        })
    }
    // 该生命周期表示组件将要卸载
    componentWillUnmount(){
        // 组件卸载之前先手动移除原生事件，避免出现内存泄漏
        document.body.removeEventListener('click');
    }
    render(){
        return(
            <div className="modal-wrapper">
            {
                this.state.isModalShow && // this.state.isModalShow为真就显示div
                <div className="modal">弹窗内容</div>
            }
            </div>
        )
    }
}
```

## 2 表单（Forms）

### 2.1 受控组件

React中强调的时对状态的可控管理，也就是对组件状态的可预知性和可测试性。表单状态时会随着用户在表单中的输入、选择或勾选等操作不断发生变化的，每当发生变化时，将它们的状态都写入组件的state中，这种组件就被称为受控组件(Controlled Component)。

```js
import  React, {Component} from 'react';
import {render} from 'react-dom';

class MyForm extends Component{
    render(){
        return(
            <input type="text" value="Hello Form" />
        )
    }
}

ReactDOM.render(
    <MyForm />, //将该组件挂载到id为root的真实DOM去渲染
    document.getElementById('root')
)
```

实时反应用户的输入示例

```js
import  React, {Component} from 'react';
import {render} from 'react-dom';

class MyForm extends Component{
    constructor(){
        super();
        this.state ={
            inputValue: "hello world"
        };
    }
    // input的change事件，每一次输入框的改变都会触发该事件
    handleChange(envent){
        this.setState({
            inputValue: event.target.value
        });
    }

    render(){
        return(
            <input type="text"
             vlaue={this.state.inputValue}
             onChange={this.handleChange.bind(this)} /> // 这里用bind()方法绑定this
        )
    }
}

ReactDOM.render(
    <MyForm />,
    document.getElementById('root')
)
```

上述代码中，将输入框输入的值绑定到组件的state，然后通过onChange事件就能让该值随用户的输入实时发生变化。

### 2.2 非受控组件

非受控组件即无约束组件，React强调的是对状态的可控管理，非受控组件是他的一种反模式。

非受控组件的示例

```js
import  React, {Component} from 'react';
import {render} from 'react-dom';

class MyForm extends Component{
    constructor(props){
        super(props);
        this.handleSubmit = this.handleSubmit.bind(this);
    }
    //表单提交事件
    handleSubmit(event){
        alert('this is a From');
        event.preventDefault();
    }

    render(){
        return(
            <form onSubmi={this.handleSubmit}>
                <lable>
                    Name:
                    <input type="text" />
                </lable>
                <lable>
                    Email:
                    <input type="email" />
                </lable>
                <button type="submit" value="Submit">submit</button>
            </form>
        );
    }
}
```

**注意**：如果想为非受控组件添加默认初始值，可以适用`defaultValue`或`defaultChecked`

### 2.3 受控组件和非受控组件对比

非受控组件本身有自己的状态缓存，而受控组件有React的state状态进行缓存。

```js
import React from 'react';

class NameForm extends React.Component {
    constructor(props) {
        super(props);
        this.handleSubmit = this.handleSubmit.bind(this);
        this.input = React.createRef();
    }

    handleSubmit(event) {
        alert('A name was submitted: ' + this.input.current.value);
        event.preventDefault();
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <label>
                    Name:
                    <input type="text" ref={this.input} />
                </label>
                <input type="submit" value="Submit" />
            </form>
        );
    }
}
```

受控组件示例

```js
import React, {Component} from 'react';

class MyForm1 extends Component{
    constructor(props){
        super(props);
        this.state={
            vlaue: ''// 在state中可以设置input的初始值，这里为空
        };
        this.onChange = this.onChange.bind(this);
    }
    onChange(event){
        this.setState({
            value: event.target.value
        })
    }
    render(){
        return(
            <div>
                <input type='text' value={this.state.value}
                onChange={this.onChange} />
                <p>您刚刚输入的内容为：{this.state.value}</p>
            </div>
        );
    }
}
```

受控组件的

- 优点是用户输入和页面显示之间做了一道可控层，可以在用户输入之后和页面显示之前对输入值进行处理。
- 缺点是需要为每个表单组件都绑定一个change事件，并且定义一个事件处理器去绑定表单值和组件的状态，而且每次表单值得改变都必定会调用一次onChange事件，带来性能上得损耗。

**注意**：Model层数据改变，View层随之同步更新就是**单向绑定**；View层代码改变，Model层随之更新就是**双向数据绑定**。

### 2.4 表单组件得几个重要属性

- value：用于`<input>`和`<textarea>`组件，类型为text
- checked：用于`<checkbox>`和`<radio>`组件，类型为Boolean
- selected：用于`<select>`组件下面得`<option>`

## 3 React的样式处理

### 3.1 基本样式设置

传统样式可以通过在标签内生命`class`和`id`名去定义，React也是这样。由于`class`在JavaScript中是保留字，为了解决“误解”，JSX语法中声明的标签属性中`class`必须使用`className`

**React中的组件的样式设计实例**：

className支持常规DOM元素和SVG元素，如：

> `<div className="button">clcik me</div>`

将定义的css文件在该组件顶部引用

```js
import React from 'react';
import './index.css';

class App extends React.Component{
    render(){
        return(
            <div className="btn">This is a demo</div>
        );
    }
}
```

行内样式style属性并不是以字符串形式被接收的，而是以一个带有驼峰命名的对象形式出现的。这个样式对象的key为驼峰命名规则的样式描述，对应的值通常是一个字符串或数字。

```js
class App extends React.Component{
    render(){
        const divStyle = {
            color: 'red',
            fontSize: 12,
            backgroundColor: "yellow"
        };
        return(
            <div style={divStyle}>This is a demo</div>
        );
    }
}
```

不能被转换成像素字符串的属性

```js
/**
 * CSS中接受数字但不以px为单位的属性
 */
var isUnitlessNumber = {
  animationIterationCount: true,
  borderImageOutset: true,
  borderImageSlice: true,
  borderImageWidth: true,
  boxFlex: true,
  boxFlexGroup: true,
  boxOrdinalGroup: true,
  columnCount: true,
  columns: true,
  flex: true,
  flexGrow: true,
  flexPositive: true,
  flexShrink: true,
  flexNegative: true,
  flexOrder: true,
  gridRow: true,
  gridRowEnd: true,
  gridRowSpan: true,
  gridRowStart: true,
  gridColumn: true,
  gridColumnEnd: true,
  gridColumnSpan: true,
  gridColumnStart: true,
  fontWeight: true,
  lineClamp: true,
  lineHeight: true,
  opacity: true,
  order: true,
  orphans: true,
  tabSize: true,
  widows: true,
  zIndex: true,
  zoom: true,
  // SVG-related 属性
  fillOpacity: true,
  floodOpacity: true,
  stopOpacity: true,
  strokeDasharray: true,
  strokeDashoffset: true,
  strokeMiterlimit: true,
  strokeOpacity: true,
  strokeWidth: true,
};
```
