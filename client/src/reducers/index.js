import { combineReducers } from 'redux';
import balanceReducer from './balanceReducer';
import infoReducer from './infoReducer';
import transactionReducer  from './transactionReducer';

export default combineReducers({
    transactions: transactionReducer,
    balance: balanceReducer,
    info: infoReducer
});