import React, {Component} from 'react';
import emitter from './events';


class ComponentA extends Component{
    constructor(props){
        super(props);
        this.state = {
            data: 'aaaa',
        };
    }
    componentDidMount(){
        // 组件加载完成以后声明一个自定义事件
        // 绑定callMe事件，处理函数为addListener()的第2个参数
        this.eventEmitter = emitter.addListener("callMe", (data)=>{
            this.setState({
                data
            })
        });
    }
    componentWillUnmount(){
        // 组件销毁前移除事件监听
        emitter.removeListener(this.eventEmitter);
    }
    render(){
        return(
            <div>
                Hello, {this.state.data}
            </div>
        );
    }
}


class ComponentB extends Component{
    render(){
        const cb = (data) =>{
            return()=>{
                // 触发自定义事件
                // 可传多个参数
                emitter.emit("callMe", "React")
            }
        }
        return(
            <div>
                <button onClick={cb("hey")}>点击</button>
            </div>
        );
    }
}


export default class NotNestedComponentCommunication extends Component{
    render(){
        return(
            <div>
                <ComponentA />
                <ComponentB />
            </div>
        );
    }
}