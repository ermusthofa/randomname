import React from 'react';
import ReactDOM from 'react-dom';
import './style/index.css';
import Main from './component/Main';
import Footer from './component/Footer';
import * as serviceWorker from './component/serviceWorker';

ReactDOM.render(
  <React.StrictMode>
    <Main />
    <Footer />
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();
