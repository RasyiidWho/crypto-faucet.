import { 
         FETCH_INFO_COUNT_OF_TRANSACTION,
         INTERVAL_GET_INFO_TRANSACTION 
       } from './constants';
import InfoTransactionApi from "../api/getInfoTransactionApi";

export const fetchInfoCountOfTransactions = (transactionsInfo) => {
    return {
      type: FETCH_INFO_COUNT_OF_TRANSACTION,
      transactionsInfo
    }
};
  
export const getInfoCountOfTransactions = () => {
    return (dispatch) => {
      return setInterval(() => InfoTransactionApi.getInfo()
            .then(response => {
            const transactionsInfo = response.data
            dispatch(fetchInfoCountOfTransactions(transactionsInfo))
            })
            .catch(error => {
            throw(error);
            }), INTERVAL_GET_INFO_TRANSACTION
        )
    };
};