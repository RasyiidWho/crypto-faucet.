import '@fortawesome/fontawesome-free/css/all.min.css';
import 'bootstrap-css-only/css/bootstrap.min.css';
import 'mdbreact/dist/css/mdb.css';
import React from 'react';
import ReactDOM from 'react-dom';
import { createStore, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import { Provider } from 'react-redux';
import App from './App';
import rootReducer from './reducers';
import { fetchAllTransactions } from "./actions/getTransactionLogsAction";
import { getInfoCountOfTransactions } from './actions/getInfoTransactionsAction'
import { getPRVBalance } from './actions/getBalanceAction';
import reportWebVitals from './reportWebVitals';

const store = createStore(rootReducer, applyMiddleware(thunk));
store.dispatch(fetchAllTransactions());
store.dispatch(getPRVBalance());
store.dispatch(getInfoCountOfTransactions());
ReactDOM.render(
    <Provider store={store}>
        <App />
    </Provider>, document.getElementById('root'));
reportWebVitals();
