import React from 'react';
import 'bootstrap/dist/css/bootstrap.css';
import './index.scss';

class App extends React.Component{
  constructor(props){
    super(props);

    this.state = {
      items: [
        { id: 1, name: 'Cras justo odio'},
        { id: 2, name: 'Dapibus ac facilisis in'},
        { id: 3, name: 'Morbi leo risus'},
        { id: 4, name: 'Porta ac consectetur ac'},
        { id: 5, name: 'Vestibulum at eros'}
      ]
    }
  }

  render(){
    var items = this.state.items;
    return(
      <div className="items-panel">
        <ul>
          {
            items.map((item, index) => (
              <li key={item.id} className="list-group-item">
                <div className="name">{item.name}</div>
                <div className="btn-group">
                  {
                    index > 0 ?
                    <button className="btn btn-primary" onClick={() => this.sortUp(index)}>Up</button>: null
                  }
                  {
                    index < items.length - 1 ?
                    <button className="btn btn-dark" onClick={() => this.sortDown(index)}>Down</button>: null
                  }
                </div>
              </li>
            ))
          }
        </ul>
      </div>
    );
  }

  sortUp(index){
    this.resort(index, -1);
  }

  sortDown(index){
    this.resort(index, 1);
  }

  resort(index, diff){
    var items = this.state.items;
    var item = items[index];
    items.splice(index, 1);
    items.splice(index+diff, 0, item);
    this.setState({items: items});
  }
}

export default App;
