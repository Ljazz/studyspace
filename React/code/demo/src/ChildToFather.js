/**
 * 子组件向父组件通信
 */
import React from 'react';

class Child extends React.Component{
    render(){
        return <input type="text" onChange={(e)=>this.props.handleChange(e.target.value)} />
    }
}

class App extends React.Component{
    constructor(props){
        super(props);
        this.state = {
            data: ''
        }
    }
    handleChange = text =>{
        this.setState({
            data: text
        })
    }
    render(){
        return(
            <div>
                <p>This message is from Child: {this.state.data}</p>
                <Child handleChange={this.handleChange} />
            </div>
        )
    }
}


export default App;