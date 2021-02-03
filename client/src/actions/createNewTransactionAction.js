import { ADD_TRANSACTION } from './constants';

import CreateTransactionsApi from '../api/createTransactionApi'

export const createTransactionRequest = ({ paymentAddress }) => {
  return (dispatch) => {
    return CreateTransactionsApi.createTransactions(paymentAddress)
            .then(response => {
                const transantion = response.data.data
                dispatch(createTransaction(transantion))
            })
            .catch(error => {
                throw(error);
            });
  };
};

export const createTransaction = (data) => {
  return {
    type: ADD_TRANSACTION,
    payload: {
      data
    }
  }
};
