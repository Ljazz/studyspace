
<!-- TOC -->

- [React入门](#react入门)
  - [1.1 React简介](#11-react简介)
    - [1.1.1 什么是React？](#111-什么是react)
    - [1.1.2 React的特点](#112-react的特点)
    - [1.1.3 React高效的原因](#113-react高效的原因)
  - [1.2 React的基本使用](#12-react的基本使用)
    - [1.2.1 相关的js库](#121-相关的js库)
    - [1.2.2 创建虚拟DOM的两种方式](#122-创建虚拟dom的两种方式)
  - [1.3 React JSX](#13-react-jsx)
    - [1.3.1 JSX](#131-jsx)
    - [1.3.2 渲染虚拟DOM（元素）](#132-渲染虚拟dom元素)
  - [1.4 模块与组件、模块化与组件化的理解](#14-模块与组件模块化与组件化的理解)

<!-- /TOC -->
# React入门

## 1.1 React简介

### 1.1.1 什么是React？

用于动态构建用户界面的JavaScript库（只关注于视图）

### 1.1.2 React的特点

- 声明式编码
- 组件化编码
- React Native编写原生应用
- 高效（优秀的Diffing算法）

### 1.1.3 React高效的原因

- 使用虚拟(virtual)DOM，不总是直接操作页面真实DOM
- DOM Diffing算法，最小化页面重绘

## 1.2 React的基本使用

### 1.2.1 相关的js库

- react.js：React核心库
- react-dom.js：提供操作DOM的react扩展库
- babel.min.js：解析JSX语法代码转为JS代码的库

### 1.2.2 创建虚拟DOM的两种方式

1. 纯JS方式

```js
<html>
<head>
    <meta charset="UTF-8">
    <title>使用JS创建虚拟DOM</title>
</head>
<body>
    <!-- 准备好一个“容器” -->
    <div id="demo"></div>

    <!-- 引入react核心库 -->
    <script src="https://cdn.staticfile.org/react/16.4.0/umd/react.development.js"></script>
    <!-- 引入react-dom 用于支持react操作DOM -->
    <script src="https://cdn.staticfile.org/react-dom/16.4.0/umd/react-dom.development.js"></script>
    <!-- 引入Babel 用于将JSX转为JS -->
    <script src="https://cdn.staticfile.org/babel-standalone/6.26.0/babel.min.js"></script>

    <script type="text/javascript"> 
        // 1. 创建虚拟DOM
        const VDOM = React.createElement('h1', {id: "title"}, "hello React")
        // 2. 渲染虚拟DOM到页面
        ReactDOM.render(VDOM, document.getElementById('demo'))
    </script>
</body>
</html>
```

2. JSX方式

```js
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>jsx语法规则</title>
</head>
<body>
    <!-- 准备好一个“容器” -->
    <div id="demo"></div>

    <!-- 引入react核心库 -->
    <script src="https://cdn.staticfile.org/react/16.4.0/umd/react.development.js"></script>
    <!-- 引入react-dom 用于支持react操作DOM -->
    <script src="https://cdn.staticfile.org/react-dom/16.4.0/umd/react-dom.development.js"></script>
    <!-- 引入Babel 用于将JSX转为JS -->
    <script src="https://cdn.staticfile.org/babel-standalone/6.26.0/babel.min.js"></script>

    <script type="text/babel"> /* 此处一定要写Babel */
        // 1. 创建虚拟DOM
        const VDOM = <h1>Hello React</h1>
        // 2. 渲染虚拟DOM到页面
        ReactDOM.render(VDOM, document.getElementById('demo'))
    </script>
</body>
</html>
```

虚拟DOM
- 本质是Object类型的对象
- 虚拟DOM属性比较少，真实DOM属性比价多。虚拟DOM只是在React内部用，无需真实DOM那么多的属性
- 虚拟DOM最终会被React转换为真实DOM，呈现在页面上

## 1.3 React JSX

### 1.3.1 JSX

1. 全称：JavaScript XML
2. react定义的一种类似于XML的JS扩展语法：JS + XML
3. 本质是React.createElement(component, props, ...children)方法的语法糖
4. 作用：用来简化创建虚拟DOM
   1. 写法：var ele = <h1>Hello JSX!</h1>
   2. 注意1：它不是字符串，也不是HTML/XML标签
   3. 注意2：它最终产生的就是一个JS对象
5. 标签名任意：HTML标签或其他标签
6. 基本语法规则
   1. 遇到`<`开头的代码，以标签的语法解析：html同名标签转化为html同名元素，其他标签需要特别解析
   2. 遇到以`{`开头的代码，以JS语法解析：标签中的js表达式必须用`{}`包含
7. babel.js的作用
   1. 浏览器不能直接解析JSX代码，需要babel转译为纯JS的代码才能运行
   2. 只要用了JSX，都要加上type="text/babel"，声明需要babel来处理

### 1.3.2 渲染虚拟DOM（元素）

1. 语法：`ReactDOM.render(virtualDOM, containerDOM)`
2. 作用：将虚拟DOM元素渲染到页面中的真实容器DOM中显示
3. 参数说明
   1. 参数1：纯JS或JSX创建的虚拟DOM对象
   2. 参数2：用来包含虚拟DOM元素的真实DOM元素对象（一般是一个div）

## 1.4 模块与组件、模块化与组件化的理解
