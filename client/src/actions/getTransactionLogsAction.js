import { 
        FETCH_TRANSACTION, 
        INTERVAL_UPDATE_TRANSACTIONS_TABLE

  } from './constants';
import TransactionsApi from "../api/allTransactionApi"

export const fetchTransactions = (transactions) => {
    return {
      type: FETCH_TRANSACTION,
      transactions
    }
};
  
export const fetchAllTransactions = () => {
    return (dispatch) => {
      return setInterval(
        () => TransactionsApi.getTransactions()
        .then(response => {
        const transactionLogs = response.data.data.transactionLogDetails
        dispatch(fetchTransactions(transactionLogs))
        })
        .catch(error => {
        throw(error);
        }), INTERVAL_UPDATE_TRANSACTIONS_TABLE)
    };
};