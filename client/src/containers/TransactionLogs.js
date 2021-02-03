import React from 'react';
import { connect } from 'react-redux';
import {
       NULL,
       ZERO,
       THREE_DOT,
       TRANSACTION_STATUS_PENDING,
       TRANSACTION_STATUS_SUCCESS,
       FIRST_CHARACTERS,
       LAST_CHARACTER,
       ENUM_STATUS_SUCCESS,
       ID,
       HASH,
       ADDRESS,
       CREATED_DATE, 
       STATUS_TRANSACTION,
       NOTIFICATION_NO_TRANSACTION
} from './constant'
import '../../node_modules/bootstrap/dist/css/bootstrap.min.css';
import "../assets/css/body-style.css"

function truncatedString (param) {
    if (param) {
        if (param.length > ZERO) {
            let firstCharacters = param.slice(0, FIRST_CHARACTERS)
            let lastCharacters = param.slice(param.length - LAST_CHARACTER)
            return firstCharacters.concat(THREE_DOT).concat(lastCharacters)
        }
        return NULL
    }
    else {
        return THREE_DOT
    }
}

function setStatus (param) {
    if (param) {
        if (param === ENUM_STATUS_SUCCESS) {
            return TRANSACTION_STATUS_SUCCESS
        }
        else {
            return TRANSACTION_STATUS_PENDING
        }
    }
    else {
        return TRANSACTION_STATUS_PENDING
    }
}

const headerTable = 
        <table className="table transaction-logs-container">
          <thead className="thead-dark div-container">
              <tr>
                <th scope="col">{ID}</th>
                <th scope="col">{ADDRESS}</th>
                <th scope="col">{HASH}</th>
                <th scope="col">{CREATED_DATE}</th>
                <th scope="col">{STATUS_TRANSACTION}</th>
              </tr>
          </thead>
        </table>

function TransactionLogs({ transactions }) {
  if(!transactions.length) {
    return (
      <div className="div-container">
        { headerTable }
        {NOTIFICATION_NO_TRANSACTION}
      </div>
    )
  }
  return (
    <div className="div-container">
        <table className="table text-center">
        <thead className="thead-dark div-container">
            <tr>
              <th scope="col">{ID}</th>
              <th scope="col">{ADDRESS}</th>
              <th scope="col">{HASH}</th>
              <th scope="col">{CREATED_DATE}</th>
              <th scope="col">{STATUS_TRANSACTION}</th>
            </tr>
        </thead>
        <tbody>
            {
              transactions.map((tx, index) => 
                  <tr key={index}>
                      <th scope="row">{tx.transactionID}</th>
                      <td>{truncatedString(tx.paymentAddress)}</td>
                      <td>{truncatedString(tx.hash)}</td>
                      <td>{tx.createdDate}</td>
                      <td>{setStatus(tx.status)}</td>
                  </tr>    
                  )
            }
            
        </tbody>
        </table>
    </div>
  );
}

const mapStateToProps = state => {
  return {
    transactions: state.transactions
  };
};

export default connect(
  mapStateToProps,
)(TransactionLogs);