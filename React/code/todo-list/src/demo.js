import React from 'react';
import axios from 'axios';


class Demo extends React.Component{

    constructor(){  // 构造方法
        super();
        this.state = store.getState();

        store.subscribe(()=>{
            this.setState(store.getState());
        })
    }

    componentDidMount(){
        axios.get('url').then((resp)=>{

        })
    }

    getList(){
        return (dispatch)=>{
            
        }
    }

    // 模拟获取数据，例如使用ajax
    getDatas(){
        let list = [4,5,6];

        return list;
    }
}