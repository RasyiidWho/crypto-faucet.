import React, { Component } from 'react';
import CreateTransaction from '../containers/CreateTransaction'
import TransactionLogs from '../containers/TransactionLogs';
import Header from '../components/Header'
import '../../node_modules/bootstrap/dist/css/bootstrap.min.css';
import '../assets/css/body-style.css'

class HomePage extends Component {
  render() {
    return (
      <div>
        <Header/>
        <div className="column div-body-container">
          <div>
            <CreateTransaction />
          </div>
          <div>
            <TransactionLogs />
          </div>
        </div>
      </div>
    );
  }
}

export default HomePage;