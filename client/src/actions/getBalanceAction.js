import {  
        FETCH_BALANCE, 
        INTERVAL_GET_BALANCE
       } from './constants';
import BalanceApi from '../api/getBalanceApi'


export const getPRVBalance = () => {
    return (dispatch) => {
            return setInterval(() => BalanceApi.getBalance()
            .then(response => {
                const balance = response.data
                dispatch(fetchPRVBalance(balance))
            })
            .catch(error => {
                throw(error);
            }), INTERVAL_GET_BALANCE)
    };
};

export const fetchPRVBalance = (balance) => {
    return {
        type: FETCH_BALANCE,
        balance
    }
};
