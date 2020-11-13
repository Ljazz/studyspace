/**
 * context使用跨级组件通信
 */

import React from "react";
import PropTypes from 'prop-types';

// 子（孙）组件
class Button extends React.Component{
    render(){
        return(
            <button style={{background: this.context.color}}>
                {this.props.children}
            </button>
        );
    }
}

// 声明contextTypes用于访问MessageList中定义的context数据
Button.contextTypes = {
    color: PropTypes.string
};

// 中间件
class Message extends React.Component{
    render(){
        return(
            <div>
                <Button>Delete</Button>
            </div>
        );
    }
}

// 父组件

class MessageList extends React.Component{
    // 定义context需要实现的方法
    getChildContext(){
        return {
            color: 'orange'
        };
    }
    render(){
        return <Message />
    }
}

MessageList.childContextTypes = {
    color: PropTypes.string
};

export default MessageList;