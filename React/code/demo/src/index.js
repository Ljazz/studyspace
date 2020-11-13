import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import reportWebVitals from './reportWebVitals';
import App from './ChildToFather';
import MessageList from './context';
import NotNestedComponentCommunication from './notNestedComponentCommunication';

ReactDOM.render(
  <React.StrictMode>
    <App />
    <hr />
    <MessageList />
    <hr />
    <NotNestedComponentCommunication />
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
